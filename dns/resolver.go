package dns

import (
	"context"
	"net"
	"strings"
	"time"
)

type Resolver struct {
	resolver *net.Resolver
}

func NewResolver(config Config) *Resolver {
	return &Resolver{
		resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Duration(config.Timeout) * time.Second,
				}
				return d.DialContext(ctx, config.Network, config.Address)
			},
		},
	}
}

func (r *Resolver) LookupIPv6s(ctx context.Context, hosts []string) []string {
	var ips []string
	for _, host := range hosts {
		hostIPs, _ := r.resolver.LookupHost(ctx, host)
		for _, hostIP := range hostIPs {
			if net.ParseIP(hostIP) != nil && strings.Contains(hostIP, ":") {
				ips = append(ips, hostIP)
			}
		}
	}
	return ips
}
