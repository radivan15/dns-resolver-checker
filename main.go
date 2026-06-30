package main

import (
	"context"
	"fmt"
	"net"
)

func main() {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}

	ips, err := resolver.LookupHost(context.Background(), "github.com")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(ips)
}
