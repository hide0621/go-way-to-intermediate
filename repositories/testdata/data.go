package testdata

import "go-way-to-intermediate/models"

var ArticleTestData = []models.Article{models.Article{
	ID:        1,
	Title:     "firstPost",
	Contetnts: "This is my first blog",
	UserName:  "saki",
	NiceNum:   2,
},
	models.Article{
		ID:        2,
		Title:     "2nd",
		Contetnts: "Second blog post",
		UserName:  "saki",
		NiceNum:   4,
	},
}
