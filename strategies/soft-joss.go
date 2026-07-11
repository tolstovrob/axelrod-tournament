package strategies

import (
	"math/rand"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

type SoftJoss struct {
	// Стратегия "Мягкий Джосс": как TitForTat, но с 10% шансом предать после
	// сотрудничества соперника
}

func (s SoftJoss) Name() string {
	return "Мягкий Джосс"
}

func (s SoftJoss) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	lastOpponentAction := opponentHistory[len(opponentHistory)-1]

	if lastOpponentAction == internal.Defect {
		return internal.Defect
	}

	if rand.Float64() < 0.1 {
		return internal.Defect
	}
	return internal.Cooperate
}

func (s SoftJoss) Reset() {}
