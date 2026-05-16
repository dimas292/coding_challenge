import { createContext, useCallback, useContext, useState } from "react";

export interface Category {
  id: number;
  name: string;
  color: string;
  created_at: string;
}

export interface Todo {
  id: number;
  title: string;
  description: string;
  category: Category;
  priority: string;
  completed: boolean;
  due_date: string | null;
  created_at: string | null;
}

interface TodoContextType {
  todos: Todo[];
  setTodos: (todos: Todo[]) => void;
  addTodo: (todo: Todo) => void;
  deleteTodo: (id: number) => void;
  todosVersion: number;
  refreshTodos: () => void;
  editingTodoId: number | null;
  drawerOpen: boolean;
  editTodo: (updatedTodo: Todo) => void;
  openDrawer: (todoId?: number, title?: string) => void;
  closeDrawer: () => void;
  categoryOption: { label: string; value: number }[];
  setCategoryOption: (
    categoryOption: { label: string; value: number }[],
  ) => void;
  priorityParam: (priority: string) => string;
  categoryParam: (category: string) => string;
  completedParam: (completed: string) => string;
}

const TodoContext = createContext<TodoContextType>({} as TodoContextType);

export const TodoProvider = ({ children }: { children: React.ReactNode }) => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [categoryOption, setCategoryOption] = useState<
    { label: string; value: number }[]
  >([]);
  const [todosVersion, setTodosVersion] = useState(0);
  const [editingTodoId, setEditingTodoId] = useState<number | null>(null);
  const [drawerOpen, setDrawerOpen] = useState(false);

  const addTodo = useCallback((newTodo: Todo) => {
    setTodos((prevTodos) => [...prevTodos, newTodo]);
  }, []);

  const priorityParam = (priority: string) => {
    if (priority === "") {
      return "";
    } else {
      return priority;
    }
  };
  const categoryParam = (category: string) => {
    if (category === "") {
      return "";
    } else {
      return category;
    }
  };
  const completedParam = (completed: string) => {
    if (completed === "") {
      return "";
    } else {
      return completed;
    }
  };

  const deleteTodo = useCallback((id: number) => {
    setTodos((prevTodos) => prevTodos.filter((todo) => todo.id !== id));
  }, []);

  const editTodo = useCallback((updatedTodo: Todo) => {
    setTodos((prevTodos) =>
      prevTodos.map((todo) =>
        todo.id === updatedTodo.id ? updatedTodo : todo,
      ),
    );
  }, []);

  const refreshTodos = useCallback(() => {
    setTodosVersion((prev) => prev + 1);
  }, []);

  const resetDrawerContext = useCallback(() => {
    setCategoryOption([]);
  }, []);

  const openDrawer = useCallback(
    (todoId?: number) => {
      resetDrawerContext();
      setEditingTodoId(todoId ?? null);
      setDrawerOpen(true);
    },
    [resetDrawerContext],
  );

  const closeDrawer = useCallback(() => {
    setDrawerOpen(false);
    setEditingTodoId(null);
    resetDrawerContext();
  }, [resetDrawerContext]);

  const value = {
    todos,
    setTodos,
    addTodo,
    deleteTodo,
    todosVersion,
    refreshTodos,
    editingTodoId,
    drawerOpen,
    editTodo,
    openDrawer,
    closeDrawer,
    categoryOption,
    setCategoryOption,
    priorityParam,
    categoryParam,
    completedParam,
  };

  return <TodoContext.Provider value={value}>{children}</TodoContext.Provider>;
};

export const UseTodo = () => {
  const context = useContext(TodoContext);
  if (!context) {
    throw new Error("useTodo must be used within a TodoProvider");
  }
  return context;
};
