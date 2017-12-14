# TimelineManagementSystem
Proyek Web Service Tugas Besar II3160 Pemrograman Integratif

Berikut ini adalah langkah pengujian dari web services timeline management system
1. Buka web browser atau postman atau bash untuk pengujian web service. Untuk BASH bisa menggunakan CURL
2. Ketikkan command request berikut ini untuk menjalankan request (menggunakan POSTMAN).
GET 
- Mendapatkan seluruh datakejaran  : 167.205.67.246:7601/timelinemanagement/
- Mendapatkan data kejaran berdasarkan nama divisi : 167.205.67.246:7601/timelinemanagement/?Nama_Divisi=MSDM, MSDM dapat dimodifikasi dengan
divisi lain
- Mendapatkan data kejaran berdasarkan penanggung jawab divisi : 167.205.67.246:7601/timelinemanagement/?Penanggung_Jawab=Fahmi, Penanggung Jawab dapat dimodifikasi
dengan nama penanggung jawab lain
- Mendapatkan data kejaran berdasarkan tanggal tenggat kejaran : 167.205.67.246:7601/timelinemanagement/?Tanggal_Tenggat=2017, Tanggal tenggat dapat dimodifikasi
dengan tanggal (format: "yyyy-mm-dd"
- Mendapatkan data kejaran berdasarkan nama pekerjaan : 167.205.67.246:7601/timelinemanagement/?Nama_Pekerjaan= Kaos, Nama pekerjaan dapat dimodifikasi dengan
nama pekerjaan lain
POST
URI: 167.205.67.246:7601/timelinemanagement/
Body: {"ID_Pekerjaan":"PKJ020", "ID_Divisi": "EXP032", "Nama_Pekerjaan":"Sudah dapat yang ngisi stand sebanyak 20 stand", "Tanggal_Mulai": "2017-01-08", "Tanggal_Tenggat":"2018-01-31", "Deskripsi_Pekerjaan":"Menghubungi kontak stand expo yang bakalan mengisi stand expo", "Penanggung_Jawab":"Arief Septian Nurhada"}
UPDATE
URI: 167.205.67.246:7601/timelinemanagement/?ID_Pekerjaan=PKJ020
Body:{"ID_Pekerjaan":"PKJ020", "ID_Divisi": "EXP032", "Nama_Pekerjaan":"Sudah dapat yang ngisi stand sebanyak 20 stand", "Tanggal_Mulai": "2017-01-08", "Tanggal_Tenggat":"2018-01-31", "Deskripsi_Pekerjaan":"Menghubungi kontak kontak stand expo yang bakalan mengisi stand expo", "Penanggung_Jawab":"Arief Septian Nurhada"}
DELETE
URI: 167.205.67.246:7601/timelinemanagement/?ID_Pekerjaan=PKJ020
