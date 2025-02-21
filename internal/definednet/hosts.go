package definednet

import (
	"encoding/json"
	"net/url"
)

func (c *Client) Hosts() ([]Host, error) {
	hosts := []Host{}

	cursor := ""
	for {
		body, err := c.get("hosts", url.Values{
			"cursor": []string{cursor},
		})

		if err != nil {
			return nil, err
		}

		response := HostsResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		hosts = append(hosts, response.Hosts...)

		if !response.Metadata.HasNextPage {
			break
		}

		cursor = response.Metadata.NextCursor
	}

	return hosts, nil
}
