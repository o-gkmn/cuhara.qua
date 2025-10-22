import { useState, useEffect } from "react"
import MobileTable from "./MobileTable"
import DesktopTable from "./DesktopTable"
import type { TableProps } from "../types/types"

export default function Table<T extends Record<string, unknown>>(props: TableProps<T>) {
    const [isMobile, setIsMobile] = useState(false)

    // Mobile detection
    useEffect(() => {
        const checkMobile = () => {
            setIsMobile(window.innerWidth < 768)
        }
        checkMobile()
        window.addEventListener('resize', checkMobile)
        return () => window.removeEventListener('resize', checkMobile)
    }, [])

    return (
        <div className="relative m-4 md:m-6 lg:m-8 border border-slate-200 rounded-lg p-2 md:p-3 bg-white">
            {isMobile ? (
                <MobileTable {...props} />
            ) : (
                <DesktopTable {...props} />
            )}
        </div>
    )
}