package endpoint

import (
	"context"
	"fmt"

	sv "rumahsakit/penyakit/penyakit/server"
)

func (pe PenyakitEndpoint) AddPenyakitService(ctx context.Context, penyakit sv.Penyakit) error {
	_, err := pe.AddPenyakitEndpoint(ctx, penyakit)
	return err
}

func (pe PenyakitEndpoint) ReadPenyakitByKodeService(ctx context.Context, Kode string) (sv.Penyakit, error) {
	req := sv.Penyakit{KodePenyakit: Kode}
	fmt.Println(req)
	resp, err := pe.ReadPenyakitByKodeEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	pen := resp.(sv.Penyakit)
	return pen, err
}

func (pe PenyakitEndpoint) ReadPenyakitService(ctx context.Context) (sv.Penyakits, error) {
	resp, err := pe.ReadPenyakitEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Penyakits), err
}

func (pe PenyakitEndpoint) UpdatePenyakitService(ctx context.Context, dok sv.Penyakit) error {
	_, err := pe.UpdatePenyakitEndpoint(ctx, dok)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (pe PenyakitEndpoint) ReadPenyakitByKeteranganService(ctx context.Context, Ket string) (sv.Penyakits, error) {
	req := sv.Penyakit{Keterangan: Ket}
	fmt.Println(req)
	resp, err := pe.ReadPenyakitByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	pen := resp.(sv.Penyakits)
	return pen, err
}
