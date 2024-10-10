package models

type Resource struct {
	Id       int    `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Url      string `json:"url" form:"url"`
	CourseId int    `json:"courseId" form:"courseId"`
}

func (Resource) TableName() string {
	return "resource"
}

// type ResourceCate struct {
// 	Id        int        `json:"id"`
// 	Title     string     `json:"title"`
// 	Resources []Resource `json:"resources" gorm:"foreignKey:CateId;references:Id"`
// }

// func (ResourceCate) TableName() string {
// 	return "resource_cate"
// }
