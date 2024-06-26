package domain

// User 领域对象，是DDD中的entity
type User struct {
	Id              int64
	Email           string
	Password        string
	NickName        string
	Birthday        string
	PersonalProfile string
	IsDel           int64
}
