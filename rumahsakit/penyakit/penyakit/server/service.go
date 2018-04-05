package server

import (
	"context"
)

type Status int32

const (
	//ServiceID is dispatch service ID

	//ini adalah konfigurasi sub domainnya
	ServiceID        = "rumahsakit"
	CreatedBy        = "Ilham"
	onAdd     Status = 1
)

type Penyakit struct {
	KodePenyakit string
	NamaPenyakit string
	Deskripsi    string
	Status       int32
	Keterangan   string
}

type Penyakits []Penyakit

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

// ini interface untuk melakukan read
type ReadWriter interface {
	AddPenyakit(Penyakit) error
	ReadPenyakitByKode(string) (Penyakit, error)
	ReadPenyakit() (Penyakits, error)
	UpdatePenyakit(Penyakit) error
	ReadPenyakitByKeterangan(string) (Penyakits, error)
}

// ini interface yang mempunyai nilai return yang berupa interfase
type PenyakitService interface {
	AddPenyakitService(context.Context, Penyakit) error
	ReadPenyakitByKodeService(context.Context, string) (Penyakit, error)
	ReadPenyakitService(context.Context) (Penyakits, error)
	UpdatePenyakitService(context.Context, Penyakit) error
	ReadPenyakitByKeteranganService(context.Context, string) (Penyakits, error)
}
