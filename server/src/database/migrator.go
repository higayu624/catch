package database

import (
	"log"

	"catch/database/table"

	"github.com/bxcodec/faker"
	"gorm.io/gorm"
)

// 　Migrate make table layout
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&table.Customer{},
		&table.Store{},
		&table.Clip{},
		&table.User{},
		&table.Location{},
		&table.Category{},
	)
}

// Seed add datas to tables TODO seedの実装
func Seed(db *gorm.DB) {
	restaurantCategories := []string{
		"和食", "寿司", "焼肉", "ラーメン", "居酒屋",
		"フレンチ", "イタリアン", "カフェ", "中華", "ベーカリー",
		"ファストフード", "ベジタリアン", "韓国料理", "タイ料理", "メキシカン",
	}
	batchSize := 1

	numberOfUser := 5
	numberOfCustomer := 5
	// numberOfStore := 1

	var users []table.User
	var categorys []table.Category
	var customers []table.Customer
	var stores []table.Store

	// create data for seed by faker
	seedCategory(restaurantCategories, &categorys)
	seedUser(numberOfUser, &users)
	seedCustomer(numberOfCustomer, &customers)

	// seed datas
	db.CreateInBatches(users, batchSize)
	db.CreateInBatches(customers, batchSize)
	db.CreateInBatches(categorys, batchSize)
	db.CreateInBatches(stores, batchSize)
}

func seedCategory(restaurantCategories []string, categorys *[]table.Category) {
	for i := range restaurantCategories {
		category := table.Category{}
		category.Name = restaurantCategories[i]
		*categorys = append(*categorys, category)
	}
}

func seedUser(numberOfUser int, users *[]table.User) {
	for range numberOfUser {
		user := table.User{}
		err := faker.FakeData(&user)
		if err != nil {
			log.Print("Seed user table can't complete:", err)
		}
		err = faker.FakeData(&user.Locations)
		if err != nil {
			log.Print("Seed Location table can't complete:", err)
		}
		*users = append(*users, user)
	}
}

func seedCustomer(numberOfCustomer int, customers *[]table.Customer) {
	for range numberOfCustomer {
		customer := table.Customer{}
		err := faker.FakeData(&customer)
		if err != nil {
			log.Print("Seed customer table can't complete:", err)
		}
		err = faker.FakeData(&customer.Store)
		if err != nil {
			log.Print("Seed store table can't complete:", err)
		}
		err = faker.FakeData(&customer.Store.Clips)
		if err != nil {
			log.Print("Seed Customer table can't complete:", err)
		}
		*customers = append(*customers, customer)
	}
}
