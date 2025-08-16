import type React from "react";

interface Props {
  bitArray: number[];
  highlightedIndexes: number[];
  flashIndexes: number[];    // NEW
  setAnimIndexes: number[];  // NEW
}

export const BinaryGrid: React.FC<Props> = ({
  bitArray,
  highlightedIndexes,
  flashIndexes,
  setAnimIndexes,
}) => {
  return (
    <div className="grid grid-cols-8 gap-2 mt-4">
      {bitArray.map((bit, i) => {
        const isHighlighted = highlightedIndexes.includes(i);
        const isFlashing = flashIndexes.includes(i);
        const isSetAnimating = setAnimIndexes.includes(i);

        // Base background derives from current bit state
        const base =
          bit === 1 ? "bg-green-500 text-white" : "bg-gray-200 text-gray-800";

        // Priority of visual states:
        // 1) set-anim (during save)
        // 2) flash (on typing)
        // 3) static highlight ring (candidate while typing)
        const animClass = isSetAnimating
          ? "set-anim"
          : isFlashing
          ? "flash"
          : isHighlighted
          ? "ring-2 ring-yellow-400"
          : "";

        return (
          <div
            key={i}
            className={`w-8 h-8 flex items-center justify-center rounded ${base} ${animClass}`}
            title={`index ${i}`}
          >
            {bit}
          </div>
        );
      })}
    </div>
  );
};

export default BinaryGrid;
