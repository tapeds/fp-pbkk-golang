import { FormProvider, useForm } from "react-hook-form";
import Layout from "../../components/Layout";
import Input from "../../components/Input";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { ApiURL } from "../../const/api";
import { useNavigate } from "react-router-dom";
import { useCookies } from "react-cookie";

type LoginFormProps = {
  email: string;
  password: string;
};

export default function Login() {
  const methods = useForm<LoginFormProps>();
  const navigate = useNavigate();

  const [_, setCookie] = useCookies(["user"]);

  const { handleSubmit } = methods;

  const { mutate } = useMutation({
    mutationFn: async (data: LoginFormProps) => {
      const res = await axios.post(ApiURL + "/user/login", data);

      return res.data;
    },
    onSuccess: (res) => {
      setCookie("user", res.data);

      if (res.data.role === "admin") {
        navigate("/admin");
      }
    },
  });

  const onSubmit = (data: LoginFormProps) => {
    mutate(data);
  };

  return (
    <Layout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Login</h1>

      <FormProvider {...methods}>
        <form
          className="flex flex-col gap-3 w-1/2"
          onSubmit={handleSubmit(onSubmit)}
        >
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
    </Layout>
  );
}
