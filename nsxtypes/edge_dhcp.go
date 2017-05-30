package nsxtypes

import (
	"encoding/xml"
)

const (
	EdgeDhcpUriFormat          = "%s/api/4.0/edges/%s/dhcp/config"
	EdgeDhcpAddIPPoolUriFormat = "%s/api/4.0/edges/%s/dhcp/config/ippools"
	EdgeDhcpDelIPPoolUriFormat = "%s/api/4.0/edges/%s/dhcp/config/ippools/%s"
)

type IPPool struct {
	XMLName             xml.Name `xml:"ipPool"`
	IPRange             string   `xml:"ipRange"`
	DefaultGw           string   `xml:"defaultGateway,omitempty"`
	SubnetMask          string   `xml:"subnetMask,omitempty"`
	DomainName          string   `xml:"domainName,omitempty"`
	PrimaryNameServer   string   `xml:"primaryNameServer,omitempty"`
	SecondaryNameServer string   `xml:"secondaryNameServer,omitempty"`
	LeaseTime           int      `xml:"leaseTime,omitempty"`
	AutoConfigureDNS    bool     `xml:"autoConfigureDNS,omitempty"`
}

type LoggingInfo struct {
	Enable   bool   `xml:"enable"`
	LogLevel string `xml:"logLevel"`
}

type ConfigDHCPServiceSpec struct {
	XMLName xml.Name    `xml:"dhcp"`
	IPPools []IPPool    `xml:"ipPools>ipPool"`
	Logging LoggingInfo `xml:"logging,omitempty"`
}

type AddIPPoolToDHCPServiceResp struct {
	Location string
}

func NewConfigDHCPServiceSpec() *ConfigDHCPServiceSpec {
	return &ConfigDHCPServiceSpec{}
}

func NewIPPool() *IPPool {
	return &IPPool{}
}
