package charger

type Charger struct {
	watt int
}

func (m *Charger) SetWatt(watt int) error {
	m.watt = watt
	return nil
}

func (m *Charger) GetWatt() int {
	return m.watt
}
