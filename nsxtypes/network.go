package nsxtypes

import (
	"encoding/xml"
)

const (
	VirtualWireUriFormat = "%s/api/2.0/vdn/scopes/%s/virtualwires"
	NetworkLableFormat   = "vxw-%s-%s-sid-%s-%s"
	CpmUnicastMode       = "UNICAST_MODE"
	CpmHybridMode        = "HYBRID_MODE"
	CpmMulticastMode     = "MULTICAST_MODE"
)

type VWCreateSpec struct {
	XMLName          xml.Name `xml:"virtualWireCreateSpec"`
	Name             string   `xml:"name,omitempty"`
	Description      string   `xml:"description,omitempty"`
	TenantId         string   `xml:"tenantId"`
	ControlPlaneMode string   `xml:"controlPlaneMode,omitempty"`
	GuestVlanAllowed bool     `xml:"guestVlanAllowed,omitempty"`
}

type VWPostResp struct {
	Location       string
	VirtualWireOID string
}

type VirtualWire struct {
	XMLName          xml.Name `xml:"virtualWire"`
	ObjectId         string   `xml:"objectId"`
	Name             string   `xml:"name"`
	TenantId         string   `xml:"tenantId"`
	SwitchOId        string   `xml:"vdsContextWithBacking>switch>objectId"`
	SwitchName       string   `xml:"vdsContextWithBacking>switch>name"`
	VdnId            string   `xml:"vdnId"`
	GuestVlanAllowed bool     `xml:"guestVlanAllowed"`
	ControlPlaneMode string   `xml:"controlPlaneMode"`
	Description      string   `xml:"description"`
}

type UpdateVirtualWire struct {
	XMLName          xml.Name `xml:"virtualWire"`
	Name             string   `xml:"name,omitempty"`
	TenantId         string   `xml:"tenantId,omitempty"`
	ControlPlaneMode string   `xml:"controlPlaneMode,omitempty"`
	Description      string   `xml:"description,omitempty"`
}

func NewVWCreateSpec() *VWCreateSpec {
	return &VWCreateSpec{
		Name: "",
	}
}

func NewVirtualWire() *VirtualWire {
	return &VirtualWire{
		ObjectId: "",
	}
}

func NewUpdateVirtualWire() *UpdateVirtualWire {
	return &UpdateVirtualWire{
		Name: "",
	}
}
