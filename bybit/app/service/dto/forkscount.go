package dto

type ForksCount struct {
	Projects Projects `json:"projects"`
}

type Projects struct {
	Nodes []Nodes `json:"nodes"`
}

type Nodes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ForksCount  int    `json:"forksCount"`
}
