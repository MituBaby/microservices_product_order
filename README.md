Cara menjalankan :
1. Buka terminal lalu ketik "docker-comppose up"

 Daftar API Endpoints (Via API Gateway - Port 8000)

Seluruh request wajib diarahkan ke API Gateway (http://localhost:8000). Gateway akan otomatis meneruskan request /products ke Product Service (Port 8001) dan /orders ke Order Service (Port 8002).

 1. Product Service (Manajemen Inventori)

 Method  Endpoint  Deskripsi 
 
 POST  /products  Membuat produk baru 
 GET  /products  Menampilkan semua daftar produk 
 GET  /products/:id  Menampilkan detail produk spesifik 
 PUT  /products/:id  Memperbarui data produk 
 DELETE  /products/:id  Menghapus produk dari database 

 2. Order Service (Transaksi & Stok)

 Method  Endpoint  Deskripsi 
 
 POST  /orders  Membuat pesanan baru (Otomatis potong stok) 
 GET  /orders  Menampilkan semua riwayat pesanan 

Berikut adalah kumpulan perintah cURL yang siap dieksekusi di terminal untuk menguji fungsionalitas sistem microservices kamu:

 A. Pengujian Product Service

1. Membuat Produk Baru (POST)

bash
curl -X POST http://localhost:8000/products \
     -H "Content-Type: application/json" \
     -d '{"name": "Amoxicillin 500mg", "price": 15000, "stock": 100}'



2. Mengambil Semua Produk (GET)

bash
curl -X GET http://localhost:8000/products



3. Mengambil Detail Produk (GET - Ganti ID sesuai MongoDB)

bash
curl -X GET http://localhost:8000/products/64f1a2b3c4d5e6f7a8b9c0d1



4. Memperbarui Data Produk (PUT - Ganti ID sesuai MongoDB)

bash
curl -X PUT http://localhost:8000/products/64f1a2b3c4d5e6f7a8b9c0d1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Amoxicillin 500mg Forte", "price": 18000, "stock": 120}'



5. Menghapus Produk (DELETE - Ganti ID sesuai MongoDB)

bash
curl -X DELETE http://localhost:8000/products/64f1a2b3c4d5e6f7a8b9c0d1



 B. Pengujian Order Service

1. Membuat Pesanan Baru (POST - Ganti productId dengan ID produk yang valid)

bash
curl -X POST http://localhost:8000/orders \
     -H "Content-Type: application/json" \
     -d '{"productId": "64f1a2b3c4d5e6f7a8b9c0d1", "quantity": 5}'



> Catatan internal untuk penguji: Jika stok kurang atau ID produk salah, sistem otomatis mengembalikan error 400/404 tanpa memotong data.

2. Mengambil Semua Riwayat Transaksi (GET)

bash
curl -X GET http://localhost:8000/orders



1. Saya memilih pola Reference untuk menyimpan data produk pada collection Order karena karakteristik data produk bersifat dinamis

2. Kelemahan utama sistem ini adalah komunikasi antar-service yang berjalan secara sinkronus. Jika memiliki waktu lebih, saya akan memperbaikinya dengan menerapkan Event-Driven Architecture menggunakan message broker seperti RabbitMQ serta menerapkan Saga Pattern untuk menjamin konsistensi data.

3. Untuk skala service yang lebih besar, saya akan menyiapkan Service Discovery dan Service Mesh untuk mengelola rute lalu lintas komunikasi antar-service secara dinamis.