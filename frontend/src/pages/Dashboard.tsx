import { useQuery } from "@tanstack/react-query";
import Layout from "../components/Layout";
import axios from "axios";
import { ApiURL } from "../const/api";
import { Link, useNavigate } from "react-router-dom";
import { useCookies } from "react-cookie";

type Flight = {
  id: string;
  no_penerbangan: string;
  jadwal_berangkat: string;
  jadwal_datang: string;
  harga: number;
  kapasitas: number;
  maskapai: {
    id: string;
    name: string;
    image: string;
  };
  bandaras: Array<{
    id: string;
    name: string;
    kode: string;
    kota: string;
    arah: string;
  }> | null;
};

export default function FlightSchedules() {
  const navigate = useNavigate();
  const [cookies] = useCookies(["authToken"]);
  const { data, isLoading, error } = useQuery({
    queryKey: ["flights"],
    queryFn: async () => {
      const token = cookies.authToken;
      if (!token) throw new Error("Unauthorized: No token found");

      const res = await axios.get(ApiURL + "/pesanan", {
        headers: {
          Authorization: `Bearer ${cookies.authToken}`,
        },
      });

      return res.data.data.map((item: any) => ({
        id: item.id,
        no_penerbangan: item.no_penerbangan,
        jadwal_berangkat: item.jadwal_berangkat,
        jadwal_datang: item.jadwal_datang,
        harga: item.harga,
        kapasitas: item.kapasitas,
        maskapai: {
          id: item.Maskapai.id,
          name: item.Maskapai.name || "Tidak diketahui",
          image: item.Maskapai.image || "https://via.placeholder.com/50",
        },
        bandaras: item.BandaraPenerbangan
          ? item.BandaraPenerbangan.map((bandara: any) => ({
              id: bandara.id,
              name: bandara.name,
              kode: bandara.kode,
              kota: bandara.kota,
              arah: bandara.arah,
            }))
          : null,
      }));
    },
  });

  const handleCheckout = (flightId: string) => {
    navigate(`/pesanan/${flightId}`);
  };

  return (
    <Layout className="flex-col gap-6 ">
      <header>
        <h1 className="text-3xl mb-6 font-bold">Jadwal Penerbangan Saya</h1>
      </header>
      {isLoading ? (
        <div className="text-center">
          <h2 className="text-xl font-semibold">Loading...</h2>
        </div>
      ) : error ? (
        <div className="text-center">
          <h2 className="text-xl font-semibold text-red-600">
            Error fetching data
          </h2>
        </div>
      ) : (
        <div className="flex flex-col gap-4 w-full max-w-4xl mx-auto ">
          {data.map((flight: Flight) => (
            <div
              key={flight.id}
              className="border rounded-lg p-4 shadow-md bg-white flex justify-between items-center"
            >
              <div>
                <h2 className="text-xl font-bold">{flight.no_penerbangan}</h2>
                <p className="text-sm">
                  Waktu Berangkat:{" "}
                  {new Date(flight.jadwal_berangkat).toLocaleString()}
                </p>
                <p className="text-sm">
                  Waktu Tiba:{" "}
                  {new Date(flight.jadwal_datang).toLocaleString()}
                </p>
                <p className="text-sm">
                  Harga: Rp {flight.harga.toLocaleString()}
                </p>
                <p className="text-sm">Kapasitas: {flight.kapasitas}</p>
                <div className="flex items-center gap-3 mt-2">
                  {/* <img
                    src={flight.maskapai.image}
                    alt={flight.maskapai.name}
                    className="w-16 h-16 rounded"
                  /> */}
                  {/* <span>{flight.maskapai.name}</span> */}
                </div>
                {/* <div>
                  <h3 className="text-sm font-semibold mt-2">Bandara:</h3>
                  {flight.bandaras ? (
                    flight.bandaras.map((bandara) => (
                      <p key={bandara.id} className="text-sm">
                        {bandara.name} ({bandara.kota}) - {bandara.arah}
                      </p>
                    ))
                  ) : (
                    <p className="text-sm">Tidak ada bandara</p>
                  )}
                </div> */}
              </div>
              {/* <button
                onClick={() => handleCheckout(flight.id)}
                className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
              >
                See Details
              </button> */}
            </div>
          ))}
        </div>
      )}
    </Layout>
  );
}
