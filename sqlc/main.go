package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"tutorial.sqlc.dev/app/product"
)

func main() {
	connstr := "postgresql://root:root@localhost:5432/sqlctest?sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	queries := product.New(db)

	createProduct := product.CreateProductParams{
		Name:      "html",
		Price:     "1000",
		Available: sql.NullBool{Bool: true, Valid: true},
	}

	newproduct, err := queries.CreateProduct(ctx, createProduct)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(newproduct)

	// err = queries.DeleteProduct(ctx, 6)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	productContent, err := queries.GetProduct(ctx, 8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(productContent)


	updateproduct:= product.UpdateProductParams{
		ID:        8,
		Name:      "css",
		Price:     "2000",
		Available: sql.NullBool{Bool: true, Valid: true},
	}
	err = queries.UpdateProduct(ctx,updateproduct)
	if err != nil {
		log.Fatal(err)
	}

	productList, err := queries.ListProducts(ctx)
	fmt.Println(productList)

	sumProductPrice,err:=queries.SumProductPrice(ctx)
	fmt.Println(sumProductPrice)

}
