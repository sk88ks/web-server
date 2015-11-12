package entity

type Person struct {
	UserName   string
	Emails     []string
	Friends    []Friend
	HasContent string
}
