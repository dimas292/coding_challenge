import { Button } from "antd";
import { PlusCircleOutlined } from "@ant-design/icons";

interface TodoHeaderProps {
  onAddTodo: () => void;
}

export const TodoHeader = ({ onAddTodo }: TodoHeaderProps) => (
  <div className="todo-header-container">
    <div className="page-title-section">
      <h2 className="page-title">Todo List</h2>
    </div>
    <div className="action-bar">
      <div className="action-left" />
      <div className="action-right">
        <Button
          type="primary"
          className="rounded-button"
          icon={<PlusCircleOutlined />}
          onClick={onAddTodo}
        >
          Add Todo
        </Button>
      </div>
    </div>
  </div>
);
