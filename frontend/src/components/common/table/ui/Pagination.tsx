import { useState } from "react"
import { FaAngleDoubleLeft, FaAngleDoubleRight, FaAngleLeft, FaAngleRight } from "react-icons/fa"

export type PaginationProps = {
    onPageChange?: (page: number) => void
    onPageSizeChange?: (pageSize: number) => void
    dataLength: number
}

export default function Pagination(props: PaginationProps) {
    const [page, setPage] = useState(1)
    const [pageSize, setPageSize] = useState(10)
    const totalPage = Math.ceil(props.dataLength / pageSize)

    function pageWindow() {
        const current = page ?? 1
        const total = totalPage ?? 1
        const pageWindow = [
            ...(current != 1 ? [current - 1] : []),
            current,
            current + 1,
            ...(current == 1 ? [current + 2] : []),
        ].filter(p => p >= 1 && p <= total)
        return pageWindow
    }

    return (
        <div className="flex flex-col sm:flex-row justify-between items-center gap-3 sm:gap-1 pt-4 px-2 sm:px-0">
        <div className="flex items-center text-xs order-2 sm:order-1">
            <span className="text-teal-700 text-xs sm:text-sm hidden sm:inline">Sayfa Başına</span>
            <span className="text-teal-700 text-xs sm:hidden">Sayfa/</span>
            <select
                className="text-teal-700 hover:bg-gray-100 rounded-md p-1 ml-1 sm:ml-2 disabled:opacity-40 disabled:cursor-not-allowed border border-teal-700 text-xs"
                onChange={(e) => {
                    setPage(1)
                    setPageSize(Number(e.target.value))
                    props.onPageSizeChange?.(Number(e.target.value))
                }}
            >
                {Array.from({ length: 10 }, (_, i) => (i + 1) * 10).map((option) => (
                    <option key={option} value={option}>{option}</option>
                ))}
            </select>
        </div>
        <div className="flex items-center gap-1 order-1 sm:order-2">
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 sm:p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === 1}
            onClick={() => {
                props.onPageChange?.(1)
                setPage(1)
            }}
        >
            <FaAngleDoubleLeft className="w-3 h-3 sm:w-4 sm:h-4" />
        </button>
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 sm:p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === 1}
            onClick={() => {
                props.onPageChange?.(page - 1)
                setPage(page - 1)
            }}
        >
            <FaAngleLeft className="w-3 h-3 sm:w-4 sm:h-4" />
        </button>
        {pageWindow().map((p) => (
            <button
                key={p}
                onClick={() => {
                    props.onPageChange?.(p)
                    setPage(p)
                }}
                className={`hover:bg-gray-100 rounded-md disabled:opacity-40 disabled:cursor-not-allowed w-6 h-6 sm:w-7 sm:h-7 text-xs sm:text-sm ${p === page ? 'bg-teal-600 text-white hover:bg-teal-600' : ''}`}
            >
                {p}
            </button>
        ))}
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 sm:p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === totalPage}
            onClick={() => {
                props.onPageChange?.(page + 1)
                setPage(page + 1)
            }}
        >
            <FaAngleRight className="w-3 h-3 sm:w-4 sm:h-4" />
        </button>
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 sm:p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === totalPage}
            onClick={() => {
                props.onPageChange?.(totalPage ?? 1)
                setPage(totalPage ?? 1)
            }}
        >
            <FaAngleDoubleRight className="w-3 h-3 sm:w-4 sm:h-4" />
        </button>
        <div className="flex items-center gap-1 sm:gap-2">
            <span className="text-teal-700 text-xs sm:text-sm hidden sm:inline">Sayfa</span>
            <input
                className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed border border-teal-700 text-xs text-center w-12 sm:w-16"
                type="number"
                value={page}
                min={1}
                max={totalPage ?? 1}
                onWheel={e => (e.target as HTMLInputElement).blur()}
                onChange={(e) => {
                    const raw = Number(e.target.value)
                    if (!Number.isFinite(raw)) return
                    const maxPage = totalPage ?? 1
                    const clamped = Math.max(1, Math.min(raw, maxPage))
                    props.onPageChange?.(clamped)
                    setPage(clamped)
                }}
            />
        </div>
        </div>
    </div>
    )
}