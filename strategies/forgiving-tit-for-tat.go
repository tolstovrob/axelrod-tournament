package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type ForgivingTitForTat struct {
	// Стратегия "Прощающее око за око": как TitForTat, но с вероятностью 20%
	// прощать предательство
	forgivenessCount int
}

func (s *ForgivingTitForTat) Name() string {
	return "Прощающее око за око"
}

func (s *ForgivingTitForTat) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	lastOpponentAction := opponentHistory[len(opponentHistory)-1]

	// Если соперник предал
	if lastOpponentAction == internal.Defect {
		s.forgivenessCount++
		// Прощаем каждое 5-е предательство
		if s.forgivenessCount%5 == 0 {
			return internal.Cooperate
		}
		return internal.Defect
	}

	// Если соперник сотрудничал - сбрасываем счетчик
	if s.forgivenessCount > 0 {
		s.forgivenessCount = 0
	}
	return internal.Cooperate
}

func (s *ForgivingTitForTat) Reset() {
	s.forgivenessCount = 0
}
