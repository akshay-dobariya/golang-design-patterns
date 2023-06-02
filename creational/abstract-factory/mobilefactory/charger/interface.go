package charger

type ChargerInterface interface {
	SetWatt(int) error
	GetWatt() int
}
