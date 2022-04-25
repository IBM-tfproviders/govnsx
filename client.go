package govnsx

import (
	//"fmt"
	"crypto/tls"

	"gopkg.in/resty.v1"
)

type NsxManagerConfig struct {
	UserName      string
	Password      string
	Uri           string
	AllowInsecssl bool
	UserAgentName string
}

type Client struct {
	Rclient   *resty.Client
	MgrConfig *NsxManagerConfig
}

// NewClient creates a new client from a URL.
func NewClient(mgrParams *NsxManagerConfig) (*Client, error) {
	rc := resty.New()

	rc.SetHeader("User-Agent", mgrParams.UserAgentName)
	if mgrParams.AllowInsecssl {
		rc.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	}

	rc.SetBasicAuth(mgrParams.UserName, mgrParams.Password)
	rc.SetHeader("Content-Type", "application/xml")

	c := &Client{
		Rclient:   rc,
		MgrConfig: mgrParams,
	}

	return c, nil
}
