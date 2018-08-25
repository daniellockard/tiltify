package tiltify_test

import "testing"

func TestGetCurrentUser(t *testing.T) {
	t.Skip("Doesn't work for public users")
	object, err := client.GetCurrentUser()
	if err != nil || object == nil {
		t.Error("Could not get current user")
	}
}

func TestGetUsers(t *testing.T) {
	t.Skip("Does not work for public users")
	object, err := client.GetUsers()
	if err != nil || object == nil {
		t.Error("Could not get user list")
	}
}

func TestGetUser(t *testing.T) {
	object, err := client.GetUser(1)
	if err != nil || object == nil || object.Data.ID != 1 {
		t.Error("Could not get user 1")
	}
}

func TestGetUserCampaigns(t *testing.T) {
	object, err := client.GetUserCampaigns(1)
	if err != nil || object == nil {
		t.Error("Could not get user 1")
	}
}

func TestGetUserCampaign(t *testing.T) {
	t.Skip("Don't know a valid ID combo here")
	object, err := client.GetUserCampaign(1, 1)
	if err != nil || object == nil {
		t.Error("Could not get user 1, campaign 1")
	}
}

func TestGetUserTeams(t *testing.T) {
	object, err := client.GetUserTeams(1)
	if err != nil || object == nil {
		t.Error("Could not get teams for user 1")
	}
}

func TestGetUserOwnedTeams(t *testing.T) {
	object, err := client.GetUserOwnedTeams(1)
	if err != nil || object == nil {
		t.Error("Could not get owned teams for user 1")
	}
}
