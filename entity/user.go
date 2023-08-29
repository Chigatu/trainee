package User

type User struct {
	Id    int       `json:"-"`
	Slugs []Segment `json:"string"`
}
