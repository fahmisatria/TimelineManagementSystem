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
	allkejaran := DataKejaran{}
	//Query untuk GetAllKejaranDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, divisi.Ketua_Divisi, pekerjaan.Nama_Pekerjaan, pekerjaan.Tanggal_Mulai, pekerjaan.Tanggal_Tenggat, pekerjaan.Deskripsi_Pekerjaan, pekerjaan.Penanggung_Jawab,divisi.Email_Ketua_Divisi from pekerjaan inner join divisi where pekerjaan.ID_Divisi = divisi.ID_Divisi")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&allkejaran.Nama_Divisi, &allkejaran.Ketua_Divisi, &allkejaran.Nama_Pekerjaan, &allkejaran.Tanggal_Mulai, &allkejaran.Tanggal_Tenggat, &allkejaran.Deskripsi_Pekerjaan, &allkejaran.Penanggung_Jawab,&allkejaran.Email_Ketua_Divisi)
		//Error Handling
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rps).Encode(&allkejaran)
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
	namadivisi := DataKejaran{}
	//Query untuk GetKejaranByNamaDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, divisi.Ketua_Divisi, pekerjaan.Nama_Pekerjaan, pekerjaan.Tanggal_Mulai, pekerjaan.Tanggal_Tenggat, pekerjaan.Deskripsi_Pekerjaan, pekerjaan.Penanggung_Jawab,divisi.Email_Ketua_Divisi from pekerjaan inner join divisi where pekerjaan.ID_Divisi = divisi.ID_Divisi and divisi.Nama_Divisi like?", "%"+Nama_Divisi+"%")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&namadivisi.Nama_Divisi, &namadivisi.Ketua_Divisi, &namadivisi.Nama_Pekerjaan, &namadivisi.Tanggal_Mulai, &namadivisi.Tanggal_Tenggat, &namadivisi.Deskripsi_Pekerjaan, &namadivisi.Penanggung_Jawab, &namadivisi.Email_Ketua_Divisi)
		//Error Handling
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rps).Encode(&namadivisi)
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
	pjb := DataKejaran{}
	//Query untuk GetKejaranByNamaDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, divisi.Ketua_Divisi, pekerjaan.Nama_Pekerjaan, pekerjaan.Tanggal_Mulai, pekerjaan.Tanggal_Tenggat, pekerjaan.Deskripsi_Pekerjaan, pekerjaan.Penanggung_Jawab,divisi.Email_Ketua_Divisi from pekerjaan inner join divisi where pekerjaan.ID_Divisi = divisi.ID_Divisi and pekerjaan.Penanggung_Jawab like?", "%"+Penanggung_Jawab+"%")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&pjb.Nama_Divisi, &pjb.Ketua_Divisi, &pjb.Nama_Pekerjaan, &pjb.Tanggal_Mulai, &pjb.Tanggal_Tenggat, &pjb.Deskripsi_Pekerjaan, &pjb.Penanggung_Jawab, &pjb.Email_Ketua_Divisi)
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
	tggltgt := DataKejaran{}
	//Query untuk GetKejaranByNamaDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, divisi.Ketua_Divisi, pekerjaan.Nama_Pekerjaan, pekerjaan.Tanggal_Mulai, pekerjaan.Tanggal_Tenggat, pekerjaan.Deskripsi_Pekerjaan, pekerjaan.Penanggung_Jawab,divisi.Email_Ketua_Divisi from pekerjaan inner join divisi where pekerjaan.ID_Divisi = divisi.ID_Divisi and pekerjaan.Tanggal_Tenggat like?", "%"+Tanggal_Tenggat+"%")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&tggltgt.Nama_Divisi, &tggltgt.Ketua_Divisi, &tggltgt.Nama_Pekerjaan, &tggltgt.Tanggal_Mulai, &tggltgt.Tanggal_Tenggat, &tggltgt.Deskripsi_Pekerjaan, &tggltgt.Penanggung_Jawab, &tggltgt.Email_Ketua_Divisi)
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

