package models

type Course struct {
	Id    int    `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
	// 加课码
	Key         string       `json:"key" form:"key"`
	Resources   []Resource   `json:"resources" form:"resources" gorm:"foreignKey:CourseId;references:Id"`
	Discussions []Discussion `json:"discussions" gorm:"foreignKey:CourseId;references:Id"`
	Videos      []Video      `json:"videos" gorm:"foreignKey:CourseId;references:Id"`
	// 练习题的知识点
	Points    []ExercisePoint `json:"points" gorm:"foreignKey:CourseId;references:Id"`
	TeacherId int             `json:"teacherId" form:"teacherId"`
	Teacher   User            `json:"tercher" form:"teacher" gorm:"foreignKey:Id;references:TeacherId"`
	Time      string          `json:"time" form:"time"`
}

func (Course) TableName() string {
	return "course"
}

type CourseUser struct {
	Id         int    `json:"id" form:"id"`
	CourseId   int    `json:"courseId" form:"courseId"`
	UserId     int    `json:"userId" form:"userId"`
	TeacherId  int    `json:"teacherId" from:"teacherId"` // 教师ID
	CourseTime string `json:"course_time" from:"course_time"`
}

func (CourseUser) TableName() string {
	return "course_user"
}

type CourseExam struct {
	Id          int    `json:"id" form:"id"`
	CourseTime  string `json:"courseTime" form:"courseTime"`
	UserId      int    `json:"userId" form:"userId"`
	CourseTitle string `json:"courseTitle" form:"courseTitle"`
	CourseKey   string `json:"courseKey" form:"courseKey"`
	Time        string `json:"time" form:"time"`
	User        User   `json:"user" form:"user" gorm:"foreignKey:UserId;references:Id"`
}

func (CourseExam) TableName() string {
	return "course_exam"
}

// CourseTeacher 结构体表示课程和教师的关联关系
type CourseTeacher struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	CourseId  int    `json:"course_id"`  // 课程ID
	TeacherId int    `json:"teacher_id"` // 教师ID
	Key       string `json:"key"`        // 加课码
}
