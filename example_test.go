package restconf

import (
	"context"
	"fmt"
	"time"
)

func ExampleNew() {
	client := New(30 * time.Second)

	fmt.Println(client.httpClient.Timeout)
	// Output: 30s
}

func ExampleClient_NewLoopbackInterface() {
	server := newRestconfServer()
	defer server.Close()

	client := New(30 * time.Second)

	config := IetfInterfaceRequest{IetfInterface: IetfInterface{
		Name:    "loopback200",
		Type:    IetfInterfaceType,
		Enabled: true,
		IPv4: Ipv4{
			Address: []Address{{
				Ip:      "20.0.0.1",
				Netmask: "255.255.255.255",
			}}}}}

	response, _ := client.NewLoopbackInterface(context.Background(), config, Server{
		URL:      server.URL,
		username: "username",
		password: "password"})

	fmt.Println(response.StatusCode)
	// Output: 200
}
