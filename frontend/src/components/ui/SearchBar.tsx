import { Input } from "antd";

interface SearchBarProps {
  onSearch?: (value: string) => void;
}

export const SearchBar = ({
  onSearch,
}: SearchBarProps & { onChange?: (value: string) => void }) => (
  <div className="filter-bar search-bar-container">
    <Input.Search
      placeholder="Cari Data..."
      allowClear
      className="search-input-box"
      onSearch={onSearch}
    />
  </div>
);
