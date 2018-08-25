package tiltify_test

import (
	"os"
	"testing"

	"github.com/daniellockard/tiltify"
)

var client = tiltify.Client{}

func TestMain(m *testing.M) {
	client.SetAuthKey("AUTH_KEY_HERE")
	client.SetURL("https://tiltify.com/api/v3/")
	os.Exit(m.Run())
}

func TestCampaignGet(t *testing.T) {
	_, err := client.GetCampaign(4814)
	if err != nil {
		t.Error("Could not get campaign for ID 4814")
	}
}

func TestCampaignDonationsGet(t *testing.T) {
	_, err := client.GetCampaignDonations(4814)
	if err != nil {
		t.Error("Could not get donations for campaign 4814")
	}
}

func TestCampaignGetBadCampaignDonations(t *testing.T) {
	_, err := client.GetCampaignDonations(42)
	if err == nil {
		t.Error("Campaign that shouldn't return did")
	}
}

func TestCampaignRewardsGet(t *testing.T) {
	_, err := client.GetCampaignRewards(4814)
	if err != nil {
		t.Error("Could not get rewards for campaign 4814")
	}
}

func TestCampaignGetBadRewards(t *testing.T) {
	_, err := client.GetCampaignRewards(42)
	if err == nil {
		t.Error("Campaign that shouldn't return did")
	}
}
func TestCampaignPollsGet(t *testing.T) {
	_, err := client.GetCampaignPolls(4814)
	if err != nil {
		t.Error("Could not get polls for campaign 4814")
	}
}

func TestCampaignChallengesGet(t *testing.T) {
	_, err := client.GetCampaignChallenges(4814)
	if err != nil {
		t.Error("Could not get challenges for campaign 4814")
	}
}

func TestCampaignSupportingCampaignsGet(t *testing.T) {
	_, err := client.GetCampaignSupportingCampaigns(4814)
	if err != nil {
		t.Error("Could not get challenges for campaign 4814")
	}
}

func TestCampaignScheduleGet(t *testing.T) {
	_, err := client.GetCampaignSchedule(4814)
	if err != nil {
		t.Error("Could not get schedule for campaign 4814")
	}
}

func TestCampaignGetBadCampaign(t *testing.T) {
	_, err := client.GetCampaign(42)
	if err == nil {
		t.Error("Campaign that shouldn't return did")
	}
}

func TestCampaignGetBadURL(t *testing.T) {
	client.SetURL("http://google.com/lolnope")
	_, err := client.GetCampaign(42)
	if err == nil {
		t.Error("Campaign that shouldn't return did")
	}
	client.SetURL("https://tiltify.com/api/v3/")
}

func TestCampaignGetBadScheme(t *testing.T) {
	client.SetURL("htxp://google.com/lolnope")
	_, err := client.GetCampaign(42)
	if err == nil {
		t.Error("Campaign that shouldn't return did")
	}
	client.SetURL("https://tiltify.com/api/v3/")
}

func TestCampaignGetBadOverall(t *testing.T) {
	client.SetURL("")
	_, err := client.GetCampaign(199999999999999999)
	if err == nil {
		t.Error("Campaign that shouldn't return did")
	}
	client.SetURL("https://tiltify.com/api/v3/")
}

func TestSetURL(t *testing.T) {
	client.SetURL("OkieDokie")
	if client.BaseURL != "OkieDokie" {
		t.Error("SetUrl didn't change URL")
	}
	client.SetURL("https://tiltify.com/api/v3/")
}
