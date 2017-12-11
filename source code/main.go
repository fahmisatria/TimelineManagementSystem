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

func main() {
	port := 8181 //port yang digunakan untuk menjalankan web service
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
				//Mendapatkan Kejaran Divisi berdasarkan Tanggal Tenggat suatu pekerjaan
				}else if rqs.URL.Query().Get("Tanggal_Tenggat") != ""{
					Tanggal_Tenggat := rqs.URL.Query().Get("Tanggal_Tenggat")
					GetKejaranByTanggalTenggat(rps,rqs,Tanggal_Tenggat)
					//Cek di bash
					log.Printf("Masukan anda telah diterima, masukan:%s", Tanggal_Tenggat)
				//Mendapatkan Seluruh Kejaran Divisi
				}else {
					GetAllKejaran(rps,rqs)
					log.Printf("Berhasil")
				}
				//Membuat data kejaran baru pada tabel memiliki 
			default:
				http.Error(rps, "invalid", 405) //mengirimkan respons error

		}
	})
	//Handle Function untuk melakukan POSTPekerjaan
	http.HandleFunc("/postpekerjaan/", func (rps http.ResponseWriter, rqs *http.Request) {
		switch rqs.Method {
			case "POST":
				//POST Pekerjaan To Tabel Pekerjaan
				PostPekerjaan(rps,rqs)
		default:
			http.Error(rps, "invalid", 405) //mengirimkan respons error
		}
	})
	//Handle Function untuk melakukan POSTDivisiMemilikiPekerjaan
	http.HandleFunc("/postmemilikipekerjaan/", func (rps http.ResponseWriter, rqs *http.Request) {
		switch rqs.Method {
			case "POST":
				//POST Pekerjaan To Tabel Pekerjaan
				PostMemilikiPekerjaan(rps,rqs)
		default:
			http.Error(rps, "invalid", 405) //mengirimkan respons error
		}
	})
	//Handle Function untuk melakukan method PUT 
	http.HandleFunc("/updatetanggaltenggat/", func(rps http.ResponseWriter, rqs *http.Request){
		switch rqs.Method {
		case "PUT":
			ID_Pekerjaan := rqs.URL.Query().Get("ID_Pekerjaan")
			UpdateTanggalTenggat(rps,rqs,ID_Pekerjaan)
		default:
			http.Error(rps, "invalid", 405)
		}
	})
	//Menampilkan pesan pada client bahwa web service dapat berjalan
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}
