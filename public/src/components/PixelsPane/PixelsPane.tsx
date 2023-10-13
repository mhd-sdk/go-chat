import React from 'react';
import { css } from '@emotion/css';
import { ColorInput } from '@mantine/core';

interface PixelProps {
  width: number;
  height: number;
  pixelColors: { x: number, y: number, color: string }[][];
  onDraw: (x: number, y: number) => void;
  onColorChange: (color: string) => void;
  defaultColor?: string;
  disabled?: boolean;
}

export const PixelsPane: React.FC<PixelProps> = ({
  width,
  height,
  pixelColors,
  onDraw,
  onColorChange,
  defaultColor = 'white',
  disabled = false,
}) => {
  return (
    <>
      <div
        className={styles.pane(width, height)}
      >
        {pixelColors.map((row, rowIndex) =>
          row.map((col, colIndex) => (
            <div
              key={`${rowIndex}-${colIndex}`}
              className={styles.pixel(col.color ?? defaultColor)}
              onClick={() => !disabled && onDraw(rowIndex, colIndex)}
            />
          ))
        )}
      </div>
      <div className={styles.width(100)}>
        <ColorInput
          radius="xl"
          label="Input label"
          description="Input description"
          placeholder="Input placeholder"
          disabled={disabled}
          onChange={(color) => onColorChange(color)}
        />
      </div>
    </>
  );
};

const styles = {
  pixel: (color?: string) => css`
    width: 1em;
    height: $1em;
    background-color: ${color};
    border-top: 0.5px solid #E0E0E0;
    border-left: 0.5px solid #E0E0E0;
    transition: transform 0.2s ease-in-out;
    &:hover {
      transform: scale(1.5);
    }
  `,
  width: (width: number) => css`
    width: ${width}%;

  `,
  pane: (width: number, height: number) => css`
    display: grid;
    grid-template-columns: repeat(${width}, 1em);
    grid-template-rows: repeat(${height}, 1em);
    margin: 0;
    border-right: 1px solid #E0E0E0;
    border-bottom: 0.5px solid #E0E0E0;
  `,
}

