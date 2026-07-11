package internal

type Action bool

const (
	Cooperate Action = true
	Defect    Action = false
)

func (a Action) String() string {
	if a == Cooperate {
		return "cooperate"
	}
	return "defect"
}

type RoundResult struct {
	P1Action Action
	P2Action Action
	P1Score  int
	P2Score  int
}

type History []Action
