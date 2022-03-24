package main

import (
	"DnsLog/Core"
	"DnsLog/Dns"
	"DnsLog/Http"
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
	"strings"
)

//GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-w -s" main.go

func main() {
	var err = gcfg.ReadFileInto(&Core.Config, "./config.ini")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, v := range strings.Split(Core.Config.HTTP.Token, ",") {
		Core.User[v] = Core.GetRandStr()
		fmt.Printf(Core.User[v])
	}
	go Dns.ListingDnsServer()
	Http.ListingHttpManagementServer()
}
