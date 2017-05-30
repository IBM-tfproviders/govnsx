package govnsx

import (
	//"fmt"

	"github.com/go-resty/resty"
)

type Client struct {
	Req	*resty.Request
}

type NsxManagerConfig	struct {
	UserName      string
	Password      string
	Uri           string
	AllowInsecssl bool
	UserAgentName string
}

// NewClient creates a new client from a URL.
func NewClient(mgrParams *NsxManagerConfig) (*Client, error) {
	restconn := resty.R()
	c := &Client{
		Req: restconn,
	}

	c.Req.SetHeader("User-Agent", mgrParams.UserAgentName)

	//ToDo: need to store the NsxManagerConfig 
	//Some of the NSX manager config can ne set into resty.Request
	//Remaining need to be stored in Client struct

	return c, nil
}
