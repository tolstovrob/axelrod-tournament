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

	// Add your strategies here:
	tm.AddStrategy(strategies.AlwaysDefect{})
	tm.AddStrategy(strategies.AlwaysCooperate{})
	tm.AddStrategy(strategies.TitForTat{})
	tm.AddStrategy(&strategies.TitForTwoTats{})
	tm.AddStrategy(&strategies.Pavlov{})
	tm.AddStrategy(&strategies.Grudger{})
	tm.AddStrategy(&strategies.Detective{})
	tm.AddStrategy(strategies.SoftJoss{})
	tm.AddStrategy(strategies.HardJoss{})
	tm.AddStrategy(&strategies.ForgivingTitForTat{})
	tm.AddStrategy(strategies.SuspiciousTitForTat{})
	tm.AddStrategy(&strategies.Spiteful{})
	tm.AddStrategy(strategies.Alternator{})
	tm.AddStrategy(strategies.Majority{})
	tm.AddStrategy(&strategies.Prober{})
	tm.AddStrategy(&strategies.Gradual{})
	tm.AddStrategy(strategies.GenerousTitForTat{})
	tm.AddStrategy(strategies.AdaptiveStrategy{Threshold: 0.3})
	tm.AddStrategy(strategies.AdaptiveStrategy{Threshold: 0.5})
	tm.AddStrategy(strategies.AdaptiveStrategy{Threshold: 0.7})
	tm.AddStrategy(&strategies.Kamikaze{BetrayalStart: 160})
	tm.AddStrategy(&strategies.Kamikaze{BetrayalStart: 120})
	tm.AddStrategy(&strategies.TwoTitsForTat{})
	// -------------------------

	tm.PlayTournament()
	// tm.PrintResults()
	if tm.ExportResultsToCSV("tournament.csv") != nil {
		fmt.Println("Не удалось экспортировать результаты в CSV файл")
	}
}
