package models

import (
	"github.com/shopspring/decimal"
)

// GORM 会自动将结构体内的成员改名为Parking_Lot_ID(就是给你加下划线)
// 所以需要打标签告诉gorm列名是什么 `gorm:"column:ParkingLotID;primaryKey"`
// 全部都要把标签打上,没有打标签的话select *都显示0
// 标签中用;隔开
type ParkingLot struct {
	ParkingLotID int             `gorm:"column:ParkingLotID;primaryKey"`
	ParkingName  string          `gorm:"column:ParkingName"`
	Longitude    decimal.Decimal `gorm:"column:Longitude;type:decimal(9,6)"`
	Latitude     decimal.Decimal `gorm:"column:Latitude;type:decimal(9,6)"`
	Capacity     int             `gorm:"column:Capacity;default:null"`
	Rates        decimal.Decimal `gorm:"column:Rates;type:decimal(10,2)"`
}

// 设置ParkingLot表名为`parkinglot`
// TableName方法是sql.Scanner接口的一部分,
// 允许为模型指定一个自定义的表名.
// 在结构体上定义了这个方法,
// GORM 会在执行数据库操作时使用这个方法返回的字符串作为表名.
func (p *ParkingLot) TableName() string {
	return "parkinglot"
}
