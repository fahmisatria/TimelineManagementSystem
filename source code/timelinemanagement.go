// Nama : Fahmi Satria Aji 
// NIM  : 18215021
// Tugas : II3160 Pemrograman Integratif: Web Service 
// Deskripsi : File "timelinemanagement.go" berisi structure data yang dibutuhkan
// untuk web service "timelinemanagementsystem"

package main 

//Struct Untuk Data Kejaran
type DataKejaran struct {
	Nama_Divisi string
	Nama_Pekerjaan string
	Tanggal_Tenggat string
	Penanggung_Jawab string
	Deskripsi_Pekerjaan string
	Kontak string
}


// Struct Untuk Data Kejaran berdasarkan divisi
type DataKejaranByDivisi struct {
	Nama_Divisi string
	Nama_Pekerjaan string
	Penanggung_Jawab string
	Tanggal_Tenggat string
	Kontak string 	
}

//Struct Untuk Data Kejaran berdasarkan penanggung jawab
type DataKejaranByPenanggungJawab struct {
	Nama_Divisi string
	Nama_Pekerjaan string
	Ketua_Divisi string
	Bidang string
	Ketua_Bidang string
	Tanggal_Tenggat string
	Kontak string
}

//Struct Untuk Data Kejaran berdasarkan tanggal tenggat
type DataKejaranByTanggalTenggat struct {
	Nama_Divisi string
	Nama_Pekerjaan string
	Ketua_Divisi string
	Bidang string
	Ketua_Bidang string
	Tanggal_Mulai string
	Kontak string
}

