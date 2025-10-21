
import { useTableContext } from "../../../context/TableContext"
import { useEffect, useState, useRef, type ReactNode } from "react"

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
    operator?: string
    value?: unknown
}

export default function Filter({ onFilter }: { onFilter: (query: string) => void }) {
    const { activeFilterColumn, setActiveFilterColumn, filterButtonRefs, filterParams, setFilterParams } = useTableContext()
    const filterRef = useRef<HTMLDivElement>(null)
    const [position, setPosition] = useState({ top: 0, left: 0 })


    const operatorTranslations = {
        // Text operators
        'contains': 'İçerir',
        'startsWith': 'İle başlar',
        'endsWith': 'İle biter',
        'equals': 'Eşittir',
        'notEquals': 'Eşit değildir',
        'empty': 'Boş',
        'notEmpty': 'Boş değil',

        // Number operators
        'eq': 'Eşittir',
        'neq': 'Eşit değildir',
        'gt': 'Büyüktür',
        'gte': 'Büyük eşittir',
        'lt': 'Küçüktür',
        'lte': 'Küçük eşittir',
        'between': 'Arasında',
        'in': 'İçinde',

        // Date operators
        'on': 'Tarihinde',
        'before': 'Öncesinde',
        'after': 'Sonrasında',

        // Boolean operators
        'isTrue': 'Doğru',
        'isFalse': 'Yanlış',
        'isNull': 'Boş',
        'isNotNull': 'Boş değil',

        // Select operators
        // 'eq' ve 'neq' zaten yukarıda tanımlı
    } as const

    // Position the filter popover relative to the filter button within header
    useEffect(() => {
        if (activeFilterColumn && filterButtonRefs.current[activeFilterColumn.key]) {
            const buttonElement = filterButtonRefs.current[activeFilterColumn.key]
            if (buttonElement) {
                const rect = buttonElement.getBoundingClientRect()
                const table = buttonElement.closest('table')
                const tableRect = table?.getBoundingClientRect()
                const container = table?.closest('.relative') // Ana container
                const containerRect = container?.getBoundingClientRect()

                if (tableRect && containerRect) {
                    const popoverWidth = window.innerWidth < 640 ? Math.min(280, window.innerWidth - 20) : 256 // Responsive width
                    const isMobile = window.innerWidth < 640

                    let leftPosition = rect.left - tableRect.left - (isMobile ? 50 : 100)

                    // Sol kenarda ise popover'ı sağa kaydır
                    if (leftPosition < 0) {
                        leftPosition = isMobile ? 5 : 10
                    }

                    // Sağ kenarda ise popover'ı sola kaydır
                    const rightEdge = leftPosition + popoverWidth
                    const tableWidth = tableRect.width
                    if (rightEdge > tableWidth) {
                        leftPosition = tableWidth - popoverWidth - (isMobile ? 5 : 10)
                    }

                    // Mobile'da center positioning
                    if (isMobile) {
                        leftPosition = Math.max(5, Math.min(leftPosition, tableWidth - popoverWidth - 5))
                    }

                    setPosition({
                        top: isMobile ? 10 : 5, // More space on mobile
                        left: leftPosition
                    })
                }
            }
        }
    }, [activeFilterColumn, filterButtonRefs])

    // Handle click outside to close filter
    useEffect(() => {
        function handleClickOutside(event: MouseEvent) {
            if (filterRef.current && !filterRef.current.contains(event.target as Node)) {
                // Check if click is on a filter button
                const isFilterButton = Object.values(filterButtonRefs.current).some(button =>
                    button && button.contains(event.target as Node)
                )
                if (!isFilterButton) {
                    setActiveFilterColumn(null)
                }
            }
        }

        if (activeFilterColumn) {
            document.addEventListener('mousedown', handleClickOutside)
            return () => document.removeEventListener('mousedown', handleClickOutside)
        }
    }, [activeFilterColumn, setActiveFilterColumn, filterButtonRefs])

    function renderFilter(): ReactNode | null {
        switch (activeFilterColumn?.filter?.type) {
            case 'text':
                return <TextFilter />
            case 'number':
                return <NumberFilter />
            case 'date':
                return <DateFilter />
            case 'boolean':
                return <BooleanFilter />
            case 'select':
                return <SelectFilter />
            default:
                return null
        }
    }

    function clearFilter(key: string) {
        setFilterParams(prev => {
            const newParams = { ...prev }
            delete newParams[key]
            return newParams
        })
    }

    function applyFilter(): string {
        const entries = Object.entries(filterParams)

        return entries.map(([columnKey, filterParam], i) => {
            if (filterParam.operator === undefined) return // operator selection must
            const part = `${columnKey} ${filterParam.operator} ${filterParam.value}`
            const hasNext = i < entries.length - 1
            return hasNext ? `${part} AND` : part
        }).join(' ')
    }

    function TextFilter() {
        return (
            <div className="flex flex-col gap-2">
                <input
                    type="text"
                    placeholder="Değer giriniz"
                    value={filterParams[activeFilterColumn?.key || '']?.value as string || ''}
                    onChange={(e) => {
                        setFilterParams(prev => ({
                            ...prev,
                            [activeFilterColumn?.key || '']: {
                                ...prev[activeFilterColumn?.key || ''],
                                value: e.target.value
                            }
                        }))
                    }}
                    className={`
                        w-full 
                        border border-slate-200 
                        rounded p-1.5 sm:p-1
                        text-xs sm:text-sm
                        ${filterParams[activeFilterColumn?.key || '']?.operator === 'empty'
                            || filterParams[activeFilterColumn?.key || '']?.operator === 'notEmpty' ? 'hidden' : ''
                        }
                    `}
                    autoFocus
                />
                <select
                    className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
                    value={filterParams[activeFilterColumn?.key || '']?.operator || ''}
                    onChange={(e) => {
                        if (e.target.value) {
                            setFilterParams(prev => ({
                                ...prev,
                                [activeFilterColumn?.key || '']: {
                                    ...prev[activeFilterColumn?.key || ''],
                                    operator: e.target.value,
                                }
                            }))
                        }
                    }}
                >
                    <option value="">Operatör seçin</option>
                    {activeFilterColumn?.filter?.operators?.map((operator) => (
                        <option key={operator} value={operator}>
                            {operatorTranslations[operator as keyof typeof operatorTranslations]}
                        </option>
                    ))}
                </select>
            </div>
        )
    }

    function NumberFilter() {
        return <div className="flex flex-col gap-2">
            <input
                type="number"
                placeholder="Değer giriniz"
                className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
                value={filterParams[activeFilterColumn?.key || '']?.value as number || ''}
                onChange={(e) => {
                    setFilterParams(prev => ({
                        ...prev,
                        [activeFilterColumn?.key || '']: {
                            ...prev[activeFilterColumn?.key || ''],
                            value: e.target.value
                        }
                    }))
                }}
                autoFocus
            />
            <select
                value={filterParams[activeFilterColumn?.key || '']?.operator || ''}
                onChange={(e) => {
                    if (e.target.value) {
                        setFilterParams(prev => ({
                            ...prev,
                            [activeFilterColumn?.key || '']: {
                                ...prev[activeFilterColumn?.key || ''],
                                operator: e.target.value,
                            }
                        }))
                    }
                }}
                className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
            >
                <option value="">Operatör seçin</option>
                {activeFilterColumn?.filter?.operators?.map((operator) => (
                    <option key={operator} value={operator}>
                        {operatorTranslations[operator as keyof typeof operatorTranslations]}
                    </option>
                ))}
            </select>
        </div>
    }

    function DateFilter() {
        return <div className="flex flex-col gap-2">
            <input
                type="date"
                placeholder="Değer giriniz"
                className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
                value={filterParams[activeFilterColumn?.key || '']?.value as string || ''}
                onChange={(e) => {
                    setFilterParams(prev => ({
                        ...prev,
                        [activeFilterColumn?.key || '']: {
                            ...prev[activeFilterColumn?.key || ''],
                            value: e.target.value,
                        }
                    }))
                }}
                autoFocus
            />
            {
                filterParams[activeFilterColumn?.key || '']?.operator === 'between' && (
                    <input
                        type="date"
                        placeholder="Değer giriniz"
                        className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
                        value={filterParams[activeFilterColumn?.key || '']?.value as string || ''}
                        onChange={(e) => {
                            setFilterParams(prev => ({
                                ...prev,
                                [activeFilterColumn?.key || '']: {
                                    ...prev[activeFilterColumn?.key || ''],
                                    value: e.target.value,
                                }
                            }))
                        }}
                        autoFocus
                    />
                )
            }
            <select
                value={filterParams[activeFilterColumn?.key || '']?.operator || ''}
                onChange={(e) => {
                    if (e.target.value) {
                        setFilterParams(prev => ({
                            ...prev,
                            [activeFilterColumn?.key || '']: {
                                ...prev[activeFilterColumn?.key || ''],
                                operator: e.target.value,
                            }
                        }))
                    }
                }}
                className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
            >
                <option value="">Operatör seçin</option>
                {activeFilterColumn?.filter?.operators?.map((operator) => (
                    <option key={operator} value={operator}>
                        {operatorTranslations[operator as keyof typeof operatorTranslations]}
                    </option>
                ))}
            </select>
        </div>
    }

    function BooleanFilter() {
        return <div className="flex flex-col gap-2">
            <select
                value={filterParams[activeFilterColumn?.key || '']?.operator || ''}
                onChange={(e) => {
                    if (e.target.value) {
                        setFilterParams(prev => ({
                            ...prev,
                            [activeFilterColumn?.key || '']: {
                                ...prev[activeFilterColumn?.key || ''],
                                operator: e.target.value,
                            }
                        }))
                    }
                }}
                className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
            >
                <option value="">Operatör seçin</option>
                {activeFilterColumn?.filter?.operators?.map((operator) => (
                    <option key={operator} value={operator}>
                        {operatorTranslations[operator as keyof typeof operatorTranslations]}
                    </option>
                ))}
            </select>
        </div>
    }

    function SelectFilter() {
        return <div className="flex flex-col gap-2">
            <select
                value={filterParams[activeFilterColumn?.key || '']?.value as string || ''}
                onChange={(e) => {
                    if (e.target.value) {
                        setFilterParams(prev => ({
                            ...prev,
                            [activeFilterColumn?.key || '']: {
                                ...prev[activeFilterColumn?.key || ''],
                                value: e.target.value,
                            }
                        }))
                    }
                }}
                className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
            >
                <option value="">Değer seçin</option>
                {
                    activeFilterColumn?.filter?.type === 'select' && activeFilterColumn.filter.options.map((option) => (
                        <option key={option.value} value={option.value}>
                            {option.label}
                        </option>
                    ))
                }
            </select>
            <select
                value={filterParams[activeFilterColumn?.key || '']?.operator || ''}
                onChange={(e) => {
                    if (e.target.value) {
                        setFilterParams(prev => ({
                            ...prev,
                            [activeFilterColumn?.key || '']: {
                                ...prev[activeFilterColumn?.key || ''],
                                operator: e.target.value,
                            }
                        }))
                    }
                }}
                className="w-full border border-slate-200 rounded p-1.5 sm:p-1 text-xs sm:text-sm"
            >
                <option value="">Operatör seçin</option>
                {activeFilterColumn?.filter?.operators?.map((operator) => (
                    <option key={operator} value={operator}>
                        {operatorTranslations[operator as keyof typeof operatorTranslations]}
                    </option>
                ))}
            </select>
        </div>
    }

    return activeFilterColumn && (
        <tr>
            <td
                colSpan={100}
                className="relative p-0"
            >
                <div
                    ref={filterRef}
                    style={{
                        position: 'absolute',
                        top: `${position.top}px`,
                        left: `${position.left}px`,
                        zIndex: 9999
                    }}
                    className="bg-white text-slate-700 border border-slate-200 rounded-md shadow-lg p-3 sm:p-4 w-72 sm:w-64 max-w-[calc(100vw-20px)] sm:max-w-none"
                    onMouseDown={(e) => e.stopPropagation()}
                >
                    <div className="mb-3">
                        <h4 className="text-xs sm:text-sm font-medium text-slate-800 mb-2">
                            {activeFilterColumn.label} - Filtre
                        </h4>
                        {renderFilter()}
                    </div>

                    <div className="flex flex-col sm:flex-row justify-end gap-2 pt-3 border-t border-slate-100">
                        <button
                            className="px-3 py-1.5 sm:py-1 text-xs text-slate-600 hover:text-slate-800 hover:bg-slate-100 rounded transition-colors"
                            onClick={() => {
                                setActiveFilterColumn(null)
                                clearFilter(activeFilterColumn?.key)
                            }}
                        >
                            İptal
                        </button>
                        <button
                            className="px-3 py-1.5 sm:py-1 text-xs text-orange-600 hover:text-orange-800 hover:bg-orange-50 rounded transition-colors"
                            onClick={() => {
                                clearFilter(activeFilterColumn?.key)
                            }}
                        >
                            Sıfırla
                        </button>
                        <button
                            className="px-3 py-1.5 sm:py-1 text-xs bg-teal-600 text-white hover:bg-teal-700 rounded transition-colors"
                            onClick={() => {
                                const query = applyFilter()
                                onFilter(query)
                                setActiveFilterColumn(null)
                            }}
                        >
                            Uygula
                        </button>
                    </div>
                </div>
            </td>
        </tr>
    )
}
