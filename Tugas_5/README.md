# Tugas 5 Pemrograman Jaringan (Kelompok)
### Anggota :
### - Alvien Ihsan Ramadhan&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&nbsp;(1301150088)
### - Chando Anggara Natanael Batubara&nbsp;(1301154390)
### - Chindy Amalia&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&nbsp;&nbsp;(1301164301)
<br></br>

## Web Server Design

Finite State Machine Web Server :

<p align="center">
  <img width="700" height="100" src="New Open Me First - Getting Started (7).png">
</p>

cara kerja :

<p align="justify">
Sesuai dengan FSM diatas cara kerja dari aplikasi web server yang kami rancang secara umum sebagai berikut :
</p>
<p align="justify">
Kondisi awal web server adalah Idle, namun pada saat client (browser) meminta data web page kepada server, maka web server aktif dan instruksi permintaan data oleh browser tersebut akan dikemas di dalam TCP yang merupakan protokol transport dan dikirim ke alamat yang dalam hal ini merupakan protokol berikutnya yaitu Hyper Text Transfer Protocol (HTTP) dan atau Hyper Text Transfer Protocol Secure (HTTPS).
</p>
<p align="justify">
Data yang diminta dari browser ke web server disebut dengan HTTP request yang kemudian akan dicarikan oleh web server di dalam komputer server. Jika ditemukan, data tersebut akan di handle atau diproses oleh aplikasi web server untuk di cari lagi di database, setelah itu data yang di minta  dikemas oleh web server dalam TCP dan dikirim kembali ke browser untuk ditampilkan.</p>
<p align="justify">
Nah, data yang dikirim dari server ke browser dikenal dengan HTTP response. Jika data yang diminta oleh browser tersebut ternyata tidak ditemukan oleh web server, maka web server akan menolak permintaan tersebut dan browser akan menampilkan notifikasi Page Not Found atau error 404.
</p>


## Web Server Implementation
<p align="center">
  "on progress"
</p>
