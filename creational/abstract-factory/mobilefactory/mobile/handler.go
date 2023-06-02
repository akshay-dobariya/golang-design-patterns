package mobile

type Mobile struct {
	hardware string
	software string
}

func (m *Mobile) SetHardware(str string) error {
	m.hardware = str
	return nil
}

func (m *Mobile) GetHardware() string {
	return m.hardware
}

func (m *Mobile) SetSoftware(str string) error {
	m.software = str
	return nil
}

func (m *Mobile) GetSoftware() string {
	return m.software
}
