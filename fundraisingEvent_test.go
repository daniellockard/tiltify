package tiltify_test

import "testing"

func TestFundraisingEventGet(t *testing.T) {
	object, err := client.GetFundraisingEvent(1)
	if err != nil {
		t.Error("Could not get fundraising-event for ID 1")
	}
	if object.Data.ID != 1 {
		t.Error("Could not get fundraising-event for ID 1")
	}
}

func TestFundraisingEventCampaignsGet(t *testing.T) {
	t.Skip("Skipping because private")
	object, err := client.GetFundraisingEventCampaigns(1)
	if err != nil || object == nil {
		t.Error("Could not get fundraising-event campaigns for ID 1")
	}
}

func TestFundraisingEventDonations(t *testing.T) {
	t.Skip("Skipping because private")
	object, err := client.GetFundraisingEventDonations(1)
	if err != nil || object == nil {
		t.Error("Could not get fundraising-event donations for ID 1")
	}
}

func TestFundraisingEventIncentives(t *testing.T) {
	t.Skip("Skipping because private")
	object, err := client.GetFundraisingEventIncentives(1)
	if err != nil || object == nil {
		t.Error("Could not get fundraising-event incentives for ID 1")
	}
}

func TestFundraisingEventLeaderboards(t *testing.T) {
	object, err := client.GetFundraisingEventLeaderboards(1)
	if err != nil || object == nil {
		t.Error("Could not get fundraising-event leaderboards for ID 1")
	}
}

func TestFundraisingEventRegistrations(t *testing.T) {
	object, err := client.GetFundraisingEventRegistrations(1)
	if err != nil || object == nil {
		t.Error("Could not get fundraising-event registrations for ID 1")
	}
}

func TestFundraisingEventRegistrationFields(t *testing.T) {
	t.Skip("Not Yet Implemented")
	object, err := client.GetFundraisingEventRegistrationFields(1)
	if err != nil || object == nil {
		t.Error(err)
		t.Error("Could not get fundraising-event registration fields for ID 1")
	}
}

func TestFundraisingEventSchedule(t *testing.T) {
	t.Skip("Not Yet Implemented")
	object, err := client.GetFundraisingEventSchedule(1)
	if err != nil || object == nil {
		t.Error(err)
		t.Error("Could not get fundraising-event schedule fields for ID 1")
	}
}

func TestFundraisingEventVisibilityOptions(t *testing.T) {
	t.Skip("Not Yet Implemented")
	object, err := client.GetFundraisingEventVisibilityOptions(1)
	if err != nil || object == nil {
		t.Error(err)
		t.Error("Could not get fundraising-event visibility-options fields for ID 1")
	}
}
