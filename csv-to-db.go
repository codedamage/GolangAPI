package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"git.qix.sx/hackathon/01-go/tank-rush/hackathon-2020/models"
	"io"
	"log"
	"os"
)

func readCsv(name string) *csv.Reader {
	//var csvToMap map[int][]string
	// Open the file
	csvfile, err := os.Open(name)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	return csv.NewReader(csvfile)
}

func ProductsToDb() {
	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	r := readCsv("products.csv")

	// Iterate through the records
	i := 0
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i > 0 {
			// Create
			db.Create(&models.Product{Code: record[1], Title: record[0]})
		}
		i++
	}

}
func ReviewsToDb() {
	// Migrate the schema
	db.AutoMigrate(&models.Review{})

	r := readCsv("reviews.csv")

	// Iterate through the records
	i := 0
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i > 0 {
			var product models.Product
			db.First(&product, "code = ?", record[0])
			// Create
			db.Create(&models.Review{Asin: record[0], Title: record[1], Content: record[2], ProdId: product.ID})
		}
		i++
	}

}

func clearDB() {
	db.Where("1 = 1").Delete(&models.Product{})
	db.Where("1 = 1").Delete(&models.Review{})
}

func initialSetup() {
	fmt.Print(`Do you want to import reviews and products from csv files located in project folder root?
  1) Import nothing
  2) Import only products
  3) Import only reviews
  4) Import all files
  5) Soft delete all existing data and place new pne from files
Please enter 1..5 and press ENTER: `)
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}
	switch result {
	case '1':
		fmt.Println("Nothing is imported, starting webserver...")
		break
	case '2':
		if _, err := os.Stat("products.csv"); os.IsNotExist(err) {
			fmt.Println("File 'products.csv' not fond...")
			break
		} else {
			fmt.Println("Only product are imported, starting webserver...")
			ProductsToDb()
			break
		}
	case '3':
		if _, err := os.Stat("reviews.csv"); os.IsNotExist(err) {
			fmt.Println("File 'reviews.csv' not fond...")
			break
		} else {
			fmt.Println("Only reviews are imported, starting webserver...")
			ReviewsToDb()
			break
		}
	case '4':
		if _, err := os.Stat("products.csv"); os.IsNotExist(err) {
			fmt.Println("File 'products.csv' not fond...")
		} else {
			fmt.Println("Product are imported, starting webserver...")
			ProductsToDb()
		}
		if _, err := os.Stat("reviews.csv"); os.IsNotExist(err) {
			fmt.Println("File 'reviews.csv' not fond...")
		} else {
			fmt.Println("Reviews are imported, starting webserver...")
			ReviewsToDb()
		}
		break
	case '5':
		fmt.Println("All data purged, inserting new one from files...")
		clearDB()
		if _, err := os.Stat("products.csv"); os.IsNotExist(err) {
			fmt.Println("File 'products.csv' not fond...")
		} else {
			fmt.Println("Product are imported, starting webserver...")
			ProductsToDb()
		}
		if _, err := os.Stat("reviews.csv"); os.IsNotExist(err) {
			fmt.Println("File 'reviews.csv' not fond...")
		} else {
			fmt.Println("Reviews are imported, starting webserver...")
			ReviewsToDb()
		}
		break
	default:
		return
	}
}
