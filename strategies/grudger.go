package strategies

import (
	"slices"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

type Grudger struct {
	// Стратегия "Злопамятный": сотрудничает, пока соперник
	// не предаст, после этого предает всегда
	hasOpponentDefected bool
}

func (s *Grudger) Name() string {
	return "Злопамятный"
}

func (s *Grudger) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	if slices.Contains(opponentHistory, internal.Defect) {
		s.hasOpponentDefected = true
	}

	if s.hasOpponentDefected {
		return internal.Defect
	}
	return internal.Cooperate
}

func (s *Grudger) Reset() {
	s.hasOpponentDefected = false
}
