package models

type Choice struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	PointId int    `json:"pointId"`
	Ans     string `json:"ans"`
	Analy   string `json:"analy"`
}

func (Choice) TableName() string {
	return "choice"
}

type ExercisePoint struct {
	Id       int      `json:"id" form:"id"`
	Title    string   `json:"title" form:"title"`
	CourseId int      `json:"CourseId" form:"CourseId"`
	Choices  []Choice `json:"choices" gorm:"foreignKey:PointId;references:Id"`
	Codes    []Code   `json:"codes" gorm:"foreignKey:PointId;references:Id"`
}

func (ExercisePoint) TableName() string {
	return "exercise_point"
}

// type ExerciseCate struct {
// 	Id             int             `json:"id" form:"id"`
// 	Title          string          `json:"title" form:"title"`
// 	ExercisePoints []ExercisePoint `json:"points" gorm:"foreignKey:CateId;references:Id"`
// }

// func (ExerciseCate) TableName() string {
// 	return "exercise_cate"
// }
