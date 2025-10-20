import { useEffect, useRef, useState } from "react";
import { FaGripVertical } from "react-icons/fa6";
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
            const rect = settingsButtonRef.current!.getBoundingClientRect()
            setSettingsPosition({ top: rect.bottom + 8, left: rect.left })
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

    return isSettingsOpen && settingsButtonRef.current && (
        <div
            ref={setSettingsPopoverRef}
            onMouseDown={(e) => e.stopPropagation()}
            className="fixed w-56 bg-white border border-slate-200 rounded-md shadow-lg p-3 text-left z-40"
            style={{
                top: settingsPosition.top,
                left: settingsPosition.left
            }}
        >
            <div className="text-sm text-slate-700 font-semibold mb-2">Sütunlar</div>
            <div className="flex flex-col gap-1 text-left text-sm text-slate-700 bg-gray-100 rounded-md p-2">
                {columns.map((column) => {
                    return (
                        <div
                            key={column.key}
                            className={`flex items-center gap-2 rounded-md px-2 py-1 ${draggingKey === column.key ? 'bg-slate-200' : 'bg-transparent'}`}
                            draggable
                            onDragStart={() => handleDragStart(column.key)}
                            onDragOver={(e) => e.preventDefault()}
                            onDrop={() => handleDrop(column.key)}
                        >
                            <span className="cursor-move select-none text-slate-500">
                                <FaGripVertical className="w-3 h-3" />
                            </span>
                            <input
                                type="checkbox"
                                className=""
                                checked={column.show}
                                onChange={() => {
                                    handleColumnChange(column)
                                }} />
                            <span className="truncate">{column.label}</span>
                        </div>
                    )
                })}
            </div>
        </div>
    )
}