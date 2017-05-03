package govnsx

import (
	//"fmt"

	"github.com/go-resty/resty"
)

type Client struct {
	*resty.Request
}

type NsxManagerParams struct {
	UserName      string
	Password      string
	Uri           string
	AllowInsecssl bool
	RestAgentName string
}

// NewClient creates a new client from a URL.
func NewClient(mgrParams *NsxManagerParams) (*Client, error) {
	restconn := resty.R()
	c := &Client{
		Request: restconn,
	}
	c.Request.SetHeader("User-Agent", mgrParams.RestAgentName)
	return c, nil
}
