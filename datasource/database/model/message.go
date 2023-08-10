package model

type Message struct {
	Id         int64
	FromUser   int64
	ToUser     int64
	Content    string
	CreateTime int64
}
