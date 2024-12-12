import { useQuery } from "@tanstack/react-query";
import Layout from "../components/Layout";
import axios from "axios";
import { ApiURL } from "../const/api";
// import { Link } from "react-router-dom";
import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";


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
  const { data, isLoading, error } = useQuery({
    queryKey: ["flights"],
    queryFn: async () => {
      const res = await axios.get(ApiURL + "/list-jadwal");
      return res.data.data;
    },
  });

  const [cookies] = useCookies(["user"]); 
  const navigate = useNavigate();

  const handleCheckout = (flightId: string) => {
    if (cookies.user && cookies.user.role === "user") {
      navigate(`/checkout/${flightId}`);

    } else {
      navigate(`/login?redirect=/checkout/${flightId}`);
    }
  };

  if (isLoading) {
    return (
      <Layout>
        <h1 className="text-3xl font-bold">Loading...</h1>
      </Layout>
    );
  }

  if (error) {
    return (
      <Layout>
        <h1 className="text-3xl font-bold">Error fetching data</h1>
      </Layout>
    );
  }

  return (
    <Layout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Jadwal Penerbangan</h1>
      <div className="flex flex-col gap-3 w-3/4">
        {data.map((flight: Flight) => (
          <div
            key={flight.id}
            className="border rounded-lg p-4 shadow-md bg-white flex justify-between items-center"
          >
            <div className="">
              <h2 className="text-xl font-bold">
                {flight.no_penerbangan}
                {/* <h2>cookie {cookies.user}</h2> */}
              </h2>
              <p className="text-sm">
                Waktu Berangkat: {new Date(flight.jadwal_berangkat).toLocaleString()}
              </p>
              <p className="text-sm">
                Waktu Tiba: {new Date(flight.jadwal_datang).toLocaleString()}
              </p>
              <p className="text-sm">Harga: Rp {flight.harga.toLocaleString()}</p>
              <p className="text-sm">Kapasitas: {flight.kapasitas}</p>
              <div className="flex gap-2">
                <img
                  src={flight.maskapai.image}
                  alt={flight.maskapai.name}
                  className="h-16 flex"
                />
                {/* <span>{flight.maskapai.name}</span> */}
              </div>
              <div>
                <h3 className="text-sm">Bandara:</h3>
                {flight.bandaras ? (
                  flight.bandaras.map((bandara) => (
                    <p key={bandara.id} className="text-sm">
                      {bandara.name} ({bandara.kota}) - {bandara.arah}
                    </p>
                  ))
                ) : (
                  <p className="text-sm">No airports listed</p>
                )}
              </div>
            </div>

            <div className="flex flex-col items-end">
              <button
                onClick={() => handleCheckout(flight.id)}
                className="bg-blue-500 text-white px-4 py-2 rounded-lg"
              >
                Rp {flight.harga.toLocaleString()}
              </button>
            </div>
          </div>
        ))}
      </div>
    </Layout>
  );
}
