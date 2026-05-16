import { Tag } from "antd";
import type { Todo } from "../context/TodoContext";
import {
  PRIORITY_COLOR_MAP,
  COMPLETED_COLOR_MAP,
  COMPLETED_DISPLAY_MAP,
} from "./constants";

export const renderPriorityTag = (priority: string) => (
  <Tag color={PRIORITY_COLOR_MAP[priority]}>{priority}</Tag>
);

export const renderCompletedTag = (completed: boolean) => (
  <Tag color={COMPLETED_COLOR_MAP[completed ? "true" : "false"]}>
    {COMPLETED_DISPLAY_MAP[completed ? "true" : "false"]}
  </Tag>
);

export const renderCategoryTag = (todo: Todo) => {
  if (!todo.category) return "-";
  return <Tag color={todo.category.color}>{todo.category.name}</Tag>;
};
