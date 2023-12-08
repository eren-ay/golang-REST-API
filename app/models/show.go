package models

type Show struct {
	ID    string `bson:"id"`
	Title string `bson:"title"`
}

/*
func (show *Show) BeforeCreate(tx *gorm.DB) error {
	return nil
}
*/
