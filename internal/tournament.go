package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	action1 := s1.Move(h2)
	action2 := s2.Move(h1)
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

func (tm *TournamentManager) getSortedStrategies() ([]Strategy, map[string]int) {
	if len(tm.strategies) == 0 {
		return nil, nil
	}

	names := make([]string, len(tm.strategies))
	for i, s := range tm.strategies {
		names[i] = s.Name()
	}

	totalScores := make(map[string]int)
	for _, s1 := range names {
		total := 0
		for _, s2 := range names {
			total += tm.results[s1][s2]
		}
		totalScores[s1] = total
	}

	sorted := make([]Strategy, len(tm.strategies))
	copy(sorted, tm.strategies)
	sort.Slice(sorted, func(i, j int) bool {
		return totalScores[sorted[i].Name()] > totalScores[sorted[j].Name()]
	})

	return sorted, totalScores
}

func (tm *TournamentManager) PrintResults() {
	if len(tm.strategies) == 0 {
		fmt.Println("Нет стратегий для отображения")
		return
	}

	sorted, scores := tm.getSortedStrategies()

	fmt.Printf("%-20s", "Стратегия")
	for _, s := range sorted {
		fmt.Printf("%-20s", s.Name())
	}
	fmt.Printf("%-20s", "ИТОГО")
	fmt.Println()

	for _, s1 := range sorted {
		fmt.Printf("%-20s", s1.Name())
		for _, s2 := range sorted {
			fmt.Printf("%-12d", tm.results[s1.Name()][s2.Name()])
		}
		fmt.Printf("%-12d", scores[s1.Name()])
		fmt.Println()
	}
}

func (tm *TournamentManager) ExportResultsToCSV(filename string) error {
	if len(tm.strategies) == 0 {
		return fmt.Errorf("нет стратегий для экспорта")
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	sorted, scores := tm.getSortedStrategies()

	row := []string{"Стратегия"}
	for _, s := range sorted {
		row = append(row, s.Name())
	}
	row = append(row, "ИТОГО")
	writer.Write(row)

	for _, s1 := range sorted {
		row := []string{s1.Name()}
		for _, s2 := range sorted {
			row = append(row, strconv.Itoa(tm.results[s1.Name()][s2.Name()]))
		}
		row = append(row, strconv.Itoa(scores[s1.Name()]))
		writer.Write(row)
	}

	return nil
}
