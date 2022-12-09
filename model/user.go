package model

const (
	MALE = iota
	FEMALE
)

const ACTIVE = 1

type User struct {
	BaseModel        // 实现了基类model
	Username  string `json:"username" gorm:"index;comment:用户登录名;not null;unique"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Sex       *uint  `json:"sex"`
	Phone     string `json:"phone"  gorm:"comment:用户手机号"`
	Status    uint64 `json:"status" gorm:"default:1;comment:用户是否冻结"`
	Avatar    string `json:"avatar"`
}

func (User) TableName() string {
	return "user"
}
