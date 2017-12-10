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
	// Handle Function untuk menjalankan fungsionalitas GETKejaranDivisiByNamaDivisi
	http.HandleFunc("/divisimanagementsystem/", func (rps http.ResponseWriter, rqs *http.Request) {
		switch rqs.Method {
			//Method GET
			case "GET":
				sx:=rqs.URL.Query().Get("Nama_Divisi") //Client request dengan masukan Nama_Divisi
				if (sx != "") {
					GetKejaranByNamaDivisi(rps,rqs,sx) //Mendapatkan kejaran berdasarkan nama divisi
				} else {
					GetAllKejaran(rps,rqs) //Mendapatkan seluruh kejaran masing-masing divisi
				}
			default:
				http.Error(rps, "invalid", 405) //mengirimkan respons error

		}
	})

	// Handle Function untuk menjalankan fungsionalitas GETKejaranDivisiByPenanggungJawab
	http.HandleFunc("/penanggungjawabdivisi/", func (rps http.ResponseWriter, rqs *http.Request) {
		switch rqs.Method {
			//Method GET
			case "GET":
				sx:=rqs.URL.Query().Get("Penanggung_Jawab") //Client request dengan masukan Nama_Divisi
				if (sx != "") {
					GetKejaranByPenanggungJawab(rps,rqs,sx) //Mendapatkan kejaran berdasarkan nama divisi
				} else {
					GetAllKejaran(rps,rqs) //Mendapatkan seluruh kejaran masing-masing divisi
				}
			default:
				http.Error(rps, "invalid", 405) //mengirimkan respons error

		}
	})

	// Handle Function untuk menjalankan fungsionalitas GETKejaranDivisiByTanggalTenggat
		http.HandleFunc("/tanggaltenggatkejaran/", func (rps http.ResponseWriter, rqs *http.Request) {
		switch rqs.Method {
			//Method GET
			case "GET":
				sx:=rqs.URL.Query().Get("Tanggal_Tenggat") //Client request dengan masukan Nama_Divisi
				if (sx != "") {
					GetKejaranByTanggalTenggat(rps,rqs,sx) //Mendapatkan kejaran berdasarkan nama divisi
				} else {
					GetAllKejaran(rps,rqs) //Mendapatkan seluruh kejaran masing-masing divisi
				}
			default:
				http.Error(rps, "invalid", 405) //mengirimkan respons error

		}
	})
	//Menampilkan pesan pada client bahwa web service dapat berjalan
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}
