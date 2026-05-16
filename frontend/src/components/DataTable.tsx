import { Table } from "antd";
import type { TableProps } from "antd";
import type { Todo } from "../context/TodoContext";
import { getTableColumns } from "./TableColumns";

interface DataTableProps {
  data: Todo[];
  pagination?: {
    current: number;
    pageSize: number;
    total: number;
  };
  onPaginationChange?: (page: number, pageSize: number) => void;
}

const DataTable: React.FC<DataTableProps> = ({
  data,
  pagination,
  onPaginationChange,
}) => {
  const onChange: TableProps<Todo>["onChange"] = () => {};

  return (
    <Table<Todo>
      columns={getTableColumns()}
      dataSource={data}
      onChange={onChange}
      rowKey="id"
      showSorterTooltip={{ target: "sorter-icon" }}
      pagination={
        pagination
          ? {
              current: pagination.current,
              pageSize: pagination.pageSize,
              total: pagination.total,
              showSizeChanger: true,
              onChange: (page, pageSize) => {
                onPaginationChange?.(page, pageSize ?? pagination.pageSize);
              },
              onShowSizeChange: (page, pageSize) => {
                onPaginationChange?.(page, pageSize);
              },
            }
          : false
      }
    />
  );
};

export default DataTable;
