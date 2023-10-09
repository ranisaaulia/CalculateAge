package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CalculateAge(w http.ResponseWriter, r *http.Request) {
	// Tangani permintaan POST dari formulir di sini
	r.ParseForm()
	nama := r.FormValue("nama")
	tanggal := r.FormValue("tanggal")
	bulan := r.FormValue("bulan")
	tahun := r.FormValue("tahun")

	// Parse tanggal lahir
	dob, err := time.Parse("02/01/2006", tanggal+"/"+bulan+"/"+tahun)
	if err != nil {
		fmt.Fprintf(w, "Format tanggal tidak valid.")
		return
	}

	// Hitung umur
	today := time.Now()
	umurTahun := today.Year() - dob.Year()
	if today.Month() < dob.Month() || (today.Month() == dob.Month() && today.Day() < dob.Day()) {
		umurTahun--
	}

	umurBulan := int(today.Month()) - int(dob.Month())
	if today.Day() < dob.Day() {
		umurBulan--
		if umurBulan < 0 {
			umurBulan += 12
		}
	}

	umurHari := today.Day() - dob.Day()
	if umurHari < 0 {
		lastDayOfLastMonth := today.AddDate(0, -1, 0).AddDate(0, 1, -today.Day())
		umurHari = lastDayOfLastMonth.Day() - dob.Day() + today.Day()
	}

	// Tentukan zodiak
	zodiak := ""
	if (dob.Month() == time.June && dob.Day() >= 21) || (dob.Month() == time.July && dob.Day() <= 22) {
		zodiak = "Cancer"
	} else if (dob.Month() == time.July && dob.Day() >= 23) || (dob.Month() == time.August && dob.Day() <= 22) {
		zodiak = "Leo"
	} else if (dob.Month() == time.August && dob.Day() >= 23) || (dob.Month() == time.September && dob.Day() <= 22) {
		zodiak = "Virgo"
	} else if (dob.Month() == time.September && dob.Day() >= 23) || (dob.Month() == time.October && dob.Day() <= 22) {
		zodiak = "Libra"
	} else if (dob.Month() == time.October && dob.Day() >= 23) || (dob.Month() == time.November && dob.Day() <= 21) {
		zodiak = "Scorpio"
	} else if (dob.Month() == time.November && dob.Day() >= 22) || (dob.Month() == time.December && dob.Day() <= 21) {
		zodiak = "Sagitarius"
	} else if (dob.Month() == time.December && dob.Day() >= 22) || (dob.Month() == time.January && dob.Day() <= 19) {
		zodiak = "Capricorn"
	} else if (dob.Month() == time.January && dob.Day() >= 20) || (dob.Month() == time.February && dob.Day() <= 18) {
		zodiak = "Aquarius"
	} else if (dob.Month() == time.February && dob.Day() >= 19) || (dob.Month() == time.March && dob.Day() <= 20) {
		zodiak = "Pisces"
	}

	// Kembalikan pesan sambutan dengan umur dan zodiak yang ditambahkan
	w.WriteHeader(http.StatusOK)
	pesanSambutan := fmt.Sprintf("Halo %s,\n", nama)
	pesanSambutan += fmt.Sprintf("Umur Anda saat ini adalah:\n")
	pesanSambutan += fmt.Sprintf("%d Tahun\n", umurTahun)
	pesanSambutan += fmt.Sprintf("%d Bulan\n", umurBulan)
	pesanSambutan += fmt.Sprintf("%d Hari\n", umurHari)

	pesanSambutan += fmt.Sprintf("\nBintang anda adalah\n%s", zodiak)

	fmt.Fprintf(w, pesanSambutan)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	router.HandleFunc("/calculate-age", CalculateAge).Methods("POST") // Menangani permintaan POST ke /calculate-age

	http.Handle("/", router)

	fmt.Println("Server started at :9999")
	http.ListenAndServe(":9999", nil)
}
