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
	results    map[int]map[int]int // [P1][P2]Score
}

func NewTournamentManager(config Config) *TournamentManager {
	return &TournamentManager{
		config:  config,
		results: make(map[int]map[int]int),
	}
}

func (tm *TournamentManager) AddStrategy(strategy Strategy) {
	tm.strategies = append(tm.strategies, strategy)
}

func (tm *TournamentManager) PlayTournament() {
	for i := range tm.strategies {
		tm.results[i] = make(map[int]int)
	}

	for i := 0; i < len(tm.strategies); i++ {
		for j := i + 1; j < len(tm.strategies); j++ {
			s1 := tm.strategies[i]
			s2 := tm.strategies[j]

			res := PlayMatch(s1, s2, tm.config, 0)

			tm.results[i][j] = res.Score1
			tm.results[j][i] = res.Score2
		}
	}
}

func (tm *TournamentManager) getSortedStrategies() ([]int, map[int]int) {
	if len(tm.strategies) == 0 {
		return nil, nil
	}

	sorted := make([]int, len(tm.strategies))
	totalScores := make(map[int]int)
	for i := range tm.strategies {
		sorted[i] = i
		for j := range tm.strategies {
			totalScores[i] += tm.results[i][j]
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return totalScores[sorted[i]] > totalScores[sorted[j]]
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
	for _, i := range sorted {
		fmt.Printf("%-20s", tm.strategies[i].Name())
	}
	fmt.Printf("%-20s", "ИТОГО")
	fmt.Println()

	for _, i := range sorted {
		fmt.Printf("%-20s", tm.strategies[i].Name())
		for _, j := range sorted {
			fmt.Printf("%-12d", tm.results[i][j])
		}
		fmt.Printf("%-12d", scores[i])
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
	for _, i := range sorted {
		row = append(row, tm.strategies[i].Name())
	}
	row = append(row, "ИТОГО")
	writer.Write(row)

	for _, i := range sorted {
		row := []string{tm.strategies[i].Name()}
		for _, j := range sorted {
			row = append(row, strconv.Itoa(tm.results[i][j]))
		}
		row = append(row, strconv.Itoa(scores[i]))
		writer.Write(row)
	}

	return nil
}
