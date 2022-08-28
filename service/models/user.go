package models

type User struct {
	Uid        uint32 `gorm:"primaryKey;column:uid;type:int(10) unsigned;not null"`                // 用户ID，自增ID
	UserName   string `gorm:"column:userName;type:varchar(32);not null"`                           // 用户名
	Password   string `gorm:"column:userName;type:varchar(32);not null"`                           // 密码
	Role       uint32 `gorm:"column:role;type:tinyint(4) unsigned;not null"`                       // 角色 1-guest 2-user 3-admin
	CreateTime string `gorm:"column:createTime;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime string `gorm:"column:updateTime;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 更新时间
}

type Token struct {
	Token     string `gorm:"column:token;type:varchar(255);not null"`                            // 用户token
	UserName  string `gorm:"column:userName;type:varchar(32);not null"`                          // 用户名
	LoginTime string `gorm:"column:loginTime;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 登录时间
}
