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

type History []Action
