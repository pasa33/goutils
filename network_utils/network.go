package network_utils

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/pasa33/goutils/string_utils"
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

// pass proxy as https://user:pass@host:port
func GetProxyIP(proxy string) (string, error) {

	prxUrl, err := url.Parse(proxy)
	if err != nil {
		return "", fmt.Errorf("invalid proxy: %w", err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(prxUrl),
	}
	client := &http.Client{
		Transport: transport,
	}

	res, err := client.Get("https://api64.ipify.org")
	if err != nil {
		return "", fmt.Errorf("req error: %w", err)
	}
	defer res.Body.Close()

	bbody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("resp error: %w", err)
	}
	return string(bbody), nil
}

func GenerateFakeIp() string {
	buf := make([]byte, 4)
	ip := rand.Uint32()
	binary.LittleEndian.PutUint32(buf, ip)
	return string(net.IP(buf))
}

// HOST:PORT:USER:PASSWORD
// USER:PASSWORD:HOST:PORT
// HOST:PORT@USER:PASSWORD
// USER:PASSWORD@HOST:PORT
func ProxyFormatter(plain string) string {

	plain = strings.ReplaceAll(plain, "@", ":")
	sP := strings.Split(plain, ":")

	if len(sP) == 2 { //host:port
		return fmt.Sprintf("http://%s:%s", sP[0], sP[1])
	}

	if len(sP) == 4 {
		if strings.Contains(sP[0], ".") && string_utils.IsDigit(sP[1]) {
			return fmt.Sprintf("http://%s:%s@%s:%s", sP[2], sP[3], sP[0], sP[1])
		}
		if strings.Contains(sP[2], ".") && string_utils.IsDigit(sP[3]) {
			return fmt.Sprintf("http://%s:%s@%s:%s", sP[0], sP[1], sP[2], sP[3])
		}
	}

	return ""
}
