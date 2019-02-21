# Tugas 4 Pemrograman Jaringan

### Nama : Chando Anggara Natanael Batubara
### Nim     &nbsp;&nbsp; : 1301154390
<br></br>

## JSON (Javascript Object Notation)

JSON atau Javascript Object Notation adalah notasi standar yang umum digunakan untuk komunikasi data via web. JSON merupakan subset dari javascript. ( sumber:https://dasarpemrogramangolang.novalagung.com/50-json.html )

### JSON MARSHAL

  Code Program :  
<p align="center">
  <img width="250" height="320" src="Soal/JSON_Marshal.png">
</p>

  Hasil Running :
<p align="center">
  <img width="470" height="100" src="Jawaban/Screenshot Output/Output_JSON_Marshal.png">
</p>

Analisis :

Cara kerja dari program diatas adalah merubah sebuah data struct yaitu person kedalam bentuk JSON dengan menggunakan fungsi Marshal.

### JSON UNMARSHAL
  
  Code Program :  
<p align="center">
  <img width="250" height="320" src="Soal/JSON_UnMarshal.png">
</p>

  Hasil Running :
<p align="center">
  <img width="470" height="90" src="Jawaban/Screenshot Output/Output_JSON_UnMarshal.png">
</p>

Analisis :

Cara kerja dari program diatas adalah membuat  data json string kedalam struct yaitu person dengan menggunakan fungsi UnMarshal.

### Flatbuffer dan Protocol Buffer
Contoh program gRPC Flatbuffer bisa di download pada repository github berikut : 
https://github.com/jonog/grpc-flatbuffers-example 

Screenshoot output contoh program gRPC Flatbuffer :
  - Client
   <p align="center">
  <img width="530" height="300" src="Jawaban/Screenshot Output/Output_Client_Flatbuffer.png">
</p>
  
  - Server
   <p align="center">
  <img width="470" height="150" src="Jawaban/Screenshot Output/Output_Server_Flatbuffer.png">
</p>
