package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type Prober struct {
	// Стратегия "Исследователь": проверяет реакцию соперника на предательство в
	// начале игры
	phase int
}

func (s *Prober) Name() string {
	return "Исследователь"
}

func (s *Prober) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		s.phase = 0
		return internal.Cooperate
	}

	s.phase++

	// Первые 3 хода: C, D, C
	if s.phase <= 3 {
		switch s.phase {
		case 1:
			return internal.Cooperate
		case 2:
			return internal.Defect
		case 3:
			return internal.Cooperate
		}
	}

	if len(opponentHistory) >= 2 && opponentHistory[1] == internal.Defect {
		return opponentHistory[len(opponentHistory)-1]
	}

	return internal.Cooperate
}

func (s *Prober) Reset() {
	s.phase = 0
}
