package domain

type Role struct {
	ID   int64  `gorm:"primary_key; auto_increment"`
	Name string `gorm:"size:100; not null; type:varchar(100);uniqueIndex" validate:"min=5"`
}
