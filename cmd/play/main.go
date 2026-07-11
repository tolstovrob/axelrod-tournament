package main

import (
	internal "github.com/tolstovrob/axelrod-tournament/internal"
)

func main() {
	config := internal.Config{
		Rounds:       200,
		PayoffMatrix: internal.DefaultPayoffMatrix(),
	}

	tm := internal.NewTournamentManager(config)

	// Add your strategies here:

	// -------------------------

	tm.PlayTournament()
	tm.PrintResults()
}
