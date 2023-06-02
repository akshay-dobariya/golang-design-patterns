package mobile

type MobileInterface interface {
	SetHardware(string) error
	SetSoftware(string) error
	GetHardware() string
	GetSoftware() string
}
