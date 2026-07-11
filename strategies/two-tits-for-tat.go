package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type TwoTitsForTat struct {
	// Стратегия "Два ока за око": на каждое предательство соперника отвечает
	// двумя предательствами подряд
	punishmentCount int
}

func (s *TwoTitsForTat) Name() string {
	return "Два ока за око"
}

func (s *TwoTitsForTat) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	if s.punishmentCount > 0 {
		s.punishmentCount--
		return internal.Defect
	}

	if opponentHistory[len(opponentHistory)-1] == internal.Defect {
		s.punishmentCount = 2
		return internal.Defect
	}

	return internal.Cooperate
}

func (s *TwoTitsForTat) Reset() {
	s.punishmentCount = 0
}
