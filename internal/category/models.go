package category

type Category struct {
	ID    uint `gorm:"unique;primaryKey;autoIncrement"`
	Title string
}

type CategoryView struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
