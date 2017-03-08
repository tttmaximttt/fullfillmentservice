package main

import (
	"os"

	"github.com/tttmaximttt/fullfillmentservice/service"
	"github.com/hudl/fargo"
	"fmt"
	"strconv"
)

func main() {
	// For a real app, you'd bind a user-provided service with eureka // credentials and URL.
	discovery := fargo.NewConn("http://localhost:8080/eureka/v2")

	port, err := strconv.Atoi((os.Getenv("PORT")))
	host := os.Getenv("HOST")

	if err != nil {
		fmt.Errorf("%s", err)
	}

	if len(string(port)) == 0 {
		port = 3000
	} else if (len(host)) == 0 {
		host = "localhost"
	}

	appInstance := fargo.Instance{
		HostName: "i-6543",
		Port: port,
		App: "FULLFILLMENT_APP",
		IPAddr: host,
		VipAddress: host,
		SecureVipAddress: host,
		DataCenterInfo: fargo.DataCenterInfo{ Name: fargo.MyOwn },
		Status: fargo.UP,
	}

	addressString := fmt.Sprintf("%s:%v", appInstance.IPAddr, appInstance.Port)

	server := service.NewServer()

	_, err = discovery.GetApp("FULLFILLMENT_APP")

	if err != nil {
		discovery.RegisterInstance(&appInstance)
	}

	if err != nil {
		fmt.Errorf("%s", err)
	}
	server.Run(addressString)
}