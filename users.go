package tiltify

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetCurrentUser() (*User, error) {
	url := fmt.Sprint(c.BaseURL + "user")
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data User
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetUsers() (*Users, error) {
	url := fmt.Sprint(c.BaseURL + "users")
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data Users
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetUser(userID int) (*User, error) {
	url := fmt.Sprintf(c.BaseURL+"users/%d", userID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data User
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetUserCampaigns(userID int) (*UserCampaigns, error) {
	url := fmt.Sprintf(c.BaseURL+"users/%d/campaigns", userID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data UserCampaigns
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetUserCampaign(userID int, campaignID int) (*UserCampaign, error) {
	url := fmt.Sprintf(c.BaseURL+"users/%d/campaigns/%d", userID, campaignID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data UserCampaign
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetUserTeams(userID int) (*UserTeams, error) {
	url := fmt.Sprintf(c.BaseURL+"users/%d/teams", userID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data UserTeams
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetUserOwnedTeams(userID int) (*UserOwnedTeams, error) {
	url := fmt.Sprintf(c.BaseURL+"users/%d/owned-teams", userID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data UserOwnedTeams
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
