package main

import (
	"errors"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name         string
	PlayersNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(homeTeam string, homeScore int, visitorTeam string, visitorScore int) error {

	_, ok := l.Wins[homeTeam]
	if !ok {
		return errors.New("team no register in league")
	}
	_, ok = l.Wins[visitorTeam]
	if !ok {
		return errors.New("team no register in league")
	}

	if homeScore == visitorScore {
		return nil
	}

	if homeScore > visitorScore {
		l.Wins[homeTeam]++
		return nil
	}

	l.Wins[visitorTeam]++
	return nil

}

func (l *League) Ranking() []string {
	type teamWins struct {
		name string
		wins int
	}

	rankedTeams := make([]teamWins, 0, len(l.Wins))
	for name, wins := range l.Wins {
		rankedTeams = append(rankedTeams, teamWins{name: name, wins: wins})
	}

	sort.Slice(rankedTeams, func(i, j int) bool {
		if rankedTeams[i].wins == rankedTeams[j].wins {
			return rankedTeams[i].name < rankedTeams[j].name
		}
		return rankedTeams[i].wins > rankedTeams[j].wins
	})

	ranking := make([]string, 0, len(rankedTeams))
	for _, team := range rankedTeams {
		ranking = append(ranking, team.name)
	}

	return ranking
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, team := range r.Ranking() {
		io.WriteString(w, team+"\n")
	}
}

func createLeague(teams []Team) League {
	wins := map[string]int{}
	for _, v := range teams {
		wins[v.Name] = 0
	}
	return League{Teams: teams, Wins: wins}
}

func main() {
	teams := []Team{
		{Name: "Real Madrid", PlayersNames: []string{"Vinicius", "Mbappe", "Valverde"}},
		{Name: "Barcelona", PlayersNames: []string{"Araujo", "Lamine", "Ferran"}},
		{Name: "Atletico de Madrid", PlayersNames: []string{"Antoaine", "Cholo", "Simeone"}},
		{Name: "Real Sociedad", PlayersNames: []string{"Oyarazabal", "Inigo", "Arda"}},
	}

	l := createLeague(teams)
	l.MatchResult("Real Madrid", 3, "Barcelona", 1)
	l.MatchResult("Atletico de Madrid", 2, "Real Sociedad", 0)
	l.MatchResult("Barcelona", 2, "Atletico de Madrid", 2) // Empate
	l.MatchResult("Real Sociedad", 1, "Real Madrid", 4)
	l.MatchResult("Real Madrid", 0, "Atletico de Madrid", 1)
	l.MatchResult("Barcelona", 3, "Real Sociedad", 0)
	l.MatchResult("Real Madrid", 3, "Real Sociedad", 0)

	RankPrinter(&l, os.Stdout)

}
