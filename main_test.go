package main

import "testing"

func TestPointsSignature(t *testing.T) {

	if PointsSignature(2) != "pts" {
		t.Fatalf("Expected pts")
	}
	if PointsSignature(0) != "pts" {
		t.Fatalf("Expected pts")
	}
	if PointsSignature(1) != "pt" {
		t.Fatalf("Expected pt")
	}

}

func TestExtractTeamAndScore(t *testing.T) {
	cases := []struct {
		input string
		team  string
		score int
	}{
		{"Team One 1", "Team One", 1},
		{"Team Two A 34", "Team Two A", 34},
		{"TeamThree 3", "TeamThree", 3},
	}

	for _, tt := range cases {
		team, score := ExtractTeamAndScore(tt.input)
		if team != tt.team {
			t.Fatalf("Expected team name: %v but received %v\n", tt.team, team)
		}

		if score != tt.score {
			t.Fatalf("Expected team: %v to have a score of: %v , but have %v\n", tt.team, tt.score, score)
		}

	}

}

func TestCalculateTeamPoints(t *testing.T) {
	cases := []struct {
		t1         string
		t1Score    int
		t2         string
		t2Score    int
		t1Expected int
		t2Expected int
	}{
		{"Team One", 2, "Team Two", 2, 1, 1},
		{"Team One", 1, "Team Two", 3, 1, 4},
		{"Team One", 3, "Team Two", 0, 4, 4},
		{"Team One", 0, "Team Two", 0, 5, 5},
	}

	var teamsMap = make(TeamsMap)

	for _, tt := range cases {
		teamsMap = CalculateTeamPoints(teamsMap, tt.t1, tt.t1Score, tt.t2, tt.t2Score)
		if teamsMap[tt.t1] != tt.t1Expected {
			t.Fatalf("Expected t1 to have: %v points, but has %v", tt.t1Expected, teamsMap[tt.t1])
		} else {
			t.Log(tt.t1, tt.t1Expected)
		}
		if teamsMap[tt.t2] != tt.t2Expected {
			t.Fatalf("Expected t2 to have: %v points, but has %v", tt.t2Expected, teamsMap[tt.t2])
		} else {
			t.Log(tt.t2, tt.t2Expected)
		}
	}

}

func TestSortTeamStanding(t *testing.T) {
	cases := []struct {
		t1         string
		t1Score    int
		t2         string
		t2Score    int
		t1Expected int
		t2Expected int
	}{
		{"Team One", 2, "Team Two", 2, 1, 1},
		{"Team One", 1, "Team Two", 3, 1, 4},
		{"Team One", 3, "Team Two", 0, 4, 4},
		{"Team One", 0, "Team Two", 0, 5, 5},
		{"AB", 0, "AC", 0, 1, 1},
		{"ADb", 0, "ADa", 0, 1, 1},
	}

	var teamsMap = make(TeamsMap)

	for _, tt := range cases {
		teamsMap = CalculateTeamPoints(teamsMap, tt.t1, tt.t1Score, tt.t2, tt.t2Score)

	}

	teams := SortTeamStanding(teamsMap)
	for key := range teams {
		t.Log(teams[key])
	}
	if teams[2].Name != "AB" {
		t.Fatalf("Expected AB at teams[2], but found: %v", teams[2])
	}
	if teams[5].Name != "ADb" {
		t.Fatalf("Expected AB at teams[5], but found: %v", teams[5])
	}

}
