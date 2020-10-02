package base_net

import (
	"net"
	"strings"
)

//获取本机所有的网卡名和对应的IP地址
func Ips() (map[string]string, error) {
	ips := make(map[string]string)
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			return nil, err
		}
		addresses, err := byName.Addrs()
		for _, v := range addresses {
			ips[byName.Name] = v.String()
		}
	}
	return ips, nil
}

//获取Linux的真实IP地址,windows下需要自己通过网卡名称进行过滤
func GetTrueLocalIp() string {
	ip := ""
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if strings.Contains(i.Name, "eth") || strings.Contains(i.Name, "ens") {
				byName, err := net.InterfaceByName(i.Name)
				if err == nil {
					addresses, err := byName.Addrs()
					if err == nil {
						for _, v := range addresses {
							ip = v.String()
							return ip
						}
					}
				}
			}
		}
	}
	return ip
}
