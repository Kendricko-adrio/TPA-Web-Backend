package model

type UnsuspendRequestType struct {
	TypeID   int    `json:"typeId" gorm:"primaryKey"`
	TypeName string `json:"typeName"`
}
