package service

type User struct {
	Id      int    `json:"-"`
	Segment string `json:"segment"`
}
