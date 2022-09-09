package restconf

// IetfInterfaceRequest is YANG data model for configuring a loopback interface
type IetfInterfaceRequest struct {
	IetfInterface IetfInterface `json:"ietf-interfaces:interface"`
}

type IetfInterface struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Enabled     bool   `json:"enabled"`
	IPv4        Ipv4   `json:"ietf-ip:ipv4"`
}
type Ipv4 struct {
	Address []Address `json:"address"`
}

type Address struct {
	Ip      string `json:"ip"`
	Netmask string `json:"netmask"`
}
