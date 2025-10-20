import Table, { type TableColumnProps, type TableProps } from "../components/common/Table";
import MainLayout from "../components/layouts/MainLayout";
import { TableProvider } from "../context/TableContext";

type Row = {
    id: number
    name: string
    email: string
    role: string
    status: "active" | "pending" | "disabled"
    createdAt: string
    updatedAt: string
    isActive: boolean
}

export default function TablePage() {
    const columns = [
        { key: 'id', label: 'Id', show: true, sortable: true, filter: { type: 'number', operators: ['eq', 'gte', 'lt', 'gt', 'lte',] } },
        { key: 'name', label: 'Name', show: true, sortable: true, filter: { type: 'text', operators: ['contains', 'equals', 'startsWith', 'endsWith', 'notEquals', 'empty', 'notEmpty'] } },
        { key: 'email', label: 'Email', show: true, sortable: true, filter: { type: 'text', operators: ['contains', 'equals', 'startsWith', 'endsWith', 'notEquals', 'empty', 'notEmpty'] } },
        { key: 'role', label: 'Role', show: true, sortable: true, filter: { type: 'text', operators: ['contains', 'equals', 'startsWith', 'endsWith', 'notEquals', 'empty', 'notEmpty'] } },
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
                return (
                    <span className={`inline-flex items-center rounded-md px-2 py-0.5 text-xs font-medium ${map[item.status]}`}>
                        {item.status}
                    </span>
                )
            }
        },
    ]

    const data: Row[] = [
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
        { id: 1, name: 'John Doe', isActive: true, email: 'john.doe@example.com', role: 'Admin', status: 'active', createdAt: '2021-01-01', updatedAt: '2021-01-01' },
    ]

    const tableProps: TableProps<Row> = {
        columns: columns as TableColumnProps<Row>[],
        data: data as Row[],
        title: 'Users',
        toolbar: {
            onRefresh: () => {
                console.log('refresh')
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


    return (
        <MainLayout>
            <TableProvider columns={columns as TableColumnProps[]}>
                <Table {...tableProps} />
            </TableProvider>
        </MainLayout >
    )
}