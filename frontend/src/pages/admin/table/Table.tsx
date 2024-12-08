"use client";

import {
  ColumnDef,
  getCoreRowModel,
  getFacetedUniqueValues,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  SortingState,
  useReactTable,
} from "@tanstack/react-table";
import * as React from "react";

import Filter from "./Filter";
import PaginationControl from "./PaginationControl";
import TBody from "./TBody";
import THead from "./THead";
import TOption from "./TOption";
import clsx from "clsx";

type TableProps<T extends object> = {
  data: T[];
  columns: ColumnDef<T>[];
  isLoading?: boolean;
  omitSort?: boolean;
  withFilter?: boolean;
  withEntries?: boolean;
  withPaginationControl?: boolean;
} & React.ComponentPropsWithoutRef<"div">;

export default function Table<T extends object>({
  className,
  columns,
  data,
  isLoading,
  omitSort = false,
  withFilter = false,
  withEntries = false,
  withPaginationControl = false,
  ...rest
}: TableProps<T>) {
  const [globalFilter, setGlobalFilter] = React.useState("");
  const [sorting, setSorting] = React.useState<SortingState>([]);

  const table = useReactTable({
    data,
    columns,
    state: {
      globalFilter,
      sorting,
    },
    onGlobalFilterChange: setGlobalFilter,
    onSortingChange: setSorting,
    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFacetedUniqueValues: getFacetedUniqueValues(),
    getPaginationRowModel: getPaginationRowModel(),
  });

  return (
    <div className={clsx("flex flex-col", className)} {...rest}>
      <div className="flex items-start justify-between gap-5">
        {withFilter && <Filter table={table} />}
        {withEntries && (
          <TOption
            value={table.getState().pagination.pageSize}
            onChange={(e) => {
              setPage((prev) => ({ ...prev, size: Number(e) }));
              table.setPageSize(Number(e));
            }}
            title="Show"
            options={[
              { value: 10, label: "10 entries" },
              { value: 25, label: "25 entries" },
              { value: 50, label: "50 entries" },
              { value: 100, label: "100 entries" },
            ]}
          />
        )}
      </div>
      <div className="-my-2 mt-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div className="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
          <div className="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
            <table className="min-w-full divide-y divide-gray-300">
              <THead table={table} omitSort={omitSort} />
              <TBody table={table} isLoading={isLoading} />
            </table>
          </div>
        </div>
      </div>
      {withPaginationControl && (
        <PaginationControl
          table={table}
          data={data}
          setParams={setPage}
          className="mt-3"
        />
      )}
    </div>
  );
}
