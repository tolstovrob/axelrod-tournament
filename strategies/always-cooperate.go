package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type AlwaysCooperate struct {
	// Стратегия "Добряк": всегда сотрудничает с соперником
}

func (s AlwaysCooperate) Name() string {
	return "Добряк"
}

func (s AlwaysCooperate) Move(_ internal.History) internal.Action {
	return internal.Cooperate
}

func (s AlwaysCooperate) Reset() {}
