import { useTableContext } from "../../../../context/TableContext"
import MobileFilter from "../filters/MobileFilter"
import type { TableProps } from "../types/types"
import MobileCardView from "../ui/MobileCardView"
import Pagination from "../ui/Pagination"
import Toolbar from "../ui/ToolBar"

export default function MobileTable<T extends Record<string, unknown>>(props: TableProps<T>) {
    const { isFilterOpen } = useTableContext()

    return (
        <>
            <Toolbar {...props.toolbar} />
            
            <div className="p-4 space-y-4">
                {isFilterOpen && <MobileFilter onFilter={props.toolbar?.onFilter ?? (() => { })} />}
                <MobileCardView data={props.data} />
            </div>

            <Pagination 
                onPageChange={props.onPageChange} 
                onPageSizeChange={props.onPageSizeChange} 
                dataLength={props.data.length} 
            />
        </>
    )
}