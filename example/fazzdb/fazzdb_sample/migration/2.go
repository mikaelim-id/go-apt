package migration

import (
	"github.com/mikaelim-id/go-apt/example/fazzdb/fazzdb_sample/seed"
	"github.com/mikaelim-id/go-apt/pkg/fazzdb"
)

var bookStatusEnum = fazzdb.NewEnum(
	"book_status",
	"BORROWED", "AVAILABLE",
)

var Version2 = fazzdb.MigrationVersion{
	Enums: []*fazzdb.MigrationEnum{
		bookStatusEnum,
	},
	Tables: []*fazzdb.MigrationTable{
		fazzdb.AlterTable("books", func(table *fazzdb.MigrationTable) {
			table.Field(fazzdb.AddEnum("status", bookStatusEnum))
			table.Field(fazzdb.AddInteger("year"))
		}),
	},
	Seeds: []fazzdb.SeederInterface{
		seed.AuthorSeeder(),
		seed.BookSeeder(),
		seed.BookObjectSeeder(),
	},
	Raw: `CREATE TABLE raw_queries (id serial primary key, name varchar);`,
}
