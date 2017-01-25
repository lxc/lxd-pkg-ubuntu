package api

// NetworksPost represents the fields of a new LXD network
//
// API extension: network
type NetworksPost struct {
	NetworkPut `yaml:",inline"`

	Managed bool   `json:"managed"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

// NetworkPost represents the fields required to rename a LXD network
//
// API extension: network
type NetworkPost struct {
	Name string `json:"name"`
}

// NetworkPut represents the modifiable fields of a LXD network
//
// API extension: network
type NetworkPut struct {
	Config map[string]string `json:"config"`
}

// Network represents a LXD network
type Network struct {
	NetworkPut `yaml:",inline"`

	Name   string   `json:"name"`
	Type   string   `json:"type"`
	UsedBy []string `json:"used_by"`

	// API extension: network
	Managed bool `json:"managed"`
}

// Writable converts a full Network struct into a NetworkPut struct (filters read-only fields)
func (network *Network) Writable() NetworkPut {
	return network.NetworkPut
}
