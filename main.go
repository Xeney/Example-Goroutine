package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

var actions = []string{
	"logget in",
	"logget out",
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
	out := fmt.Sprintf("ID:%d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i+1, item.action, item.timestamp)
	}

	return out
}

func main() {
	t := time.Now()
	rand.Seed(time.Now().Unix())

	users := generateUsers(1000)
	for _, user := range users {
		wg.Add(1)
		go saveUserInfo(user)
	}

	wg.Wait()

	fmt.Println("TIME:", time.Since(t))
}

func saveUserInfo(u User) error {
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("WRITTING FILE FOR USER ID: %d\n", u.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", u.id)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = file.WriteString(u.getActivityInfo())
	if err != nil {
		return err
	}

	wg.Done()

	return nil
}

func generateUsers(count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmai.com", i+1),
			logs:  generateLogs(10 + rand.Intn(990)),
		}
	}

	return users
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
