import { Input } from "antd";

interface SearchBarProps {
  onSearch?: (value: string) => void;
}

export const SearchBar = ({
  onSearch,
}: SearchBarProps & { onChange?: (value: string) => void }) => (
  <div className="filter-bar">
    <Input.Search
      placeholder="Cari Data..."
      allowClear
      style={{ maxWidth: 400 }}
      onSearch={onSearch}
    />
  </div>
);
