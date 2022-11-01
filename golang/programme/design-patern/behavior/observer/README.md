# HOW TO IMPLEMENT

1. Lihat logika bisnis Anda dan coba bagi menjadi dua bagian: fungsionalitas inti, terpisah dari kode lain, akan bertindak sebagai penerbit; sisanya akan berubah menjadi satu set kelas pelanggan.
2. Deklarasikan antarmuka pelanggan. Minimal, itu harus mendeklarasikan metode pembaruan tunggal.
3. Deklarasikan antarmuka penerbit dan jelaskan sepasang metode untuk menambahkan objek pelanggan dan menghapusnya dari daftar. Ingat bahwa penerbit harus bekerja dengan pelanggan hanya melalui antarmuka pelanggan.
4. Decide where to put the actual subscription list and the implementation of subscription methods. Usually, this code looks the same for all types of publishers, so the obvious place to put it is in an abstract class derived directly from the publisher interface. Concrete publishers extend that class, inheriting the subscription behavior. </br>
Namun, jika Anda menerapkan pola ke hierarki kelas yang ada, pertimbangkan pendekatan berdasarkan komposisi: letakkan logika langganan ke objek terpisah, dan buat semua penerbit nyata menggunakannya.
5. Buat kelas penerbit konkret. Setiap kali sesuatu yang penting terjadi di dalam penerbit, itu harus memberi tahu semua pelanggannya.
6. Menerapkan metode notifikasi pembaruan di kelas pelanggan konkret. Sebagian besar pelanggan akan memerlukan beberapa data konteks tentang acara tersebut. Itu dapat diteruskan sebagai argumen dari metode notifikasi. </br>
Tapi ada pilihan lain. Setelah menerima notifikasi, pelanggan dapat mengambil data apa pun langsung dari notifikasi. Dalam hal ini, penerbit harus meneruskan dirinya sendiri melalui metode pembaruan. Opsi yang kurang fleksibel adalah menautkan penerbit ke pelanggan secara permanen melalui konstruktor.
7. Klien harus membuat semua pelanggan yang diperlukan dan mendaftarkannya ke penerbit yang tepat.
