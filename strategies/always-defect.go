package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type AlwaysDefect struct {
	// Стратегия "Предатель": всегда предаёт соперника
}

func (s AlwaysDefect) Name() string {
	return "Предатель"
}

func (s AlwaysDefect) Move(_ internal.History) internal.Action {
	return internal.Defect
}

func (s AlwaysDefect) Reset() {}
