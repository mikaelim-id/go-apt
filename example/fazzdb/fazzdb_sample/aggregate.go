package fazzdb_sample

import (
	"fmt"
	"github.com/mikaelim-id/go-apt/example/fazzdb/fazzdb_sample/model"
	"github.com/mikaelim-id/go-apt/pkg/fazzdb"
)

func Sum(query *fazzdb.Query) {
	sum, err := query.Use(model.BookModel()).
		Sum("stock")

	if nil != err {
		panic(err)
	}

	fmt.Println("--------------")
	fmt.Println("Total stock of books:", *sum)
}

func Count(query *fazzdb.Query) {
	count, err := query.Use(model.BookModel()).
		Count()

	if nil != err {
		panic(err)
	}

	fmt.Println("--------------")
	fmt.Println("Book Count:", *count)
}

func Max(query *fazzdb.Query) {
	max, err := query.Use(model.BookModel()).
		WhereOp("year", fazzdb.OP_LESS_THAN, 1960).
		Max("stock")

	if nil != err {
		panic(err)
	}

	fmt.Println("--------------")
	fmt.Println("Max stock of book released before 1960:", *max)
}

func Min(query *fazzdb.Query) {
	min, err := query.Use(model.BookModel()).
		WhereOp("year", fazzdb.OP_LESS_THAN, 1960).
		Min("stock")

	if nil != err {
		panic(err)
	}

	fmt.Println("--------------")
	fmt.Println("Min stock of book released before 1960:", *min)
}

func Avg(query *fazzdb.Query) {
	avg, err := query.Use(model.BookModel()).
		WhereOp("year", fazzdb.OP_LESS_THAN, 1960).
		Max("stock")

	if nil != err {
		panic(err)
	}

	fmt.Println("--------------")
	fmt.Println("Average stock of book released before 1960:", *avg)
}
