package models

type Code struct {
	Id      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Code    string `json:"code" form:"code"`
	PointId int    `json:"pointId" form:"pointId"`
}

func (Code) TableName() string {
	return "code"
}

// type CodeCate struct {
// 	Id    int    `json:"id"`
// 	Title string `json:"title"`
// 	Codes []Code `json:"codes" gorm:"foreignKey:CateId;references:Id"`
// }

// func (CodeCate) TableName() string {
// 	return "code_cate"
// }