//Fungsi Untuk Mendapatkan Kejaran Divisi Berdasarkan Pekerjaan Divisi
func GetKejaranByPekerjaan(rps http.ResponseWriter, rqs *http.Request, Nama_Pekerjaan string) {
	//Membuka koneksi ke database "timeline_management_system"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer dbms.Close()
	//Inisialisasi Data Kejaran By Tanggal Tenggat
	pekerjaan := DataKejaran{}
	//Query untuk GetKejaranByNamaDivisi
	baris, err := dbms.Query("select divisi.Nama_Divisi, divisi.Ketua_Divisi, pekerjaan.Nama_Pekerjaan, pekerjaan.Tanggal_Mulai, pekerjaan.Tanggal_Tenggat, pekerjaan.Deskripsi_Pekerjaan, pekerjaan.Penanggung_Jawab,divisi.Email_Ketua_Divisi from pekerjaan inner join divisi where pekerjaan.ID_Divisi = divisi.ID_Divisi and pekerjaan.Nama_Pekerjaan like?", "%"+Nama_Pekerjaan+"%")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	defer baris.Close()
	//Scanning Rows
	for baris.Next(){
		err:= baris.Scan(&pekerjaan.Nama_Divisi, &pekerjaan.Ketua_Divisi, &pekerjaan.Nama_Pekerjaan, &pekerjaan.Tanggal_Mulai, &pekerjaan.Tanggal_Tenggat, &pekerjaan.Deskripsi_Pekerjaan, &pekerjaan.Penanggung_Jawab, &pekerjaan.Email_Ketua_Divisi)
		//Error Handling
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(rps).Encode(&pekerjaan)
	}
	err=baris.Err()
	if err != nil {
		log.Fatal(err)
	}
}
//Fungsi Untuk melakukan POST Data Kejaran Divisi ke tabel memiliki
func PostPekerjaan (rps http.ResponseWriter, rqs *http.Request) {
	var pekerjaanbaru Pekerjaan
	x := json.NewDecoder(rqs.Body)
	err:= x.Decode(&pekerjaanbaru)
	if err != nil {
		log.Fatal(err)
	}
	defer rqs.Body.Close()
	//Membuka koneksi ke database "timelinemanagementsystem"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil{
		log.Fatal(err)
	}
	//Query untuk melakukan insert into tabel pekerjaan 
	baris,err := dbms.Prepare("insert into pekerjaan (ID_Pekerjaan,ID_Divisi,Nama_Pekerjaan,Tanggal_Mulai,Tanggal_Tenggat,Deskripsi_Pekerjaan,Penanggung_Jawab) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_,err = baris.Exec(pekerjaanbaru.ID_Pekerjaan, pekerjaanbaru.ID_Divisi,pekerjaanbaru.Nama_Pekerjaan,pekerjaanbaru.Tanggal_Mulai,pekerjaanbaru.Tanggal_Tenggat,pekerjaanbaru.Deskripsi_Pekerjaan,pekerjaanbaru.Penanggung_Jawab)
}

//Fungsi Untuk melakukan UPDATE tabel data pekerjaan / kejaran divisi
func UpdateTabelPekerjaan(rps http.ResponseWriter, rqs *http.Request, ID_Pekerjaan string) {
	var memilikipekerjaan Pekerjaan
	idpekerjaan := ID_Pekerjaan
	memilikipekerjaandecoder := json.NewDecoder(rqs.Body)
	err := memilikipekerjaandecoder.Decode(&memilikipekerjaan)
	if err != nil {
		log.Fatal(err)
	}
	defer rqs.Body.Close()
	//Membuka koneksi ke database "timelinemanagementsystem"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil{
		log.Fatal(err)
	}
	//Query untuk melakukan update data kejaran pada tabel pekerjaan
	baris,err := dbms.Prepare("UPDATE pekerjaan SET pekerjaan.Nama_Pekerjaan = ?, pekerjaan.Tanggal_Tenggat = ?, pekerjaan.Deskripsi_Pekerjaan = ?, pekerjaan.Penanggung_Jawab = ? where pekerjaan.ID_Pekerjaan like?")
	if err != nil {
		log.Fatal(err)
	}
	_,err = baris.Exec(memilikipekerjaan.Nama_Pekerjaan, memilikipekerjaan.Tanggal_Tenggat, memilikipekerjaan.Deskripsi_Pekerjaan, memilikipekerjaan.Penanggung_Jawab, idpekerjaan)

}

//Fungsi untuk melakukan delete kejaran by tanggal tenggat
func DeleteKejaranDivisi(rps http.ResponseWriter, rqs *http.Request, ID_Pekerjaan string) {
	memilikipekerjaan := ID_Pekerjaan
	//Membuka koneksi ke database "timelinemanagementsystem"
	dbms, err := sql.Open("mysql",
			  "root:@tcp(127.0.0.1:3306)/timeline_management_system")
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}
	
	//Query untuk melakukan delete kejaran pada tabel memiliki berdasarkan tanggal tenggat
	baris,err := dbms.Query("DELETE FROM pekerjaan WHERE ID_Pekerjaan=?",memilikipekerjaan)

	defer baris.Close()
}
