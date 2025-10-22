import { useTableContext } from "../../../../context/TableContext"
import { TextFilter, NumberFilter, DateFilter, BooleanFilter, SelectFilter, applyFilter } from "./components"

export default function MobileFilter({ onFilter }: { onFilter: (query: string) => void }) {
    const { columns, filterParams, setFilterParams, setIsFilterOpen, activeFilterColumn, setActiveFilterColumn } = useTableContext()

    return (
        <div className={`
            h-fit
            bg-gray-100
            rounded-lg
            shadow-sm
            p-4
            border
            border-gray-200
            animate-in
            slide-in-from-bottom-2
            duration-200
            space-y-2
        `}>
            {/* Sütunların üzerine tıklanıldıktan sonra activeFilterColumn değiştirilir ve filtreler görünür hale gelir */}
            {/* activeFilterColumn state'i kaldırmaya çalışırsan tüm inputların state'leri aynı anda değişecektir. */}
            {columns.map((column) => (
                <div key={column.key}
                    className="flex flex-col gap-2 p-2 bg-white rounded-lg border border-gray-200"
                    onClick={() => {
                        setActiveFilterColumn(column)
                    }}
                >
                    <h3 className="text-sm font-medium">{column.key}</h3>
                    {activeFilterColumn?.key === column.key && (
                        <div className="flex flex-col gap-2">
                            {activeFilterColumn?.filter?.type === 'text' && <TextFilter />}
                            {activeFilterColumn?.filter?.type === 'number' && <NumberFilter />}
                            {activeFilterColumn?.filter?.type === 'date' && <DateFilter />}
                            {activeFilterColumn?.filter?.type === 'boolean' && <BooleanFilter />}
                            {activeFilterColumn?.filter?.type === 'select' && <SelectFilter />}
                        </div>
                    )}
                </div>
            ))}

            <div className="flex flex-col sm:flex-row justify-end gap-2 pt-3 border-t border-slate-100">
                <button
                    className="px-3 py-1.5 sm:py-1 text-xs bg-teal-600 text-white hover:bg-teal-700 rounded transition-colors"
                    onClick={() => {
                        const query = applyFilter(filterParams)
                        onFilter(query)
                        setIsFilterOpen(false)
                    }}
                >
                    Uygula
                </button>
                <button
                    className="px-3 py-1.5 sm:py-1 text-xs text-orange-600 hover:text-orange-800 hover:bg-orange-50 rounded transition-colors"
                    onClick={() => {
                        setFilterParams(() => {
                            return {}
                        })
                        setIsFilterOpen(false)
                    }}
                >
                    Sıfırla
                </button>
                <button
                    className="px-3 py-1.5 sm:py-1 text-xs text-slate-600 hover:text-slate-800 hover:bg-slate-100 rounded transition-colors"
                    onClick={() => {
                        setIsFilterOpen(false)
                        setFilterParams(() => {
                            return {}
                        })
                    }}
                >
                    İptal
                </button>
            </div>
        </div>

    )
}