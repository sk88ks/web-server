package body

// AliveRes is response for alive API
type AliveRes struct {
	OK bool `json:"ok"`
}

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []Friend
}
