package nsxtypes

import (
	"encoding/xml"
)

const (
	EdgeUriFormat = "%s/api/4.0/edges/"
)

type Appliance struct {
	ResourcePoolId string `xml:"resourcePoolId"`
	DatastoreId    string `xml:"datastoreId"`
}

type AddressGroup struct {
	PrimaryAddress string `xml:"primaryAddress"`
	SubnetMask     string `xml:"subnetMask"`
}

type Vnic struct {
	Index         string         `xml:"index"`
	PortgroupId   string         `xml:"portgroupId"`
	AddressGroups []AddressGroup `xml:"addressGroups>addressGroup"`
	IsConnected   bool           `xml:"isConnected"`
}

type EdgeSGWInstallSpec struct {
	XMLName        xml.Name    `xml:"edge"`
	Name           string      `xml:"name"`
	Description    string      `xml:"description"`
	Tenant         string      `xml:"tenant"`
	AppliancesList []Appliance `xml:"appliances>appliance"`
	Vnics          []Vnic      `xml:"vnics>vnic"`
}

type EdgePostResp struct {
	EdgeId string
}

type Edge struct {
	XMLName        xml.Name    `xml:"edge"`
	Id             string      `xml:"id"`
	Version        string      `xml:"version"`
	Description    string      `xml:"description"`
	Status         string      `xml:"status"`
	Tenant         string      `xml:"tenant"`
	Name           string      `xml:"name"`
	AppliancesList []Appliance `xml:"appliances>appliance"`
	Vnics          []Vnic      `xml:"vnics>vnic"`
	Type           string      `xml:"type"`
}

func NewEdge() *Edge {
	return &Edge{
		Id: "",
	}
}
