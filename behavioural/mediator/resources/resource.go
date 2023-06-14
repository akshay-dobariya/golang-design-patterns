package resources

type Resource interface {
	IsAvailable(Application) bool
	NotifyFree(Application)
}
