package main

import "fmt"

/* interfaceの定義 */
type Greetable interface {
	greet()
}

type Member struct {
	name    string
	age     int
	message string
}

type MemberEx struct {
	name  string
	grade string
}

func (self Member) greet() {
	fmt.Printf("%sです。%s\n", self.name, self.message)
}

func (self MemberEx) greet() {
	fmt.Printf("Grade:%s Name:%s\n", self.grade, self.name)
}

func main() {
	members := []Greetable{
		Member{"minaki", 27, "beginning golang"},
		MemberEx{"ex-minaki", "user"},
	}

	for _, m := range members {
		m.greet()
	}
}
