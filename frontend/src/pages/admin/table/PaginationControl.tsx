import { RowData, Table } from "@tanstack/react-table";
import * as React from "react";
import { HiChevronLeft, HiChevronRight } from "react-icons/hi";

import clsx from "clsx";

type PaginationState = {
  page: number;
  size: number;
};

type PaginationControlProps<T extends RowData> = {
  data: T[];
  table: Table<T>;
  setParams: React.Dispatch<React.SetStateAction<PaginationState>>;
} & React.ComponentPropsWithoutRef<"div">;

/**
 *
 * @see https://javascript.plainenglish.io/create-a-pagination-in-a-react-way-df5c6fe1e0c7
 */
export default function PaginationControl<T extends RowData>({
  className,
  data,
  table,
  setParams,
  ...rest
}: PaginationControlProps<T>) {
  const currentPage = table.getState().pagination.pageIndex + 1;
  const pageCount = table.getPageCount();
  const paginationControl = buildPaginationControl(currentPage, pageCount);

  const handlePageControlClick = (page: string | number) => {
    if (page !== "...") {
      table.setPageIndex((page as number) - 1);
    }
  };

  return (
    <div
      className={clsx(
        "flex items-center justify-between gap-x-2 md:justify-end",
        className,
      )}
      {...rest}
    >
      <div className="flex items-center gap-1">
        <button
          className={clsx(
            "flex min-w-[38px] justify-center rounded-md !border-none bg-transparent !p-2 !font-semibold text-primary-info-dark drop-shadow-sm",
            "disabled:cursor-not-allowed",
          )}
          disabled={!table.getCanPreviousPage()}
          onClick={() => {
            setParams((params) => ({
              ...params,
              page: Number(params.page) - 1,
            }));
            table.previousPage();
          }}
        >
          <HiChevronLeft size={20} />
        </button>
        {paginationControl.map((pageIndex, index) => (
          <button
            key={index}
            className={clsx(
              "flex min-w-[38px] justify-center rounded-md border-2 border-primary-info-dark bg-white !p-2 !font-semibold text-primary-info-dark drop-shadow-sm",
              currentPage === pageIndex &&
                "bg-primary-info-hover text-typo-white",
            )}
            onClick={() => {
              setParams((params) => ({
                ...params,
                page: pageIndex as number,
              }));
              handlePageControlClick(pageIndex);
            }}
          >
            {pageIndex}
          </button>
        ))}
        <button
          color="basic"
          className={clsx(
            "flex min-w-[38px] justify-center rounded-md !border-none bg-transparent !p-2 !font-semibold text-primary-info-dark drop-shadow-sm ",
            "disabled:cursor-not-allowed",
          )}
          disabled={
            !table.getCanNextPage() ||
            data.length < table.getState().pagination.pageSize
          }
          onClick={() => {
            setParams((params) => ({
              ...params,
              page: Number(params.page) + 1,
            }));
            table.nextPage();
          }}
        >
          <HiChevronRight size={20} />
        </button>
      </div>
    </div>
  );
}

export function buildPaginationControl(
  currentPage: number,
  pageCount: number,
  delta = 1,
) {
  const rangeWithDots: (number | string)[] = [];

  const range = [...Array(pageCount)]
    .map((_, i) => i + 1)
    .map((page) => {
      if (
        Math.abs(page - 1) <= delta ||
        Math.abs(pageCount - page) <= delta ||
        Math.abs(currentPage - page) <= delta
      )
        return page;

      return -1;
    })
    .filter((page) => page !== -1);

  range.forEach((page, i) => {
    const previousPage = range[i - 1];
    if (page - previousPage > 1) rangeWithDots.push("...");
    rangeWithDots.push(page);
  });

  return rangeWithDots;
}
