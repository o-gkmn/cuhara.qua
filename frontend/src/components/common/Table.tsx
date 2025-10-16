import { useEffect, useState, type ReactNode, useRef } from "react"
import { FaFilter, FaPlus, FaTrash } from 'react-icons/fa';
import { FaArrowRotateRight, FaEllipsis, FaGear } from "react-icons/fa6";
import { FaAngleRight, FaAngleLeft } from 'react-icons/fa';
import { FaAngleDoubleRight, FaAngleDoubleLeft } from 'react-icons/fa';
import { FaArrowDown } from 'react-icons/fa';
import { FaGripVertical } from 'react-icons/fa6';
import { FaSortUp } from 'react-icons/fa6';
import { FaSortDown } from 'react-icons/fa6';

export type TextOperator = 'contains' | 'startsWith' | 'endsWith' | 'equals' | 'notEquals' | 'empty' | 'notEmpty'
export type NumberOperator = 'eq' | 'neq' | 'gt' | 'gte' | 'lt' | 'lte' | 'between' | 'in'
export type DateOperator = 'on' | 'before' | 'after' | 'between'
export type BooleanOperator = 'isTrue' | 'isFalse' | 'isNull' | 'isNotNull'
export type SelectOperator = 'eq' | 'neq'
export type SortOperator = 'asc' | 'desc'

type BaseFilter = {
    disabled?: boolean
    visible?: boolean
    operator?: string
    operators?: string[]
}

export type ColumnFilterProps =
    | (BaseFilter & {
        type: 'text'
        operator?: TextOperator
        operators?: TextOperator[]
        caseSensitive?: boolean
    })
    | (BaseFilter & {
        type: 'number'
        operator?: NumberOperator
        operators?: NumberOperator[]
    })
    | (BaseFilter & {
        type: 'date'
        operator?: DateOperator
        operators?: DateOperator[]
        format?: string
        utc?: string
    })
    | (BaseFilter & {
        type: 'boolean'
        operator?: BooleanOperator
        operators?: BooleanOperator[]
    })
    | (BaseFilter & {
        type: 'select'
        operator?: SelectOperator
        operators?: SelectOperator[]
        multiple?: boolean
        options: Array<{ label: string; value: string | number }>
    })

export type FilterParam = {
    field: string
    operator: string
    value: unknown
}

export interface TableColumnProps<T extends Record<string, unknown> = Record<string, unknown>> {
    key: keyof T & string
    label: string
    sortable?: boolean
    show?: boolean
    filter?: ColumnFilterProps
    render?: (item: T) => ReactNode
}

export interface TableProps<T extends Record<string, unknown> = Record<string, unknown>> {
    columns: TableColumnProps<T>[]
    data: T[],
    title?: string,
    onNew?: () => void
    onDelete?: (item: T) => void
    onRefresh?: () => void
    onExport?: () => void
    onFilter?: () => void
    onPageChange?: (page: number) => void
    onPageSizeChange?: (pageSize: number) => void
    onSorted?: (field: string, operator: SortOperator) => void
}

