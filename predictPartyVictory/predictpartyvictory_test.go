package main

import (
	"testing"
)

func TestPredictPartyVictory1(t *testing.T) {
	senate := "RD"
	result := predictPartyVictory(senate)
	expected := "Radiant"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestPredictPartyVictory2(t *testing.T) {
	senate := "RDD"
	result := predictPartyVictory(senate)
	expected := "Dire"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestPredictPartyVictory3(t *testing.T) {
	senate := "RRDDDD"
	result := predictPartyVictory(senate)
	expected := "Radiant"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestPredictPartyVictory4(t *testing.T) {
	senate := "DDRRR"
	result := predictPartyVictory(senate)
	expected := "Radiant"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
