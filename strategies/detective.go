package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type Detective struct {
	// Стратегия "Детектив": анализирует соперника первые 4 хода, затем
	// адаптируется
	phase        int
	analyzed     bool
	opponentType string // "cooperative" или "defective"
}

func (s *Detective) Name() string {
	return "Детектив"
}

func (s *Detective) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		s.phase = 0
		return internal.Cooperate
	}

	s.phase++

	// Первые 4 хода: C, D, C, C
	if s.phase <= 4 {
		switch s.phase {
		case 1:
			return internal.Cooperate
		case 2:
			return internal.Defect
		case 3:
			return internal.Cooperate
		case 4:
			return internal.Cooperate
		}
	}

	// После 4-го хода анализируем ответы соперника
	if !s.analyzed && len(opponentHistory) >= 4 {
		s.analyzed = true
		// Если соперник ответил предательством на наш Defect
		if opponentHistory[1] == internal.Defect {
			s.opponentType = "defective"
		} else {
			s.opponentType = "cooperative"
		}
	}

	// Если соперник кооперативный - всегда сотрудничаем
	if s.opponentType == "cooperative" {
		return internal.Cooperate
	}

	// Если соперник предатель - играем TitForTat
	if len(opponentHistory) > 0 {
		return opponentHistory[len(opponentHistory)-1]
	}

	return internal.Cooperate
}

func (s *Detective) Reset() {
	s.phase = 0
	s.analyzed = false
	s.opponentType = ""
}
