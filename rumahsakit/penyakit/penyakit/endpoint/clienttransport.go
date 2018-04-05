package endpoint

import (
	"context"
	"time"

	svc "rumahsakit/penyakit/penyakit/server"

	pb "rumahsakit/penyakit/penyakit/grpc"

	util "rumahsakit/penyakit/util/grpc"
	disc "rumahsakit/penyakit/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.PenyakitService"
)

func NewGRPCPenyakitClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.PenyakitService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addPenyakitEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddPenyakitEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addPenyakitEp = retry
	}

	var readPenyakitByKodeEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadPenyakitByKodeEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readPenyakitByKodeEp = retry
	}

	var readPenyakitByKeteranganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadPenyakitByKeteranganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readPenyakitByKeteranganEp = retry
	}

	var readPenyakitEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadPenyakitEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readPenyakitEp = retry
	}

	var updatePenyakitEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdatePenyakit, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updatePenyakitEp = retry
	}

	return PenyakitEndpoint{AddPenyakitEndpoint: addPenyakitEp, ReadPenyakitByKodeEndpoint: readPenyakitByKodeEp, ReadPenyakitByKeteranganEndpoint: readPenyakitByKeteranganEp,
		ReadPenyakitEndpoint: readPenyakitEp, UpdatePenyakitEndpoint: updatePenyakitEp}, nil
}

func encodeAddPenyakitRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Penyakit)
	return &pb.AddPenyakitReq{
		KodePenyakit: req.KodePenyakit,
		NamaPenyakit: req.NamaPenyakit,
		Deskripsi:    req.Deskripsi,
	}, nil
}

func encodeReadPenyakitByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Penyakit)
	return &pb.ReadPenyakitByKodeReq{Kode: req.KodePenyakit}, nil
}

func encodeReadPenyakitByKeteranganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Penyakit)
	return &pb.ReadPenyakitByKeteranganReq{Keterangan: req.Keterangan}, nil
}

func encodeReadPenyakitRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdatePenyakitRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Penyakit)
	return &pb.UpdatePenyakitReq{
		KodePenyakit: req.KodePenyakit,
		NamaPenyakit: req.NamaPenyakit,
		Deskripsi:    req.Deskripsi,
		Status:       req.Status,
		Keterangan:   req.Keterangan,
	}, nil
}

func decodePenyakitResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadPenyakitByKodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadPenyakitByKodeResp)
	return svc.Penyakit{
		KodePenyakit: resp.KodePenyakit,
		NamaPenyakit: resp.NamaPenyakit,
		Deskripsi:    resp.Deskripsi,
		Keterangan:   resp.Keterangan,
	}, nil
}

func decodeReadPenyakitByKeteranganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadPenyakitByKeteranganResp)
	var rsp svc.Penyakits

	for _, v := range resp.AllKeterangan {
		itm := svc.Penyakit{
			KodePenyakit: v.KodePenyakit,
			NamaPenyakit: v.NamaPenyakit,
			Deskripsi:    v.Deskripsi,
			Keterangan:   v.Keterangan,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadPenyakitResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadPenyakitResp)
	var rsp svc.Penyakits

	for _, v := range resp.Allkode {
		itm := svc.Penyakit{
			KodePenyakit: v.KodePenyakit,
			NamaPenyakit: v.NamaPenyakit,
			Deskripsi:    v.Deskripsi,
			Keterangan:   v.Keterangan}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddPenyakitEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddPenyakit",
		encodeAddPenyakitRequest,
		decodePenyakitResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddPenyakit")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddPenyakit",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadPenyakitByKodeEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadPenyakitByKode",
		encodeReadPenyakitByKodeRequest,
		decodeReadPenyakitByKodeResponse,
		pb.ReadPenyakitByKodeResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadPenyakitByKode")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadPenyakitByKode",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadPenyakitByKeteranganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadPenyakitByKeterangan",
		encodeReadPenyakitByKeteranganRequest,
		decodeReadPenyakitByKeteranganResponse,
		pb.ReadPenyakitByKeteranganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadPenyakitByKeterangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadPenyakitByKeterangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadPenyakitEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadPenyakit",
		encodeReadPenyakitRequest,
		decodeReadPenyakitResponse,
		pb.ReadPenyakitResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadPenyakit")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadPenyakit",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdatePenyakit(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdatePenyakit",
		encodeUpdatePenyakitRequest,
		decodePenyakitResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdatePenyakit")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdatePenyakit",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
