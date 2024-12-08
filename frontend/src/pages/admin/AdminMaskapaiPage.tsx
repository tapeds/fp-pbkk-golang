import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { ApiResponse, ApiURL } from "../../const/api";
import AdminLayout from "./component/AdminLayout";
import Table from "../../components/table/Table";
import { ColumnDef } from "@tanstack/react-table";
import { useState } from "react";
import { useCookies } from "react-cookie";
import { FormProvider, useForm } from "react-hook-form";
import Modal from "../../components/Modal";
import Input from "../../components/Input";
import DeleteModal from "../../components/DeleteModal";
import EditModal from "../../components/EditModal";

type AirlineProps = {
  id: string;
  name: string;
  image: string;
};

type MaskapaiFormProps = {
  name: string;
  image: string;
};

export default function AdminMaskapai() {
  const [isTambahOpen, setIsTambahOpen] = useState(false);
  const queryClient = useQueryClient();
  const [cookie] = useCookies(["user"]);

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
    {
      accessorKey: "action",
      enableSorting: false,
      header: " ",
      cell(row) {
        const onDelete = () => {
          DeleteMutation(row.row.original.id);
        };

        const onSubmit = (data: MaskapaiFormProps) => {
          const payload: MaskapaiFormProps & {
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
              title={`Edit Maskapai ${row.row.original.name}`}
              data={{ ...row.row.original }}
              onSubmit={onSubmit}
            >
              <Input
                id="name"
                label="Nama Maskapai"
                placeholder="Masukan nama maskapai"
                className="w-full"
                validation={{
                  required: "Nama maskapai harus diisi",
                }}
              />
              <Input
                id="image"
                label="Link Gambar Maskapai"
                className="w-full"
                placeholder="Masukan link gambar maskapai"
                validation={{
                  required: "Link gambar maskapai harus diisi",
                }}
              />
            </EditModal>
            <DeleteModal
              title={`Apakah anda yakin untuk menghapus maskapai ${row.row.original.name}`}
              onPositive={onDelete}
            />
          </div>
        );
      },
    },
  ];

  const methods = useForm<MaskapaiFormProps>();

  const { handleSubmit } = methods;

  const { mutate: DeleteMutation } = useMutation({
    mutationFn: async (id: string) => {
      await axios.delete(ApiURL + "/admin/maskapai/" + id, {
        headers: {
          Authorization: "Bearer " + cookie.user.token,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["maskapai"],
      });
    },
  });

  const { mutate: EditMutation } = useMutation({
    mutationFn: async (
      data: MaskapaiFormProps & {
        id: string;
      },
    ) => {
      await axios.patch(ApiURL + "/admin/maskapai", data, {
        headers: {
          Authorization: "Bearer " + cookie.user.token,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["maskapai"],
      });
    },
  });

  const { mutate } = useMutation({
    mutationFn: async (data: MaskapaiFormProps) => {
      await axios.post(ApiURL + "/admin/maskapai", data, {
        headers: {
          Authorization: "Bearer " + cookie.user.token,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["maskapai"],
      });
      setIsTambahOpen(false);
    },
  });

  const onSubmit = (data: MaskapaiFormProps) => {
    mutate(data);
  };

  return (
    <AdminLayout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Daftar Maskapai</h1>

      <Modal
        buttonText="Tambah Maskapai"
        isOpen={isTambahOpen}
        setIsOpen={setIsTambahOpen}
        title="Tambah Maskapai"
      >
        <FormProvider {...methods}>
          <form
            className="flex flex-col gap-3"
            onSubmit={handleSubmit(onSubmit)}
          >
            <Input
              id="name"
              label="Nama Maskapai"
              placeholder="Masukan nama maskapai"
              className="w-full"
              validation={{
                required: "Nama maskapai harus diisi",
              }}
            />
            <Input
              id="image"
              label="Link Gambar Maskapai"
              className="w-full"
              placeholder="Masukan link gambar maskapai"
              validation={{
                required: "Link gambar maskapai harus diisi",
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
