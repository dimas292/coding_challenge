import { Select } from "antd";
import { useEffect, useState } from "react";
import { GetCategory } from "../../services/CategoryService";

interface FilterBarProps {
  onCategoryChange: (value: number | null) => void;
  onPriorityChange: (value: string) => void;
  onStatusChange: (value: string) => void;
}

const FilterBar = ({
  onCategoryChange,
  onPriorityChange,
  onStatusChange,
}: FilterBarProps) => {
  const [data, setData] = useState<{ id: number; name: string }[]>([]);
  useEffect(() => {
    // Fetch categories and priorities here if needed
    GetCategory().then((data) => {
      // Handle category data if needed
      console.log("Categories:", data.data);
      setData(data.data);
    });
  }, []);

  const handleCategoryChange = (value: number | undefined) => {
    onCategoryChange(value ?? null);
  };

  const handlePriorityChange = (value: string | undefined) => {
    onPriorityChange(value ?? "");
  };

  const handleStatusChange = (value: string | undefined) => {
    onStatusChange(value ?? "");
  };

  return (
    <div className="filter-bar">
      <Select
        placeholder="Select Priority"
        style={{ width: 200 }}
        allowClear
        onChange={handlePriorityChange}
      >
        <Select.Option value="high">High</Select.Option>
        <Select.Option value="medium">Medium</Select.Option>
        <Select.Option value="low">Low</Select.Option>
      </Select>
      <Select
        placeholder="Select Category"
        style={{ width: 200 }}
        allowClear
        onChange={handleCategoryChange}
      >
        {/* Add category options here */}
        {data.map((category) => (
          <Select.Option key={category.id} value={category.id}>
            {category.name}
          </Select.Option>
        ))}
      </Select>
      <Select
        placeholder="Select Status"
        style={{ width: 200 }}
        allowClear
        onChange={handleStatusChange}
      >
        <Select.Option value="false">Pending</Select.Option>
        <Select.Option value="true">Done</Select.Option>
      </Select>
      {/* Add filter options here */}
    </div>
  );
};

export default FilterBar;
