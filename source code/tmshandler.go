// Nama : Fahmi Satria Aji 
// NIM  : 18215021
// Tugas : II3160 Pemrograman Integratif: Web Service 
// Deskripsi : File "tmshandler.go adalah file handler yang berisi fungsi 
//dibutuhkan untuk menjalankan fungsi web service "timelinemanagementsystem" 

package main 

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
//Fungsi Untuk Mendapatkan Seluruh Kejaran Divisi
func GetAllKejaran(rps http.ResponseWriter, rqs *http.Request) {
	//Membuka koneksi ke database "timeline_management_system"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer dbms.Close()
	//Inisialisasi Data Kejaran 
	div := DataKejaran{}
	//Query untuk GetAllKejaranDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, pekerjaan.Nama_Pekerjaan, memiliki.Tanggal_Tenggat, memiliki.Penanggung_Jawab, memiliki.Deskripsi_Pekerjaan, memiliki.Kontak from memiliki inner join divisi, pekerjaan where memiliki.ID_Divisi = divisi.ID_Divisi and memiliki.ID_Pekerjaan = pekerjaan.ID_Pekerjaan")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&div.Nama_Divisi, &div.Nama_Pekerjaan, &div.Tanggal_Tenggat, &div.Penanggung_Jawab, &div.Deskripsi_Pekerjaan,&div.Kontak)
		//Error Handling
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rps).Encode(&div)
	}
	err=baris.Err()
	if err != nil {
		log.Fatal(err)
	}
}
//Fungsi Untuk Mendapatkan Kejaran Divisi Berdasarkan Nama Divisi
func GetKejaranByNamaDivisi(rps http.ResponseWriter, rqs *http.Request, Nama_Divisi string) {
	//Membuka koneksi ke database "timeline_management_system"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer dbms.Close()
	//Inisialisasi Data Kejaran Divisi
	divisi := DataKejaranByDivisi{}
	//Query untuk GetKejaranByNamaDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, pekerjaan.Nama_Pekerjaan, memiliki.Penanggung_Jawab, memiliki.Tanggal_Tenggat, memiliki.Kontak from memiliki inner join divisi, pekerjaan where memiliki.ID_Divisi = divisi.ID_Divisi and memiliki.ID_Pekerjaan = pekerjaan.ID_Pekerjaan and Nama_Divisi like?", Nama_Divisi)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&divisi.Nama_Divisi, &divisi.Nama_Pekerjaan, &divisi.Penanggung_Jawab, &divisi.Tanggal_Tenggat, &divisi.Kontak)
		//Error Handling
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rps).Encode(&divisi)
	}
	err=baris.Err()
	if err != nil {
		log.Fatal(err)
	}
}
//Fungsi Untuk Mendapatkan Kejaran Divisi Berdasarkan Penanggung Jawab
func GetKejaranByPenanggungJawab(rps http.ResponseWriter, rqs *http.Request, Penanggung_Jawab string) {
	//Membuka koneksi ke database "timeline_management_system"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer dbms.Close()
	//Inisialisasi Data Kejaran By PenanggungJawab
	pjb := DataKejaranByPenanggungJawab{}
	//Query untuk GetKejaranByNamaDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, pekerjaan.Nama_Pekerjaan, divisi.Ketua_Divisi, divisi.Bidang, divisi.Ketua_Bidang, memiliki.Tanggal_Tenggat, memiliki.Kontak from memiliki inner join divisi, pekerjaan where memiliki.ID_Divisi = divisi.ID_Divisi and memiliki.ID_Pekerjaan = pekerjaan.ID_Pekerjaan and memiliki.Penanggung_Jawab like?", Penanggung_Jawab)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&pjb.Nama_Divisi, &pjb.Nama_Pekerjaan, &pjb.Ketua_Divisi, &pjb.Bidang, &pjb.Ketua_Bidang, &pjb.Tanggal_Tenggat, &pjb.Kontak)
		//Error Handling
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rps).Encode(&pjb)
	}
	err=baris.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//Fungsi Untuk Mendapatkan Kejaran Divisi Berdasarkan Tanggal Tenggat
func GetKejaranByTanggalTenggat(rps http.ResponseWriter, rqs *http.Request, Tanggal_Tenggat string) {
	//Membuka koneksi ke database "timeline_management_system"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer dbms.Close()
	//Inisialisasi Data Kejaran By Tanggal Tenggat
	tggltgt := DataKejaranByTanggalTenggat{}
	//Query untuk GetKejaranByNamaDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, pekerjaan.Nama_Pekerjaan, divisi.Ketua_Divisi, divisi.Bidang, divisi.Ketua_Bidang, pekerjaan.Tanggal_Mulai, memiliki.Kontak from memiliki inner join divisi, pekerjaan where memiliki.ID_Divisi = divisi.ID_Divisi and memiliki.ID_Pekerjaan = pekerjaan.ID_Pekerjaan and memiliki.Tanggal_Tenggat like?", Tanggal_Tenggat)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&tggltgt.Nama_Divisi, &tggltgt.Nama_Pekerjaan, &tggltgt.Ketua_Divisi, &tggltgt.Bidang, &tggltgt.Ketua_Bidang, &tggltgt.Tanggal_Mulai, &tggltgt.Kontak)
		//Error Handling
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rps).Encode(&tggltgt)
	}
	err=baris.Err()
	if err != nil {
		log.Fatal(err)
	}
}


