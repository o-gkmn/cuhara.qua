import { useState } from "react"
import type { TableProps } from "../Table"
import { FaArrowDown, FaArrowRotateRight, FaEllipsis, FaFilter, FaGear, FaPlus, FaTrash } from "react-icons/fa6"
import { TableProvider, useTableContext } from "../../../context/TableContext"

export default function TableToolbar(props: TableProps) {
    const PRIMARY_BUTTONS = ["refresh", "export", "filter"]
    const SECONDARY_BUTTONS = ["new", "delete", "settings"]

    const [activeButtons, setActiveButtons] = useState<string[]>(PRIMARY_BUTTONS)
    const isExpanded = activeButtons.length > PRIMARY_BUTTONS.length

    const { isSettingsOpen, setIsSettingsOpen, settingsButtonRef } = useTableContext();

    return (
        <TableProvider>
            <div className="flex items-center gap-1 overflow-visible">
                <div className={`flex items-center gap-1 transition-transform duration-300 ${isExpanded ? '-translate-x-1' : 'translate-x-0'}`}>
                    {props.onRefresh && activeButtons.includes("refresh") && (
                        <button
                            key="refresh"
                            onClick={() => { }}
                            className="inline-flex items-center bg-sky-600 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-sky-700"
                            aria-label="Yenile"
                        >
                            <FaArrowRotateRight className="w-3 h-3" />
                            <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                Yenile
                            </span>
                        </button>
                    )}
                    {props.onExport && activeButtons.includes("export") && (
                        <button
                            key="export"
                            onClick={() => { }}
                            className="inline-flex items-center bg-emerald-600 hover:bg-emerald-700 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group"
                            aria-label="Dışa Aktar"
                        >
                            <FaArrowDown className="w-3 h-3" />
                            <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                Dışa Aktar
                            </span>
                        </button>
                    )}
                    {props.onFilter && activeButtons.includes("filter") && (
                        <button
                            key="filter"
                            onClick={() => {
                                setIsFilterOpen(!isFilterOpen)
                                setFilters({})
                                setOpenFilterKey(null)
                            }}
                            className="inline-flex items-center bg-yellow-500 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-yellow-600"
                            aria-label="Filtrele"
                        >
                            <FaFilter className="w-3 h-3" />
                            <span className={`overflow-hidden whitespace-nowrap transition-all duration-200 text-sm ${isFilterOpen ? 'ml-2 max-w-[80px] opacity-100' : 'ml-0 max-w-0 opacity-0'} group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100`}>
                                Filtrele
                            </span>
                        </button>
                    )}
                </div>

                <div
                    className={`flex items-center gap-1 overflow-visible transition-all duration-300 ${isExpanded ? 'max-w-[999px] opacity-100 translate-x-0 border-l pl-2 border-slate-200' : 'max-w-0 opacity-0 translate-x-4 pl-0 border-l-0'}`}
                >
                    {props.onNew && (
                        <button
                            key="new"
                            onClick={() => { }}
                            className="inline-flex items-center bg-teal-600 hover:bg-teal-700 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group"
                            aria-label="Ekle"
                        >
                            <FaPlus className="w-3 h-3" />
                            <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                Yeni
                            </span>
                        </button>
                    )}
                    {props.onDelete && (
                        <button
                            key="delete"
                            onClick={() => { }}
                            className="inline-flex items-center bg-rose-600 hover:bg-rose-700 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group"
                            aria-label="Sil"
                        >
                            <FaTrash className="w-3 h-3" />
                            <span className="ml-0 max-w-0 overflow-hidden whitespace-nowrap opacity-0 transition-all duration-200 text-sm group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100">
                                Sil
                            </span>
                        </button>
                    )}
                    {(
                        <div className="relative inline-block">
                            <button
                                key="settings"
                                ref={settingsButtonRef}
                                onMouseDown={(e) => e.stopPropagation()}
                                onClick={() => {
                                    setIsSettingsOpen(!isSettingsOpen)
                                }}
                                className="inline-flex items-center bg-slate-500 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-slate-600"
                                aria-label="Ayarlar"
                            >
                                <FaGear className="w-3 h-3" />
                                <span className={`overflow-hidden whitespace-nowrap transition-all duration-200 text-sm ${isSettingsOpen ? 'ml-2 max-w-[80px] opacity-100' : 'ml-0 max-w-0 opacity-0'} group-hover:ml-2 group-hover:max-w-[80px] group-hover:opacity-100`}>
                                    Ayarlar
                                </span>
                            </button>
                        </div>
                    )}
                </div>
                <button
                    key="other"
                    onClick={() => {
                        if (activeButtons.length === PRIMARY_BUTTONS.length) {
                            setActiveButtons([...PRIMARY_BUTTONS, ...SECONDARY_BUTTONS])
                        }
                        else {
                            setActiveButtons(PRIMARY_BUTTONS)
                        }
                    }}
                    className="inline-flex items-center z-0 bg-slate-500 text-white px-2 py-2 h-8 rounded-md transition-all duration-200 group hover:bg-slate-600"
                    aria-label="Diğer İşlemler"
                >
                    <FaEllipsis className="w-3 h-3" />
                </button>
            </div>
        </TableProvider>
    )
}