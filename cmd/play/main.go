package main

import (
	"fmt"

	"github.com/tolstovrob/axelrod-tournament/internal"
	"github.com/tolstovrob/axelrod-tournament/strategies"
)

func main() {
	config := internal.Config{
		Rounds:       200,
		PayoffMatrix: internal.DefaultPayoffMatrix(),
	}

	tm := internal.NewTournamentManager(config)

	for _, spec := range strategies.AllClassicSpecs() {
		tm.AddStrategy(strategies.New(spec))
	}

	tm.PlayTournament()
	if err := tm.ExportResultsToCSV("tournament.csv"); err != nil {
		fmt.Println("Не удалось экспортировать результаты в CSV файл")
	}
}
