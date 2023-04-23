package models

type Post struct {
	Id    uint `gorm:"primaryKey"`
	Title string
	Body  string
}

func init() {
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(Post{})
}

func Get() []Post {
	var allPost []Post
	result := Db.Find(&allPost)
	if result.Error != nil {
		panic(result.Error)
	}

	return allPost
}

func (data Post) Create() {
	result := Db.Create(data)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
