// Package tournament solves the Tournament problem on Exercism.
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/template"
)

// Tally creates a scoring table based on match results.
func Tally(reader io.Reader, writer io.Writer) error {
	tallies := make(map[string]*Team)
	scanner := bufio.NewScanner(reader)
	// read lines
	for scanner.Scan() {
		line := scanner.Text()
		// filter empty lines, comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// parse team1;team2;result
		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return fmt.Errorf("invalid line: %s", line)
		}
		team1 := parts[0]
		team2 := parts[1]
		result := parts[2]

		// create team entries if it's the first time we've seen it
		if _, ok := tallies[team1]; !ok {
			tallies[team1] = &Team{Name: team1}
		}
		if _, ok := tallies[team2]; !ok {
			tallies[team2] = &Team{Name: team2}
		}

		// update scores
		switch result {
		case "win":
			tallies[team1].Wins++
			tallies[team2].Losses++
		case "draw":
			tallies[team1].Draws++
			tallies[team2].Draws++
		case "loss":
			tallies[team1].Losses++
			tallies[team2].Wins++
		default:
			return fmt.Errorf("invalid result: %s", result)
		}
	}

	// extract tallies as a slice
	result := make([]Team, 0, len(tallies))
	for _, team := range tallies {
		result = append(result, *team)
	}

	// sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].Points() == result[j].Points() {
			// break ties by team name
			return result[i].Name < result[j].Name
		}
		// sort descending
		return result[i].Points() > result[j].Points()
	})

	// output
	template := template.Must(template.New("tally").
		Funcs(template.FuncMap{"mp": Team.Matches, "p": Team.Points}).
		Parse(`Team                           | MP |  W |  D |  L |  P{{range .}}
{{ printf "%-30s" .Name }} |  {{mp .}} |  {{.Wins}} |  {{.Draws}} |  {{.Losses}} |  {{p .}}{{ end}}
`))
	template.Execute(writer, result)
	return nil
}

// Team represents a team's name and results.
type Team struct {
	Name                string
	Wins, Draws, Losses int
}

// Points returns the total number of points for a team.
func (t Team) Points() int {
	return t.Wins*3 + t.Draws
}

// Matches returns a team's total number of matches played.
func (t Team) Matches() int {
	return t.Wins + t.Draws + t.Losses
}
