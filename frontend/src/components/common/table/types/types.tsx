import type { ReactNode } from "react"
import type { ToolbarProps } from "../ui/ToolBar"

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

export interface TableColumnProps<T extends Record<string, unknown> = Record<string, unknown>> {
    key: keyof T & string
    label: string
    sortable?: boolean
    show?: boolean
    filter?: ColumnFilterProps
    render?: (item: T) => ReactNode
}

export interface TableProps<T extends Record<string, unknown> = Record<string, unknown>> {
    data: T[],
    title?: string,
    toolbar?: ToolbarProps
    onPageChange?: (page: number) => void
    onPageSizeChange?: (pageSize: number) => void
    onSorted?: (field: string, operator: SortOperator) => void
}