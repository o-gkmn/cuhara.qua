import { useTableContext } from "../../../../context/TableContext"
import { useEffect, useState, useRef, type ReactNode } from "react"
import { TextFilter, NumberFilter, DateFilter, BooleanFilter, SelectFilter, applyFilter } from "./components"

export default function Filter({ onFilter }: { onFilter: (query: string) => void }) {
    const { activeFilterColumn, setActiveFilterColumn, filterButtonRefs, filterParams, setFilterParams } = useTableContext()
    const filterRef = useRef<HTMLDivElement>(null)
    const [position, setPosition] = useState({ top: 0, left: 0 })

    // Position the filter popover relative to the filter button within header
    useEffect(() => {
        if (activeFilterColumn && filterButtonRefs.current[activeFilterColumn.key]) {
            const buttonElement = filterButtonRefs.current[activeFilterColumn.key]
            if (buttonElement) {
                const rect = buttonElement.getBoundingClientRect()
                const table = buttonElement.closest('table')
                const tableRect = table?.getBoundingClientRect()

                const thead = table?.querySelector('thead')
                const theadRect = thead?.getBoundingClientRect()

                if (tableRect && theadRect) {
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
                        top: theadRect.top,
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

    return activeFilterColumn && (
        <div
            ref={filterRef}
            style={{
                position: 'absolute',
                top: position.top,
                left: position.left,
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
                        const query = applyFilter(filterParams)
                        onFilter(query)
                        setActiveFilterColumn(null)
                    }}
                >
                    Uygula
                </button>
            </div>
        </div>
    )
}
