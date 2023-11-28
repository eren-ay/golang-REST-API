package models

type Show struct {
	ID    string
	Title string `gorm:"type:text"`
}

/*
func (show *Show) BeforeCreate(tx *gorm.DB) error {
	return nil
}
*/
