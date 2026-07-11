package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type Gradual struct {
	// Стратегия "Постепенный": начинает с сотрудничества, после предательства
	// соперника предает несколько раз подряд
	defectionStreak int
	punishmentCount int
}

func (s *Gradual) Name() string {
	return "Постепенный"
}

func (s *Gradual) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	lastOpponentAction := opponentHistory[len(opponentHistory)-1]

	// Если соперник предал
	if lastOpponentAction == internal.Defect {
		s.defectionStreak++
		// Наказываем столько же раз, сколько было предательств подряд
		if s.punishmentCount < s.defectionStreak {
			s.punishmentCount++
			return internal.Defect
		}
		// После наказания возвращаемся к сотрудничеству
		s.punishmentCount = 0
		s.defectionStreak = 0
		return internal.Cooperate
	}

	// Если соперник сотрудничает, но мы еще наказываем
	if s.punishmentCount > 0 {
		s.punishmentCount--
		return internal.Defect
	}

	s.defectionStreak = 0
	return internal.Cooperate
}

func (s *Gradual) Reset() {
	s.defectionStreak = 0
	s.punishmentCount = 0
}
