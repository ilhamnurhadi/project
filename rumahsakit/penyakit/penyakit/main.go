package main

import (
	"fmt"

	// ini adalah tools dari dalam namanya local modul
	ep "rumahsakit/penyakit/penyakit/endpoint"
	pb "rumahsakit/penyakit/penyakit/grpc"
	svc "rumahsakit/penyakit/penyakit/server"

	// ini adalah tools tambahan namanya utility modul
	cfg "rumahsakit/penyakit/util/config"
	run "rumahsakit/penyakit/util/grpc"
	util "rumahsakit/penyakit/util/microservice"

	// ini adalah library dari luar namanya global modul
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// log ini digunakan dari utility,
	// dari sini
	logger := util.Logger()

	ok := cfg.AppConfig.LoadConfig()
	if !ok {
		logger.Log(util.LogError, "failed to load configuration")
		return
	}

	// ini konfigurasi yang pertama
	discHost := cfg.GetA("discoveryhost", "127.0.0.1:2181")
	ip := cfg.Get("serviceip", "127.0.0.1")
	port := cfg.Get("serviceport", "2181")
	address := fmt.Sprintf("%s:%s", ip, port)

	registrar, err := util.ServiceRegistry(discHost, svc.ServiceID, address, logger)
	if err != nil {
		logger.Log(util.LogError, "cannot find registrar")
		return
	}
	registrar.Register()
	defer registrar.Deregister()

	tracerHost := cfg.Get("tracerhost", "127.0.0.1:9999")
	tracer := util.Tracer(tracerHost)
	//sampe sini

	var server pb.PenyakitServiceServer
	{
		//chHost := cfg.Get("chhost", "127.0.0.1:6379")
		//cacher := svc.NewRedisCache(chHost)

		//gmapKey := cfg.Get("gmapkey", "AIzaSyD9tm3UVfxRWeaOy_MQ7tsCj1fVCLfG8Bo")
		//locator := svc.NewLocator(gmapKey)

		// ini konfigurasi yang kedua jika yang pertama gak nyambung jadi pake yang iini
		dbHost := cfg.Get(cfg.DBhost, "127.0.0.1:3306")
		dbName := cfg.Get(cfg.DBname, "rumahsakit")
		dbUser := cfg.Get(cfg.DBuid, "toor")
		dbPwd := cfg.Get(cfg.DBpwd, "toor")

		//brokers := cfg.GetA("mqbrokers", "127.0.0.1:9092")

		// svc = server untuk mau ke server harus ke folder server
		// Sebelum mulai code yang bawah pastikan servicenya selesai
		// kita harus kelarin disevicenya di folder server
		dbReadWriter := svc.NewDBReadWriter(dbHost, dbName, dbUser, dbPwd)
		//dbRuler := svc.NewDBDispatchRuler(dbReadWriter, locator)
		//notifier := mq.NewAsyncProducer(brokers, nil)

		//auctioneer := svc.NewAuctioneer(dbReadWriter, cacher)
		service := svc.NewPenyakit(dbReadWriter)
		endpoint := ep.NewPenyakitEndpoint(service)
		fmt.Println(endpoint)
		server = ep.NewGRPCPenyakitServer(endpoint, tracer, logger)
	}

	// ca := cfg.Get("capath", "cert/cacert.pem")
	// cert := cfg.Get("certpath", "cert/server.pem")
	// prikey := cfg.Get("keypath", "cert/server.key")
	//
	// tls, err := run.TLSCredentialFromFile(ca, cert, prikey, true)
	// if err != nil {
	// 	logger.Log("tlserr", err)
	// 	return
	// }
	//grpcServer := grpc.NewServer(append(run.Recovery(logger), grpc.Creds(tls))...)
	grpcServer := grpc.NewServer(run.Recovery(logger)...)
	pb.RegisterPenyakitServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	exit := make(chan bool, 1)
	go run.Serve(address, grpcServer, exit, logger)

	<-exit
}
