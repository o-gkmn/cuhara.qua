import { useState, type ReactNode } from "react"
import { useTableContext } from "../../../../context/TableContext"

interface MobileCardViewProps<T extends Record<string, unknown>> {
    data: T[]
}

export default function MobileCardView<T extends Record<string, unknown>>({ data }: MobileCardViewProps<T>) {
    const [expandedCards, setExpandedCards] = useState<Set<number>>(new Set())
    const { columns } = useTableContext()

    const visibleColumns = columns.filter(column => column.show !== false)

    const toggleCard = (index: number) => {
        setExpandedCards(prev => {
            const newSet = new Set(prev)
            if (newSet.has(index)) {
                newSet.delete(index)
            } else {
                newSet.add(index)
            }
            return newSet
        })
    }

    const defaultRenderCell = (value: unknown): ReactNode => {
        if (value === null || value === undefined) return "-"
        if (typeof value === 'object') return JSON.stringify(value)
        return String(value)
    }

    return (
        <div className="space-y-2">
            {data.map((item, rowIndex) => {
                const isExpanded = expandedCards.has(rowIndex)
                const firstColumn = visibleColumns[0]
                const firstValue = firstColumn ? (firstColumn.render ? firstColumn.render(item) : defaultRenderCell(item[firstColumn.key])) : ''
                
                return (
                    <div
                        key={`card-${rowIndex}`}
                        className="bg-white border border-slate-200 rounded-lg shadow-sm hover:shadow-md transition-all duration-200"
                    >
                        {/* Compact Card Header */}
                        <div 
                            className="flex items-center justify-between p-3 cursor-pointer hover:bg-slate-50 transition-colors"
                            onClick={() => toggleCard(rowIndex)}
                        >
                            <div className="flex items-center space-x-3 flex-1 min-w-0">
                                <div className="w-2 h-2 bg-teal-500 rounded-full flex-shrink-0"></div>
                                <div className="flex-1 min-w-0">
                                    <div className="text-sm font-medium text-slate-900 truncate">
                                        {firstValue}
                                    </div>
                                    <div className="text-xs text-slate-500">
                                        Kayıt #{rowIndex + 1}
                                    </div>
                                </div>
                            </div>
                            <div className="flex items-center">
                                <div className={`w-5 h-5 flex items-center justify-center transition-transform duration-200 ${isExpanded ? 'rotate-180' : ''}`}>
                                    <svg className="w-3 h-3 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                                    </svg>
                                </div>
                            </div>
                        </div>

                        {/* Expandable Content */}
                        {isExpanded && (
                            <div className="border-t border-slate-100 p-3 space-y-3 animate-in slide-in-from-top-2 duration-200">
                                {visibleColumns.slice(1).map((column) => (
                                    <div key={column.key} className="flex flex-col space-y-1">
                                        <div className="text-xs font-medium text-slate-600 uppercase tracking-wide">
                                            {column.label}
                                        </div>
                                        <div className="text-sm text-slate-900 bg-slate-50 p-2 rounded">
                                            {column.render
                                                ? column.render(item)
                                                : defaultRenderCell(item[column.key])
                                            }
                                        </div>
                                    </div>
                                ))}
                                
                                {/* Action Buttons */}
                                <div className="flex items-center justify-end space-x-2 pt-2 border-t border-slate-100">
                                    <button className="px-3 py-1.5 text-xs bg-teal-50 text-teal-700 rounded-md hover:bg-teal-100 transition-colors">
                                        Düzenle
                                    </button>
                                    <button className="px-3 py-1.5 text-xs bg-slate-50 text-slate-700 rounded-md hover:bg-slate-100 transition-colors">
                                        Görüntüle
                                    </button>
                                </div>
                            </div>
                        )}
                    </div>
                )
            })}
        </div>
    )
}