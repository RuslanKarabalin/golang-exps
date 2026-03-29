package model

type SomeType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateSomeRequest struct {
	Name string `json:"name"`
}
