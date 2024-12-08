import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { ApiResponse, ApiURL } from "../../const/api";
import AdminLayout from "./component/AdminLayout";
import Table from "../../components/table/Table";
import { ColumnDef } from "@tanstack/react-table";

type AirlineProps = {
  id: string;
  name: string;
  image: string;
};

export default function AdminMaskapai() {
  const { data } = useQuery<ApiResponse<AirlineProps>>({
    queryKey: ["maskapai"],
    queryFn: async () => {
      const res = await axios.get(ApiURL + "/admin/maskapai");
      return res.data;
    },
  });

  const columns: ColumnDef<AirlineProps>[] = [
    {
      accessorKey: "no",
      header: "No",
      size: 50,
      accessorFn: (_, index) => index + 1,
      enableSorting: false,
      enableColumnFilter: false,
      enableGlobalFilter: false,
    },
    {
      accessorKey: "name",
      header: "Nama",
      cell: (row) => (
        <div className="flex flex-row gap-5 items-center justify-center">
          <figure className="w-32 rounded-full">
            <img src={row.row.original.image} className="flex-none" />
          </figure>
          <p>{row.getValue() as string}</p>
        </div>
      ),
    },
  ];

  return (
    <AdminLayout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Daftar Maskapai</h1>
      <div className="w-full px-20 overflow-hidden">
        <Table data={data?.data ?? []} columns={columns} className="w-full" />
      </div>
    </AdminLayout>
  );
}
