package endpoint

import (
	"context"

	scv "rumahsakit/penyakit/penyakit/server"

	pb "rumahsakit/penyakit/penyakit/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcPenyakitServer struct {
	addPenyakit              grpctransport.Handler
	readPenyakitByKode       grpctransport.Handler
	readPenyakit             grpctransport.Handler
	updatePenyakit           grpctransport.Handler
	readPenyakitByKeterangan grpctransport.Handler
}

func NewGRPCPenyakitServer(endpoints PenyakitEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.PenyakitServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcPenyakitServer{
		addPenyakit: grpctransport.NewServer(endpoints.AddPenyakitEndpoint,
			decodeAddPenyakitRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddPenyakit", logger)))...),
		readPenyakitByKode: grpctransport.NewServer(endpoints.ReadPenyakitByKodeEndpoint,
			decodeReadPenyakitByKodeRequest,
			encodeReadPenyakitByKodeResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadPenyakitByKode", logger)))...),
		readPenyakit: grpctransport.NewServer(endpoints.ReadPenyakitEndpoint,
			decodeReadPenyakitRequest,
			encodeReadPenyakitResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadPenyakit", logger)))...),
		updatePenyakit: grpctransport.NewServer(endpoints.UpdatePenyakitEndpoint,
			decodeUpdatePenyakitRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdatePenyakit", logger)))...),
		readPenyakitByKeterangan: grpctransport.NewServer(endpoints.ReadPenyakitByKeteranganEndpoint,
			decodeReadPenyakitByKeteranganRequest,
			encodeReadPenyakitByKeteranganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadPenyakitByKeterangan", logger)))...),
	}
}

func decodeAddPenyakitRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddPenyakitReq)
	return scv.Penyakit{KodePenyakit: req.KodePenyakit, NamaPenyakit: req.NamaPenyakit, Deskripsi: req.Deskripsi}, nil
}

func decodeReadPenyakitByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadPenyakitByKodeReq)
	return scv.Penyakit{KodePenyakit: req.Kode}, nil
}

func decodeReadPenyakitRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdatePenyakitRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdatePenyakitReq)
	return scv.Penyakit{KodePenyakit: req.KodePenyakit, NamaPenyakit: req.NamaPenyakit, Deskripsi: req.Deskripsi, Status: req.Status, Keterangan: req.Keterangan}, nil
}

func decodeReadPenyakitByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadPenyakitByKeteranganReq)
	return scv.Penyakit{Keterangan: req.Keterangan}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadPenyakitByKodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Penyakit)
	return &pb.ReadPenyakitByKodeResp{KodePenyakit: resp.KodePenyakit, NamaPenyakit: resp.NamaPenyakit, Deskripsi: resp.Deskripsi, Keterangan: resp.Keterangan}, nil
}

func encodeReadPenyakitResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Penyakits)

	rsp := &pb.ReadPenyakitResp{}

	for _, v := range resp {
		itm := &pb.ReadPenyakitByKodeResp{
			KodePenyakit: v.KodePenyakit,
			NamaPenyakit: v.NamaPenyakit,
			Deskripsi:    v.Deskripsi,
			Keterangan:   v.Keterangan,
		}
		rsp.Allkode = append(rsp.Allkode, itm)
	}
	return rsp, nil
}

func encodeReadPenyakitByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Penyakits)

	rsp := &pb.ReadPenyakitByKeteranganResp{}

	for _, v := range resp {
		itm := &pb.ReadPenyakitByKodeResp{
			KodePenyakit: v.KodePenyakit,
			NamaPenyakit: v.NamaPenyakit,
			Deskripsi:    v.Deskripsi,
			Keterangan:   v.Keterangan,
		}
		rsp.AllKeterangan = append(rsp.AllKeterangan, itm)
	}
	return rsp, nil
}

func (s *grpcPenyakitServer) AddPenyakit(ctx oldcontext.Context, penyakit *pb.AddPenyakitReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addPenyakit.ServeGRPC(ctx, penyakit)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcPenyakitServer) ReadPenyakitByKode(ctx oldcontext.Context, kode *pb.ReadPenyakitByKodeReq) (*pb.ReadPenyakitByKodeResp, error) {
	_, resp, err := s.readPenyakitByKode.ServeGRPC(ctx, kode)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadPenyakitByKodeResp), nil
}

func (s *grpcPenyakitServer) ReadPenyakitByKeterangan(ctx oldcontext.Context, keterangan *pb.ReadPenyakitByKeteranganReq) (*pb.ReadPenyakitByKeteranganResp, error) {
	_, resp, err := s.readPenyakitByKeterangan.ServeGRPC(ctx, keterangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadPenyakitByKeteranganResp), nil
}

func (s *grpcPenyakitServer) ReadPenyakit(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadPenyakitResp, error) {
	_, resp, err := s.readPenyakit.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadPenyakitResp), nil
}

func (s *grpcPenyakitServer) UpdatePenyakit(ctx oldcontext.Context, cus *pb.UpdatePenyakitReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updatePenyakit.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
