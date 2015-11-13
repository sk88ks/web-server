package entity

type Friend struct {
	Fname string
}

type Person struct {
	UserName   string
	Emails     []string
	Friends    []Friend
	HasContent string
}
