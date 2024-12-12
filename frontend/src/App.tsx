import { Route, Routes } from "react-router-dom";
import AdminPage from "./pages/admin/AdminPage";
import AdminPenerbangan from "./pages/admin/AdminPenerbanganPage";
import AdminMaskapai from "./pages/admin/AdminMaskapaiPage";
import AdminBandara from "./pages/admin/AdminBandaraPage";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import JadwalPenerbangan from "./pages/JadwalPenerbangan";
import Checkout from "./pages/CheckoutPage";
import SuccessPage from "./pages/CheckoutResponse";
import Dashboard from "./pages/Dashboard";
import ProfilePage from "./pages/ProfilePage";
import Home from "./pages/Home";

function App() {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />

      <Route path="/" element={<Home />}></Route>
      <Route path="/pesanan" element={<Dashboard />} />
      <Route path="/booking" element={<JadwalPenerbangan />}></Route>
      <Route path="/checkout/:id" element={<Checkout />} />
      <Route path="/success" element={<SuccessPage />} />
      <Route path="/user" element={<Dashboard/>}></Route>
      <Route path="/profil" element={<ProfilePage />} />


      <Route path="/admin" element={<AdminPage />} />
      <Route path="/admin/penerbangan" element={<AdminPenerbangan />} />
      <Route path="/admin/maskapai" element={<AdminMaskapai />} />
      <Route path="/admin/bandara" element={<AdminBandara />} />
    </Routes>
  );
}

export default App;
