package main

import (
	"errors"
	"fmt"
	"net"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	charonAddr, err := lookupServiceWithConsul("charon")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Charon url: %s\n", charonAddr)
}

func lookupServiceWithConsul(serviceName string) ([]string, error) {
	var addresses = make([]string, 0)

	config := consulapi.DefaultConfig()
	config.Address = "http://10.0.4.236:8500"

	consul, err := consulapi.NewClient(config)
	if err != nil {
		return addresses, err
	}

	services, err := consul.Agent().Services()
	if err != nil {
		return addresses, err
	}

	for _, svc := range services {
		if svc.Service == serviceName {
			address := fmt.Sprintf("%s:%d", svc.Address, svc.Port)
			addresses = append(addresses, address)
		}
	}
	return addresses, nil
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
