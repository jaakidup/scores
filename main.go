package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TeamsMap = map[string]int
type Team struct {
	Name   string
	Points int
}

func main() {
	inputFileName := flag.String("input", "input.txt", "Scores to check and sort")
	outputFileName := flag.String("output", "output.txt", "The score calculated and sorted")
	flag.Parse()

	teamsMap, err := ExtractTeamScoreMap(*inputFileName)
	if err != nil {
		log.Fatalln(err)
	}

	sortedTeams := SortTeamStanding(teamsMap)

	err = WriteToFile(*outputFileName, sortedTeams)
	if err != nil {
		log.Fatalln("Failed writing team results to file.")
	}

	fmt.Println("Completed the team calculations using: ")
	fmt.Printf("input:\t %v\n", *inputFileName)
	fmt.Printf("output:\t %v\n", *outputFileName)

}

// PointsSignature calculates whether plural is needed
func PointsSignature(points int) string {
	if points == 0 || points >= 2 {
		return "pts"
	}
	return "pt"
}

func WriteToFile(outputFile string, teams []Team) error {
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Failed to open file for writing")
	}
	defer file.Close()

	for _, team := range teams {
		_, err := file.WriteString(fmt.Sprintf("%v, %v %v\n", team.Name, team.Points, PointsSignature(team.Points)))
		if err != nil {
			return err
		}
	}
	return nil
}

// SortTeamStandings sorts the teams according to total points
// teams with equal points get sorted alphabetically
func SortTeamStanding(teamsMap TeamsMap) []Team {

	teams := []Team{}
	for name, points := range teamsMap {
		teams = append(teams, Team{Name: name, Points: points})
	}

	// Let's sort according to points first
	sort.SliceStable(teams, func(i, j int) bool {
		return teams[i].Points > teams[j].Points
	})

	// Let's sort the teams with the same points
	sort.SliceStable(teams, func(i, j int) bool {
		if teams[i].Points == teams[j].Points {
			if strings.Compare(teams[i].Name, teams[j].Name) == -1 {
				return true
			}
		}
		return false
	})

	return teams
}

func ExtractTeamScoreMap(inputFileName string) (TeamsMap, error) {

	var teamsMap = make(TeamsMap)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {

		teams := strings.Split(scanner.Text(), ",")

		t1, t1Score := ExtractTeamAndScore((teams[0]))
		t2, t2Score := ExtractTeamAndScore((teams[1]))

		teamsMap = CalculateTeamPoints(teamsMap, t1, t1Score, t2, t2Score)

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return teamsMap, nil
}

// CalculateTeamPoints takes a two team match and calculates the league standings based on scores
func CalculateTeamPoints(teamsMap TeamsMap, t1 string, t1Score int, t2 string, t2Score int) TeamsMap {
	if t1Score == t2Score {
		teamsMap[t1] = teamsMap[t1] + 1
		teamsMap[t2] = teamsMap[t2] + 1
	}
	if t1Score > t2Score {
		teamsMap[t1] = teamsMap[t1] + 3
	}
	if t1Score < t2Score {
		teamsMap[t2] = teamsMap[t2] + 3
	}
	if t1Score == 0 {
		teamsMap[t1] = teamsMap[t1] + 0
	}
	if t2Score == 0 {
		teamsMap[t2] = teamsMap[t2] + 0
	}
	return teamsMap
}

// ExtractTeamAndScore takes a team name and a score as a string combo
// then returns a seperated name including spaces plus a score
func ExtractTeamAndScore(input string) (string, int) {
	teamData := strings.Split(input, " ")
	team := strings.TrimLeft(strings.Join(teamData[0:len(teamData)-1], " "), " ")

	scoreString := teamData[len(teamData)-1]
	score, _ := strconv.Atoi(scoreString)

	return team, score
}
