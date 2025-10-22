
import { useTableContext } from "../../../../../context/TableContext"
import { operatorTranslations } from "./constants"

export function DateFilter() {
    const { filterParams, setFilterParams, activeFilterColumn } = useTableContext()
    
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