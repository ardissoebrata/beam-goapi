package main

import (
	"api_golang/internal/app"
	"api_golang/internal/users"
	"fmt"

	"github.com/jaswdr/faker"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := app.InitDb(app.DB_URL)
	defer app.CloseDb(db)

	fake := faker.New()

	fmt.Println("Seeding...")

	fmt.Print("Seeding users...")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	for i := 0; i < 100; i++ {
		user := users.User{
			Full_name: fake.Person().Name(),
			Email:     "user" + fmt.Sprintf("%d", i+1) + "@example.com",
			Password:  string(hashedPassword),
		}

		db.Create(&user)
	}
	fmt.Println("done")

	fmt.Println("Seed complete")
}
