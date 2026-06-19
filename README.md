1. Saya memilih pola Reference untuk menyimpan data produk pada collection Order karena karakteristik data produk bersifat dinamis

2. Kelemahan utama sistem ini adalah komunikasi antar-service yang berjalan secara sinkronus. Jika memiliki waktu lebih, saya akan memperbaikinya dengan menerapkan Event-Driven Architecture menggunakan message broker seperti RabbitMQ serta menerapkan Saga Pattern untuk menjamin konsistensi data.

3. Untuk skala service yang lebih besar, saya akan menyiapkan Service Discovery dan Service Mesh untuk mengelola rute lalu lintas komunikasi antar-service secara dinamis.