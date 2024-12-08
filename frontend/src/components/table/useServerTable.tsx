import {
  // FiltersTableState,
  PaginationState,
  SortingState,
} from "@tanstack/react-table";
import * as React from "react";

import queryString, { StringifyOptions } from "query-string";

type useServerTableProps<T extends object> = {
  pageSize?: number;
  sort?: {
    key: Extract<keyof T, string>;
    type: "asc" | "desc";
  };
  // find?: FiltersTableState;
};

export type ServerTableState = {
  pagination: PaginationState;
  sorting: SortingState;
  globalFilter: string;
};

export default function useServerTable<T extends object>({
  pageSize = 10,
  sort,
}: useServerTableProps<T>) {
  const [sorting, setSorting] = React.useState<SortingState>(
    sort
      ? [
          {
            id: sort.key,
            desc: sort.type === "desc",
          },
        ]
      : [],
  );

  const [pagination, setPagination] = React.useState<PaginationState>({
    pageIndex: 0,
    pageSize,
  });

  const [globalFilter, setGlobalFilter] = React.useState("");

  return {
    tableState: {
      pagination,
      sorting,
      globalFilter,
    },
    setTableState: {
      setPagination,
      setSorting,
      setGlobalFilter,
    },
  };
}

type BuildPaginationTableParam = {
  /** API Base URL, with / on the front */
  baseUrl: string;
  tableState: ServerTableState;
  /** Parameter addition
   * @example ['include=user,officer']
   */
  option?: StringifyOptions;
  additionalParam?: Record<string, unknown>;
};
type BuildPaginationTableURL = (props: BuildPaginationTableParam) => string;

export const buildPaginatedTableURL: BuildPaginationTableURL = ({
  baseUrl,
  tableState,
  option,
  additionalParam,
}) => {
  const queryParams = queryString.stringify(
    {
      per_page: tableState.pagination.pageSize,
      page: tableState.pagination.pageIndex + 1,
      ...additionalParam,
    },
    {
      arrayFormat: "bracket",
      skipEmptyString: true,
      ...option,
    },
  );

  return `${baseUrl}?${queryParams}`;
};
