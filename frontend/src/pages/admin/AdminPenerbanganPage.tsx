import { useQuery } from "@tanstack/react-query";
import axios from "axios";
import { ApiResponse, ApiURL } from "../../const/api";
import AdminLayout from "./component/AdminLayout";
import Table from "../../components/table/Table";
import { ColumnDef } from "@tanstack/react-table";

type PenerbanganProps = {
  id: string;
  no_penerbangan: string;
  jadwal_berangkat: Date;
  jadwal_datang: Date;
  harga: number;
  kapasitas: number;
  maskapai: AirlineProps;
  bandaras: AirportProps[];
};

type AirlineProps = {
  id: string;
  name: string;
  image: string;
};

type AirportProps = {
  id: string;
  name: string;
  kode: string;
  kota: string;
  arah: "BERANGKAT" | "DATANG";
};

export default function AdminPenerbangan() {
  const { data } = useQuery<ApiResponse<PenerbanganProps>>({
    queryKey: ["penerbangan"],
    queryFn: async () => {
      const res = await axios.get(ApiURL + "/admin/penerbangan");
      return res.data;
    },
  });

  const columns: ColumnDef<PenerbanganProps>[] = [
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
      accessorKey: "no_penerbangan",
      header: "No Penerbangan",
    },
    {
      accessorKey: "jadwal_berangkat",
      header: "Jadwal Keberangkatan",
      cell: (row) => {
        const date = new Date(row.row.original.jadwal_berangkat);
        return date.toLocaleString("id-ID", {
          day: "2-digit",
          month: "long",
          year: "numeric",
          hour: "2-digit",
          minute: "2-digit",
        });
      },
      size: 500,
    },
    {
      accessorKey: "jadwal_datang",
      header: "Jadwal Kedatangan",
      cell: (row) => {
        const date = new Date(row.row.original.jadwal_berangkat);
        return date.toLocaleString("id-ID", {
          day: "2-digit",
          month: "long",
          year: "numeric",
          hour: "2-digit",
          minute: "2-digit",
        });
      },
      size: 500,
    },
    {
      accessorKey: "harga",
      header: "Harga",
      cell: (row) => {
        return new Intl.NumberFormat("id-ID", {
          style: "currency",
          currency: "IDR",
        }).format(row.row.original.harga);
      },
      size: 400,
    },
    {
      header: "Bandara Keberangkatan",
      cell: (row) => {
        return row.row.original.bandaras.filter(
          (bandara) => bandara.arah === "BERANGKAT",
        )[0].name;
      },
      size: 500,
    },
    {
      header: "Bandara Kedatangan",
      cell: (row) => {
        return row.row.original.bandaras.filter(
          (bandara) => bandara.arah === "DATANG",
        )[0].name;
      },
      size: 500,
    },
    {
      accessorKey: "kapasitas",
      header: "Kapasitas",
    },
    {
      accessorKey: "maskapai.name",
      header: "Maskapai",
    },
  ];

  return (
    <AdminLayout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Daftar Penerbangan</h1>
      <div className="w-3/4 overflow-hidden">
        <Table data={data?.data ?? []} columns={columns} className="w-full" />
      </div>
    </AdminLayout>
  );
}
