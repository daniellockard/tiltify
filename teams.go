package tiltify

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetTeams() (*Teams, error) {
	url := fmt.Sprint(c.BaseURL + "teams")
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data Teams
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetTeam(teamNumber int) (*Team, error) {
	url := fmt.Sprintf(c.BaseURL+"teams/%d", teamNumber)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data Team
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetTeamCampaigns(teamNumber int) (*TeamCampaigns, error) {
	url := fmt.Sprintf(c.BaseURL+"teams/%d/campaigns", teamNumber)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data TeamCampaigns
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetTeamCampaign(teamNumber int, campaignNumber int) (*TeamCampaign, error) {
	url := fmt.Sprintf(c.BaseURL+"teams/%d/campaigns/%d", teamNumber, campaignNumber)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data TeamCampaign
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
