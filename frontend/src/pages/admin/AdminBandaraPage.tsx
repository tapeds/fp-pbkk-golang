import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import axios, { AxiosError } from "axios";
import { ApiResponse, ApiURL } from "../../const/api";
import AdminLayout from "./component/AdminLayout";
import Table from "../../components/table/Table";
import { ColumnDef } from "@tanstack/react-table";
import Modal from "../../components/Modal";
import { useCookies } from "react-cookie";
import { useState } from "react";
import Input from "../../components/Input";
import { FormProvider, useForm } from "react-hook-form";
import DeleteModal from "../../components/DeleteModal";
import EditModal from "../../components/EditModal";
import toast from "react-hot-toast";

type AirportProps = {
  id: string;
  name: string;
  kode: string;
  kota: string;
};

type BandaraFormProps = {
  name: string;
  kode: string;
  kota: string;
};

export default function AdminBandara() {
  const [isTambahOpen, setIsTambahOpen] = useState(false);
  const queryClient = useQueryClient();
  const [cookie] = useCookies(["user"]);

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
    {
      accessorKey: "action",
      enableSorting: false,
      header: " ",
      cell(row) {
        const onDelete = () => {
          DeleteMutation(row.row.original.id);
        };
        const onSubmit = (data: BandaraFormProps) => {
          const payload: BandaraFormProps & {
            id: string;
          } = {
            id: row.row.original.id,
            ...data,
          };

          EditMutation(payload);
        };
        return (
          <div className="flex flex-row items-center gap-5">
            <EditModal
              title={`Edit Bandara ${row.row.original.name}`}
              data={{
                ...row.row.original,
              }}
              onSubmit={onSubmit}
            >
              <Input
                id="name"
                label="Nama Bandara"
                placeholder="Masukan nama bandara"
                className="w-full"
                validation={{
                  required: "Nama bandara harus diisi",
                }}
              />
              <Input
                id="kode"
                label="Kode Bandara"
                className="w-full"
                placeholder="Masukan kode bandara"
                validation={{
                  required: "Kode bandara harus diisi",
                }}
              />
              <Input
                id="kota"
                label="Kota Bandara"
                placeholder="Masukan kota bandara"
                className="w-full"
                validation={{
                  required: "Kota bandara harus diisi",
                }}
              />
            </EditModal>
            <DeleteModal
              title={`Apakah anda yakin untuk menghapus bandara ${row.row.original.name}`}
              onPositive={onDelete}
            />
          </div>
        );
      },
    },
  ];

  const methods = useForm<BandaraFormProps>();

  const { handleSubmit } = methods;

  const { mutate: DeleteMutation } = useMutation({
    mutationFn: async (id: string) => {
      await axios.delete(ApiURL + "/admin/bandara/" + id, {
        headers: {
          Authorization: "Bearer " + cookie.user.token,
        },
      });
    },
    onSuccess: () => {
      toast.success("Bandara berhasil dihapus");
      queryClient.invalidateQueries({
        queryKey: ["bandara"],
      });
    },
    onError: (err: AxiosError) => {
      toast.error((err.response?.data as { error: string }).error);
    },
  });

  const { mutate: EditMutation } = useMutation({
    mutationFn: async (
      data: BandaraFormProps & {
        id: string;
      },
    ) => {
      await axios.patch(ApiURL + "/admin/bandara", data, {
        headers: {
          Authorization: "Bearer " + cookie.user.token,
        },
      });
    },
    onSuccess: () => {
      toast.success("Bandara berhasil diedit");
      queryClient.invalidateQueries({
        queryKey: ["bandara"],
      });
    },
    onError: (err: AxiosError) => {
      toast.error((err.response?.data as { error: string }).error);
    },
  });

  const { mutate } = useMutation({
    mutationFn: async (data: BandaraFormProps) => {
      await axios.post(ApiURL + "/admin/bandara", data, {
        headers: {
          Authorization: "Bearer " + cookie.user.token,
        },
      });
    },
    onSuccess: () => {
      toast.success("Bandara berhasil ditambahkan");
      queryClient.invalidateQueries({
        queryKey: ["bandara"],
      });
      setIsTambahOpen(false);
    },
    onError: (err: AxiosError) => {
      toast.error((err.response?.data as { error: string }).error);
    },
  });

  const onSubmit = (data: BandaraFormProps) => {
    mutate(data);
  };

  return (
    <AdminLayout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Daftar Bandara</h1>

      <Modal
        buttonText="Tambah Bandara"
        isOpen={isTambahOpen}
        setIsOpen={setIsTambahOpen}
        title="Tambah Bandara"
      >
        <FormProvider {...methods}>
          <form
            className="flex flex-col gap-3"
            onSubmit={handleSubmit(onSubmit)}
          >
            <Input
              id="name"
              label="Nama Bandara"
              placeholder="Masukan nama bandara"
              className="w-full"
              validation={{
                required: "Nama bandara harus diisi",
              }}
            />
            <Input
              id="kode"
              label="Kode Bandara"
              className="w-full"
              placeholder="Masukan kode bandara"
              validation={{
                required: "Kode bandara harus diisi",
              }}
            />
            <Input
              id="kota"
              label="Kota Bandara"
              placeholder="Masukan kota bandara"
              className="w-full"
              validation={{
                required: "Kota bandara harus diisi",
              }}
            />
            <button
              className="border px-3 py-1.5 rounded-lg bg-blue-400 text-white"
              type="submit"
            >
              Submit
            </button>
          </form>
        </FormProvider>
      </Modal>
      <div className="w-full px-20 overflow-hidden mt-10">
        <Table data={data?.data ?? []} columns={columns} className="w-full" />
      </div>
    </AdminLayout>
  );
}
