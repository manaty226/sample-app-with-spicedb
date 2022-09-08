package entity

type Object struct {
	Type ObjectType
	ID   string
}

type ObjectType string

const (
	BlogType ObjectType = "blog"
	UserType ObjectType = "user"
)

type Action string

const (
	Read  Action = "read"
	Write Action = "write"
)
