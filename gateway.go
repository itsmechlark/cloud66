package cloud66

import (
	"fmt"
	"time"
)

type Gateway struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Address   string    `json:"address"`
	Ttl       string    `json:"ttl"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at_iso"`
	UpdatedAt time.Time `json:"updated_at_iso"`
}

func (c *Client) ListGateways(accountId int) ([]Gateway, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/accounts/%d/gateways.json", accountId), nil, nil)
	if err != nil {
		return nil, err
	}

	var result []Gateway
	return result, c.DoReq(req, &result, nil)
}

func (c *Client) AddGateway(accountId int, name string, address string, username string) error {
	params := struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Username string `json:"username"`
	}{
		Name:     name,
		Address:  address,
		Username: username,
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("/accounts/%d/gateways.json", accountId), params, nil)
	if err != nil {
		return err
	}

	err = c.DoReq(req, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

/*
func (c *Client) RemoveGatewayKey(accountId int) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("/accounts/%d/gateway_key/0.json", accountId), nil, nil)
	if err != nil {
		return err
	}

	err = c.DoReq(req, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
*/
