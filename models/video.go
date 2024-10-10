package models

type Video struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	CourseId int    `json:"courseId"`
}

func (Video) TableName() string {
	return "video"
}

// type VideoCate struct {
// 	Id     int     `json:"id"`
// 	Title  string  `json:"title"`
// 	Videos []Video `json:"videos" gorm:"foreignKey:CateId;references:Id"`
// }

// func (VideoCate) TableName() string {
// 	return "video_cate"
// }
