package dns_lookup

import (
	"errors"
	"fmt"
	"net"
)

func ServiceDiscovery(serviceName string, namespace string) ([]string, error) {
	domain := fmt.Sprintf("%s.%s.svc.cluster.local", serviceName, namespace)
	ips, err := net.LookupHost(domain)
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, errors.New("empty addr")
	}
	addresses := make([]string, len(ips))
	for i, ip := range ips {
		addresses[i] = ip
	}
	return addresses, nil
}

func ServiceDiscoveryOverDocker(serviceName string) ([]string, error) {
	domain := fmt.Sprintf("%v", serviceName)
	ips, err := net.LookupHost(domain)
	if err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, errors.New("empty addr")
	}
	addresses := make([]string, len(ips))
	for i, ip := range ips {
		addresses[i] = ip
	}
	return addresses, nil
}
