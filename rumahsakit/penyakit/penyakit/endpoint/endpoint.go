package endpoint

import (
	"context"

	svc "rumahsakit/penyakit/penyakit/server"

	kit "github.com/go-kit/kit/endpoint"
)

type PenyakitEndpoint struct {
	AddPenyakitEndpoint              kit.Endpoint
	ReadPenyakitByKodeEndpoint       kit.Endpoint
	ReadPenyakitEndpoint             kit.Endpoint
	UpdatePenyakitEndpoint           kit.Endpoint
	ReadPenyakitByKeteranganEndpoint kit.Endpoint
}

func NewPenyakitEndpoint(service svc.PenyakitService) PenyakitEndpoint {
	addPenyakitEp := makeAddPenyakitEndpoint(service)
	readPenyakitByKodeEp := makeReadPenyakitByKodeEndpoint(service)
	readPenyakitEp := makeReadPenyakitEndpoint(service)
	updatePenyakitEp := makeUpdatePenyakitEndpoint(service)
	readPenyakitByKeteranganEp := makeReadPenyakitByKeteranganEndpoint(service)
	return PenyakitEndpoint{AddPenyakitEndpoint: addPenyakitEp,
		ReadPenyakitByKodeEndpoint:       readPenyakitByKodeEp,
		ReadPenyakitEndpoint:             readPenyakitEp,
		UpdatePenyakitEndpoint:           updatePenyakitEp,
		ReadPenyakitByKeteranganEndpoint: readPenyakitByKeteranganEp,
	}
}

func makeAddPenyakitEndpoint(service svc.PenyakitService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Penyakit)
		err := service.AddPenyakitService(ctx, req)
		return nil, err
	}
}

func makeReadPenyakitByKodeEndpoint(service svc.PenyakitService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Penyakit)
		result, err := service.ReadPenyakitByKodeService(ctx, req.KodePenyakit)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadPenyakitEndpoint(service svc.PenyakitService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadPenyakitService(ctx)
		return result, err
	}
}

func makeUpdatePenyakitEndpoint(service svc.PenyakitService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Penyakit)
		err := service.UpdatePenyakitService(ctx, req)
		return nil, err
	}
}

func makeReadPenyakitByKeteranganEndpoint(service svc.PenyakitService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Penyakit)
		result, err := service.ReadPenyakitByKeteranganService(ctx, req.Keterangan)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}
