package main

import (
	"api_golang/internal/app"
	"api_golang/internal/users"
	"fmt"
)

func main() {
	db := app.InitDb(app.DB_URL)
	defer app.CloseDb(db)

	fmt.Println("Migrating...")
	db.AutoMigrate(users.User{})
	fmt.Println("Migration complete")

}
