package migration

import (
	"github.com/mikaelim-id/go-apt/pkg/esfazz/eventstore/eventpostgres"
	"github.com/mikaelim-id/go-apt/pkg/fazzdb"
)

var Version1 = fazzdb.MigrationVersion{
	Tables: []*fazzdb.MigrationTable{
		eventpostgres.CreateEventsTable("events"),
		//snappostgres.CreateSnapshotsTable("snapshots"),
		fazzdb.CreateTable("account", func(table *fazzdb.MigrationTable) {
			table.Field(fazzdb.CreateUuid("id").Primary())
			table.Field(fazzdb.CreateBigInteger("version"))
			table.Field(fazzdb.CreateString("name"))
			table.Field(fazzdb.CreateBigInteger("balance"))
			table.Field(fazzdb.CreateTimestampTz("deleted_time", 0).Nullable())
		}),
	},
}
