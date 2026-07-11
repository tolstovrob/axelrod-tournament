package internal

type Strategy interface {
	Name() string
	Move(history History) Action
	Reset()
}
