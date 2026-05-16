import "./App.css";
import { Spin } from "antd";
import { useEffect, useState } from "react";
import DataTable from "./components/DataTable";
import { UseTodo } from "./context/TodoContext";
import CreateTodo from "./components/CreateTodo";
import EditTodo from "./components/EditTodo";
import { useTodoFetch } from "./hooks/useTodoFetch";
import { TodoHeader } from "./components/ui/TodoHeader";
import { SearchBar } from "./components/ui/SearchBar";
import FilterBar from "./components/ui/FilterBar";

const App = () => {
  const { openDrawer, todosVersion, todos } = UseTodo();
  const { loading, fetchTodos } = useTodoFetch();

  const [pagination, setPagination] = useState({
    current_page: 1,
    per_page: 10,
    total: 0,
    total_pages: 0,
  });
  const [search, setSearch] = useState("");
  const [priority, setPriority] = useState("");
  const [category, setCategory] = useState("");
  const [completed, setCompleted] = useState("");

  useEffect(() => {
    fetchTodos(
      pagination.current_page,
      pagination.per_page,
      search,
      priority,
      category,
      completed,
    ).then((responsePagination) => {
      if (responsePagination) {
        setPagination((prev) => ({ ...prev, ...responsePagination }));
      }
    });
  }, [
    fetchTodos,
    pagination.current_page,
    pagination.per_page,
    search,
    priority,
    category,
    completed,
    todosVersion,
  ]);

  const handlePaginationChange = (page: number, pageSize: number) => {
    setPagination((prev) => ({
      ...prev,
      current_page: page,
      per_page: pageSize,
    }));
  };

  const handleSearch = (value: string) => {
    setTimeout(() => {
      setPagination((prev) => ({ ...prev, current_page: 1 }));
    }, 0);
    setSearch(value);
    setPagination((prev) => ({ ...prev, current_page: 1 }));
  };

  if (loading) {
    return (
      <div
        style={{ display: "flex", justifyContent: "center", padding: "50px" }}
      >
        <Spin />
      </div>
    );
  }

  return (
    <>
      <TodoHeader onAddTodo={() => openDrawer()} />
      <SearchBar onSearch={handleSearch} />
      <FilterBar
        onPriorityChange={setPriority}
        onCategoryChange={(value) => setCategory(value ? String(value) : "")}
        onStatusChange={setCompleted}
      />
      <CreateTodo />
      <EditTodo />
      <DataTable
        data={todos}
        pagination={{
          current: pagination.current_page,
          pageSize: pagination.per_page,
          total: pagination.total,
        }}
        onPaginationChange={handlePaginationChange}
      />
    </>
  );
};

export default App;
