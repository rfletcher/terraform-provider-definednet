package definednet

import (
	"encoding/json"
)

func (c *Client) Hosts() ([]Host, error) {
	body, err := c.get("hosts")

	if err != nil {
		return nil, err
	}

	response := HostsResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Hosts, nil
}
