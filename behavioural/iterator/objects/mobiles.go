package objects

type Mobile struct {
	name string
}

func CreateMobile(name string) Mobile {
	return Mobile{
		name: name,
	}
}

type MobileIterator struct {
	index   int
	Mobiles *[]Mobile
}

func (mi *MobileIterator) HasNext() bool {
	if mi.Mobiles == nil {
		return false
	} else if mi.index >= len(*mi.Mobiles) {
		return false
	}
	return true
}

func (mi *MobileIterator) GetNext() any {
	if mi.HasNext() {
		user := (*mi.Mobiles)[mi.index]
		mi.index++
		return user
	}
	return nil
}

func (mi *MobileIterator) RestartInteration() {
	mi.index = 0
}
