package nsxtypes

import (
	"encoding/xml"
)

const (
	EdgeUriFormat    = "%s/api/4.0/edges/"
	EdgeUriLocFormat = "%s/api/4.0/edges/%s"
)

type Appliance struct {
	ResourcePoolId string `xml:"resourcePoolId"`
	DatastoreId    string `xml:"datastoreId"`
	Deployed       bool   `xml:"deployed,omitempty"`
}

type Appliances struct {
	AppliancesList   []Appliance `xml:"appliance"`
	DeployAppliances bool        `xml:"deployAppliances"`
	ApplianceSize    string      `xml:"applianceSize,omitempty"`
}

type AddressGroup struct {
	PrimaryAddress string `xml:"primaryAddress"`
	SubnetMask     string `xml:"subnetMask"`
}

type Vnic struct {
	Index         string         `xml:"index"`
	PortgroupId   string         `xml:"portgroupId,omitempty"`
	AddressGroups []AddressGroup `xml:"addressGroups>addressGroup,omitempty"`
	IsConnected   bool           `xml:"isConnected"`
	Mtu           string         `xml:"mtu,omitempty"`
	Type          string         `xml:"type,omitempty"`
}

type EdgeInstallSpec struct {
	XMLName     xml.Name   `xml:"edge"`
	Datacenter  string     `xml:"datacenterName"`
	Name        string     `xml:"name,omitempty"`
	Description string     `xml:"description,omitempty"`
	Type        string     `xml:"type,omitempty"`
	Tenant      string     `xml:"tenant,omitempty"`
	Fqdn        string     `xml:"fqdn,omitempty"`
	VseLogLevel string     `xml:"vseLogLevel,omitempty"`
	EnableAesni bool       `xml:"enableAesni,omitempty"`
	EnableFips  bool       `xml:"enableFips,omitempty"`
	Appliances  Appliances `xml:"appliances"`
	Vnics       []Vnic     `xml:"vnics>vnic,omitempty"`
	Features    Features   `xml:"features"`
}

type EdgePostResp struct {
	EdgeId   string
	Location string
}

type Features struct {
	Dhcp DHCPConfig `xml:"dhcp"`
}

type Edge struct {
	XMLName     xml.Name   `xml:"edge"`
	Id          string     `xml:"id"`
	Version     string     `xml:"version"`
	Description string     `xml:"description"`
	Status      string     `xml:"status"`
	Tenant      string     `xml:"tenant"`
	Name        string     `xml:"name"`
	Appliances  Appliances `xml:"appliances"`
	Vnics       []Vnic     `xml:"vnics>vnic"`
	Type        string     `xml:"type"`
	Features    Features   `xml:"features"`
}

func NewEdge() *Edge {
	return &Edge{
		Id: "",
	}
}
