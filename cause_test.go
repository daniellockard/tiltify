package tiltify_test

import (
	"testing"
)

func TestCauseGet(t *testing.T) {
	object, err := client.GetCause(35)
	if err != nil {
		t.Error("Could not get cause for ID 35")
	}
	if object.Data.ID != 35 {
		t.Error("Could not get cause for ID 35")
	}
}

func TestCauseGetBySlug(t *testing.T) {
	object, err := client.GetCauseBySlug("st-jude-children-s-research-hospital")
	if err != nil {
		t.Error("Could not get cause for st-jude-children-s-research-hospital")
	}
	if object.Data.Slug != "st-jude-children-s-research-hospital" {
		t.Error("Could not get cause for st-jude-children-s-research-hospital")
	}
}

func TestCauseCampaigns(t *testing.T) {
	t.Skip()
	object, err := client.GetCauseCampaigns(35)
	if err != nil {
		t.Error(err)
	}
	if object == nil || len(object.Data) == 0 || object.Data[0].CauseID != 35 {
		t.Error("Could not get causeCampaigns for ID 35")
	}
}

func TestCauseDonations(t *testing.T) {
	t.Skip()
	object, err := client.GetCauseDonations(35)
	if err != nil {
		t.Error(err)
	}
	if object == nil || len(object.Data) == 0 {
		t.Error("Could not get causeDonations for ID 35")
	}
}

func TestCauseEvents(t *testing.T) {
	object, err := client.GetCauseEvents(35)
	if err != nil {
		t.Logf("%#v", object)
		t.Error(err)
	}
	if object == nil {
		t.Error("Could not get causeEvents for ID ")
	}
}

func TestCauseLeaderboards(t *testing.T) {
	object, err := client.GetCauseLeaderboards(35)
	if err != nil {
		t.Error(err)
	}
	if object == nil {
		t.Error("Could not get causeLeaderboards for ID 35")
	}
}

func TestCauseVisibilityOptions(t *testing.T) {
	object, err := client.GetCauseVisibilityOptions(35)
	if err != nil {
		t.Error(err)
	}
	if object == nil || !object.Data.Donate.Visible {
		t.Error("Could not get causeVisibilityOptions for ID 35")
	}
}

func TestCausePermissons(t *testing.T) {
	object, err := client.GetCausePermissions(35)
	if err != nil {
		t.Error(err)
	}
	if object == nil || !object.Data.Activated {
		t.Error("Could not get causePermissions for ID 35")
	}
}
