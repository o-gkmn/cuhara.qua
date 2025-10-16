/* eslint-disable react-refresh/only-export-components */

import { createContext, useContext, useRef, useState } from "react";

interface TableContextValue {
    isSettingsOpen: boolean
    setIsSettingsOpen: React.Dispatch<React.SetStateAction<boolean>>
    settingsButtonRef: React.RefObject<HTMLButtonElement | null>
}

export const TableContext = createContext<TableContextValue | null>(null)

export function TableProvider({ children }: { children: React.ReactNode }) {
    const [isSettingsOpen, setIsSettingsOpen] = useState(false)
    const settingsButtonRef = useRef<HTMLButtonElement | null>(null);

    return (
        <TableContext.Provider value={{ isSettingsOpen, setIsSettingsOpen, settingsButtonRef}}>
            {children}
        </TableContext.Provider>
    )
}

export function useTableContext() {
    const ctx = useContext(TableContext)
    if (!ctx) throw new Error("useTableContext must be used within TableProvider")
    return ctx
}