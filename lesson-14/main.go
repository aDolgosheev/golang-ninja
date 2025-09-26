package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, item.action, item.timestamp)
	}

	return out
}

func main() {
	// u := User{
	// 	id:    1,
	// 	email: "vasya@ninja.go",
	// 	logs: []logItem{
	// 		{actions[0], time.Now()},
	// 		{actions[3], time.Now()},
	// 		{actions[2], time.Now()},
	// 		{actions[1], time.Now()},
	// 		{actions[0], time.Now()},
	// 		{actions[3], time.Now()},
	// 	},
	// }
	// fmt.Println(u.getActivityInfo())

	rand.New(rand.NewSource(time.Now().Unix()))

	users := generateUsers(1000)

	for _, user := range users {
		saveUserInfo(user)
	}
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@ninja.go", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
	}

	return users
}

func saveUserInfo(u User) error {
	fmt.Printf("WRITING FILE FOR USER ID: %d\n", u.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", u.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(u.getActivityInfo())
	return err
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}

// func main() {
// 	go func() {
// 		time.Sleep(time.Second)
// 		fmt.Println("Concurent ninja wait")
// 	}()
// 	go fmt.Println("Concurent ninja")
// 	go fmt.Println("Concurent ninja")

// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Non concurent ninja")
// }
