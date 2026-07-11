package strategies

import (
	"fmt"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

type AdaptiveStrategy struct {
	// Стратегия "Адаптивная": сотрудничает, если соперник сотрудничает чаще
	// заданного порога
	Threshold float64
}

func (s AdaptiveStrategy) Name() string {
	return fmt.Sprintf("Адаптивная (%.2f)", s.Threshold)
}

func (s AdaptiveStrategy) Move(opponentHistory internal.History) internal.Action {
	if len(opponentHistory) == 0 {
		return internal.Cooperate
	}

	cooperationCount := 0
	for _, action := range opponentHistory {
		if action == internal.Cooperate {
			cooperationCount++
		}
	}

	cooperationRate := float64(cooperationCount) / float64(len(opponentHistory))

	if cooperationRate >= s.Threshold {
		return internal.Cooperate
	}
	return internal.Defect
}

func (s AdaptiveStrategy) Reset() {}
