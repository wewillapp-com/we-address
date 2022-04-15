package address

import (
	"github.com/wewillapp-com/we-address/internal/database"
)

type AmphurModel struct {
	ID         uint
	NameTh     string `gorm:"type:varchar(255);not null"`
	NameEn     string `gorm:"type:varchar(255);not null"`
	ProvinceID uint
	Province   ProvinceModel
	District   []DistrictModel `gorm:"foreignkey:AmphurID"`
}

func (AmphurModel) TableName() string {
	return "amphurs"
}
func GetAmphurList() (*[]AmphurModel, error) {
	amphurs := []AmphurModel{}
	if err := database.DB.Find(&amphurs).Error; err != nil {
		return nil, err
	}
	return &amphurs, nil
}

func GetAmphurById(id uint) (*AmphurModel, error) {
	amphur := AmphurModel{}
	if err := database.DB.First(&amphur, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &amphur, nil
}
