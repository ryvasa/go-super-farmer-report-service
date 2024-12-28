package domain

type City struct {
	ID         int64     `gorm:"primary_key; auto_increment"`
	Name       string    `gorm:"size:100; not null; type:varchar(100);" validate:"min=5"`
	ProvinceID int64     `gorm:"not null"`
	Province   *Province `gorm:"foreignkey:ProvinceID" json:"province,omitempty"`
}
