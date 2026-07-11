package strategies

import (
	"math/rand"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

type GenerousTitForTat struct {
	// Стратегия "Великодушное око за око": как TitForTat, но с вероятностью 50%
	// прощает предательство
}

func (s GenerousTitForTat) Name() string {
	return "Великодушное око за око"
}

func (s GenerousTitForTat) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	lastOpponentAction := opponentHistory[len(opponentHistory)-1]

	if lastOpponentAction == internal.Defect {
		if rand.Float64() < 0.5 {
			return internal.Cooperate
		}
		return internal.Defect
	}
	return internal.Cooperate
}

func (s GenerousTitForTat) Reset() {}
