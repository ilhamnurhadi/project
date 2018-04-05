package main

import (
	"context"
	"fmt"
	"time"

	cli "rumahsakit/penyakit/penyakit/endpoint"
	//svc "rumahsakit/penyakit/penyakit/server"
	opt "rumahsakit/penyakit/util/grpc"
	util "rumahsakit/penyakit/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	// dicobanya 3x, timeoutnya, dan balancingnya
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCPenyakitClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	// Add penyakit
	//addPny := client.AddPenyakitService(context.Background(), svc.Penyakit{KodePenyakit: "PNY006", NamaPenyakit: "Demam", Deskripsi: "-"})
	//fmt.Println("Add penyakit:", addPny)

	//Get penyakit By Keterangan No
	ketPny, _ := client.ReadPenyakitByKeteranganService(context.Background(), "%K%")
	fmt.Println("Penyakit based on Kode:", ketPny)

	//List Penyakit
	//peny, _ := client.ReadPenyakitService(context.Background())
	//fmt.Println("All penyakit:", peny)

	//Update Penyakit
	//client.UpdatePenyakitService(context.Background(), svc.Penyakit{NamaPenyakit: "Kanker", Deskripsi: "-", Status: 0, Keterangan: "Kronis", KodePenyakit: "PNY001"})

}
