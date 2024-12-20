import { FormProvider, useForm } from "react-hook-form";
import Layout from "../../components/Layout";
import Input from "../../components/Input";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { ApiURL } from "../../const/api";
import { useNavigate, useLocation } from "react-router-dom";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";

type LoginFormProps = {
  email: string;
  password: string;
};

export default function Login() {
  const methods = useForm<LoginFormProps>();
  const navigate = useNavigate();
  const location = useLocation();
  const [_, setCookie] = useCookies(["user", "authToken"]);

  const { handleSubmit } = methods;

  const { mutate } = useMutation({
    mutationFn: async (data: LoginFormProps) => {
      const res = await axios.post(ApiURL + "/user/login", data);
      return res.data;
    },
    onSuccess: (res) => {
      setCookie("user", res.data, { path: "/" });
      const token = res.data.token;
      setCookie("authToken", token, { path: "/" });

      // Handle redirectUrl or default route
      const redirectUrl =
        new URLSearchParams(location.search).get("redirect") || "/";

      if (res.data.role === "admin") {
        navigate("/admin");
      } else if (res.data.role === "user") {
        navigate(redirectUrl);
      }
    },
    onError: (err) => {
      console.error("Login failed:", err);
      alert("Login failed. Please check your email and password.");
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
      <p className="mt-4">
        Belum punya akun?{" "}
        <Link to="/register" className="text-blue-500 underline">
          Daftar di sini
        </Link>
      </p>
    </Layout>
  );
}
