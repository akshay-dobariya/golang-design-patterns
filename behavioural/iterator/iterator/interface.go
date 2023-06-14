package iterator

type Interator interface {
	HasNext() bool
	GetNext() any
	RestartInteration()
}
