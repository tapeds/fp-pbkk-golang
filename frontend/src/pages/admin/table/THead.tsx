"use client";
import { flexRender, RowData, Table } from "@tanstack/react-table";
import * as React from "react";
import { AiFillCaretDown } from "react-icons/ai";

import clsx from "clsx";

type THeadProps<T extends RowData> = {
  omitSort: boolean;
  table: Table<T>;
} & React.ComponentPropsWithoutRef<"div">;

export default function THead<T extends RowData>({
  className,
  omitSort,
  table,
  ...rest
}: THeadProps<T>) {
  return (
    <thead
      className={clsx(
        "border-b-2 border-typo-inline bg-blue-500 text-white",
        className,
      )}
      {...rest}
    >
      {table.getHeaderGroups().map((headerGroup) => (
        <tr key={headerGroup.id} className="divide-x-2 divide-typo-inline">
          {headerGroup.headers.map((header) => (
            <th
              key={header.id}
              scope="col"
              className={clsx(
                "group py-1 pr-3 text-left text-sm font-semibold sm:text-base",
                !omitSort && header.column.getCanSort() ? "pl-4" : "pl-4",
              )}
            >
              {header.isPlaceholder ? null : (
                <div
                  className={clsx(
                    "relative flex items-center justify-center gap-2 py-1",
                    !omitSort && header.column.getCanSort()
                      ? "cursor-pointer select-none"
                      : "",
                  )}
                  onClick={
                    omitSort
                      ? () => null
                      : header.column.getToggleSortingHandler()
                  }
                >
                  <p className="text-white">
                    {flexRender(
                      header.column.columnDef.header,
                      header.getContext(),
                    )}
                  </p>
                </div>
              )}
            </th>
          ))}
        </tr>
      ))}
    </thead>
  );
}
