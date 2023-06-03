package iterator

import "fmt"

type User struct {
	name string
	age  int
}

type Iterator interface {
	hasNext() bool
	next() *User
}

type UserIterator struct {
	users []User
	index int
}

func (ui *UserIterator) hasNext() bool {
	return ui.index < len(ui.users)
}
func (ui *UserIterator) next() *User {
	if ui.hasNext() {
		user := &ui.users[ui.index]
		ui.index++
		return user
	}
	return nil
}

type Collection interface {
	getIterator() Iterator
}

type UserCollection struct {
	users []User
}

func (c UserCollection) getIterator() Iterator {
	return &UserIterator{users: c.users}
}

func Main() {
	users := []User{{"JanessaTech1", 10}, {"JanessaTech2", 11}, {"JanessaTech3", 12}}

	c := UserCollection{users: users}
	it := c.getIterator()
	for it.hasNext() {
		user := it.next()
		fmt.Printf("name=%s, age=%d\n", user.name, user.age)
	}
}
