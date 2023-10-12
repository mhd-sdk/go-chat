import React from 'react';
import { css } from '@emotion/css';

interface PixelProps {
  width: number;
  height: number;
  pixelColors: {x:number,y:number,color:string}[][];
  onDraw: (x: number, y: number) => void;
  defaultColor?: string;
}

export const PixelsPane: React.FC<PixelProps> = ({
  width,
  height,
  pixelColors,
  onDraw,
  defaultColor = 'white',
}) => {
  return (
      <div
        className={styles.pane(width, height)}
      >
        {pixelColors.map((row, rowIndex) =>
          row.map((col, colIndex) => (
            <div
              key={`${rowIndex}-${colIndex}`}
              className={styles.pixel(col.color ?? defaultColor)}
              onClick={() => onDraw(rowIndex, colIndex)}
            />
          ))
        )}
      </div>
  );
};

const styles = {
  pixel: (color?:string) => css`
    width: 1em;
    height: $1em;
    background-color: ${color};
    border-top: 0.5px solid #E0E0E0;
    border-left: 0.5px solid #E0E0E0;
    // hover effet
    &:hover {
      background-color: #E0E0E0;
    }
  `,
  pane:(width:number, height:number)=> css`
    display: grid;
    grid-template-columns: repeat(${width}, 1em);
    grid-template-rows: repeat(${height}, 1em);
    margin: 0;
    border-right: 1px solid #E0E0E0;
    border-bottom: 0.5px solid #E0E0E0;
  `,
}

