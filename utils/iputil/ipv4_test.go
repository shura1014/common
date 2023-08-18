package iputil

import "testing"

func TestIpv4ToUint32(t *testing.T) {
	ip := "192.168.1.1"
	ipInt := Ipv4ToUint32(ip)
	t.Log(ipInt)
	t.Log(Uint32ToIpv4(ipInt))
}
