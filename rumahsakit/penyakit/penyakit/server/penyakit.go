package server

import (
	"context"
)

// langkah ke lima
type penyakit struct {
	writer ReadWriter
}

func NewPenyakit(writer ReadWriter) PenyakitService {
	return &penyakit{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (d *penyakit) AddPenyakitService(ctx context.Context, penyakit Penyakit) error {
	//fmt.Println("customer")
	err := d.writer.AddPenyakit(penyakit)
	if err != nil {
		return err
	}

	return nil
}

func (d *penyakit) ReadPenyakitByKodeService(ctx context.Context, mob string) (Penyakit, error) {
	pen, err := d.writer.ReadPenyakitByKode(mob)
	//fmt.Println(pen)
	if err != nil {
		return pen, err
	}
	return pen, nil
}

func (d *penyakit) ReadPenyakitService(ctx context.Context) (Penyakits, error) {
	pny, err := d.writer.ReadPenyakit()
	//fmt.Println("customer", pen)
	if err != nil {
		return pny, err
	}
	return pny, nil
}

func (d *penyakit) UpdatePenyakitService(ctx context.Context, pny Penyakit) error {
	err := d.writer.UpdatePenyakit(pny)
	if err != nil {
		return err
	}
	return nil
}

func (d *penyakit) ReadPenyakitByKeteranganService(ctx context.Context, ket string) (Penyakits, error) {
	pen, err := d.writer.ReadPenyakitByKeterangan(ket)
	//fmt.Println(pen)
	if err != nil {
		return pen, err
	}
	return pen, nil
}
