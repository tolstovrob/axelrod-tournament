package internal

type Config struct {
	Rounds       int
	PayoffMatrix map[[2]Action][2]int
}

func DefaultPayoffMatrix() map[[2]Action][2]int {
	// [P1Action][P2Action] -> (P1Score, P2Score)
	return map[[2]Action][2]int{
		{Cooperate, Cooperate}: {3, 3},
		{Cooperate, Defect}:    {0, 5},
		{Defect, Cooperate}:    {5, 0},
		{Defect, Defect}:       {1, 1},
	}
}
