package tiltify_test

import "testing"

func TestTeamsGet(t *testing.T) {
	t.Skip("Not Yet Implemented")
	object, err := client.GetTeams()
	if err != nil || object == nil {
		t.Error("Could not get Teams listing")
	}
}

func TestTeamGet(t *testing.T) {
	object, err := client.GetTeam(35)
	if err != nil || object == nil || object.Data.ID != 35 {
		t.Error("Could not get Team 35")
	}
}

func TestTeamGetCampaigns(t *testing.T) {
	t.Skip("Doesn't work for public users?")
	object, err := client.GetTeamCampaigns(35)
	if err != nil || object == nil {
		t.Error("Could not get Team 35 Campaigns")
	}
}

func TestTeamGetCampaign(t *testing.T) {
	t.Skip("Doesn't work for public users?")
	object, err := client.GetTeamCampaign(35, 1)
	if err != nil || object == nil {
		t.Error("Could not get Team 35 Campaign 1")
	}
}
