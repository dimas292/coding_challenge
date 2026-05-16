import {
  Col,
  DatePicker,
  Drawer,
  Form,
  Button,
  Input,
  Row,
  Select,
  Space,
} from "antd";
import type { Todo } from "../context/TodoContext";
import { UseTodo } from "../context/TodoContext";
import dayjs from "dayjs";
import { useEffect } from "react";
import { useFormSubmit } from "../hooks/useFormSubmit";
import { FORM_MODE, type FormMode } from "../utils/constants";
import { GetCategory } from "../services/CategoryService";
import type { Category } from "../context/TodoContext";

interface DrawerFormProps {
  open: boolean;
  onClose: () => void;
  title: string;
  mode: FormMode;
  categoryOption?: { label: string; value: number }[];
  initialData?: Todo;
}

const DrawerForm: React.FC<DrawerFormProps> = ({
  open,
  onClose,
  title,
  mode,
  categoryOption = [],
  initialData,
}) => {
  const { refreshTodos, setCategoryOption } = UseTodo();
  const { handleFormSubmit } = useFormSubmit();
  const [form] = Form.useForm();

  // Load categories on drawer open
  useEffect(() => {
    if (open && categoryOption.length === 0) {
      GetCategory().then((data) => {
        const options = data.data.map((cat: Category) => ({
          label: cat.name,
          value: cat.id,
        }));
        setCategoryOption(options);
      });
    }
  }, [open, categoryOption.length, setCategoryOption]);

  // Initialize form with data or reset
  useEffect(() => {
    if (initialData && mode === FORM_MODE.EDIT) {
      form.setFieldsValue(getFormValues(initialData));
    } else {
      form.resetFields();
    }
  }, [initialData, form, mode]);

  const handleSubmit = async () => {
    form.submit();
    await handleFormSubmit({
      form,
      mode,
      initialData,
      onSuccess: () => {
        refreshTodos();
        onClose();
      },
    });
  };

  return (
    <Drawer
      title={title}
      size={720}
      onClose={onClose}
      open={open}
      styles={{ body: { paddingBottom: 80 } }}
      extra={
        <Space>
          <Button onClick={onClose}>Cancel</Button>
          <Button onClick={handleSubmit} type="primary">
            {title}
          </Button>
        </Space>
      }
    >
      <Form layout="vertical" requiredMark={false} form={form}>
        <FormFields initialData={initialData} categoryOption={categoryOption} />
      </Form>
    </Drawer>
  );
};

// Helper component for form fields
const FormFields: React.FC<{
  initialData?: Todo;
  categoryOption: { label: string; value: number }[];
}> = ({ initialData, categoryOption }) => {
  const dueDate = initialData?.due_date
    ? dayjs(initialData.due_date)
    : undefined;
  const createdAt = initialData?.created_at
    ? dayjs(initialData.created_at)
    : undefined;

  return (
    <>
      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            name="title"
            label="Title"
            rules={[{ required: true, message: "Please enter title" }]}
          >
            <Input placeholder="Please enter title" />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={24}>
          <Form.Item
            name="description"
            label="Description"
            rules={[{ required: true, message: "Please enter description" }]}
          >
            <Input.TextArea rows={2} placeholder="Please enter description" />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            name="category_id"
            label="Category"
            rules={[{ required: true, message: "Please select a category" }]}
          >
            <Select
              placeholder="Please select a category"
              options={categoryOption}
            />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            name="priority"
            label="Priority"
            rules={[{ required: true, message: "Please choose priority" }]}
          >
            <Select
              placeholder="Please choose priority"
              options={[
                { label: "Low", value: "low" },
                { label: "Medium", value: "medium" },
                { label: "High", value: "high" },
              ]}
            />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            name="completed"
            label="Completed"
            rules={[{ required: true, message: "Please choose status" }]}
          >
            <Select
              placeholder="Please choose status"
              options={[
                { label: "True", value: true },
                { label: "False", value: false },
              ]}
            />
          </Form.Item>
        </Col>
        <Col span={12}>
          <Form.Item
            name="due_date"
            label="Due Date"
            initialValue={dueDate}
            rules={[{ required: true, message: "Please choose due date" }]}
          >
            <DatePicker
              style={{ width: "100%" }}
              getPopupContainer={(trigger) => trigger.parentElement!}
            />
          </Form.Item>
        </Col>
      </Row>

      <Row gutter={16}>
        <Col span={12}>
          <Form.Item
            name="created_at"
            label="Created At"
            initialValue={createdAt}
            rules={[{ required: true, message: "Please choose created date" }]}
          >
            <DatePicker
              style={{ width: "100%" }}
              getPopupContainer={(trigger) => trigger.parentElement!}
            />
          </Form.Item>
        </Col>
      </Row>
    </>
  );
};

// Helper to extract form values from initial data
const getFormValues = (initialData: Todo) => ({
  title: initialData.title,
  description: initialData.description,
  category_id: initialData.category.id,
  priority: initialData.priority,
  completed: initialData.completed,
  due_date: dayjs(initialData.due_date),
  created_at: dayjs(initialData.created_at),
});

export default DrawerForm;
