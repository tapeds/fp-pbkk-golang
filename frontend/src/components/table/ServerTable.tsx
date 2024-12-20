"use client";

import {
  ColumnDef,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  PaginationState,
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
import { PaginatedApiResponse } from "../../const/api";

type ServerTableState = {
  pagination: PaginationState;
  sorting: SortingState;
  globalFilter: string;
};

type SetServerTableState = {
  setPagination: React.Dispatch<React.SetStateAction<PaginationState>>;
  setSorting: React.Dispatch<React.SetStateAction<SortingState>>;
  setGlobalFilter: React.Dispatch<React.SetStateAction<string>>;
};

type ServerTableProps<T extends object> = {
  columns: ColumnDef<T>[];
  data: T[];
  header?: React.ReactNode;
  isLoading?: boolean;
  metadata: PaginatedApiResponse<T>["meta"] | undefined;
  tableState: ServerTableState;
  setTableState: SetServerTableState;
  omitSort?: boolean;
  withFilter?: boolean;
  placeholderFilter?: string;
} & React.ComponentPropsWithoutRef<"div">;

export default function ServerTable<T extends object>({
  className,
  columns,
  data,
  header: Header,
  metadata,
  isLoading,
  tableState,
  setTableState,
  omitSort = false,
  withFilter = false,
  placeholderFilter = "Cari Data",
  ...rest
}: ServerTableProps<T>) {
  const columnResizeMode = "onEnd";

  const table = useReactTable({
    data,
    columns,
    columnResizeMode,
    pageCount: metadata?.max_page,
    state: {
      ...tableState,
    },
    onGlobalFilterChange: setTableState.setGlobalFilter,
    onPaginationChange: setTableState.setPagination,
    onSortingChange: setTableState.setSorting,
    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    manualPagination: true,
    manualSorting: true,
    manualFiltering: true,
  });

  return (
    <div className={clsx("flex flex-col", className)} {...rest}>
      <div
        className={clsx(
          "flex flex-col flex-wrap items-stretch gap-3",
          "md:mx-3 md:flex-row md:flex-nowrap md:items-end md:justify-between",
        )}
      >
        {withFilter && (
          <Filter
            placeholder={placeholderFilter}
            table={table}
            className="mx-auto w-full md:mx-0 md:w-auto"
          />
        )}
        <div className="mx-auto flex w-full items-center gap-3 max-[350px]:flex-col md:mx-0 md:w-auto">
          {Header}
          <TOption
            value={table.getState().pagination.pageSize}
            onChange={(e) => {
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
        </div>
      </div>
      <div className="-my-2 mx-0 mt-2 overflow-x-auto md:-mx-6 lg:-mx-8">
        <div className="inline-block min-w-full py-2 align-middle max-md:w-[200vw] md:px-6 lg:px-8">
          <div className="overflow-hidden rounded-2xl shadow ring-1 ring-black ring-opacity-5">
            <table className="min-w-full divide-y divide-gray-300">
              <colgroup>
                {table.getAllColumns().map((column) => (
                  <col
                    key={column.id}
                    span={1}
                    style={{
                      width: column.columnDef.size
                        ? column.columnDef.size
                        : "auto",
                    }}
                    className="border-[1px]"
                  />
                ))}
              </colgroup>
              <THead table={table} omitSort={omitSort} />
              <TBody
                table={table}
                isLoading={isLoading}
                className="text-center"
              />
            </table>
          </div>
        </div>
      </div>

      <PaginationControl table={table} data={data} className="mt-4" />
    </div>
  );
}
