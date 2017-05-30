package nsxtypes

import (
	"encoding/xml"
)

type VirtualWireCreateSpec struct {
	XMLName          xml.Name `xml:"virtualWireCreateSpec"`
	Name             string   `xml:"name"`
	Description      string   `xml:"description"`
	TenantId         string   `xml:"tenantId"`
	ControlPlaneMode string   `xml: "controlPlaneMode"`
	GuestVlanAllowed string   `xml: "guestVlanAllowed"`
}

type VirtualWire struct {
	XMLName          xml.Name `xml:"virtualWire"`
	ObjectId         string   `xml:"objectId"`
	TenantId         string   `xml:"tenantId"`
	SwitchOId        string   `xml:"vdsContextWithBacking>switch>objectId"`
	SwitchName       string   `xml:"vdsContextWithBacking>switch>name"`
	VdnId            string   `xml:"vdnId"`
	GuestVlanAllowed string   `xml:"guestVlanAllowed"`
	ControlPlaneMode string   `xml:"controlPlaneMode"`
}

func NewVirtualWire() *VirtualWire {
	return &VirtualWire{
		ObjectId: "",
	}
}
