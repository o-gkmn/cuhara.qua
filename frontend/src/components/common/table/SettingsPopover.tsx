import { useEffect, useRef, useState } from "react";
import { FaGripVertical, FaArrowUp, FaArrowDown } from "react-icons/fa6";
import type { TableColumnProps } from "../Table";
import { useTableContext } from "../../../context/TableContext";

export default function SettingsPopover() {
    const [settingsPopoverRef, setSettingsPopoverRef] = useState<HTMLDivElement | null>(null)
    const [settingsPosition, setSettingsPosition] = useState<{ top: number; left: number }>({ top: 0, left: 0 })
    const settingsCloseTimerRef = useRef<number | null>(null)

    const { isSettingsOpen, setIsSettingsOpen, settingsButtonRef, columns, setColumns } = useTableContext();
    const [draggingKey, setDraggingKey] = useState<string | null>(null)

    // Keep open when hovering button or popover; close when leaving both with a small grace period
    useEffect(() => {
        if (window.innerWidth < 640) return // Mobile'da hareket kontrolü yapmaya gerek yok
        if (!isSettingsOpen) return
        function handlePointerMove(e: PointerEvent) {
            if (draggingKey) return
            const target = e.target as Node | null
            if (!target) return
            const overPopover = !!settingsPopoverRef && settingsPopoverRef.contains(target)
            const overButton = !!settingsButtonRef.current && settingsButtonRef.current.contains(target)
            if (overPopover || overButton) {
                if (settingsCloseTimerRef.current != null) {
                    window.clearTimeout(settingsCloseTimerRef.current)
                    settingsCloseTimerRef.current = null
                }
                return
            }
            if (settingsCloseTimerRef.current == null) {
                settingsCloseTimerRef.current = window.setTimeout(() => {
                    setIsSettingsOpen(false)
                    settingsCloseTimerRef.current = null
                }, 200)
            }
        }
        document.addEventListener('pointermove', handlePointerMove)
        return () => {
            document.removeEventListener('pointermove', handlePointerMove)
            if (settingsCloseTimerRef.current != null) {
                window.clearTimeout(settingsCloseTimerRef.current)
                settingsCloseTimerRef.current = null
            }
        }
    }, [isSettingsOpen, settingsPopoverRef, draggingKey, settingsButtonRef, setIsSettingsOpen])

    // Compute and track settings popover position relative to the button
    useEffect(() => {
        if (!isSettingsOpen || !settingsButtonRef.current) return
        const updatePosition = () => {
            // Calculate popover width for responsive design
            const popoverWidth = window.innerWidth < 640 ? Math.min(280, window.innerWidth - 20) : 256
            const rect = settingsButtonRef.current!.getBoundingClientRect()
            let left = rect.left
            // If popover is too wide, move it to the left
            if (rect.left + popoverWidth > window.innerWidth) {
                left = window.innerWidth - popoverWidth
            }
            setSettingsPosition({ top: rect.bottom + 8, left: left })
        }
        updatePosition()
        // Listen to scroll on capture to catch scrolls in any ancestor
        window.addEventListener('scroll', updatePosition, true)
        window.addEventListener('resize', updatePosition)
        return () => {
            window.removeEventListener('scroll', updatePosition, true)
            window.removeEventListener('resize', updatePosition)
        }
    }, [isSettingsOpen, settingsButtonRef])

    function handleColumnChange(column: TableColumnProps) {
        setColumns(columns.map((c) => c.key === column.key ? { ...c, show: !c.show } : c))
    }

    function handleDragStart(key: string) {
        setDraggingKey(key)
    }

    function handleDrop(targetKey: string) {
        if (!draggingKey || draggingKey === targetKey) return
        const current = [...columns]
        const fromIndex = current.findIndex((c) => c.key === draggingKey)
        const toIndex = current.findIndex((c) => c.key === targetKey)
        if (fromIndex === -1 || toIndex === -1) {
            setDraggingKey(null)
            return
        }
        current.splice(fromIndex, 1)
        current.splice(toIndex, 0, columns.find((c) => c.key === draggingKey)!)
        setColumns(current)
        setDraggingKey(null)
    }

    // Mobile-friendly up/down functions
    function moveColumnUp(key: string) {
        const current = [...columns]
        const index = current.findIndex((c) => c.key === key)
        if (index > 0) {
            [current[index - 1], current[index]] = [current[index], current[index - 1]]
            setColumns(current)
        }
    }

    function moveColumnDown(key: string) {
        const current = [...columns]
        const index = current.findIndex((c) => c.key === key)
        if (index < current.length - 1) {
            [current[index], current[index + 1]] = [current[index + 1], current[index]]
            setColumns(current)
        }

    }

    return isSettingsOpen && settingsButtonRef.current && (
        <div
            ref={setSettingsPopoverRef}
            onMouseDown={(e) => e.stopPropagation()}
            style={{
                top: settingsPosition.top,
                left: settingsPosition.left
            }}
            className={`
                sm:fixed
                w-full sm:w-56 
                bg-white 
                border 
                border-slate-200 
                rounded-md 
                shadow-none sm: shadow-lg 
                p-3 sm:p-3 
                text-left 
                z-40
                max-w-[calc(100vw-20px)] sm:max-w-none
            `}
        >
            <div className="text-xs sm:text-sm text-slate-700 font-semibold mb-2">Sütunlar</div>
            <div className="flex flex-col gap-1 text-left text-xs sm:text-sm text-slate-700 bg-gray-100 rounded-md p-2">
                {columns.map((column) => {
                    return (
                        <div
                            key={column.key}
                            className={`flex items-center gap-1.5 sm:gap-2 rounded-md px-2 py-1.5 sm:py-1 ${draggingKey === column.key ? 'bg-slate-200' : 'bg-transparent'}`}
                            draggable
                            onDragStart={() => handleDragStart(column.key)}
                            onDragOver={(e) => e.preventDefault()}
                            onDrop={() => handleDrop(column.key)}
                        >
                            {/* Desktop drag handle */}
                            <span className="hidden sm:block cursor-move select-none text-slate-500">
                                <FaGripVertical className="w-3 h-3" />
                            </span>
                            
                            {/* Mobile up/down buttons */}
                            <div className="flex flex-col sm:hidden gap-0.5">
                                <button
                                    onClick={() => moveColumnUp(column.key)}
                                    disabled={columns.findIndex(c => c.key === column.key) === 0}
                                    className="p-0.5 text-slate-500 hover:text-slate-700 disabled:opacity-30 disabled:cursor-not-allowed"
                                >
                                    <FaArrowUp className="w-2 h-2" />
                                </button>
                                <button
                                    onClick={() => moveColumnDown(column.key)}
                                    disabled={columns.findIndex(c => c.key === column.key) === columns.length - 1}
                                    className="p-0.5 text-slate-500 hover:text-slate-700 disabled:opacity-30 disabled:cursor-not-allowed"
                                >
                                    <FaArrowDown className="w-2 h-2" />
                                </button>
                            </div>
                            
                            <input
                                type="checkbox"
                                className="w-4 h-4 sm:w-3 sm:h-3"
                                checked={column.show}
                                onChange={() => {
                                    handleColumnChange(column)
                                }} />
                            <span className="truncate text-xs sm:text-sm">{column.label}</span>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}
