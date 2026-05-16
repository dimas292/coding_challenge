import DrawerForm from "./DrawerForm";
import { UseTodo } from "../context/TodoContext";
import { useEffect } from "react";
import { GetCategory } from "../services/CategoryService";
import type { Category } from "../context/TodoContext";
import type { Todo } from "../context/TodoContext";
import { useState } from "react";

const EditTodo = () => {
  const {
    todos,
    drawerOpen,
    closeDrawer,
    editingTodoId,
    categoryOption,
    setCategoryOption,
  } = UseTodo();

  const [dataEdit, setDataEdit] = useState<Todo[]>([]);

  useEffect(() => {
    if (drawerOpen && editingTodoId !== null) {
      GetCategory().then((categoryData) => {
        const options = categoryData.data.map((category: Category) => ({
          label: category.name,
          value: category.id,
        }));
        setCategoryOption(options);
      });
      const dataEdit = () => {
        const todo = todos.filter((todo: Todo) => todo.id === editingTodoId);
        setDataEdit(todo);
      };
      dataEdit();
    }
  }, [drawerOpen, editingTodoId, setCategoryOption, todos]);

  return (
    <>
      <DrawerForm
        open={drawerOpen && editingTodoId !== null}
        onClose={closeDrawer}
        categoryOption={categoryOption}
        title={"Edit"}
        mode="edit"
        initialData={dataEdit[0]}
      />
    </>
  );
};

export default EditTodo;
