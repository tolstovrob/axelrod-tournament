package strategies

import (
	"math/rand/v2"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

type Random struct {
	// Стратегия "Рандом": случайный выбор на каждом ходу
}

func (s Random) Name() string {
	return "Рандом"
}

func (s Random) Move(_ internal.History) internal.Action {
	if rand.Float64() < 0.5 {
		return internal.Cooperate
	}
	return internal.Defect
}

func (s Random) Reset() {}
