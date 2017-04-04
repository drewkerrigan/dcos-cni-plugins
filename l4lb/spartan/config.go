package spartan

import (
	"net"

	"github.com/containernetworking/cni/pkg/types"
)

type SpartanIPAM struct {
	Type       string      `json:"type"`
	RangeStart net.IP      `json:"rangeStart"`
	RangeEnd   net.IP      `json:"rangeEnd"`
	Subnet     types.IPNet `json:"subnet"`
}

type SpartanNetwork struct {
	Name string      `json:"name"`
	IPAM SpartanIPAM `json:"ipam"`
}

// TODO(asridharan): This needs to be derived from the spartan
// configuration.
var SpartanIPs = []net.IPNet{
	net.IPNet{
		IP:   net.IPv4(198, 50, 100, 1),
		Mask: net.IPv4Mask(0xff, 0xff, 0xff, 0xff),
	},
	net.IPNet{
		IP:   net.IPv4(198, 50, 100, 2),
		Mask: net.IPv4Mask(0xff, 0xff, 0xff, 0xff),
	},
	net.IPNet{
		IP:   net.IPv4(198, 50, 100, 3),
		Mask: net.IPv4Mask(0xff, 0xff, 0xff, 0xff),
	},
}

// TODO(asridharan): This needs to be derived from the spartan
// configuration.
var SpartanConfig = SpartanNetwork{
	Name: "spartan-network",
	IPAM: SpartanIPAM{
		Type:       "host-local",
		RangeStart: net.IPv4(198, 50, 100, 10),
		RangeEnd:   net.IPv4(198, 50, 100, 253),
		Subnet: types.IPNet{
			IP:   net.IPv4(198, 50, 100, 0),
			Mask: net.IPv4Mask(0xff, 0xff, 0xff, 0),
		},
	},
}
