import type { FilterParam } from "../../types/types"

export function applyFilter(filterParams: Record<string, FilterParam>): string {
    const entries = Object.entries(filterParams)

    return entries.map(([columnKey, filterParam], i) => {
        if (filterParam.operator === undefined) return // operator selection must
        const part = `${columnKey} ${filterParam.operator} ${filterParam.value}`
        const hasNext = i < entries.length - 1
        return hasNext ? `${part} AND` : part
    }).join(' ')
}
