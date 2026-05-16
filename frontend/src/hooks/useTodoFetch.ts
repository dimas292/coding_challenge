import { useCallback, useRef, useState } from "react";
import { notification } from "antd";
import { UseTodo } from "../context/TodoContext";
import type { Todo } from "../context/TodoContext";

interface PaginationResponse {
  current_page: number;
  per_page: number;
  total: number;
  total_pages: number;
}

interface TodoResponse<T> {
  data: T[];
  pagination?: PaginationResponse;
}

export const useTodoFetch = () => {
  const { setTodos } = UseTodo();
  const [loading, setLoading] = useState(false);
  const loadingRef = useRef(false);

  const fetchTodos = useCallback(
    async (
      page: number,
      limit: number,
      search: string,
      priority: string,
      category: string,
      completed: string,
    ) => {
      if (loadingRef.current) return null; // Prevent multiple simultaneous fetches
      loadingRef.current = true;
      setLoading(true);
      try {
        const response = await fetch(
          `${import.meta.env.VITE_API_URL}/todos/?page=${page}&limit=${limit}&search=${search}&priority=${priority}&category=${category}&completed=${completed}`,
        );
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        const data: TodoResponse<Todo> = await response.json();
        setTodos(data.data ?? []);
        return data.pagination ?? null;
      } catch (error) {
        notification.error({
          message: "Error",
          description: error instanceof Error ? error.message : "Unknown error",
        });
        return null;
      } finally {
        loadingRef.current = false;
        setLoading(false);
      }
    },
    [setTodos],
  );

  return { loading, fetchTodos };
};
