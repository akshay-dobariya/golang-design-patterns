package objects

type User struct {
	name string
	age  int
}

func CreateUser(name string, age int) User {
	return User{
		name: name,
		age:  age,
	}
}

type UserIterator struct {
	index int
	Users *[]User
}

func (ui *UserIterator) HasNext() bool {
	if ui.Users == nil {
		return false
	} else if ui.index >= len(*ui.Users) {
		return false
	}
	return true
}

func (ui *UserIterator) GetNext() any {
	if ui.HasNext() {
		user := (*ui.Users)[ui.index]
		ui.index++
		return user
	}
	return nil
}

func (ui *UserIterator) RestartInteration() {
	ui.index = 0
}
