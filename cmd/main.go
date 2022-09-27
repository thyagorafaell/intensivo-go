package main

import (
	"database/sql"
	"fmt"
	"intensivo-go/internal/order/infra/database"
	"intensivo-go/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 100,
		Tax:   10,
	}

	output, err := uc.Execute(input)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The final price is %f.", output.FinalPrice)

}
