package models

type User struct {
	BaseModel
	Username string `json:"username" gorm:"unique;not null;comment:用户名"` // 用户名
	Password string `json:"password" gorm:"not null;comment:密码"`         // 密码
	Email    string `json:"email" gorm:"unique;not null;comment:邮箱"`      // 邮箱
	Mobile    string `json:"mobile" gorm:"index:idx_mobile;unique;type:varchar(11);not null;comment:手机号"`      // 手机号
	Nickname string `json:"nickname,omitempty" gorm:"comment:昵称"`          // 昵称
	Avatar   string `json:"avatar,omitempty" gorm:"comment:头像"`            // 头像
	Gender string `json:"gender,omitempty" gorm:"type:tinyint;default:0;comment:性别 0-未知 1-男 2-女"` // 性别
}