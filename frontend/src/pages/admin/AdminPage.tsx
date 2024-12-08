import { useNavigate } from "react-router-dom";
import AdminLayout from "./component/AdminLayout";

export default function AdminPage() {
  const navigate = useNavigate();

  const handleNavigation = (route: string) => {
    navigate(route);
  };
  return (
    <AdminLayout className="flex-col">
      <h1 className="text-3xl font-bold mb-10">Admin Dashboard</h1>
      <div className="space-x-4">
        <button
          onClick={() => handleNavigation("/admin/penerbangan")}
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          Penerbangan
        </button>
        <button
          onClick={() => handleNavigation("/admin/bandara")}
          className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
        >
          Bandara
        </button>
        <button
          onClick={() => handleNavigation("/admin/maskapai")}
          className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
        >
          Maskapai
        </button>
      </div>
    </AdminLayout>
  );
}
