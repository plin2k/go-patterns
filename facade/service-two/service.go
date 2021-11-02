package service_two

import (
	"time"
)

type Post struct {
	Title     string
	CreatedAt time.Time
	Publish   bool
}

func NewPost(title string) *Post {
	return &Post{
		Title:     title,
		CreatedAt: time.Now(),
	}
}

func (p *Post) IsPublic() bool {
	if p.CreatedAt.Before(time.Now()) {
		return true
	}
	return false
}

func (p *Post) SetPublic() {
	p.Publish = true
}
