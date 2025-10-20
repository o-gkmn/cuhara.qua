import { useState, type ReactNode } from "react"
import { FaFilter, FaSortUp } from 'react-icons/fa6';
import { FaSortDown } from 'react-icons/fa6';
import TableSettingsPopover from "./table/SettingsPopover";
import Toolbar, { type ToolbarProps } from "./table/ToolBar";
import { useTableContext } from "../../context/TableContext";
import type { ColumnFilterProps, SortOperator } from "./table/Filter";
import Filter from "./table/Filter";
import Pagination from "./table/Pagination";

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
    toolbar?: ToolbarProps
    onPageChange?: (page: number) => void
    onPageSizeChange?: (pageSize: number) => void
    onSorted?: (field: string, operator: SortOperator) => void
}

export default function Table<T extends Record<string, unknown>>(props: TableProps<T>) {
    const [sortColumn, setSortColumn] = useState<{ key: string, operator: SortOperator }>()
    const { columns, activeFilterColumn, setActiveFilterColumn, isFilterOpen, filterButtonRefs } = useTableContext()

    const defaultRenderCell = (value: unknown): ReactNode => {
        if (value === null || value === undefined) return "-"
        if (typeof value === 'object') return JSON.stringify(value)
        return String(value)
    }

    return (
        <div className="relative m-50 border border-slate-200 rounded-lg p-3 bg-white">
            <Toolbar {...props.toolbar} />
            <div
                className="w-full overflow-x-auto overflow-y-visible"
            >
                <table className="table-fixed rounded-lg">
                    <thead className="bg-teal-600 text-white border-b border-teal-700 sticky relative">
                        <tr className="text-left text-sm">
                            {columns.map((column, idx) => (
                                <th
                                    className={`
                                        relative 
                                        py-3 
                                        px-10 
                                        whitespace-nowrap 
                                        overflow-visible 
                                        text-ellipsis 
                                        ${props.onSorted && column.sortable ? 'cursor-pointer select-none' : ''}
                                    `}
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
                                        {isFilterOpen && column.filter &&
                                            <button
                                                key={`filterbtn-${column.key}`}
                                                ref={(e) => {
                                                    filterButtonRefs.current[column.key] = e
                                                }}
                                                onClick={(e) => {
                                                    e.stopPropagation()
                                                    if (activeFilterColumn?.key === column.key) {
                                                        setActiveFilterColumn(null)
                                                    } else {
                                                        setActiveFilterColumn(column)
                                                    }
                                                }}
                                                className={`p-1 inline-flex items-center hover:bg-teal-700 rounded transition-colors ${
                                                    activeFilterColumn?.key === column.key ? 'bg-teal-700' : ''
                                                }`}
                                            >
                                                <FaFilter className={`w-3 h-3 ${
                                                    activeFilterColumn?.key === column.key ? 'text-yellow-400' : 'text-teal-100'
                                                }`} />
                                            </button>
                                        }
                                    </div>
                                    {idx < columns.length - 1 && (
                                        <div className="absolute top-1/4 right-0 w-px h-1/2 bg-teal-500"></div>
                                    )}
                                </th>
                            ))}
                        </tr>
                        <Filter onFilter={props.toolbar?.onFilter ?? (() => {})} />
                    </thead>
                    <tbody>
                        {props.data.map((item, rowIndex) => (
                            <tr
                                key={`row-${rowIndex}`}
                                className={`
                                    text-left
                                    text-sm
                                    border-b border-gray-100
                                    hover:bg-gray-50
                                    transition-colors duration-100
                                    whitespace-nowrap
                                    overflow-hidden
                                    text-ellipsis
                                `}
                            >
                                {columns.map((column) => (
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

            <Pagination onPageChange={props.onPageChange} onPageSizeChange={props.onPageSizeChange} dataLength={props.data.length} />

            <TableSettingsPopover />

        </div >
    )
}
