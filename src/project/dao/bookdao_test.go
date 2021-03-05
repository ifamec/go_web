package dao

import (
	_ "database/sql"
	"fmt"
	"go_web/src/project/model"
	"testing"
)

func TestBook (t *testing.T) {
	fmt.Println("Test - bookdao.go")
	// t.Run("Get Books", testGetBooks)
	// t.Run("Add Book", testAddBook)
	// t.Run("Delete Book", testDeleteBook)
	// t.Run("Get Book By Id", testGetBookById)
	// t.Run("Update Book", testUpdateBook)
	t.Run("Get Page Books", testGetPageBooks)
}

func testGetBooks(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	for k, v := range books {
		fmt.Println(k, v)
	}
}
func testAddBook(t *testing.T) {
	err := AddBook(&model.Book{
		Title: "Me Talk Pretty One Day - 2",
		Author: "David Sedaris",
		Price: 16.99,
		Sales: 100,
		Stock: 100,
		ImgPath: "512DFet3JbL",
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
func testDeleteBook(t *testing.T) {
	err := DeleteBook(34)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
func testGetBookById(t *testing.T)  {
	book, err := GetBookById(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(book)
	}
}
func testUpdateBook(t *testing.T)  {
	err := UpdateBook(&model.Book{
		Id: 47,
		Title: "Me Talk Pretty One Day - 47",
		Author: "David Sedaris",
		Price: 18.99,
		Sales: 200,
		Stock: 200,
		ImgPath: "512DFet3JbL",
	})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
func testGetPageBooks(t *testing.T)  {
	page, err := GetPageBooks(9)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(page.Books, page.PageNumber, page.PageSize, page.TotalPages, page.TotalRecords)
	}
}

