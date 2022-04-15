package address

import (
	"github.com/wewillapp-com/we-address/internal/database"
)

type ProvinceModel struct {
	ID     uint
	NameTh string        `gorm:"type:varchar(255);not null"`
	NameEn string        `gorm:"type:varchar(255);not null"`
	Amphur []AmphurModel `gorm:"foreignkey:ProvinceID"`
}

func (ProvinceModel) TableName() string {
	return "provinces"
}

//Get all provinces
func GetProvinceList() (*[]ProvinceModel, error) {
	provinces := []ProvinceModel{}
	if err := database.DB.Find(&provinces).Error; err != nil {
		return nil, err
	}
	return &provinces, nil
}

//Get specific province by id
func GetProvinceById(id uint) (*ProvinceModel, error) {
	province := ProvinceModel{}
	if err := database.DB.First(&province, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &province, nil
}

func SearchProvince(q string) (*[]ProvinceModel, error) {
	provinces := []ProvinceModel{}
	if err := database.DB.Where("name LIKE ?", "%"+q+"%").Find(&provinces).Error; err != nil {
		return nil, err
	}
	return &provinces, nil
}
