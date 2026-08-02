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

func (tm *TournamentManager) PlayTournament() {
	for _, s := range tm.strategies {
		tm.results[s.Name()] = make(map[string]int)
	}

	for i := 0; i < len(tm.strategies); i++ {
		for j := i + 1; j < len(tm.strategies); j++ {
			s1 := tm.strategies[i]
			s2 := tm.strategies[j]

			res := PlayMatch(s1, s2, tm.config, 0)

			tm.results[s1.Name()][s2.Name()] = res.Score1
			tm.results[s2.Name()][s1.Name()] = res.Score2
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
