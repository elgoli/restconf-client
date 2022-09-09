package restconf

import (
	"context"
	"fmt"
	"time"
)

func ExampleNew() {
	client := New(30*time.Second, "URL", "username", "password")

	fmt.Println(client.restconfURL)
	// Output: URL
}

func ExampleClient_NewLoopbackInterface() {
	agent := newRestconfAgent()
	defer agent.Close()

	client := New(30*time.Second, agent.URL, "username", "password")

	config := IetfInterfaceRequest{IetfInterface: IetfInterface{
		Name:    "loopback200",
		Type:    IetfInterfaceType,
		Enabled: true,
		IPv4: Ipv4{
			Address: []Address{{
				Ip:      "20.0.0.1",
				Netmask: "255.255.255.255",
			}}}}}

	response, _ := client.NewLoopbackInterface(context.Background(), config)

	fmt.Println(response.StatusCode)
	// Output: 200
}
