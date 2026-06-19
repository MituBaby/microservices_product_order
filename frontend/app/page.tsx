"use client";

import { useState, useEffect } from "react";

interface Product {
  id?: string;
  name: string;
  price: number;
  stock: number;
}

export default function Dashboard() {
  const [products, setProducts] = useState<Product[]>([]);
  const [name, setName] = useState("");
  const [price, setPrice] = useState(0);
  const [stock, setStock] = useState(0);
  
  // State untuk Order
  const [selectedProductId, setSelectedProductId] = useState("");
  const [quantity, setQuantity] = useState(1);
  const [message, setMessage] = useState("");

  const GATEWAY_URL = "http://localhost:8000";

  // 1. Ambil data produk saat halaman pertama kali dibuka
  const fetchProducts = async () => {
    try {
      const res = await fetch(`${GATEWAY_URL}/products`);
      const data = await res.json();
      if (Array.isArray(data)) setProducts(data);
    } catch (err) {
      console.error("Gagal mengambil data produk:", err);
    }
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  // 2. Tambah Produk Baru
  const handleAddProduct = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await fetch(`${GATEWAY_URL}/products`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, price: Number(price), stock: Number(stock) }),
      });
      if (res.ok) {
        setName(""); setPrice(0); setStock(0);
        fetchProducts(); // Refresh list
        setMessage("Produk berhasil ditambahkan!");
      }
    } catch (err) {
      setMessage("Gagal menambahkan produk.");
    }
  };

  // 3. Fungsi Buat Order (Pengurangan Stok Otomatis)
  const handleCreateOrder = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedProductId) {
      setMessage("Pilih produk terlebih dahulu!");
      return;
    }
    try {
      const res = await fetch(`${GATEWAY_URL}/orders`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ productId: selectedProductId, quantity: Number(quantity) }),
      });
      const data = await res.json();
      
      if (res.ok) {
        setMessage(`Order Sukses! ID Order: ${data.id}`);
        fetchProducts(); // Refresh stok produk yang berkurang
      } else {
        setMessage(`Gagal: ${data.error || "Stok tidak mencukupi"}`);
      }
    } catch (err) {
      setMessage("Terjadi kesalahan koneksi ke server.");
    }
  };

  return (
    <div className="min-h-screen bg-gray-100 p-8 text-gray-800">
      <header className="mb-8 border-b-2 border-blue-500 pb-4">
        <h1 className="text-3xl font-bold text-blue-600">BIP ERP - Inventory & Order Dashboard</h1>
        <p className="text-sm text-gray-500">Take Home Test System PT. Bharata</p>
      </header>

      {message && (
        <div className="mb-6 p-4 bg-blue-100 text-blue-800 font-medium rounded shadow">
          {message}
        </div>
      )}

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* FORM TAMBAH PRODUK */}
        <div className="bg-white p-6 rounded-lg shadow">
          <h2 className="text-xl font-semibold mb-4 text-gray-700">Tambah Produk Baru</h2>
          <form onSubmit={handleAddProduct} className="space-y-4">
            <div>
              <label className="block text-sm font-medium mb-1">Nama Produk</label>
              <input type="text" value={name} onChange={(e) => setName(e.target.value)} className="w-full border p-2 rounded" required />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Harga (Rp)</label>
              <input type="number" value={price} onChange={(e) => setPrice(Number(e.target.value))} className="w-full border p-2 rounded" required />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Stok Awal</label>
              <input type="number" value={stock} onChange={(e) => setStock(Number(e.target.value))} className="w-full border p-2 rounded" required />
            </div>
            <button type="submit" className="w-full bg-blue-600 text-white p-2 rounded font-bold hover:bg-blue-700">Simpan Produk</button>
          </form>
        </div>

        {/* LIST INVENTORY / STOK PRODUK */}
        <div className="bg-white p-6 rounded-lg shadow lg:col-span-2">
          <h2 className="text-xl font-semibold mb-4 text-gray-700">Daftar Stok Produk (Product Service)</h2>
          <div className="overflow-x-auto">
            <table className="w-full table-auto border-collapse text-left">
              <thead>
                <tr className="bg-gray-200">
                  <th className="p-3 border">ID Produk</th>
                  <th className="p-3 border">Nama</th>
                  <th className="p-3 border">Harga</th>
                  <th className="p-3 border">Stok</th>
                </tr>
              </thead>
              <tbody>
                {products.length === 0 ? (
                  <tr>
                    <td colSpan={4} className="p-3 text-center text-gray-400">Belum ada produk di database.</td>
                  </tr>
                ) : (
                  products.map((p) => (
                    <tr key={p.id} className="hover:bg-gray-50">
                      <td className="p-3 border text-xs font-mono">{p.id}</td>
                      <td className="p-3 border font-semibold">{p.name}</td>
                      <td className="p-3 border">Rp {p.price.toLocaleString("id-ID")}</td>
                      <td className={`p-3 border font-bold ${p.stock < 5 ? "text-red-500" : "text-green-600"}`}>{p.stock}</td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        </div>

        {/* FORM BUAT PESANAN (ORDER SERVICE) */}
        <div className="bg-white p-6 rounded-lg shadow lg:col-span-3">
          <h2 className="text-xl font-semibold mb-4 text-gray-700">Buat Pesanan Baru (Order Service)</h2>
          <form onSubmit={handleCreateOrder} className="flex flex-col md:flex-row md:items-end gap-4">
            <div className="flex-1">
              <label className="block text-sm font-medium mb-1">Pilih Produk</label>
              <select value={selectedProductId} onChange={(e) => setSelectedProductId(e.target.value)} className="w-full border p-2 rounded bg-white" required>
                <option value="">-- Pilih Produk yang Tersedia --</option>
                {products.map((p) => (
                  <option key={p.id} value={p.id}>{p.name} (Sisa Stok: {p.stock})</option>
                ))}
              </select>
            </div>
            <div className="w-full md:w-32">
              <label className="block text-sm font-medium mb-1">Jumlah</label>
              <input type="number" min="1" value={quantity} onChange={(e) => setQuantity(Number(e.target.value))} className="w-full border p-2 rounded" required />
            </div>
            <button type="submit" className="w-full md:w-48 bg-green-600 text-white p-2 rounded font-bold hover:bg-green-700 h-11">Kirim Order</button>
          </form>
        </div>
      </div>
    </div>
  );
}