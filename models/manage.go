package models

type Prop struct {
	Id         uint32 `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"`                 // 属性ID，自增ID
	Label      string `gorm:"column:label;type:varchar(10);not null"`                              // 属性名
	Value      uint32 `gorm:"column:value;type:tinyint(4) unsigned;not null"`                      // 属性值
	MType      uint32 `gorm:"column:mType;type:tinyint(4) unsigned;not null"`                      // 类型 1-类型 2-制片国家/地区 3-语言 4-上映日期
	CreateTime string `gorm:"column:createTime;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 创建时间
	UpdateTime string `gorm:"column:updateTime;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 更新时间
}
