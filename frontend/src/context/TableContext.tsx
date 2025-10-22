/* eslint-disable react-refresh/only-export-components */

import { createContext, useContext, useRef, useState } from "react";
import type { TableColumnProps } from "../components/common/table/types/types";
import type { FilterParam } from "../components/common/table/types/types";

interface TableContextValue {
    columns: TableColumnProps[]
    filterParams: Record<string, FilterParam>
    isFilterOpen: boolean
    isSettingsOpen: boolean
    activeFilterColumn: TableColumnProps | null
    settingsButtonRef: React.RefObject<HTMLButtonElement | null>
    filterButtonRefs: React.RefObject<Record<string, HTMLButtonElement | null>>
    setColumns: React.Dispatch<React.SetStateAction<TableColumnProps[]>>
    setFilterParams: React.Dispatch<React.SetStateAction<Record<string, FilterParam>>>
    setIsFilterOpen: React.Dispatch<React.SetStateAction<boolean>>
    setIsSettingsOpen: React.Dispatch<React.SetStateAction<boolean>>
    setActiveFilterColumn: React.Dispatch<React.SetStateAction<TableColumnProps | null>>
}

export const TableContext = createContext<TableContextValue | null>(null)

export function TableProvider({ children, columns: initialColumns }: { children: React.ReactNode, columns?: TableColumnProps[] }) {
    const [columns, setColumns] = useState<TableColumnProps[]>(initialColumns ?? [])
    const [activeFilterColumn, setActiveFilterColumn] = useState<TableColumnProps | null>(null)
    
    const [filterParams, setFilterParams] = useState<Record<string, FilterParam>>({})
    
    const [isFilterOpen, setIsFilterOpen] = useState(false)
    const [isSettingsOpen, setIsSettingsOpen] = useState(false)

    const settingsButtonRef = useRef<HTMLButtonElement | null>(null);
    const filterButtonRefs = useRef<Record<string, HTMLButtonElement | null>>({})

    return (
        <TableContext.Provider value={
            {
                columns,
                filterParams,
                isFilterOpen,
                isSettingsOpen,
                activeFilterColumn,
                settingsButtonRef,
                filterButtonRefs,
                setColumns,
                setFilterParams,
                setIsFilterOpen,
                setIsSettingsOpen,
                setActiveFilterColumn,
            }}>
            {children}
        </TableContext.Provider>
    )
}

export function useTableContext() {
    const ctx = useContext(TableContext)
    if (!ctx) throw new Error("useTableContext must be used within TableProvider")
    return ctx
}