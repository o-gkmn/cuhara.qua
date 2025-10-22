import { useState } from "react"
import { FaFilter, FaSortUp, FaSortDown } from 'react-icons/fa6'
import type { SortOperator, TableProps } from "../types/types"
import { useTableContext } from "../../../../context/TableContext"
import Filter from "../filters/Filter"
import Pagination from "../ui/Pagination"
import Toolbar from "../ui/ToolBar"

export default function DesktopTable<T extends Record<string, unknown>>(props: TableProps<T>) {
    const [sortColumn, setSortColumn] = useState<{ key: string, operator: SortOperator }>()
    const { activeFilterColumn, setActiveFilterColumn, isFilterOpen, filterButtonRefs, columns } = useTableContext()

    const visibleColumns = columns.filter(column => column.show !== false)

    const defaultRenderCell = (value: unknown) => {
        if (value === null || value === undefined) return "-"
        if (typeof value === 'object') return JSON.stringify(value)
        return String(value)
    }

    return (
        <>
            <Toolbar {...props.toolbar} />

            <div className="w-full overflow-x-auto overflow-y-visible">
                <table className="table-auto md:table-fixed rounded-lg w-full">
                    <thead className="bg-teal-600 text-white border-b border-teal-700 sticky relative">
                        {/* <table className="table-fixed rounded-lg w-full min-w-[800px]">
                    <thead className="bg-teal-600 text-white border-b border-teal-700 sticky relative"> */}
                        <tr className="text-left text-xs md:text-sm">
                            {visibleColumns.map((column, idx) => (
                                <th
                                    key={column.key}
                                    className={`
                                        relative 
                                        py-2 md:py-3 
                                        px-2 md:px-4 lg:px-6 
                                        whitespace-nowrap 
                                        overflow-visible 
                                        text-ellipsis 
                                        text-xs md:text-sm
                                        ${props.onSorted && column.sortable ? 'cursor-pointer select-none' : ''}
                                    `}
                                    onClick={props.onSorted && column.sortable ? () => {
                                        let nextOp: SortOperator
                                        if (sortColumn?.key == column.key) {
                                            nextOp = sortColumn?.operator == 'asc' ? 'desc' : 'asc'
                                        } else {
                                            nextOp = 'asc'
                                        }
                                        setSortColumn({ key: column.key, operator: nextOp })
                                        props.onSorted?.(column.key, nextOp)
                                    } : undefined}
                                >
                                    <div className="inline-flex items-center gap-1 md:gap-2">
                                        <span className="truncate max-w-[80px] md:max-w-none">{column.label}</span>
                                        {/* Sort icons */}
                                        {props.onSorted && column.sortable && sortColumn?.key == column.key && sortColumn?.operator == 'asc' &&
                                            <FaSortDown className="w-2.5 h-2.5 md:w-3 md:h-3 -translate-y-[1px]" />
                                        }
                                        {props.onSorted && column.sortable && sortColumn?.key == column.key && sortColumn?.operator == 'desc' &&
                                            <FaSortUp className="w-2.5 h-2.5 md:w-3 md:h-3 translate-y-[1px]" />
                                        }
                                        {/* Filter button */}
                                        {isFilterOpen && column.filter &&
                                            <button
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
                                                className={`p-0.5 md:p-1 inline-flex items-center hover:bg-teal-700 rounded transition-colors ${activeFilterColumn?.key === column.key ? 'bg-teal-700' : ''}`}
                                            >
                                                <FaFilter className={`w-2.5 h-2.5 md:w-3 md:h-3 ${activeFilterColumn?.key === column.key ? 'text-yellow-400' : 'text-teal-100'}`} />
                                            </button>
                                        }
                                    </div>
                                    {idx < visibleColumns.length - 1 && (
                                        <div className="absolute top-1/4 right-0 w-px h-1/2 bg-teal-500"></div>
                                    )}
                                </th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {props.data.map((item, rowIndex) => (
                            <tr
                                key={`row-${rowIndex}`}
                                className="text-left text-xs md:text-sm border-b border-gray-100 hover:bg-gray-50 transition-colors duration-100 whitespace-nowrap overflow-hidden text-ellipsis"
                            >
                                {visibleColumns.map((column) => (
                                    <td key={column.key} className="py-2 md:py-3 px-2 md:px-4 lg:px-10 w-1/3 overflow-hidden text-ellipsis text-xs md:text-sm">
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

                <Filter onFilter={props.toolbar?.onFilter ?? (() => { })} />
            </div>


            <Pagination
                onPageChange={props.onPageChange}
                onPageSizeChange={props.onPageSizeChange}
                dataLength={props.data.length}
            />
        </>
    )
}