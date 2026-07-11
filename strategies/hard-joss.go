package strategies

import (
	"math/rand"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

type HardJoss struct {
	// Стратегия "Жесткий Джосс": как TitForTat, но с 30% шансом
	// предать после сотрудничества соперника
}

func (s HardJoss) Name() string {
	return "Жесткий Джосс"
}

func (s HardJoss) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	lastOpponentAction := opponentHistory[len(opponentHistory)-1]

	if lastOpponentAction == internal.Defect {
		return internal.Defect
	}

	if rand.Float64() < 0.3 {
		return internal.Defect
	}
	return internal.Cooperate
}

func (s HardJoss) Reset() {}
