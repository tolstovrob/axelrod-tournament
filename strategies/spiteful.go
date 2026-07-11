package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type Spiteful struct {
	// Стратегия "Мстительный": сотрудничает, пока соперник сотрудничает, но после
	// двух предательств подряд предает до конца игры
	defectionCount int
	retaliating    bool
}

func (s *Spiteful) Name() string {
	return "Мстительный"
}

func (s *Spiteful) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	if s.retaliating {
		return internal.Defect
	}

	if opponentHistory[len(opponentHistory)-1] == internal.Defect {
		s.defectionCount++
		if s.defectionCount >= 2 {
			s.retaliating = true
			return internal.Defect
		}
	} else {
		s.defectionCount = 0
	}

	return internal.Cooperate
}

func (s *Spiteful) Reset() {
	s.defectionCount = 0
	s.retaliating = false
}
