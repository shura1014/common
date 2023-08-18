package iputil

import (
	"encoding/binary"
	"net"
)

func GetIpArray() (ips []string, err error) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips, nil
}

func Ipv4ToUint32(ip string) uint32 {
	netIp := net.ParseIP(ip)
	if netIp == nil {
		return 0
	}
	// netIp.To4() 如果是一个ipv6形式的地址，使用该方法获取到ipv4
	return binary.BigEndian.Uint32(netIp.To4())
}

func Uint32ToIpv4(ip uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ip)
	return net.IP(ipByte).String()
}
