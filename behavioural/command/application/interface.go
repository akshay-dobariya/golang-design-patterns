package application

type Application interface {
	Save()
	Undo()
	Work(string)
	DisplayWork()
}
