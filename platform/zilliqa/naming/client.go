package naming

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) LookupName(name string) (response ZNSResponse, err error) {
	err = c.Get(&response, "/"+name, nil)
	return
}
