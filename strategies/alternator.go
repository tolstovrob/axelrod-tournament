package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type Alternator struct {
	// Стратегия "Переключатель": чередует C и D каждый ход
}

func (s Alternator) Name() string {
	return "Переключатель"
}

func (s Alternator) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	if opponentHistory[len(opponentHistory)-1] == internal.Cooperate {
		return internal.Defect
	}
	return internal.Cooperate
}

func (s Alternator) Reset() {}
