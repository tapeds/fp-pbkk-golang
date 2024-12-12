import React, { useState } from "react";
import { useCookies } from "react-cookie";
import Layout from "../components/Layout";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { ApiURL } from "../const/api";
import { useNavigate } from "react-router-dom";

export default function ProfilePage() {
  const [cookies, , removeCookie] = useCookies(["authToken"]);
  const [isEditModalOpen, setEditModalOpen] = useState(false);
  const [editData, setEditData] = useState({ name: "", email: "", telp_number: "" });
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const { data: userData, isLoading, error } = useQuery({
    queryKey: ["userProfile"],
    queryFn: async () => {
      if (!cookies.authToken) {
        throw new Error("Unauthorized: No token found");
      }
      const res = await axios.get(ApiURL + "/user/me", {
        headers: {
          Authorization: `Bearer ${cookies.authToken}`,
        },
      });
      return res.data.data;
    },
  });

  const editUserMutation = useMutation({
    mutationFn: async (updatedData) => {
      if (!cookies.authToken) {
        throw new Error("Unauthorized: No token found");
      }
      await axios.patch(ApiURL + "/user", updatedData, {
        headers: {
          Authorization: `Bearer ${cookies.authToken}`,
        },
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries(["userProfile"]);
      setEditModalOpen(false);
    },
  });

  const deleteUserMutation = useMutation({
    mutationFn: async () => {
      if (!cookies.authToken) {
        throw new Error("Unauthorized: No token found");
      }
      await axios.delete(ApiURL + "/user", {
        headers: {
          Authorization: `Bearer ${cookies.authToken}`,
        },
      });
    },
    onSuccess: () => {
      removeCookie("authToken");
      navigate("/login");
    },
  });

  const handleEdit = () => {
    setEditData({
      name: userData.name,
      email: userData.email,
      telp_number: userData.telp_number || "",
    });
    setEditModalOpen(true);
  };

  const handleModalChange = (e) => {
    const { name, value } = e.target;
    setEditData((prev) => ({ ...prev, [name]: value }));
  };

  const handleModalSubmit = (e) => {
    e.preventDefault();
    editUserMutation.mutate(editData);
  };

  const handleDelete = () => {
    if (window.confirm("Apakah Anda yakin ingin menghapus akun ini?")) {
      deleteUserMutation.mutate();
    }
  };

  const handleLogout = () => {
    removeCookie("authToken");
    navigate("/login");
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
        <h1 className="text-3xl font-bold">Error fetching user data</h1>
      </Layout>
    );
  }

  return (
    <Layout className="flex-col">
      <h1 className="text-3xl mb-10 font-bold">Profil Saya</h1>
      <div className="bg-white p-6 rounded-lg shadow-md w-3/4">
        <div className="flex flex-col gap-3">
          <p>
            <strong>Nama:</strong> {userData.name}
          </p>
          <p>
            <strong>Email:</strong> {userData.email}
          </p>
          <p>
            <strong>Nomor Telepon:</strong> {userData.telp_number || "Tidak tersedia"}
          </p>
        </div>
        <div className="flex gap-4 mt-6">
          <button
            onClick={handleEdit}
            className="bg-green-500 text-white px-4 py-2 rounded-lg"
          >
            Edit
          </button>
          <button
            onClick={handleDelete}
            className="bg-red-500 text-white px-4 py-2 rounded-lg"
          >
            Delete
          </button>
          <button
            onClick={handleLogout}
            className="bg-blue-500 text-white px-4 py-2 rounded-lg"
          >
            Logout
          </button>
        </div>
      </div>

      {isEditModalOpen && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center">
          <div className="bg-white p-6 rounded-lg shadow-lg w-1/3">
            <h2 className="text-2xl font-bold mb-4">Edit Profil</h2>
            <form onSubmit={handleModalSubmit} className="flex flex-col gap-4">
              <div>
                <label className="block text-sm font-medium">Nama</label>
                <input
                  type="text"
                  name="name"
                  value={editData.name}
                  onChange={handleModalChange}
                  className="w-full border rounded px-3 py-2"
                />
              </div>
              <div>
                <label className="block text-sm font-medium">Email</label>
                <input
                  type="email"
                  name="email"
                  value={editData.email}
                  onChange={handleModalChange}
                  className="w-full border rounded px-3 py-2"
                />
              </div>
              <div>
                <label className="block text-sm font-medium">Nomor Telepon</label>
                <input
                  type="text"
                  name="telp_number"
                  value={editData.telp_number}
                  onChange={handleModalChange}
                  className="w-full border rounded px-3 py-2"
                />
              </div>
              <div className="flex justify-end gap-2">
                <button
                  type="button"
                  onClick={() => setEditModalOpen(false)}
                  className="bg-gray-500 text-white px-4 py-2 rounded-lg"
                >
                  Batal
                </button>
                <button
                  type="submit"
                  className="bg-blue-500 text-white px-4 py-2 rounded-lg"
                >
                  Simpan
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </Layout>
  );
}
