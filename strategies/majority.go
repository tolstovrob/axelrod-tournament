package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type Majority struct {
	// Стратегия "Большинство": повторяет то действие, которое соперник
	// использовал чаще
}

func (s Majority) Name() string {
	return "Большинство"
}

func (s Majority) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	cooperationCount := 0
	for _, action := range opponentHistory {
		if action == internal.Cooperate {
			cooperationCount++
		}
	}

	if cooperationCount > len(opponentHistory)/2 {
		return internal.Cooperate
	}
	return internal.Defect
}

func (s Majority) Reset() {}
