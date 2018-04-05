package controllers

import (
	"day15/item/data"
	"encoding/json"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	// ambil datanya
	// untuk ambil data perlu koneksi
	// 1.c buat koneksi
	context := Context{}
	d := DBInitial(context.DB)
	defer d.Close()
	// ambil data dari repositori
	// buat perintah di repositori
	// 1.d buat repo petugas
	repo := &data.ItemRepository{d}
	item := data.GetAll(repo)
	// olah datanya, olah errornya
	// tampilkan datanya
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}
