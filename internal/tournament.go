package internal

import (
	"fmt"
	"sort"
)

type TournamentManager struct {
	config     Config
	strategies []Strategy
	results    map[string]map[string]int // [P1][P2]Score
}

func NewTournamentManager(config Config) *TournamentManager {
	return &TournamentManager{
		config:  config,
		results: make(map[string]map[string]int),
	}
}

func (tm *TournamentManager) AddStrategy(strategy Strategy) {
	tm.strategies = append(tm.strategies, strategy)
}

func (tm *TournamentManager) PlayRound(s1, s2 Strategy, h1, h2 History) (Action, Action, int, int) {
	action1 := s1.Move(h1)
	action2 := s2.Move(h2)
	score1, score2 := tm.calculateScores(action1, action2)
	return action1, action2, score1, score2
}

func (tm *TournamentManager) calculateScores(action1, action2 Action) (int, int) {
	key := [2]Action{action1, action2}
	if val, ok := tm.config.PayoffMatrix[key]; ok {
		return val[0], val[1]
	}
	return 0, 0
}

func (tm *TournamentManager) PlayMatch(s1, s2 Strategy) (int, int) {
	var h1 History
	var h2 History
	score1, score2 := 0, 0

	for round := 0; round < tm.config.Rounds; round++ {
		action1, action2, s1score, s2score := tm.PlayRound(s1, s2, h1, h2)
		h1 = append(h1, action1)
		h2 = append(h2, action2)
		score1 += s1score
		score2 += s2score
	}

	return score1, score2
}

func (tm *TournamentManager) PlayTournament() {
	for _, s := range tm.strategies {
		tm.results[s.Name()] = make(map[string]int)
	}

	for i := 0; i < len(tm.strategies); i++ {
		for j := i + 1; j < len(tm.strategies); j++ {
			s1 := tm.strategies[i]
			s2 := tm.strategies[j]

			score1, score2 := tm.PlayMatch(s1, s2)

			tm.results[s1.Name()][s2.Name()] = score1
			tm.results[s2.Name()][s1.Name()] = score2
		}
	}
}

func (tm *TournamentManager) PrintResults() {
	if len(tm.strategies) == 0 {
		println("Нет первого среди равных... равных нулю, хах")
		return
	}

	names := make([]string, len(tm.strategies))
	for i, s := range tm.strategies {
		names[i] = s.Name()
	}

	scores := make(map[string]int)
	for _, s1 := range names {
		total := 0
		for _, s2 := range names {
			total += tm.results[s1][s2]
		}
		scores[s1] = total
	}

	sort.Slice(tm.strategies, func(i, j int) bool {
		return scores[tm.strategies[i].Name()] > scores[tm.strategies[j].Name()]
	})

	fmt.Printf("%-20s", "Стратегия")
	for _, s := range tm.strategies {
		fmt.Printf("%-12s", s.Name())
	}
	fmt.Printf("%-12s", "ИТОГО")
	fmt.Println()

	for _, s1 := range tm.strategies {
		fmt.Printf("%-20s", s1.Name())
		for _, s2 := range tm.strategies {
			fmt.Printf("%-12d", tm.results[s1.Name()][s2.Name()])
		}
		fmt.Printf("%-12d", scores[s1.Name()])
		fmt.Println()
	}

}
