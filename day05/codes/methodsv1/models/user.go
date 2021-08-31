package models

type User struct {
	id   int
	name string
	age  int
}

func NewUser(id, age int, name string) *User {
	return &User{id: id, name: name, age: age}
}
func AddAge(user User) {
	user.age += 1
}
func GrtName(user User) string {
	return user.name
}

// 不能修改调用者的age
func (user *User) AddAge() {
	user.age += 1
}
func (user *User) GetName() string {
	return user.name

}
