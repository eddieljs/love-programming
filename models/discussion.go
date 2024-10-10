package models

type Discussion struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserId     int       `json:"userId"`
	UploadTime string    `json:"uploadTime"`
	CourseId   int       `json:"CourseId"`
	Comments   []Comment `json:"comments" gorm:"foreignKey:DisId;references:Id"`
	User       User      `json:"user" gorm:"foreignKey:UserId;references:Id"`
}

func (Discussion) TableName() string {
	return "discussion"
}

type Comment struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	UserId     int    `json:"userId"`
	UploadTime string `json:"uploadTime"`
	DisId      int    `json:"disId"`
	User       User   `json:"user" gorm:"foreignKey:UserId;references:Id"`
}

func (Comment) TableName() string {
	return "comment"
}
