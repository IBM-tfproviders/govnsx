package nsxtypes

import (
        "encoding/xml"
)

const (
	EdgeDLRAddInterfacesUriFormat = "%s/api/4.0/edges/%s/interfaces/?action=patch"
	EdgeDLRDelAllInterfacesUriFormat = "%s/api/4.0/edges/%s/interfaces"
	EdgeDLRDelbyIndexInterfacesUriFormat = "%s/api/4.0/edges/%s/interfaces/?index=%s"
)

type Interface struct {
	Label           string         `xml:"label,omitempty"`
	Name            string         `xml:"name,omitempty"`
	AddressGroups   []AddressGroup `xml:"addressGroups>addressGroup,omitempty"`
        Mtu             string         `xml:"mtu,omitempty"`
        Type            string         `xml:"type,omitempty"`
	IsConnected     bool           `xml:"isConnected,omitempty"`
	IsSharedNetwork bool           `xml:"isSharedNetwork,omitempty"`
	Index           string         `xml:"index,omitempty"`
	ConnectedToId   string         `xml:"connectedToId"`
	ConnectedToName string         `xml:"connectedToName"`
}

type EdgeDLRInterfaces struct {
	XMLName      xml.Name      `xml:"interfaces"`
	Interfaces   []Interface   `xml:"interface"`
}

type EdgeDLRAddInterfacesSpec EdgeDLRInterfaces
type EdgeDLRAddInterfacesResp EdgeDLRInterfaces

func NewEdgeDLRAddInterfacesSpec() *EdgeDLRAddInterfacesSpec {
	return &EdgeDLRAddInterfacesSpec{}
}

func NewEdgeDLRAddInterfacesResp() *EdgeDLRAddInterfacesResp {
	return &EdgeDLRAddInterfacesResp{}
}
