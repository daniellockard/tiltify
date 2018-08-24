package tiltify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AuthKey string
	BaseURL string
}

func (c *Client) SetURL(url string) {
	c.BaseURL = url
}

func (c *Client) SetAuthKey(key string) {
	c.AuthKey = key
}

func (c *Client) GetCampaign(campaignNumber int) (*Campaign, error) {
	url := fmt.Sprintf(c.BaseURL+"campaigns/%d", campaignNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Campaign
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCampaignDonations(campaignNumber int) (*Donations, error) {
	url := fmt.Sprintf(c.BaseURL+"campaigns/%d/donations", campaignNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Donations
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCampaignRewards(campaignNumber int) (*Rewards, error) {
	url := fmt.Sprintf(c.BaseURL+"campaigns/%d/rewards", campaignNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Rewards
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCampaignPolls(campaignNumber int) (*Polls, error) {
	url := fmt.Sprintf(c.BaseURL+"campaigns/%d/polls", campaignNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Polls
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCampaignChallenges(campaignNumber int) (*Challenges, error) {
	url := fmt.Sprintf(c.BaseURL+"campaigns/%d/challenges", campaignNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Challenges
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCampaignSchedule(campaignNumber int) (*Schedule, error) {
	url := fmt.Sprintf(c.BaseURL+"campaigns/%d/schedule", campaignNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Schedule
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetCampaignSupportingCampaigns(campaignNumber int) (*SupportingCampaigns, error) {
	url := fmt.Sprintf(c.BaseURL+"campaigns/%d/schedule", campaignNumber)
	fmt.Println("Making request to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data SupportingCampaigns
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	req.Header.Add("Authorization", "Bearer "+c.AuthKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
