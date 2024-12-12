import React from "react";
import { useQuery } from "@tanstack/react-query";
import Layout from "../components/Layout";
import axios from "axios";
import { ApiURL } from "../const/api";
import { useNavigate } from "react-router-dom";
import { useCookies } from "react-cookie";

type PenumpangInput = {
  name: string;
  nik: string;
};

export default function Checkout() {
  const navigate = useNavigate();
  const [cookies] = useCookies(["authToken"]);

  // Handle the authentication token and create an axios instance with auth header
  const axiosInstance = axios.create({
    baseURL: ApiURL,
    headers: {
      Authorization: `Bearer ${cookies.authToken}`,
    },
  });

  // Redirect to login page if no authToken is present
  React.useEffect(() => {
    if (!cookies.authToken) {
      navigate(`/login?redirect=${window.location.pathname}`);
    }
  }, [cookies, navigate]);

  const { data: tiketData, isLoading, error } = useQuery({
    queryKey: ["ticketDetails"],
    queryFn: async () => {
      const tiketId = window.location.pathname.split("/")[2];
      const res = await axiosInstance.get(`/checkout/${tiketId}`);
      return res.data;
    },
  });

  const [penumpang, setPenumpang] = React.useState<PenumpangInput[]>([
    { name: "", nik: "" },
  ]);

  const handleInputChange = (
    index: number,
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    const { name, value } = e.target;
    const values = [...penumpang];
    values[index][name as keyof PenumpangInput] = value;
    setPenumpang(values);
  };

  const handleAddPenumpang = () => {
    setPenumpang([...penumpang, { name: "", nik: "" }]);
  };

  const handleRemovePenumpang = (index: number) => {
    const values = [...penumpang];
    values.splice(index, 1);
    setPenumpang(values);
  };
  
  const handleGoHome = async () => { navigate("/"); };
  const handleSubmit = async () => {
    console.log("Request Data: ", penumpang);
    console.log("Request Payload: ", {
      penumpangs: penumpang,
    });

    

    // Validasi input penumpang
    const isValid = penumpang.every((p) => p.name && p.nik);
    if (!isValid) {
      alert("Please fill in all passenger details.");
      return;
    }

    const tiketId = window.location.pathname.split("/")[2];
    try {
      // Menggunakan axiosInstance untuk memastikan header autentikasi diterapkan
      
      const response = await axiosInstance.post(`/checkout/${tiketId}`, {
        penumpang: penumpang,
      });
      // console.log({
      //   penumpangs: penumpang,
      // });
      alert("Checkout successful!");
      navigate("/success");
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error("Axios error response data:", error.response?.data);
        alert(
          `Failed to checkout: ${error.response?.data?.message || "Unknown error"}`
        );
      } else {
        console.error("Unknown error:", error);
        alert("An unknown error occurred. Please try again.");
      }
    }
  };

  if (isLoading) {
    return <Layout><h1>Loading...</h1></Layout>;
  }

  if (error) {
    return <Layout><h1>Error fetching ticket data</h1></Layout>;
  }

  return (
    <Layout>
<div className="flex-col">
      <h1 className="text-3xl font-bold mb-6">Checkout Ticket</h1>

      <div className="mb-6">
        <h2 className="text-xl font-semibold mb-4 ">Flight Information</h2>
        <div className="bg-white p-4 rounded shadow-md">
          <p><strong>No Penerbangan:</strong> {tiketData.data.no_penerbangan}</p>
          <p><strong>Departure:</strong> {tiketData.data.jadwal_berangkat}</p>
          <p><strong>Arrival:</strong> {tiketData.data.jadwal_datang}</p>
          <p><strong>Price:</strong> IDR {tiketData.data.harga} </p>
        </div>
      </div>

      <div className="mb-6">
        <h2 className="text-xl font-semibold mb-4">Passenger Information</h2>
        {penumpang.map((penumpangData, index) => (
          <div key={index} className="bg-white p-4 mb-4 rounded shadow-md">
            <label className="block mb-2">Name</label>
            <input
              type="text"
              name="name"
              value={penumpangData.name}
              onChange={(e) => handleInputChange(index, e)}
              className="w-full p-2 border rounded"
              placeholder="Enter name"
            />
            <label className="block mt-4 mb-2">NIK</label>
            <input
              type="text"
              name="nik"
              value={penumpangData.nik}
              onChange={(e) => handleInputChange(index, e)}
              className="w-full p-2 border rounded"
              placeholder="Enter NIK"
            />
            <button
              type="button"
              onClick={handleAddPenumpang}
              className="mt-4 text-blue-500"
            >
              Add Passenger
            </button>
            <button
              type="button"
              onClick={() => handleRemovePenumpang(index)}
              className=" mx-4 mt-4 text-red-500"
            >
              Remove Passenger
            </button>
          </div>
        ))}
        {/* <button
          type="button"
          onClick={handleAddPenumpang}
          className="mt-4 p-2 bg-blue-500 text-white rounded"
        >
          Add Passenger
        </button> */}
      </div>

      <div className="mt-4">
        <button
          onClick={handleSubmit}
          className="p-3 bg-green-500 text-white rounded"
        >
          Complete Checkout
        </button>

        <button
          onClick={handleGoHome}
          className="p-3 mx-4 bg-blue-500 text-white rounded"
        >
          Back
        </button>
      </div>
</div>
    </Layout>
  );
}
