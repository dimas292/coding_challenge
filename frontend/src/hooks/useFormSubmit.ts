import { notification } from "antd";
import type { FormInstance } from "antd";
import type { Todo } from "../context/TodoContext";
import { CreateTodo, UpdateTodo } from "../services/TodoService";
import { FORM_MODE, type FormMode } from "../utils/constants";

interface SubmitParams {
  form: FormInstance;
  mode: FormMode;
  initialData?: Todo;
  onSuccess: () => void;
}

export const useFormSubmit = () => {
  const handleFormSubmit = async ({
    form,
    mode,
    initialData,
    onSuccess,
  }: SubmitParams) => {
    try {
      const values = form.getFieldsValue();

      if (mode === FORM_MODE.CREATE) {
        await CreateTodo(values);
        notification.success({
          message: "Success",
          description: "Todo created successfully",
        });
      } else {
        if (!initialData) return;
        await UpdateTodo(initialData.id, values);
        notification.success({
          message: "Success",
          description: "Todo updated successfully",
        });
      }

      onSuccess();
    } catch (error) {
      notification.error({
        message: "Error",
        description: error instanceof Error ? error.message : "Unknown error",
      });
    }
  };

  return { handleFormSubmit };
};
