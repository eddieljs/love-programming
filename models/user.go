package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	// Password  string `json:"password"`
	Auth      int      `json:"auth"`
	Name      string   `json:"realName"`
	StudentId string   `json:"studentId"`
	Avatar    string   `json:"avatar"`
	Openid    string   `json:"openid"`
	VipTime   int      `json:"vipTime"`
	Courses   []Course `json:"courses" gorm:"many2many:course_user;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:CourseId"`
}

func (User) TableName() string {
	return "user"
}

type License struct {
	Id        int
	SecretKey string
}

func (License) TableName() string {
	return "license"
}

type Record struct {
	Id     int    `json:"id" form:"id"`
	UserId int    `json:"userId" form:"userId"`
	Type   string `json:"type" form:"type"`
	Title  string `json:"title" form:"title"`
	Url    string `json:"url" form:"url"`
	Time   string `json:"time" form:"time"`
}

func (Record) TableName() string {
	return "record"
}

type ChoiceRecord struct {
	Id       int    `json:"id" form:"id"`
	ChoiceId int    `json:"ChoiceId" form:"ChoiceId"`
	UserId   int    `json:"uId" form:"uId"`
	UserAns  string `json:"userAns" form:"userAns"`
}

func (ChoiceRecord) TableName() string {
	return "choice_record"
}

type VIPKey struct {
	Id       int    `json:"id" form:"id"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" forn:"password"`
}

func (VIPKey) TableName() string {
	return "vip_key"
}

type AdminAccount struct {
	Id       int    `json:"id" form:"id"`
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

func (AdminAccount) TableName() string {
	return "admin_account"
}
