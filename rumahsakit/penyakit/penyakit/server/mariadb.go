package server

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addPenyakit                = `INSERT INTO penyakit(kode_penyakit,nama_penyakit,deskripsi,keterangan) VALUES (?,?,?,?)`
	selectPenyakitByKode       = `SELECT nama_penyakit,deskripsi FROM penyakit WHERE kode_penyakit = ?`
	selectPenyakit             = `SELECT kode_penyakit, nama_penyakit, deskripsi, keterangan FROM penyakit`
	updatePenyakit             = `UPDATE penyakit SET nama_penyakit =?, deskripsi =?, status =?, keterangan=? WHERE kode_penyakit =?`
	selectPenyakitByKeterangan = `SELECT nama_penyakit,deskripsi,status,keterangan FROM penyakit WHERE keterangan LIKE ?`
)

// Langkah ke 4

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddPenyakit(penyakit Penyakit) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addPenyakit, penyakit.KodePenyakit, penyakit.NamaPenyakit, penyakit.Deskripsi, penyakit.Keterangan)
	fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadPenyakitByKode(kode string) (Penyakit, error) {
	penyakit := Penyakit{KodePenyakit: kode}
	err := rw.db.QueryRow(selectPenyakitByKode, kode).Scan(&penyakit.NamaPenyakit, &penyakit.Deskripsi, &penyakit.Keterangan)

	if err != nil {
		return Penyakit{}, err
	}

	return penyakit, nil
}

func (rw *dbReadWriter) ReadPenyakit() (Penyakits, error) {
	penyakit := Penyakits{}
	rows, _ := rw.db.Query(selectPenyakit)
	defer rows.Close()
	for rows.Next() {
		var d Penyakit
		err := rows.Scan(&d.KodePenyakit, &d.NamaPenyakit, &d.Deskripsi, &d.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return penyakit, err
		}
		penyakit = append(penyakit, d)
	}
	//fmt.Println("db nya:", penyakit)
	return penyakit, nil
}

func (rw *dbReadWriter) UpdatePenyakit(pny Penyakit) error {
	fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updatePenyakit, pny.NamaPenyakit, pny.Deskripsi, pny.Status, pny.KodePenyakit, pny.Keterangan)

	//fmt.Println("name:", pen.Name, pen.KodePenyakit)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadPenyakitByKeterangan(ket string) (Penyakits, error) {
	penyakit := Penyakits{}
	rows, _ := rw.db.Query(selectPenyakitByKeterangan, ket)
	defer rows.Close()
	for rows.Next() {
		var d Penyakit
		err := rows.Scan(&d.KodePenyakit, &d.NamaPenyakit, &d.Deskripsi, &d.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return penyakit, err
		}
		penyakit = append(penyakit, d)
	}
	//fmt.Println("db nya:", penyakit)
	return penyakit, nil
}
