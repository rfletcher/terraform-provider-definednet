package definednet

type Host struct {
	ID              string       `json:"id"`
	OrganizationId  string       `json:"organizationId"`
	NetworkId       string       `json:"networkId"`
	RoleId          string       `json:"roleId"`
	Name            string       `json:"name"`
	IpAddress       string       `json:"ipAddress"`
	StaticAddresses []string     `json:"staticAddresses"`
	ListenPort      int          `json:"listenPort"`
	IsBlocked       bool         `json:"isBlocked"`
	IsLighthouse    bool         `json:"isLighthouse"`
	IsRelay         bool         `json:"isRelay"`
	CreatedAt       string       `json:"createdAt"`
	Metadata        HostMetadata `json:"metadata"`
	Tags            []string     `json:"tags"`
}

type HostMetadata struct {
	LastSeenAt      string `json:"lastSeenAt"`
	Version         string `json:"version"`
	Platform        string `json:"platform"`
	UpdateAvailable bool   `json:"updateAvailable"`
}

type HostsResponse struct {
	Hosts    []Host           `json:"data"`
	Metadata ResponseMetadata `json:"metadata"`
}

type ResponseMetadata struct {
	TotalCount  int                  `json:"totalCount"`
	HasNextPage bool                 `json:"hasNextPage"`
	HasPrevPage bool                 `json:"hasPrevPage"`
	NextCursor  string               `json:"nextCursor"`
	PrevCursor  string               `json:"prevCursor"`
	Page        ResponsePageMetadata `json:"page"`
}

type ResponsePageMetadata struct {
	Count int `json:"count"`
	Start int `json:"start"`
}
