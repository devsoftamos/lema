package database

import (
	"lema/src/models"
	"log"
)

var users = []models.User{
	{
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
	},

	{
		Name:  "John Smith",
		Email: "johnsmith@gmail.com",
	},
}

func SeedUsers() {

	for i, _ := range users {
		var newUser models.User
		query := SqliteDB.Debug().Model(&models.User{}).
			Where(models.User{Email: users[i].Email}).
			FirstOrCreate(&users[i]).Scan(&newUser)

		if query.Error != nil {
			log.Fatalf("cannot seed users table: %v", query.Error)
		}

		go func(user models.User) {
			addressData := models.Address{
				UserID: user.ID,
				Street: "Ikeja",
				City:   "Ikeja",
				State:  "Lagos",
				Zip:    "10006",
			}
			SqliteDB.Where("user_id = ?", user.ID).FirstOrCreate(&addressData)
		}(newUser)
	}

}
