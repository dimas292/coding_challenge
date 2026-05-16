import type { ColumnProps } from "antd/es/table";
import { Popconfirm } from "antd";
import { notification } from "antd";
import type { Todo } from "../context/TodoContext";
import { formatDate, compareDates } from "../utils/dateUtils";
import {
  renderPriorityTag,
  renderCompletedTag,
  renderCategoryTag,
} from "../utils/tagRenderer";
import * as TodoService from "../services/TodoService";
import { UseTodo } from "../context/TodoContext";
import { DeleteOutlined, EditOutlined } from "@ant-design/icons/es/icons/index";

export const getTableColumns = (): ColumnProps<Todo>[] => {
  const { refreshTodos, openDrawer } = UseTodo();

  const handleDelete = async (id: number) => {
    try {
      await TodoService.DeleteTodo(id);
      notification.success({
        message: "Success",
        description: "Todo deleted successfully",
      });
      refreshTodos();
    } catch (error) {
      notification.error({
        message: "Error",
        description:
          error instanceof Error ? error.message : "Failed to delete todo",
      });
    }
  };

  return [
    {
      title: "Title",
      dataIndex: "title",
      key: "title",
      width: 160,
      ellipsis: true,
    },
    {
      title: "Description",
      dataIndex: "description",
      key: "description",
      width: 250,
      ellipsis: true,
    },
    {
      title: "Priority",
      dataIndex: "priority",
      key: "priority",
      width: 150,
      render: renderPriorityTag,
    },
    {
      title: "Category",
      dataIndex: "category",
      key: "category",
      width: 150,
      render: (_, todo) => renderCategoryTag(todo),
    },
    {
      title: "Completed",
      dataIndex: "completed",
      key: "completed",
      width: 150,
      render: renderCompletedTag,
    },
    {
      title: "Created At",
      dataIndex: "created_at",
      sorter: (a, b) => compareDates(a.created_at, b.created_at),
      render: (date: string | null) => formatDate(date),
    },
    {
      title: "Action",
      width: 150,
      align: "center",
      render: (_, todo) => {
        return (
          <div
            style={{ display: "flex", gap: "8px", justifyContent: "center" }}
          >
            <a
              href="#"
              onClick={(e) => {
                e.preventDefault();
                openDrawer(todo.id);
              }}
            >
              Edit <EditOutlined />
            </a>
            <Popconfirm
              title="Delete Todo"
              description="Are you sure you want to delete this todo?"
              onConfirm={() => handleDelete(todo.id)}
              okText="Yes"
              cancelText="No"
            >
              <a
                href="#"
                onClick={(e) => e.preventDefault()}
                style={{ color: "#ff4d4f" }}
              >
                Delete <DeleteOutlined />
              </a>
            </Popconfirm>
          </div>
        );
      },
    },
  ];
};
