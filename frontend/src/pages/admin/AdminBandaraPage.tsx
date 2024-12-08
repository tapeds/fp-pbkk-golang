import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { ApiResponse, ApiURL } from "../../const/api";
import AdminLayout from "./component/AdminLayout";
import Table from "../../components/table/Table";
import { ColumnDef } from "@tanstack/react-table";

type AirportProps = {
  id: string;
  name: string;
  kode: string;
  kota: string;
};

export default function AdminBandara() {
  const { data } = useQuery<ApiResponse<AirportProps>>({
    queryKey: ["bandara"],
    queryFn: async () => {
      const res = await axios.get(ApiURL + "/admin/bandara");
      return res.data;
    },
  });

  const columns: ColumnDef<AirportProps>[] = [
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
    },
    {
      accessorKey: "kode",
      header: "Kode",
    },
    {
      accessorKey: "kota",
      header: "Kota",
    },
  ];

  return (
    <AdminLayout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Daftar Bandara</h1>
      <div className="w-full px-20 overflow-hidden">
        <Table data={data?.data ?? []} columns={columns} className="w-full" />
      </div>
    </AdminLayout>
  );
}
