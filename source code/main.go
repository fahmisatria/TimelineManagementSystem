// Nama : Fahmi Satria Aji 
// NIM  : 18215021
// Tugas : II3160 Pemrograman Integratif: Web Service 
// Deskripsi : File "main.go adalah file utama yang dibutuhkan
// untuk menjalankan fungsi web service "timelinemanagementsystem" 

package main 

import (
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

//Main untuk menjalankan fungsionalitas web service 
func main() {
	port := 7601 //port yang digunakan untuk menjalankan web service
	// Handle Function untuk menjalankan fungsionalitas GETKejaranDivisi By Nama Divisi, Penanggung Jawab, dan Tanggal_Tenggat
	http.HandleFunc("/timelinemanagementsystem/", func (rps http.ResponseWriter, rqs *http.Request) {
		switch rqs.Method {
			//Method GET
			case "GET":
				//Mendapatkan Kejaran Divisi berdasarkan nama divisi
				if rqs.URL.Query().Get("Nama_Divisi") != ""{
					Nama_Divisi := rqs.URL.Query().Get("Nama_Divisi")
					GetKejaranByNamaDivisi(rps,rqs,Nama_Divisi)
					//Cek di bash
					log.Printf("Masukan anda telah diterima, masukan:%s",Nama_Divisi)
				//Mendapatkan Kejaran Divisi berdasarkan penanggung jawab 
				}else if rqs.URL.Query().Get("Penanggung_Jawab") != ""{
					Penanggung_Jawab := rqs.URL.Query().Get("Penanggung_Jawab")
					GetKejaranByPenanggungJawab(rps,rqs,Penanggung_Jawab)
					//Cek di bash
					log.Printf("Masukan anda telah diterima, masukan:%s",Penanggung_Jawab)
				//Mendapatkan Kejaran Divisi berdasarkan tanggal_tenggat pekerjaan
				}else if rqs.URL.Query().Get("Tanggal_Tenggat") != ""{
					Tanggal_Tenggat := rqs.URL.Query().Get("Tanggal_Tenggat")
					GetKejaranByTanggalTenggat(rps,rqs,Tanggal_Tenggat)
					//Cek di bash
					log.Printf("Masukan anda telah diterima, masukan:%s", Tanggal_Tenggat)
				//Mendapatkan Kejaran Divisi berdasarkan pekerjaan divisi
				}else if rqs.URL.Query().Get("Nama_Pekerjaan") != "" {
					Nama_Pekerjaan := rqs.URL.Query().Get("Nama_Pekerjaan") 
					GetKejaranByPekerjaan(rps,rqs,Nama_Pekerjaan)
					//Cek di bash
					log.Printf("Masukan anda telah diterima, masukan:%s",Nama_Pekerjaan)
				//Mendapatkan seluruh data kejaran divisi apabila masukan get klien tanpa syarat apapun/kosong
				}else {
					GetAllKejaran(rps,rqs)
					log.Printf("Berhasil")
				}
			//Method POST
			case "POST":
				//Memasukkan Data Kejaran Ke Tabel Pekerjaan
				PostPekerjaan(rps,rqs)
			//Method PUT
			case "PUT":
				//Update Tanggal Tenggat di Tabel Pekerjaan
				ID_Pekerjaan := rqs.URL.Query().Get("ID_Pekerjaan")
				UpdateTabelPekerjaan(rps,rqs,ID_Pekerjaan)
			//Method DELETE
			case "DELETE":
				ID_Pekerjaan := rqs.URL.Query().Get("ID_Pekerjaan")
				DeleteKejaranDivisi(rps,rqs,ID_Pekerjaan)
				//Delete Kejaran Divisi Berdasarkan ID_Pekerjaan
			default:
				http.Error(rps, "method request tidak valid", 405) //mengirimkan respons error
			//Method PUT

		}
	})
	//Menampilkan pesan pada client bahwa web service dapat berjalan
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}
