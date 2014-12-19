package models

type Post struct {
	Id        int    `db:"id"`
	Title     string `db:"title"`
	body      string `db:"body"`
	Timestamp int32  `db:"timestamp"`
}

func (p *Post) Save() error {
	return Db.Insert(p)
}

func (p *Post) FindByTimestamp(timestamp int32) error {
	err := Db.SelectOne(p, "SELECT * FROM posts WHERE timestamp = :ts", map[string]interface{}{"ts": timestamp})
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) Remove() error {
	_, err := Db.Delete(p)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPosts() []Post {
	var posts []Post
	_, err := Db.Select(&posts, "select * from posts")
	if err != nil {
		panic(err)
	}
	return posts
}
