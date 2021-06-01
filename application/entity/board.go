package entity

type Board struct {
	Audit
	Id          int64
	Name        string
	Description string
}

type Category struct {
	Audit
	Id      int64
	Name    string
	BoardId int64
}

type Article struct {
	Audit
	Id         int64
	Content    string
	BoardId    int64
	CategoryId int64
}

type Comment struct {
	Audit
	Id        int64
	Content   string
	ArticleId int64
}
