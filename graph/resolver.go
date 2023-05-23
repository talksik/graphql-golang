package graph

import "github.com/talksik/graphql-golang/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.


var todos = []*model.Todo{
	{
		ID:   "1",
    Text: "laundry",
		Done: false,
	},
	{
		ID: "2",
    Text: "woah",
    Done: false,
	},
	{
		ID:   "3",
		Text: "Hello",
		Done: false,
		User: &model.User{
			ID:   "1",
			Name: "Talksik",
		},
	},
}

type Resolver struct{
  Todos []*model.Todo
}
