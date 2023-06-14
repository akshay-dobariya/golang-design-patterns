package resources

type Application interface {
	MemoryRequired() MemorySizeInMB
	GetAppName() string
	RequestMemory()
	UseAndFreeMemory()
	GotMemory()
}

func GetApplication(appType, name string, memoryRequired, timeRequired int) Application {
	switch appType {
	case "photoshop":
		return &Photoshop{
			name:           name,
			memoryRequired: MemorySizeInMB(memoryRequired),
			timeRequired:   timeRequired,
			resource:       GetMemory(),
		}
	case "illustrator":
		return &Illustrator{
			name:           name,
			memoryRequired: MemorySizeInMB(memoryRequired),
			timeRequired:   timeRequired,
			resource:       GetMemory(),
		}
	default:
		return nil
	}
}
