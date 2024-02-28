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

/*
1. Выключить докер компоус
2. Удалить докер компоус
3. Удалить старые имеджи
4. Перебилдить все имеджи
5. Запустить докер компоус
*/
