package facade

import (
	service_one "github.com/plin2k/go-patterns/facade/service-one"
	service_two "github.com/plin2k/go-patterns/facade/service-two"
)

type Facade struct {
	User *service_one.User
	Post *service_two.Post
}

func NewFacade(userName string, postTitle string) *Facade {
	facade := &Facade{
		User: service_one.NewUser(userName),
		Post: service_two.NewPost(postTitle),
	}
	return facade
}

func (f *Facade) PublishPost() bool {
	if f.User.IsAdmin() || f.User.IsManager() {
		f.Post.SetPublic()
		return true
	}
	return false
}

func (f *Facade) ChangeUserName(newName string) {
	f.User.Name = newName
}
