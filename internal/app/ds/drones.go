package ds

import (
	"gorm.io/datatypes"
)

type Region struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Details    string
	AreaKm     float64
	Population int
}

type UserRole struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type FlightStatus struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type User struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	UserRoleRefer int
	Role          UserRole `gorm:"foreignKey:UserRoleRefer"`
}

type Flight struct {
	ID                uint `gorm:"primaryKey"`
	UserRefer         int
	FlightStatusRefer int
	Status            FlightStatus `gorm:"foreignKey:FlightStatusRefer"`
	DateCreated       datatypes.Date
	DateProcessed     datatypes.Date
	DateFinished      datatypes.Date
	Moderator         User `gorm:"foreignKey:UserRefer"`
	TakeoffDate       datatypes.Date
	ArrivalDate       datatypes.Date
}

type FlightToRegion struct {
	ID          uint `gorm:"primaryKey"`
	FlightRefer int
	RegionRefer int
	Flight      Flight `gorm:"foreignKey:FlightRefer"`
	Region      Region `gorm:"foreignKey:RegionRefer"`
}