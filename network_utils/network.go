package network_utils

import (
	"encoding/binary"
	"io"
	"math/rand"
	"net"
	"net/http"
)

func MyLocalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
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
			return ""
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
			return ip.String()
		}
	}
	return ""
}

func MyPublicIP() string {
	res, err := http.Get("https://api64.ipify.org")
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	bbody, err := io.ReadAll(res.Body)
	if err != nil {
		return ""
	}
	return string(bbody)
}

func GenerateFakeIp() string {
	buf := make([]byte, 4)
	ip := rand.Uint32()
	binary.LittleEndian.PutUint32(buf, ip)
	return string(net.IP(buf))
}
