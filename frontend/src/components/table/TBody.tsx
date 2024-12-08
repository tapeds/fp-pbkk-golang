"use client";

import { flexRender, RowData, Table } from "@tanstack/react-table";
import * as React from "react";

import clsx from "clsx";

type TBodyProps<T extends RowData> = {
  table: Table<T>;
  isLoading?: boolean;
} & React.ComponentPropsWithoutRef<"div">;

export default function TBody<T extends RowData>({
  className,
  isLoading,
  table,
  ...rest
}: TBodyProps<T>) {
  return (
    <tbody
      className={clsx("divide-y-2 divide-typo-inline bg-white", className)}
      {...rest}
    >
      {isLoading ? (
        <tr>
          <td
            className="col-span-full truncate whitespace-nowrap px-3 py-3 text-center text-typo-main"
            colSpan={table.getAllColumns().length}
          >
            Loading...
          </td>
        </tr>
      ) : table.getRowModel().rows.length == 0 ? (
        <tr>
          <td
            className="col-span-full truncate whitespace-nowrap px-3 py-3 text-center text-typo-main"
            colSpan={table.getAllColumns().length}
          >
            No Data
          </td>
        </tr>
      ) : (
        table.getRowModel().rows.map((row) => (
          <tr
            key={row.id}
            className={clsx("divide divide-x-2 divide-typo-inline bg-white")}
          >
            {row.getVisibleCells().map((cell) => {
              return (
                <td
                  key={cell.id}
                  className="truncate whitespace-nowrap px-3 py-2 text-center"
                  style={{ maxWidth: cell.column.getSize() }}
                >
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </td>
              );
            })}
          </tr>
        ))
      )}
    </tbody>
  );
}
