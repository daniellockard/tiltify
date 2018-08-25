package tiltify

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetFundraisingEvent(eventID int) (*FundraisingEvent, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEvent
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventCampaigns(eventID int) (*FundraisingEventCampaigns, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/campaigns", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventCampaigns
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventDonations(eventID int) (*FundraisingEventDonations, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/donations", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventDonations
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventIncentives(eventID int) (*FundraisingEventIncentives, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/incentives", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventIncentives
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventLeaderboards(eventID int) (*FundraisingEventLeaderboards, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/leaderboards", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventLeaderboards
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventRegistrations(eventID int) (*FundraisingEventRegistratons, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/registrations", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventRegistratons
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventRegistrationFields(eventID int) (*FundraisingEventRegistrationFields, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/registration-fields", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventRegistrationFields
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventSchedule(eventID int) (*FundraisingEventSchedule, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/schedule", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventSchedule
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetFundraisingEventVisibilityOptions(eventID int) (*FundraisingEventVisibilityOptions, error) {
	url := fmt.Sprintf(c.BaseURL+"fundraising-events/%d/visibility-options", eventID)
	fmt.Println("Making request to " + url)
	bytes, err := c.SetupAndDo(url)
	if err != nil {
		return nil, err
	}
	var data FundraisingEventVisibilityOptions
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