export default function Table<T extends Record<string, unknown>>(props: TableProps<T>) {
    const DEFAULT_BUTTONS = ["refresh", "export", "filter"]
    const OTHER_BUTTONS = ["new", "delete", "settings"]
    const ALL_BUTTONS = [...DEFAULT_BUTTONS, ...OTHER_BUTTONS]

    const [page, setPage] = useState(1)
    const [pageSize, setPageSize] = useState(10)
    const totalPage = Math.ceil(props.data.length / pageSize)

    const [activeButtons, setActiveButtons] = useState<string[]>(DEFAULT_BUTTONS)
    const isExpanded = activeButtons.length > DEFAULT_BUTTONS.length

    const [isSettingsOpen, setIsSettingsOpen] = useState(false)
    const [settingsButtonRef, setSettingsButtonRef] = useState<HTMLButtonElement | null>(null)
    const [settingsPopoverRef, setSettingsPopoverRef] = useState<HTMLDivElement | null>(null)
    const [settingsPosition, setSettingsPosition] = useState<{ top: number; left: number }>({ top: 0, left: 0 })
    const settingsCloseTimerRef = useRef<number | null>(null)

    const [sortColumn, setSortColumn] = useState<{ key: string, operator: SortOperator }>()

    type ColumnFilterState = {
        operator?: string
        value?: unknown
        value2?: unknown // used for between
    }
    const [isFilterOpen, setIsFilterOpen] = useState(false)
    const [filters, setFilters] = useState<Record<string, ColumnFilterState>>({})
    const [openFilterKey, setOpenFilterKey] = useState<string | null>(null)
    const [filterPopoverRef, setFilterPopoverRef] = useState<HTMLDivElement | null>(null)

    const [activeColumns, setActiveColumns] = useState<TableColumnProps<T>[]>(props.columns.filter((column) => column.show))
    const [columnOrder, setColumnOrder] = useState<string[]>(props.columns.map((c) => c.key))
    const [draggingKey, setDraggingKey] = useState<string | null>(null)

    useEffect(() => {
        if (!openFilterKey) return
        function handleDocumentClick(e: MouseEvent) {
            const target = e.target as Node | null
            if (!target) return
            if (filterPopoverRef && filterPopoverRef.contains(target)) return
            setOpenFilterKey(null)
        }
        document.addEventListener('mousedown', handleDocumentClick)
        return () => {
            document.removeEventListener('mousedown', handleDocumentClick)
        }
    }, [openFilterKey, filterPopoverRef])

    // Close settings popover on outside click, but don't close while dragging
    useEffect(() => {
        if (!isSettingsOpen) return
        function handleSettingsOutsideMouseDown(e: MouseEvent) {
            if (draggingKey) return
            const target = e.target as Node | null
            if (!target) return
            if (settingsPopoverRef && settingsPopoverRef.contains(target)) return
            if (settingsButtonRef && settingsButtonRef.contains(target)) return
            setIsSettingsOpen(false)
        }
        document.addEventListener('mousedown', handleSettingsOutsideMouseDown)
        return () => {
            document.removeEventListener('mousedown', handleSettingsOutsideMouseDown)
        }
    }, [isSettingsOpen, settingsPopoverRef, settingsButtonRef, draggingKey])

    // Close settings popover on Escape for usability
    useEffect(() => {
        if (!isSettingsOpen) return
        function handleEscape(e: KeyboardEvent) {
            if (draggingKey) return
            if (e.key === 'Escape') {
                setIsSettingsOpen(false)
            }
        }
        document.addEventListener('keydown', handleEscape)
        return () => {
            document.removeEventListener('keydown', handleEscape)
        }
    }, [isSettingsOpen, draggingKey])

    // Keep open when hovering button or popover; close when leaving both with a small grace period
    useEffect(() => {
        if (!isSettingsOpen) return
        function handlePointerMove(e: PointerEvent) {
            if (draggingKey) return
            const target = e.target as Node | null
            if (!target) return
            const overPopover = !!settingsPopoverRef && settingsPopoverRef.contains(target)
            const overButton = !!settingsButtonRef && settingsButtonRef.contains(target)
            if (overPopover || overButton) {
                if (settingsCloseTimerRef.current != null) {
                    window.clearTimeout(settingsCloseTimerRef.current)
                    settingsCloseTimerRef.current = null
                }
                return
            }
            if (settingsCloseTimerRef.current == null) {
                settingsCloseTimerRef.current = window.setTimeout(() => {
                    setIsSettingsOpen(false)
                    settingsCloseTimerRef.current = null
                }, 200)
            }
        }
        document.addEventListener('pointermove', handlePointerMove)
        return () => {
            document.removeEventListener('pointermove', handlePointerMove)
            if (settingsCloseTimerRef.current != null) {
                window.clearTimeout(settingsCloseTimerRef.current)
                settingsCloseTimerRef.current = null
            }
        }
    }, [isSettingsOpen, settingsPopoverRef, settingsButtonRef, draggingKey])

    // Compute and track settings popover position relative to the button
    useEffect(() => {
        if (!isSettingsOpen || !settingsButtonRef) return
        const updatePosition = () => {
            const rect = settingsButtonRef.getBoundingClientRect()
            setSettingsPosition({ top: rect.bottom + 8, left: rect.left })
        }
        updatePosition()
        // Listen to scroll on capture to catch scrolls in any ancestor
        window.addEventListener('scroll', updatePosition, true)
        window.addEventListener('resize', updatePosition)
        return () => {
            window.removeEventListener('scroll', updatePosition, true)
            window.removeEventListener('resize', updatePosition)
        }
    }, [isSettingsOpen, settingsButtonRef])

    const defaultRenderCell = (value: unknown): ReactNode => {
        if (value === null || value === undefined) return "-"
        if (typeof value === 'object') return JSON.stringify(value)
        return String(value)
    }

    function pageWindow() {
        const current = page ?? 1
        const total = totalPage ?? 1
        const pageWindow = [
            ...(current != 1 ? [current - 1] : []),
            current,
            current + 1,
            ...(current == 1 ? [current + 2] : []),
        ].filter(p => p >= 1 && p <= total)
        return pageWindow
    }

    function handleColumnChange(column: TableColumnProps<T>) {
        props.columns.find(c => c.key === column.key)!.show = !column.show
        setActiveColumns(props.columns.filter((c) => c.show))
    }

    function handleDragStart(key: string) {
        setDraggingKey(key)
    }

    function handleDrop(targetKey: string) {
        if (!draggingKey || draggingKey === targetKey) return
        const current = [...columnOrder]
        const fromIndex = current.indexOf(draggingKey)
        const toIndex = current.indexOf(targetKey)
        if (fromIndex === -1 || toIndex === -1) {
            setDraggingKey(null)
            return
        }
        current.splice(fromIndex, 1)
        current.splice(toIndex, 0, draggingKey)
        setColumnOrder(current)
        setDraggingKey(null)
    }

    function orderedActiveColumns(): TableColumnProps<T>[] {
        const keyToColumn = new Map(activeColumns.map((c) => [c.key, c]))
        return columnOrder
            .map((k) => keyToColumn.get(k))
            .filter((c): c is TableColumnProps<T> => Boolean(c))
    }

    function buildQuery(params: FilterParam[]): string {
        const query = params.map((param, i) => {
            const part = `${param.field} ${param.operator} ${param.value}`
            const hasNext = i < params.length - 1
            return hasNext ? `${part} AND` : part
        })
            .join(' ')

        console.log(query)
        return query
    }

    function applyFiltersFrom(source: Record<string, ColumnFilterState>) {
        const params: FilterParam[] = []
        for (const key of Object.keys(source)) {
            const state = source[key]
            if (!state) continue
            const col = props.columns.find(c => c.key === key)
            if (!col || !col.filter) continue

            const operator = state.operator ?? col.filter.operators?.[0] ?? ''
            if (!operator) continue


            // Skip if no value provided (except boolean-only operator without value)
            if (operator === 'between') {
                const v1 = state.value
                const v2 = state.value2
                if (v1 == null || v1 === '' || v2 == null || v2 === '') continue
                params.push({ field: key, operator, value: `${v1} AND ${v2}` })
            } else if (col.filter.type == 'boolean') {
                params.push({ field: key, operator, value: '' })
            }
            else {
                // for select/number/text/date single-value
                const v = state.value
                if (v == null || v === '') continue
                params.push({ field: key, operator, value: v })
            }
        }

        console.log(params)

        if (params.length) {
            buildQuery(params)
        }
    }

    function applyFilters() {
        applyFiltersFrom(filters)
    }

    function setFilterState(key: string, next: ColumnFilterState) {
        setFilters(prev => ({ ...prev, [key]: { ...prev[key], ...next } }))
    }

    function clearFilter(key: string) {
        setFilters(prev => {
            const copy = { ...prev }
            delete copy[key]
            return copy
        })
    }

    function isColumnFilterActive(column: TableColumnProps<T>): boolean {
        if (!column.filter) return false
        const state = filters[column.key]
        if (!state) return false

        const op = state.operator ?? column.filter.operators?.[0] ?? ''

        if (op === 'between') {
            const v1 = state.value
            const v2 = state.value2
            return v1 != null && v1 !== '' && v2 != null && v2 !== ''
        }

        if (column.filter.type === 'boolean') {
            return op !== ''
        }

        const v = state.value
        return v != null && v !== ''
    }

    function renderFilter(column: TableColumnProps<T>) {
        if (!column.filter) return null
        const f = column.filter
        const current = filters[column.key] || {}

        switch (f.type) {
            case 'text': {
                return (
                    <div ref={setFilterPopoverRef} onMouseDown={(e) => e.stopPropagation()} onClick={(e) => e.stopPropagation()} className="absolute top-full left-1/2 -translate-x-1/2 mt-1 z-20 bg-white text-slate-700 border border-slate-200 rounded shadow p-2 w-56">
                        <div className="flex flex-col gap-2">
                            <input
                                type="text"
                                placeholder="Ara"
                                className="w-full block rounded border border-slate-200 px-2 py-1 text-sm"
                                value={(current.value as string) ?? ''}
                                onChange={(e) => setFilterState(column.key, { value: e.target.value })}
                            />
                            {f.operators?.length ? (
                                <select
                                    className="w-full block border border-slate-200 rounded px-2 py-1 text-sm"
                                    value={current.operator ?? f.operators?.[0] ?? ''}
                                    onChange={(e) => setFilterState(column.key, { operator: e.target.value })}
                                >
                                    {f.operators.map(op => (
                                        <option className="text-sm" key={op} value={op}>{op}</option>
                                    ))}
                                </select>
                            ) : null}
                        </div>
                        <div className="flex justify-between gap-2 pt-2">
                            <button
                                className="text-xs text-teal-700 hover:underline"
                                onClick={() => setOpenFilterKey(null)}
                            >
                                Kapat
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    clearFilter(column.key)
                                }}
                            >
                                Sıfırla
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    applyFilters()
                                    setOpenFilterKey(null)
                                }}
                            >
                                Uygula
                            </button>
                        </div>
                    </div>
                )
            }

            case 'number': {
                return (
                    <div ref={setFilterPopoverRef} onMouseDown={(e) => e.stopPropagation()} onClick={(e) => e.stopPropagation()} className="absolute top-full left-1/2 -translate-x-1/2 mt-1 z-20 bg-white text-slate-700 border border-slate-200 rounded shadow p-2 w-56">

                        <div className="flex flex-col gap-2">
                            <input
                                type="number"
                                placeholder="Ara"
                                className="w-full block rounded border border-slate-200 px-2 py-1 text-sm"
                                value={(current.value as number | undefined) ?? ''}
                                onChange={(e) => setFilterState(column.key, { value: e.target.value === '' ? undefined : Number(e.target.value) })}
                            />
                            {f.operators?.length ? (
                                <select
                                    className="w-full block border border-slate-200 rounded px-2 py-1 text-sm"
                                    value={current.operator ?? f.operators?.[0] ?? ''}
                                    onChange={(e) => setFilterState(column.key, { operator: e.target.value })}
                                >
                                    {f.operators.map(op => (
                                        <option className="text-sm" key={op} value={op}>{op}</option>
                                    ))}
                                </select>
                            ) : null}
                        </div>
                        <div className="flex justify-between gap-2 pt-2">
                            <button
                                className="text-xs text-teal-700 hover:underline"
                                onClick={() => setOpenFilterKey(null)}
                            >
                                Kapat
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    clearFilter(column.key)
                                }}
                            >
                                Sıfırla
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    applyFilters()
                                    setOpenFilterKey(null)
                                }}
                            >
                                Uygula
                            </button>
                        </div>
                    </div>
                )
            }

            case 'boolean': {
                return (
                    <div ref={setFilterPopoverRef} onMouseDown={(e) => e.stopPropagation()} onClick={(e) => e.stopPropagation()} className="absolute top-full left-1/2 -translate-x-1/2 mt-1 z-20 bg-white text-slate-700 border border-slate-200 rounded shadow p-2 w-56">
                        <div className="flex flex-col gap-2">
                            {f.operators?.length ? (
                                <select
                                    className="w-full block border border-slate-200 rounded px-2 py-1 text-sm"
                                    value={current.operator ?? f.operators?.[0] ?? ''}
                                    onChange={(e) => {
                                        setFilterState(column.key, { operator: e.target.value })
                                    }}
                                >
                                    {f.operators.map(op => (
                                        <option className="text-sm" key={op} value={op}>{op}</option>
                                    ))}
                                </select>
                            ) : null}
                        </div>
                        <div className="flex justify-between gap-2 pt-2">
                            <button
                                className="text-xs text-teal-700 hover:underline"
                                onClick={() => setOpenFilterKey(null)}
                            >
                                Kapat
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    clearFilter(column.key)
                                }}
                            >
                                Sıfırla
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    if (!filters[column.key]) {
                                        const op = current.operator ?? f.operators?.[0] ?? ''
                                        const next = { ...current, operator: op }
                                        setFilters(prev => ({ ...prev, [column.key]: next }))
                                        const snapshot = { ...filters, [column.key]: next }
                                        applyFiltersFrom(snapshot)
                                    } else {
                                        applyFilters()
                                    }
                                    setOpenFilterKey(null)
                                }}
                            >
                                Uygula
                            </button>
                        </div>
                    </div>
                )
            }

            case 'select': {
                return (
                    <div ref={setFilterPopoverRef} onMouseDown={(e) => e.stopPropagation()} onClick={(e) => e.stopPropagation()} className="absolute top-full left-1/2 -translate-x-1/2 mt-1 z-20 bg-white text-slate-700 border border-slate-200 rounded shadow p-2 w-56">
                        <div className="flex flex-col gap-2">
                            <select
                                className="w-full block border border-slate-200 rounded px-2 py-1 text-sm"
                                value={(current.value as string | number | undefined) ?? ''}
                                onChange={(e) => setFilterState(column.key, { value: e.target.value })}
                            >
                                {f.options.map(op => (
                                    <option className="text-sm" key={op.value} value={op.value}>{op.label}</option>
                                ))}
                            </select>
                            {f.operators?.length ? (
                                <select
                                    className="w-full block border border-slate-200 rounded px-2 py-1 text-sm"
                                    value={current.operator ?? f.operators?.[0] ?? ''}
                                    onChange={(e) => setFilterState(column.key, { operator: e.target.value })}
                                >
                                    {f.operators.map(op => (
                                        <option className="text-sm" key={op} value={op}>{op}</option>
                                    ))}
                                </select>
                            ) : null}
                        </div>
                        <div className="flex justify-between gap-2 pt-2">
                            <button
                                className="text-xs text-teal-700 hover:underline"
                                onClick={() => setOpenFilterKey(null)}
                            >
                                Kapat
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    clearFilter(column.key)
                                }}
                            >
                                Sıfırla
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    if (!filters[column.key]) {
                                        const op = current.operator ?? f.operators?.[0] ?? ''
                                        const defaultVal = (current.value as string | number | undefined) ?? (column.filter && column.filter.type === 'select' ? column.filter.options?.[0]?.value : undefined)
                                        const next = { ...current, operator: op, value: defaultVal }
                                        setFilters(prev => ({ ...prev, [column.key]: next }))
                                        const snapshot = { ...filters, [column.key]: next }
                                        applyFiltersFrom(snapshot)
                                    } else {
                                        applyFilters()
                                    }
                                    setOpenFilterKey(null)
                                }}
                            >
                                Uygula
                            </button>
                        </div>
                    </div>
                )
            }

            case 'date': {
                return (
                    <div ref={setFilterPopoverRef} onMouseDown={(e) => e.stopPropagation()} onClick={(e) => e.stopPropagation()} className="absolute top-full left-1/2 -translate-x-1/2 mt-1 z-20 bg-white text-slate-700 border border-slate-200 rounded shadow p-2 w-56">
                        <DateFilterUI
                            operators={f.operators}
                            value={(current.value as string | undefined) ?? ''}
                            value2={(current.value2 as string | undefined) ?? ''}
                            operator={current.operator}
                            onChange={(next) => setFilterState(column.key, next)}
                        />
                        <div className="flex justify-between gap-2 pt-2">
                            <button
                                className="text-xs text-teal-700 hover:underline"
                                onClick={() => setOpenFilterKey(null)}
                            >
                                Kapat
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    clearFilter(column.key)
                                }}
                            >
                                Sıfırla
                            </button>
                            <button
                                className="text-xs text-emerald-700 hover:underline"
                                onClick={() => {
                                    if (!filters[column.key]) {
                                        const op = current.operator ?? f.operators?.[0] ?? ''
                                        const next = { ...current, operator: op }
                                        setFilters(prev => ({ ...prev, [column.key]: next }))
                                        const snapshot = { ...filters, [column.key]: next }
                                        applyFiltersFrom(snapshot)
                                    } else {
                                        applyFilters()
                                    }
                                    setOpenFilterKey(null)
                                }}
                            >
                                Uygula
                            </button>
                        </div>
                    </div>
                )
            }
        }
    }

    function DateFilterUI({ operators, value, value2, operator, onChange }: { operators?: string[]; value?: string; value2?: string; operator?: string; onChange: (next: ColumnFilterState) => void }) {
        const [op, setOp] = useState<string>(operator ?? operators?.[0] ?? '')

        return (
            <div className="flex flex-col gap-2">
                {operators?.length ? (
                    <select
                        className="w-full block border border-slate-200 rounded px-2 py-1 text-sm"
                        value={op}
                        onChange={(e) => {
                            setOp(e.target.value)
                            onChange({ operator: e.target.value })
                        }}
                    >
                        {operators.map((o) => (
                            <option className="text-sm" key={o} value={o}>{o}</option>
                        ))}
                    </select>
                ) : null}

                <input
                    type="date"
                    placeholder="Ara"
                    className="w-full block rounded border border-slate-200 px-2 py-1 text-sm"
                    value={value ?? ''}
                    onChange={(e) => onChange({ value: e.target.value })}
                />

                {op === 'between' ? (
                    <input
                        type="date"
                        placeholder="Ara"
                        className="w-full block rounded border border-slate-200 px-2 py-1 text-sm"
                        value={value2 ?? ''}
                        onChange={(e) => onChange({ value2: e.target.value })}
                    />
                ) : null}
            </div>
        )
    }

    return (
        <div className="relative m-50 border border-slate-200 rounded-lg p-3 bg-white">
            <div className="flex items-center justify-between pl-2 pt-2 text-right mb-2">
                <div className="text-sm font-semibold text-gray-700">
                    {props.title}
                </div>

                <div className="flex items-center gap-1 overflow-visible">
                    <div className={`flex items-center gap-1 transition-transform duration-300 ${isExpanded ? '-translate-x-1' : 'translate-x-0'}`}>
                        {props.onRefresh && activeButtons.includes("refresh") && (
                            <button
                                key="refresh"
                                onClick={() => { }}
                                className="inline-flex items-center bg-sky-600 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-sky-700"
                                aria-label="Yenile"
                            >
                                <FaArrowRotateRight className="w-3 h-3" />
                                <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                    Yenile
                                </span>
                            </button>
                        )}
                        {props.onExport && activeButtons.includes("export") && (
                            <button
                                key="export"
                                onClick={() => { }}
                                className="inline-flex items-center bg-emerald-600 hover:bg-emerald-700 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group"
                                aria-label="Dışa Aktar"
                            >
                                <FaArrowDown className="w-3 h-3" />
                                <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                    Dışa Aktar
                                </span>
                            </button>
                        )}
                        {props.onFilter && activeButtons.includes("filter") && (
                            <button
                                key="filter"
                                onClick={() => {
                                    setIsFilterOpen(!isFilterOpen)
                                    setFilters({})
                                    setOpenFilterKey(null)
                                }}
                                className="inline-flex items-center bg-yellow-500 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-yellow-600"
                                aria-label="Filtrele"
                            >
                                <FaFilter className="w-3 h-3" />
                                <span className={`overflow-hidden whitespace-nowrap transition-all duration-200 text-sm ${isFilterOpen ? 'ml-2 max-w-[80px] opacity-100' : 'ml-0 max-w-0 opacity-0'} group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100`}>
                                    Filtrele
                                </span>
                            </button>
                        )}
                    </div>

                    <div
                        className={`flex items-center gap-1 overflow-visible transition-all duration-300 ${isExpanded ? 'max-w-[999px] opacity-100 translate-x-0 border-l pl-2 border-slate-200' : 'max-w-0 opacity-0 translate-x-4 pl-0 border-l-0'}`}
                    >
                        {props.onNew && (
                            <button
                                key="new"
                                onClick={() => { }}
                                className="inline-flex items-center bg-teal-600 hover:bg-teal-700 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group"
                                aria-label="Ekle"
                            >
                                <FaPlus className="w-3 h-3" />
                                <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                    Yeni
                                </span>
                            </button>
                        )}
                        {props.onDelete && (
                            <button
                                key="delete"
                                onClick={() => { }}
                                className="inline-flex items-center bg-rose-600 hover:bg-rose-700 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group"
                                aria-label="Sil"
                            >
                                <FaTrash className="w-3 h-3" />
                                <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                    Sil
                                </span>
                            </button>
                        )}
                        {(
                            <div className="relative inline-block">
                                <button
                                    key="settings"
                                    ref={setSettingsButtonRef}
                                    onMouseDown={(e) => e.stopPropagation()}
                                    onClick={() => {
                                        setIsSettingsOpen(!isSettingsOpen)
                                    }}
                                    className="inline-flex items-center bg-slate-500 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-slate-600"
                                    aria-label="Ayarlar"
                                >
                                    <FaGear className="w-3 h-3" />
                                    <span className={`overflow-hidden whitespace-nowrap transition-all duration-200 text-sm ${isSettingsOpen ? 'ml-2 max-w-[80px] opacity-100' : 'ml-0 max-w-0 opacity-0'} group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100`}>
                                        Ayarlar
                                    </span>
                                </button>
                            </div>
                        )}
                    </div>
                    <button
                        key="other"
                        onClick={() => {
                            if (activeButtons.length === DEFAULT_BUTTONS.length) {
                                setActiveButtons(ALL_BUTTONS)
                            }
                            else {
                                setActiveButtons(DEFAULT_BUTTONS)
                            }
                        }}
                        className="inline-flex items-center z-0 bg-slate-500 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-slate-600"
                        aria-label="Diğer İşlemler"
                    >
                        <FaEllipsis className="w-3 h-3" />
                    </button>
                </div>
            </div>
            <div
                className="w-full overflow-x-auto"
            >
                <table className="table-fixed rounded-lg">
                    <thead className="bg-teal-600 text-white border-b border-teal-700">
                        <tr className="text-left text-sm">
                            {orderedActiveColumns().map((column, idx) => (
                                <th
                                    className={`relative py-3 px-10 whitespace-nowrap overflow-visible text-ellipsis ${props.onSorted && column.sortable ? 'cursor-pointer select-none' : ''}`}
                                    onClick={props.onSorted && column.sortable ? () => {
                                        let nextOp: SortOperator
                                        if (sortColumn?.key == column.key) {
                                            nextOp = sortColumn.operator == 'asc' ? 'desc' : 'asc'
                                        } else {
                                            nextOp = 'asc'
                                        }
                                        setSortColumn({ key: column.key, operator: nextOp })
                                        props.onSorted?.(column.key, nextOp)
                                    } : undefined}
                                >
                                    <div className="inline-flex items-center gap-2">
                                        <span>{column.label}</span>
                                        {props.onSorted && column.sortable && sortColumn?.key == column.key && sortColumn.operator == 'asc' &&
                                            <button
                                                className="p-1 inline-flex items-center"
                                            >
                                                <FaSortDown className={`w-3 h-3 -translate-y-[1px]`} />
                                            </button>
                                        }
                                        {props.onSorted && column.sortable && sortColumn?.key == column.key && sortColumn.operator == 'desc' &&
                                            <button
                                                className="p-1 inline-flex items-center"
                                            >
                                                <FaSortUp className={`w-3 h-3 translate-y-[1px]`} />
                                            </button>
                                        }
                                        {isFilterOpen && column.filter && (
                                            <>
                                                <button
                                                    key={`filterbtn-${column.key}`}
                                                    onMouseDown={(e) => e.stopPropagation()}
                                                    onClick={(e) => {
                                                        e.stopPropagation()
                                                        setOpenFilterKey(openFilterKey === column.key ? null : column.key)
                                                    }}
                                                    className="p-1 inline-flex items-center"
                                                >
                                                    <FaFilter className={`w-3 h-3 ${isColumnFilterActive(column) ? 'text-yellow-400' : ''}`} />
                                                </button>
                                                {openFilterKey === column.key && (
                                                    renderFilter(column)
                                                )}
                                            </>
                                        )}
                                    </div>
                                    {idx < orderedActiveColumns().length - 1 && (
                                        <div className="absolute top-1/4 right-0 w-px h-1/2 bg-teal-500"></div>
                                    )}
                                </th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {props.data.slice((page - 1) * pageSize, page * pageSize).map((item, rowIndex) => (
                            <tr key={`row-${rowIndex}`} className="text-left text-sm border-b border-gray-100 hover:bg-gray-50 transition-colors duration-100 whitespace-nowrap overflow-hidden text-ellipsis">
                                {orderedActiveColumns().map((column) => (
                                    <td key={column.key} className="py-3 px-10 w-1/3 overflow-hidden text-ellipsis">
                                        {column.render
                                            ? column.render(item)
                                            : defaultRenderCell(item[column.key])
                                        }
                                    </td>
                                ))}
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>

            <div className="flex justify-end items-center gap-1 pt-4 dasdas">
                <div className="flex items-center mr-auto ml-2 text-xs">
                    <span className="text-teal-700 text-sm">Sayfa Başına</span>
                    <select
                        className="text-teal-700 hover:bg-gray-100 rounded-md p-1 ml-2 disabled:opacity-40 disabled:cursor-not-allowed border border-teal-700"
                        onChange={(e) => {
                            setPage(1)
                            setPageSize(Number(e.target.value))
                            props.onPageSizeChange?.(Number(e.target.value))
                        }}
                    >
                        {Array.from({ length: 10 }, (_, i) => (i + 1) * 10).map((option) => (
                            <option key={option} value={option}>{option}</option>
                        ))}
                    </select>
                </div>
                <button
                    className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
                    disabled={page === 1}
                    onClick={() => {
                        props.onPageChange?.(1)
                        setPage(1)
                    }}
                >
                    <FaAngleDoubleLeft className="w-4 h-4" />
                </button>
                <button
                    className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
                    disabled={page === 1}
                    onClick={() => {
                        props.onPageChange?.(page - 1)
                        setPage(page - 1)
                    }}
                >
                    <FaAngleLeft className="w-4 h-4" />
                </button>
                {pageWindow().map((p) => (
                    <button
                        key={p}
                        onClick={() => {
                            props.onPageChange?.(p)
                            setPage(p)
                        }}
                        className={`hover:bg-gray-100 rounded-md disabled:opacity-40 disabled:cursor-not-allowed w-6 h-6 ${p === page ? 'bg-teal-600 text-white hover:bg-teal-600' : ''}`}
                    >
                        {p}
                    </button>
                ))}
                <button
                    className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
                    disabled={page === totalPage}
                    onClick={() => {
                        props.onPageChange?.(page + 1)
                        setPage(page + 1)
                    }}
                >
                    <FaAngleRight className="w-4 h-4" />
                </button>
                <button
                    className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
                    disabled={page === totalPage}
                    onClick={() => {
                        props.onPageChange?.(totalPage ?? 1)
                        setPage(totalPage ?? 1)
                    }}
                >
                    <FaAngleDoubleRight className="w-4 h-4" />
                </button>
                <span className="text-teal-700 ml-3 text-sm">Sayfa</span>
                <input
                    className="text-teal-700 hover:bg-gray-100 rounded-md p-1 mr-3 disabled:opacity-40 disabled:cursor-not-allowed border border-teal-700 text-xs text-center"
                    type="number"
                    value={page}
                    min={1}
                    max={totalPage ?? 1}
                    onWheel={e => (e.target as HTMLInputElement).blur()}
                    onChange={(e) => {
                        const raw = Number(e.target.value)
                        if (!Number.isFinite(raw)) return
                        const maxPage = totalPage ?? 1
                        const clamped = Math.max(1, Math.min(raw, maxPage))
                        props.onPageChange?.(clamped)
                        setPage(clamped)
                    }}
                />
            </div>

            {/* Settings Popover - positioned outside table to avoid z-index conflicts */}
            {
                isSettingsOpen && settingsButtonRef &&
                <div
                    ref={setSettingsPopoverRef}
                    onMouseDown={(e) => e.stopPropagation()}
                    className="fixed w-56 bg-white border border-slate-200 rounded-md shadow-lg p-3 text-left z-40"
                    style={{
                        top: settingsPosition.top,
                        left: settingsPosition.left
                    }}
                >
                    <div className="text-sm text-slate-700 font-semibold mb-2">Sütunlar</div>
                    <div className="flex flex-col gap-1 text-left text-sm text-slate-700 bg-gray-100 rounded-md p-2">
                        {columnOrder.map((key) => {
                            const column = props.columns.find((c) => c.key === key)!
                            return (
                                <div
                                    key={column.key}
                                    className={`flex items-center gap-2 rounded-md px-2 py-1 ${draggingKey === column.key ? 'bg-slate-200' : 'bg-transparent'}`}
                                    draggable
                                    onDragStart={() => handleDragStart(column.key)}
                                    onDragOver={(e) => e.preventDefault()}
                                    onDrop={() => handleDrop(column.key)}
                                >
                                    <span className="cursor-move select-none text-slate-500">
                                        <FaGripVertical className="w-3 h-3" />
                                    </span>
                                    <input
                                        type="checkbox"
                                        className=""
                                        checked={column.show}
                                        onChange={() => {
                                            handleColumnChange(column)
                                        }} />
                                    <span className="truncate">{column.label}</span>
                                </div>
                            )
                        })}
                    </div>
                </div>
            }
        </div >
    )
}
