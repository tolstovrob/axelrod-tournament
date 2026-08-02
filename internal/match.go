package internal

import "math/rand/v2"

type MatchResult struct {
	Score1            int
	Score2            int
	CooperationRounds int
	TotalRounds       int
}

func PlayMatch(s1, s2 Strategy, cfg Config, noise float64) MatchResult {
	s1.Reset()
	s2.Reset()

	var h1, h2 History
	score1, score2 := 0, 0
	coopRounds := 0

	for round := 0; round < cfg.Rounds; round++ {
		a1 := s1.Move(h2)
		a2 := s2.Move(h1)

		if noise > 0 {
			if rand.Float64() < noise {
				a1 = !a1
			}
			if rand.Float64() < noise {
				a2 = !a2
			}
		}

		sc1, sc2 := scorePair(a1, a2, cfg.PayoffMatrix)
		score1 += sc1
		score2 += sc2

		if a1 == Cooperate && a2 == Cooperate {
			coopRounds++
		}

		h1 = append(h1, a1)
		h2 = append(h2, a2)
	}

	return MatchResult{
		Score1:            score1,
		Score2:            score2,
		CooperationRounds: coopRounds,
		TotalRounds:       cfg.Rounds,
	}
}

func scorePair(a1, a2 Action, matrix map[[2]Action][2]int) (int, int) {
	if val, ok := matrix[[2]Action{a1, a2}]; ok {
		return val[0], val[1]
	}
	return 0, 0
}
