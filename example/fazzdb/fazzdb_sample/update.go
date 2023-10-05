package fazzdb_sample

import (
	"github.com/mikaelim-id/go-apt/example/fazzdb/fazzdb_sample/model"
	"github.com/mikaelim-id/go-apt/pkg/fazzdb"
)

func UpdateAuthor(query *fazzdb.Query) {
	row, err := query.Use(model.AuthorModel()).
		Where("id", 1).
		First()
	if nil != err {
		panic(err)
	}

	author := row.(*model.Author)
	author.Name = "J.K Rowling"
	_, err = query.Use(author).Update()
	if nil != err {
		panic(err)
	}
}

func UpdateBook(query *fazzdb.Query) {
	row, err := query.Use(model.BookModel()).
		First()
	if nil != err {
		panic(err)
	}

	book := row.(*model.Book)
	book.Title = "Harry Potter"
	_, err = query.Use(book).Update()
	if nil != err {
		panic(err)
	}
}
