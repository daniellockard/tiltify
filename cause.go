package tiltify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetCause(causeNumber int) (*Cause, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Cause
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCauseBySlug(causeSlug string) (*Cause, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%s", causeSlug)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Cause
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCauseCampaigns(causeNumber int) (*CauseCampaigns, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d/campaigns", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CauseCampaigns
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCauseDonations(causeNumber int) (*CauseDonations, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d/donations", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CauseDonations
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCauseEvents(causeNumber int) (*CauseEvents, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d/fundraising-events", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CauseEvents
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCauseLeaderboards(causeNumber int) (*CauseLeaderboard, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d/leaderboards", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CauseLeaderboard
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCauseVisibilityOptions(causeNumber int) (*CauseVisibilityOptions, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d/visibility-options", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CauseVisibilityOptions
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// I have no idea if this works or what the response is
func (c *Client) PatchCauseVisibilityOptions(causeNumber int, patchData []byte) (*CauseVisibilityOptions, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d/visibility-options", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(patchData))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CauseVisibilityOptions
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCausePermissions(causeNumber int) (*CausePermissions, error) {
	url := fmt.Sprintf(c.BaseURL+"causes/%d/visibility-options", causeNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CausePermissions
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
