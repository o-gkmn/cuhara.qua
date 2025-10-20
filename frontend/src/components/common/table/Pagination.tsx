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
        <div className="flex justify-end items-center gap-1 pt-4 dasdas">
        <div className="flex items-center mr-auto ml-2 text-xs">
            <span className="text-teal-700 text-sm">Sayfa Başına</span>
            <select
                className="text-teal-700 hover:bg-gray-100 rounded-md p-1 ml-2 disabled:opacity-40 disabled:cursor-not-allowed border border-teal-700"
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
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === 1}
            onClick={() => {
                props.onPageChange?.(1)
                setPage(1)
            }}
        >
            <FaAngleDoubleLeft className="w-4 h-4" />
        </button>
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === 1}
            onClick={() => {
                props.onPageChange?.(page - 1)
                setPage(page - 1)
            }}
        >
            <FaAngleLeft className="w-4 h-4" />
        </button>
        {pageWindow().map((p) => (
            <button
                key={p}
                onClick={() => {
                    props.onPageChange?.(p)
                    setPage(p)
                }}
                className={`hover:bg-gray-100 rounded-md disabled:opacity-40 disabled:cursor-not-allowed w-6 h-6 ${p === page ? 'bg-teal-600 text-white hover:bg-teal-600' : ''}`}
            >
                {p}
            </button>
        ))}
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === totalPage}
            onClick={() => {
                props.onPageChange?.(page + 1)
                setPage(page + 1)
            }}
        >
            <FaAngleRight className="w-4 h-4" />
        </button>
        <button
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 disabled:opacity-40 disabled:cursor-not-allowed"
            disabled={page === totalPage}
            onClick={() => {
                props.onPageChange?.(totalPage ?? 1)
                setPage(totalPage ?? 1)
            }}
        >
            <FaAngleDoubleRight className="w-4 h-4" />
        </button>
        <span className="text-teal-700 ml-3 text-sm">Sayfa</span>
        <input
            className="text-teal-700 hover:bg-gray-100 rounded-md p-1 mr-3 disabled:opacity-40 disabled:cursor-not-allowed border border-teal-700 text-xs text-center"
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
    )
}