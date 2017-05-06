package govnsx

import (
	//"fmt"
	"crypto/tls"

	"github.com/go-resty/resty"
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

	c := &Client{
		Rclient:   rc,
		MgrConfig: mgrParams,
	}

	return c, nil
}
