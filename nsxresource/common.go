package nsxresource

import (
	"github.com/IBM-tfproviders/govnsx"
)

type Common struct {
	Nsxc     *govnsx.Client
	Location string
}

func NewCommon(c *govnsx.Client) Common {
	return Common{
		Nsxc: c,
	}
}
