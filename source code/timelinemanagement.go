// Nama : Fahmi Satria Aji 
// NIM  : 18215021
// Tugas : II3160 Pemrograman Integratif: Web Service 
// Deskripsi : File "timelinemanagement.go" berisi structure data yang dibutuhkan
// untuk web service "timelinemanagementsystem"

package main

//Struct Untuk Data Kejaran
type DataKejaran struct {
	Nama_Divisi string
	Ketua_Divisi string
	Nama_Pekerjaan string 
	Tanggal_Mulai string
	Tanggal_Tenggat string
	Deskripsi_Pekerjaan string
	Penanggung_Jawab string
	Email_Ketua_Divisi string
}

//Struct Untuk Tabel Pekerjaan
type Pekerjaan struct {
	ID_Pekerjaan string
	ID_Divisi string
	Nama_Pekerjaan string
	Tanggal_Mulai string
	Tanggal_Tenggat string
	Deskripsi_Pekerjaan string
	Penanggung_Jawab string
}
