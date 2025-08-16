
import type React from "react";

interface InputText {
  onTextChange: (val: string) => void;
  value: string;
}

export const SearchBar: React.FC<InputText> = ({ onTextChange, value }) => {
  return (
    <input
    type="text"
      value={value}
      placeholder="Enter a username"
      onChange={(e) => onTextChange(e.target.value)}
      className="border p-2 rounded w-full"
    />
  );
};
