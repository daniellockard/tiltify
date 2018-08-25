package tiltify_test

import (
	"testing"
)

func TestCauseGet(t *testing.T) {
	object, err := client.GetCause(35)
	if err != nil {
		t.Error("Could not get campaign for ID 35")
	}
	if object.Data.ID != 35 {
		t.Error("Could not get campaign for ID 35")
	}
}

func TestCauseGetBySlug(t *testing.T) {
	object, err := client.GetCauseBySlug("st-jude-children-s-research-hospital")
	if err != nil {
		t.Error("Could not get campaign for st-jude-children-s-research-hospital")
	}
	if object.Data.Slug != "st-jude-children-s-research-hospital" {
		t.Error("Could not get campaign for st-jude-children-s-research-hospital")
	}
}
