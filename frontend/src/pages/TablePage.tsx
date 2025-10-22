
import Table from "../components/common/table/core/Table";
import type { TableColumnProps, TableProps } from "../components/common/table/types/types";
import { TableProvider } from "../context/TableContext";
import { useUsers } from "../hooks/useUsers";

type Row = {
    id?: number
    name?: string
    email?: string
    vscAccount?: string
    role?: {
        id?: number
        name?: string
    }
    // Additional fields for display purposes
    status?: "active" | "pending" | "disabled"
    createdAt?: string
    updatedAt?: string
    isActive?: boolean
}

export default function TablePage() {
    const { users, loading, error, refetch } = useUsers();

    const columns = [
        { key: 'id', label: 'Id', show: true, sortable: true, filter: { type: 'number', operators: ['eq', 'gte', 'lt', 'gt', 'lte',] } },
        { key: 'name', label: 'Name', show: true, sortable: true, filter: { type: 'text', operators: ['contains', 'equals', 'startsWith', 'endsWith', 'notEquals', 'empty', 'notEmpty'] } },
        { key: 'email', label: 'Email', show: true, sortable: true, filter: { type: 'text', operators: ['contains', 'equals', 'startsWith', 'endsWith', 'notEquals', 'empty', 'notEmpty'] } },
        { key: 'vscAccount', label: 'VSC Account', show: true, sortable: true, filter: { type: 'text', operators: ['contains', 'equals', 'startsWith', 'endsWith', 'notEquals', 'empty', 'notEmpty'] } },
        { 
            key: 'role', 
            label: 'Role', 
            show: true, 
            sortable: true, 
            filter: { type: 'text', operators: ['contains', 'equals', 'startsWith', 'endsWith', 'notEquals', 'empty', 'notEmpty'] },
            render: (item: Row) => item.role?.name || 'N/A'
        },
        { key: 'createdAt', label: 'Created At', sortable: true, show: true, filter: { type: 'date', operators: ['on', 'before', 'after', 'between'] } },
        { key: 'updatedAt', label: 'Updated At', sortable: true, show: true, filter: { type: 'date', operators: ['on', 'before', 'after', 'between'] } },
        { key: 'isActive', label: "Aktif", show: true, sortable: true, filter: { type: 'boolean', operators: ['isTrue', 'isFalse', 'isNotNull', 'isNull'] } },
        {
            key: 'status',
            label: 'Status',
            show: true,
            filter: {
                type: 'select', options: [
                    { label: 'Active', value: 'active' },
                    { label: 'Pending', value: 'pending' },
                    { label: 'Disabled', value: 'disabled' },
                ],
                operators: ['eq', 'neq']
            },
            render: (item: Row) => {
                const map = {
                    active: "bg-green-100 text-green-700",
                    pending: "bg-yellow-100 text-yellow-700",
                    disabled: "bg-gray-100 text-gray-600",
                } as const
                const status = item.status || 'active'
                return (
                    <span className={`inline-flex items-center rounded-md px-2 py-0.5 text-xs font-medium ${map[status]}`}>
                        {status}
                    </span>
                )
            }
        },
    ]

    // Transform API data to match our Row type
    const data: Row[] = users.map(user => ({
        ...user,
        // Add default values for fields that might not exist in API response
        status: 'active' as const,
        createdAt: new Date().toISOString().split('T')[0],
        updatedAt: new Date().toISOString().split('T')[0],
        isActive: true,
    }))

    const tableProps: TableProps<Row> = {
        data: data as Row[],
        title: 'Users',
        toolbar: {
            title: 'Users',
            onRefresh: () => {
                refetch()
            },
            onExport: () => {
                console.log('export')
            },
            onFilter: (query) => {
                console.log(query)
            },
            onNew: () => {
                console.log('new')
            },
            onDelete: () => {
                console.log('delete')
            },
            onSettings: () => {
                console.log('settings')
            },
        },
        onSorted: () => {
            console.log("sorded")
        },
        onPageChange: (page) => {
            console.log('page', page)
        },
        onPageSizeChange: (pageSize) => {
            console.log('pageSize', pageSize)
        },
    }


    // Show loading state
    if (loading) {
        return (
            <div className="flex items-center justify-center h-64">
                <div className="text-lg">Loading users...</div>
            </div>
        )
    }

    // Show error state
    if (error) {
        return (
            <div className="flex items-center justify-center h-64">
                <div className="text-center">
                    <div className="text-red-600 text-lg mb-4">Error loading users</div>
                    <div className="text-gray-600 mb-4">{error}</div>
                    <button 
                        onClick={refetch}
                        className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                    >
                        Try Again
                    </button>
                </div>
            </div>
        )
    }

    return (
        <TableProvider columns={columns as TableColumnProps[]}>
            <Table {...tableProps} />
        </TableProvider>
    )
}