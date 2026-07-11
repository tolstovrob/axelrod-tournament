package strategies

import (
	"fmt"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

type Kamikaze struct {
	// Стратегия "Камикадзе": сотрудничает до последних раундов, затем начинает
	// предавать, чтобы другая стратегия не успела отомстить
	currentRound  int
	BetrayalStart int
}

func (s *Kamikaze) Name() string {
	return fmt.Sprintf("Камикадзе (%d)", s.BetrayalStart)
}

func (s *Kamikaze) Move(opponentHistory internal.History) internal.Action {
	s.currentRound++

	if s.currentRound >= s.BetrayalStart {
		return internal.Defect
	}

	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}
	return opponentHistory[len(opponentHistory)-1]
}

func (s *Kamikaze) Reset() {
	s.currentRound = 0
}
