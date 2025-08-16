import { useState, useMemo, useEffect } from "react";
import "./App.css";
import { SearchBar } from "./components/search";
import { BinaryGrid } from "./components/BinaryGrid";

const hashFunctions = [
  (str: string, size: number) =>
    str.split("").reduce((acc, c) => acc + c.charCodeAt(0), 0) % size,
  (str: string, size: number) =>
    str.split("").reduce((acc, c) => acc + c.charCodeAt(0) * 7, 0) % size,
  (str: string, size: number) =>
    str.split("").reduce((acc, c) => acc ^ c.charCodeAt(0) * 13, 0) % size,
];

export default function App() {
  const size = 32;
  const [text, setText] = useState("");
  const [bitArray, setBitArray] = useState<number[]>(Array(size).fill(0));
  const [savedUsers, setSavedUsers] = useState<string[]>([]);
  const [showUsers, setShowUsers] = useState(false);

  const [flashIndices, setFlashIndices] = useState<number[]>([]);
  const [setAnimIndices, setSetAnimIndices] = useState<number[]>([]);

  const highlightedIndexes = useMemo(() => {
    if (!text) return [];
    return hashFunctions.map((fn) => fn(text, size));
  }, [text]);

  useEffect(() => {
    if (!text) {
      setFlashIndices([]);
      return;
    }
    setFlashIndices(highlightedIndexes);
    const t = setTimeout(() => setFlashIndices([]), 350);
    return () => clearTimeout(t);
  }, [text, highlightedIndexes]);


  const canSave = useMemo(() => {
    if (!text) return false;
    return highlightedIndexes.some((i) => bitArray[i] === 0);
  }, [highlightedIndexes, bitArray, text]);

  const handleSave = () => {
    if (!canSave) return;

    
    setSetAnimIndices(highlightedIndexes);

    setTimeout(() => {
      const newBits = [...bitArray];
      highlightedIndexes.forEach((i) => (newBits[i] = 1));
      setBitArray(newBits);
      setSavedUsers((prev) => [...prev, text]);
      setText("");
      setSetAnimIndices([]);
    }, 600);
  };

  return (
    <div className="p-6 max-w-lg mx-auto space-y-6">
      {text && !canSave && (
        <p className="text-red-600 font-medium">
          Username not available (probably already taken)
        </p>
      )}
      <SearchBar onTextChange={setText} value={text} />

      <BinaryGrid
        bitArray={bitArray}
        highlightedIndexes={highlightedIndexes}
        flashIndexes={flashIndices}
        setAnimIndexes={setAnimIndices}
      />

      <button
        className={`px-4 py-2 rounded ${
          canSave ? "bg-blue-500 text-white" : "bg-gray-400 text-gray-700"
        }`}
        onClick={handleSave}
        disabled={!canSave}
      >
        Save Username
      </button>

      <button
        className="ml-4 px-4 py-2 rounded bg-green-500 text-white"
        onClick={() => setShowUsers((s) => !s)}
      >
        {showUsers ? "Hide Users" : "Show Users"}
      </button>

      {showUsers && (
        <ul className="mt-4 border p-2 rounded">
          {savedUsers.length === 0 && <li>No users saved yet.</li>}
          {savedUsers.map((u, i) => (
            <li key={i}>{u}</li>
          ))}
        </ul>
      )}
    </div>
  );
}
