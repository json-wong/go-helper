package helper

import (
	"errors"
	"net"
	"strings"
)

var contentType = "application/json"
var userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36"

// GetHostByName
// Get the IPv4 address corresponding to a given Internet host name
func GetHostByName(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}

// GetHostByAddr
// Get the Internet host name corresponding to a given IP address
func GetHostByAddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// getMacAddress
func getMacAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}

	mac, macErr := "", errors.New("cannot found mac address")
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags&net.FlagUp) == 0 || (netInterfaces[i].Flags&net.FlagLoopback) != 0 {
			continue
		}

		adders, _ := netInterfaces[i].Addrs()
		for _, address := range adders {
			ipNet, ok := address.(*net.IPNet)
			if !ok || !ipNet.IP.IsGlobalUnicast() {
				continue
			}

			mac = netInterfaces[i].HardwareAddr.String()
			return mac, nil
		}
	}

	return mac, macErr
}
