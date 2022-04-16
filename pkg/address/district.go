package address

import (
	"github.com/wewillapp-com/we-address/internal/database"
)

type DistrictModel struct {
	ID       uint
	ZipCode  string `gorm:"type:varchar(20);"`
	NameTh   string `gorm:"type:varchar(255);not null"`
	NameEn   string `gorm:"type:varchar(255);not null"`
	AmphurID uint
	Amphur   AmphurModel
}

//force gorm to use plural table name
func (DistrictModel) TableName() string {
	return "districts"
}
func GetDistrictList() (*[]DistrictModel, error) {
	districts := []DistrictModel{}
	if err := database.DB.Find(&districts).Error; err != nil {
		return nil, err
	}
	return &districts, nil
}

func GetDistrictById(id uint) (*DistrictModel, error) {
	district := DistrictModel{}
	if err := database.DB.First(&district, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &district, nil
}
