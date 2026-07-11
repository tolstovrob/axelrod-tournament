package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type TitForTat struct {
	// Стратегия "Око за око": начинает с сотрудничества, затем повторяет
	// предыдущий ход соперника
}

func (s TitForTat) Name() string {
	return "Око за око"
}

func (s TitForTat) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}
	return opponentHistory[len(opponentHistory)-1]
}

func (s TitForTat) Reset() {}
