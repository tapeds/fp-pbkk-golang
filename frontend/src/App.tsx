import { Route, Routes } from "react-router-dom";
import AdminPage from "./pages/admin/AdminPage";
import AdminPenerbangan from "./pages/admin/AdminPenerbanganPage";
import AdminMaskapai from "./pages/admin/AdminMaskapaiPage";
import AdminBandara from "./pages/admin/AdminBandaraPage";

function App() {
  return (
    <Routes>
      <Route path="/admin" element={<AdminPage />} />
      <Route path="/admin/penerbangan" element={<AdminPenerbangan />} />
      <Route path="/admin/maskapai" element={<AdminMaskapai />} />
      <Route path="/admin/bandara" element={<AdminBandara />} />
    </Routes>
  );
}

export default App;
