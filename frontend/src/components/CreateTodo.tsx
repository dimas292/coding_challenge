import DrawerForm from "./DrawerForm";
import { UseTodo } from "../context/TodoContext";
import { useEffect } from "react";
import { GetCategory } from "../services/CategoryService";
import type { Category } from "../context/TodoContext";

const CreateTodo = () => {
  const {
    drawerOpen,
    closeDrawer,
    editingTodoId,
    categoryOption,
    setCategoryOption,
  } = UseTodo();

  useEffect(() => {
    if (drawerOpen && editingTodoId === null) {
      GetCategory().then((categoryData) => {
        const options = categoryData.data.map((category: Category) => ({
          label: category.name,
          value: category.id,
        }));
        setCategoryOption(options);
      });
    }
  }, [drawerOpen, editingTodoId, setCategoryOption]);

  return (
    <>
      <DrawerForm
        title={"Tambah"}
        mode="create"
        open={drawerOpen && editingTodoId === null}
        onClose={closeDrawer}
        categoryOption={categoryOption}
      />
    </>
  );
};

export default CreateTodo;
