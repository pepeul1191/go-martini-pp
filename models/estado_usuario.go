package models

type Estado struct {
    ID      uint `gorm:"primary_key"` 
    Nombre   string
}

func (Estado) TableName() string {
    return "estado_usuarios"
}