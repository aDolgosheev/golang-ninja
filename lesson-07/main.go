package main

import "fmt"

type Age int

func(a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func (u *User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

func (u *User) setUName(name string) {
	u.name = name
}

func (u User) getName() string {
	return u.name
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name: name,
		sex: sex,
		age: Age(age),
		weight: weight,
		height: height,
	}
}

func main() {

	user1 := NewUser("Vasya", "Male", 23, 75, 185)
	user2 := User{"Petya", 31, "Male", 84, 197}
	// user1 := User{"Vasya", 23, "Male", 75, 185}

	// fmt.Println(user1.name)

	// fmt.Printf("%+v\n", user1)
	// fmt.Printf("%+v\n", user2)

	fmt.Println(user1.age.isAdult())

	user1.printUserInfo()
	user2.printUserInfo()
}


