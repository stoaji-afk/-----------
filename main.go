package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"example/solid/model"
	"example/solid/service"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var repo model.RepositoryWriter = model.NewSQLiteRepository(db)

	if initializer, ok := repo.(model.RepositoryInitializer); ok {
		if err = initializer.Initialize(); err != nil {
			log.Fatal(err)
		}
	}

	// --- Вариант 1: EmailSender ---
	fmt.Println(model.MainOption1)
	svc1 := service.NewOrderService(repo, model.NewEmailSender())
	err = svc1.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}

	// --- Вариант 2: SMSSender ---
	fmt.Println("\n" + model.MainOption2)
	svc2 := service.NewOrderService(repo, model.NewSMSSender())
	err = svc2.CreateOrder("Мария", []string{"orange", "grape"}, 25.0)
	if err != nil {
		log.Fatal(err)
	}

	// --- Вариант 3: TelegramNotifier ---
	fmt.Println("\n" + model.MainOption3)
	svc3 := service.NewOrderService(repo, model.NewTelegramNotifier())
	err = svc3.CreateOrder("Олег", []string{"coffee", "cake"}, 18.0)
	if err != nil {
		log.Fatal(err)
	}
}
