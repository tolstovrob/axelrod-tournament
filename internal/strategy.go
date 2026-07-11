package internal

type Strategy interface {
	Name() string
	Move(opponentHistory History) Action
	Reset()
}
