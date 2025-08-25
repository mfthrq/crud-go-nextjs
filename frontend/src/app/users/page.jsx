"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function UsersPage() {
  const router = useRouter();
  const [users, setUsers] = useState([]);

  const fetchUsers = () => {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/users`)
      .then((res) => res.json())
      .then((data) => {
        if (Array.isArray(data)) {
          setUsers(data);
        } else {
          setUsers([]); // fallback jika bukan array
          console.warn("Respons bukan array:", data);
        }
      })
      .catch((err) => console.error("Gagal fetch users:", err));
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleDelete = async (id) => {
    const confirm = window.confirm("Yakin ingin menghapus user ini?");
    if (!confirm) return;

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/users/${id}`, {
        method: "DELETE",
      });

      if (res.status === 204) {
        fetchUsers(); // refresh data
      } else {
        const errMsg = await res.text();
        alert("Gagal menghapus user: " + errMsg);
      }
    } catch (error) {
      console.error("Delete error:", error);
      alert("Terjadi kesalahan");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-white p-6">
      <div className="w-full max-w-4xl bg-gray-50 rounded-xl shadow-md p-8">
        <h1 className="text-3xl font-semibold mb-6 text-center">
          Daftar Users
        </h1>

        <button
          onClick={() => router.push(`/users/create`)}
          className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
        >
          Tambah Data
        </button>

        {!Array.isArray(users) || users.length === 0 ? (
          <p className="text-center text-gray-600">Belum ada data user.</p>
        ) : (
          <div className="overflow-x-auto">
            <table className="min-w-full border-collapse border border-gray-300 mx-auto mt-6">
              <thead>
                <tr>
                  <th className="border border-gray-300 px-4 py-2">ID</th>
                  <th className="border border-gray-300 px-4 py-2">Nama</th>
                  <th className="border border-gray-300 px-4 py-2">Email</th>
                  <th className="border border-gray-300 px-4 py-2">Aksi</th>
                </tr>
              </thead>
              <tbody>
                {users.map((user) => (
                  <tr key={user.id}>
                    <td className="border border-gray-300 px-4 py-2">
                      {user.id}
                    </td>
                    <td className="border border-gray-300 px-4 py-2">
                      {user.name}
                    </td>
                    <td className="border border-gray-300 px-4 py-2">
                      {user.email}
                    </td>
                    <td className="border border-gray-300 px-4 py-2">
                      <button
                        className="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600"
                        onClick={() => router.push(`/users/edit/${user.id}`)}
                      >
                        Edit
                      </button>
                      <button
                        className="bg-red-500 text-white ml-2 px-3 py-1 rounded hover:bg-red-600"
                        onClick={() => handleDelete(user.id)}
                      >
                        Delete
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}
