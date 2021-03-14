package resolver

import "github.com/kendricko-adrio/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen --verbose
type Resolver struct{
	ChatSockets map[int]chan *model.Chat
}

func NewResolver() *Resolver{
	return &Resolver{
		ChatSockets: map[int]chan *model.Chat{},
	}
}
