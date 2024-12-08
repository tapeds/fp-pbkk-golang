export const ApiURL = "http://127.0.0.1:8888/api";

export type PaginatedApiResponse<T> = {
  status: boolean;
  message: string;
  data: T[];
  meta: {
    page: number;
    per_page: number;
    max_page: number;
    count: number;
  };
};

export type ApiResponse<T> = {
  status: boolean;
  message: string;
  data: T[];
};
