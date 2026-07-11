package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type SuspiciousTitForTat struct {
	// Стратегия "Подозрительное око за око": как TitForTat, но начинает с
	// предательства
}

func (s SuspiciousTitForTat) Name() string {
	return "Подозрительное око за око"
}

func (s SuspiciousTitForTat) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Defect
	}
	return opponentHistory[len(opponentHistory)-1]
}

func (s SuspiciousTitForTat) Reset() {}
