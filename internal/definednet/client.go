package definednet

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const BASE_URL string = "https://api.defined.net/v1"
const RECORDS_PER_PAGE string = "100"

type Client struct {
	apiKey     string
	baseUrl    string
	httpClient *http.Client
}

func NewClient(apiKey string) (*Client, error) {
	c := Client{
		apiKey:     apiKey,
		baseUrl:    BASE_URL,
		httpClient: &http.Client{},
	}

	return &c, nil
}

func (c *Client) get(path string, params url.Values) ([]byte, error) {
	return c.request(http.MethodGet, path, params)
}

func (c *Client) request(method string, path string, params url.Values) ([]byte, error) {
	params.Set("pageSize", RECORDS_PER_PAGE)

	url := fmt.Sprintf("%s/%s?%s", c.baseUrl, path, params.Encode())

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, nil
}
