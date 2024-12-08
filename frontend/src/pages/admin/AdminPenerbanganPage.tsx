import {
  useQuery,
  keepPreviousData,
  useMutation,
  useQueryClient,
} from "@tanstack/react-query";
import axios from "axios";
import { ApiResponse, ApiURL, PaginatedApiResponse } from "../../const/api";
import AdminLayout from "./component/AdminLayout";
import { ColumnDef } from "@tanstack/react-table";
import ServerTable from "../../components/table/ServerTable";
import useServerTable, {
  buildPaginatedTableURL,
} from "../../components/table/useServerTable";
import Modal from "../../components/Modal";
import { useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import Input from "../../components/Input";
import SelectInput from "../../components/SelectInput";
import { useCookies } from "react-cookie";

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

type PenerbanganFormProps = {
  no_penerbangan: string;
  jadwal_berangkat: Date;
  jadwal_datang: Date;
  harga: number;
  kapasitas: number;
  bandara_berangkat: string;
  bandara_datang: string;
  maskapai: string;
};

export default function AdminPenerbangan() {
  const [isTambahOpen, setIsTambahOpen] = useState(false);
  const queryClient = useQueryClient();
  const [cookie] = useCookies(["user"]);

  const { tableState, setTableState } = useServerTable({
    pageSize: 10,
  });

  const url = buildPaginatedTableURL({
    baseUrl: `/admin/penerbangan`,
    tableState,
    option: {
      arrayFormat: "none",
    },
  });

  const { data } = useQuery<PaginatedApiResponse<PenerbanganProps>>({
    queryKey: ["penerbangan", url],
    queryFn: async () => {
      const res = await axios.get(ApiURL + url);
      return res.data;
    },
    placeholderData: keepPreviousData,
  });

  const { data: DataMaskapai } = useQuery<ApiResponse<AirlineProps>>({
    queryKey: ["maskapai"],
    queryFn: async () => {
      const res = await axios.get(ApiURL + "/admin/maskapai");
      return res.data;
    },
  });

  const { data: DataBandara } = useQuery<ApiResponse<AirportProps>>({
    queryKey: ["bandara"],
    queryFn: async () => {
      const res = await axios.get(ApiURL + "/admin/bandara");
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

  const methods = useForm<PenerbanganFormProps>();

  const { handleSubmit, watch } = methods;

  const { mutate } = useMutation({
    mutationFn: async (data: PenerbanganFormProps) => {
      await axios.post(ApiURL + "/admin/penerbangan", data, {
        headers: {
          Authorization: "Bearer " + cookie.user.token,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["penerbangan"],
      });
      setIsTambahOpen(false);
    },
  });

  const onSubmit = (data: PenerbanganFormProps) => {
    const payload: PenerbanganFormProps = {
      ...data,
      harga: Number(data.harga),
      kapasitas: Number(data.kapasitas),
      jadwal_berangkat: new Date(data.jadwal_berangkat),
      jadwal_datang: new Date(data.jadwal_datang),
    };
    mutate(payload);
  };

  const BandaraBerangkatValue = watch("bandara_berangkat");
  const BandaraDatangValue = watch("bandara_datang");

  return (
    <AdminLayout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Daftar Penerbangan</h1>
      <Modal
        buttonText="Tambah Penerbangan"
        isOpen={isTambahOpen}
        setIsOpen={setIsTambahOpen}
        title="Tambah Penerbangan"
      >
        <FormProvider {...methods}>
          <form
            className="flex flex-col gap-3"
            onSubmit={handleSubmit(onSubmit)}
          >
            <Input
              id="no_penerbangan"
              label="No Penerbangan"
              placeholder="Masukan no penerbangan"
              className="w-full"
              validation={{
                required: "No penerbangan harus diisi",
              }}
            />
            <Input
              id="jadwal_berangkat"
              label="Jadwal Keberangkatan"
              type="datetime-local"
              className="w-full"
              validation={{
                required: "Jadwal keberangkatan harus diisi",
              }}
            />
            <Input
              id="jadwal_datang"
              label="Jadwal Kedatangan"
              type="datetime-local"
              className="w-full"
              validation={{
                required: "Jadwal kedatangan harus diisi",
              }}
            />
            <Input
              id="harga"
              label="Harga"
              placeholder="Masukan harga penerbangan"
              type="number"
              className="w-full"
              validation={{
                required: "Harga harus diisi",
              }}
            />
            <Input
              id="kapasitas"
              label="Kapasitas"
              placeholder="Masukan kapasitas penerbangan"
              type="number"
              className="w-full"
              validation={{
                required: "Kapasitas harus diisi",
              }}
            />

            <SelectInput
              id="bandara_berangkat"
              label="Bandara Keberangkatan"
              placeholder="Pilih bandara keberangkatan"
              validation={{
                required: "Bandara keberangkatan harus diisi",
              }}
            >
              {DataBandara?.data.map((bandara) => (
                <option
                  key={bandara.id}
                  value={bandara.id}
                  disabled={bandara.id === BandaraDatangValue}
                >
                  {bandara.name}
                </option>
              ))}
            </SelectInput>

            <SelectInput
              id="bandara_datang"
              label="Bandara Kedatangan"
              placeholder="Pilih bandara kedatangan"
              validation={{
                required: "Bandara kedatangan harus diisi",
              }}
            >
              {DataBandara?.data.map((bandara) => (
                <option
                  key={bandara.id}
                  value={bandara.id}
                  disabled={bandara.id === BandaraBerangkatValue}
                >
                  {bandara.name}
                </option>
              ))}
            </SelectInput>

            <SelectInput
              id="maskapai"
              label="Maskapai"
              placeholder="Pilih maskapai"
              validation={{
                required: "Maskapai harus diisi",
              }}
            >
              {DataMaskapai?.data.map((maskapai) => (
                <option key={maskapai.id} value={maskapai.id}>
                  {maskapai.name}
                </option>
              ))}
            </SelectInput>
            <button
              className="border px-3 py-1.5 rounded-lg bg-blue-400 text-white"
              type="submit"
            >
              Submit
            </button>
          </form>
        </FormProvider>
      </Modal>
      <div className="w-full px-20 overflow-hidden">
        <ServerTable
          data={data?.data ?? []}
          columns={columns}
          metadata={data?.meta}
          tableState={tableState}
          setTableState={setTableState}
          withFilter={false}
          className="w-full"
        />
      </div>
    </AdminLayout>
  );
}
