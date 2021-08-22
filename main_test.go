package main

import "testing"

func TestMain(t *testing.M) {
	main()
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
