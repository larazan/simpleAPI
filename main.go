package main

import (
	"fmt"
	"log"
	"simpleAPI/book"
	"simpleAPI/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Db connection error")
	}

	fmt.Println("Database connection success")

	db.AutoMigrate(&book.Book{})

	//
	// Repository
	//
	bookRepository := book.NewRepository(db)
	// bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// bookRequest := book.BookRequest{
	// 	Title:       "$100 Startup",
	// 	Description: "Good book for entreprenuer",
	// 	Price:       "110000",
	// 	Rating:      4,
	// 	Discount:    0,
	// }

	// bookService.Create(bookRequest)

	// repository find
	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title :", book.Title)
	// }

	// repository create
	// book := book.Book{
	// 	Title:       "$100 Startup",
	// 	Description: "Good book for entreprenuer",
	// 	Price:       110000,
	// 	Rating:      4,
	// 	Discount:    0,
	// }

	// bookRepository.Create(book)

	//
	//CRUD
	//

	// book := book.Book{}
	// book.Title = "Man Tiger"
	// book.Price = 90000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "buku bagus dari eka kurnia"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book")
	// }

	// Read
	// var book book.Book

	// err = db.Debug().First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error find book record")
	// }

	// fmt.Println("Title :", book.Title)
	// fmt.Println("book object %v", book)

	// var books []book.Book

	// err = db.Debug().Where("title = ?", "Man Tiger").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("Error find book record")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	// Update
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error find book record")
	// }

	// book.Title = "Man Tiger (Revised edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("Error updating book record")
	// }

	// Delete
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("Error find book record")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("Error delete book record")
	// }

	//

	router := gin.Default()

	v1 := router.Group("/v1")

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/books/:id", bookHandler.BooksHandler)
	// v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	// v2 := router.Group("/v2")

	router.Run()
}

// main
// handler
// service
// repository
// db
// mysql
