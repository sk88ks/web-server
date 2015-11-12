package entity

// AliveRes is response for alive API
type AliveRes struct {
	OK bool `json:"ok"`
}

type Friend struct {
	Fname string
}
