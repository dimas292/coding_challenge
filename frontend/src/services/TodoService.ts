import type { Todo } from "../context/TodoContext";

export const CreateTodo = async (data: Todo) => {
  const response = await fetch(`${import.meta.env.VITE_API_URL}/todos/`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  const createdData = await response.json();
  return createdData;
};

export const UpdateTodo = async (id: number, data: Todo) => {
  if (typeof id !== "number") {
    throw new Error("ID must be a number");
  }
  const response = await fetch(`${import.meta.env.VITE_API_URL}/todos/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  const updatedData = await response.json();
  return updatedData;
};

export const DeleteTodo = async (id: number) => {
  if (typeof id !== "number") {
    throw new Error("ID must be a number");
  }
  const response = await fetch(`${import.meta.env.VITE_API_URL}/todos/${id}`, {
    method: "DELETE",
  });
  const deletedData = await response.json();
  return deletedData;
};
