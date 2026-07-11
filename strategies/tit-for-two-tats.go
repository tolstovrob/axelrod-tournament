package strategies

import "github.com/tolstovrob/axelrod-tournament/internal"

type TitForTwoTats struct {
	// Стратегия "Око за два ока": предает только после двух предательств подряд
	// от соперника
}

func (s TitForTwoTats) Name() string {
	return "Око за два ока"
}

func (s TitForTwoTats) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	if len(opponentHistory) >= 2 {
		if opponentHistory[len(opponentHistory)-1] == internal.Defect &&
			opponentHistory[len(opponentHistory)-2] == internal.Defect {
			return internal.Defect
		}
	}
	return internal.Cooperate
}

func (s TitForTwoTats) Reset() {}
