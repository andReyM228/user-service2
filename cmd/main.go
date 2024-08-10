package main

import (
	"embed"
	"user_service/internal/app"

	_ "github.com/lib/pq"
)

const serviceName = "user_service"

//go:embed dbschema/migrations
var dbMigrationFS embed.FS

func main() {
	a := app.New(serviceName)
	a.Run(dbMigrationFS)
}
