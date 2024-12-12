import { FormProvider, useForm } from "react-hook-form";
import Layout from "../../components/Layout";
import Input from "../../components/Input";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { ApiURL } from "../../const/api";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";

type RegisterFormProps = {
    name: string;
    telp_number: string;
    email: string;
    password: string;
    confirmPassword: string;
};

export default function Register() {
  const methods = useForm<RegisterFormProps>();
  const navigate = useNavigate();

  const { handleSubmit, watch } = methods;

  const { mutate } = useMutation({
    mutationFn: async (data: RegisterFormProps) => {
      await axios.post(ApiURL + "/user/register", {
        name: data.name,
        telp_number: data.telp_number,
        email: data.email,
        password: data.password,
      });
    },
    onSuccess: () => {
      navigate("/");
    },
  });

  const onSubmit = (data: RegisterFormProps) => {
    mutate(data);
  };

  return (
    <Layout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Register</h1>

      <FormProvider {...methods}>
        <form
          className="flex flex-col gap-3 w-1/2"
          onSubmit={handleSubmit(onSubmit)}
        >
          <Input
            id="name"
            label="Nama"
            placeholder="Masukkan nama"
            className="w-full"
            validation={{
              required: "Nama harus diisi",
            }}
          />
          <Input
            id="telp_number"
            label="Nomor Telepon"
            placeholder="Masukkan no telepon"
            className="w-full"
            validation={{
              required: "No telepon harus diisi",
            }}
          />
          <Input
            id="email"
            label="Email"
            placeholder="Masukkan email"
            className="w-full"
            validation={{
              required: "Email harus diisi",
            }}
          />
          <Input
            id="password"
            label="Password"
            placeholder="Masukkan password"
            type="password"
            validation={{
              required: "Password harus diisi",
              minLength: {
                value: 6,
                message: "Password minimal 6 karakter",
              },
            }}
          />
          <Input
            id="confirmPassword"
            label="Konfirmasi Password"
            placeholder="Konfirmasi password"
            type="password"
            validation={{
              required: "Konfirmasi password harus diisi",
              validate: (value) =>
                value === watch("password") || "Password tidak cocok",
            }}
          />
          <button
            className="border px-3 py-1.5 rounded-lg bg-green-500 text-white"
            type="submit"
          >
            Daftar
          </button>
        </form>
      </FormProvider>
      <p className="mt-4">
        Sudah punya akun?{" "}
        <Link to="/login" className="text-blue-500 underline">
          Login di sini
        </Link>
      </p>
    </Layout>
  );
}
