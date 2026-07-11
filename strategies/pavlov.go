package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type Pavlov struct {
	// Стратегия "Павлов": если последний ход был успешным, повторяет его, иначе
	// меняет на противоположный
	lastAction internal.Action
}

func (s *Pavlov) Name() string {
	return "Павлов"
}

func (s *Pavlov) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		s.lastAction = internal.Cooperate
		return internal.Cooperate
	}

	lastOpponentAction := opponentHistory[len(opponentHistory)-1]

	wasSuccessful := (s.lastAction == internal.Cooperate && lastOpponentAction == internal.Cooperate) ||
		(s.lastAction == internal.Defect && lastOpponentAction == internal.Defect)

	if wasSuccessful {
		return s.lastAction
	}

	if s.lastAction == internal.Cooperate {
		s.lastAction = internal.Defect
	} else {
		s.lastAction = internal.Cooperate
	}
	return s.lastAction
}

func (s *Pavlov) Reset() {
	s.lastAction = internal.Cooperate
}
