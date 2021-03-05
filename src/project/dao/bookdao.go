package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"go_web/src/project/model"
	"go_web/src/project/utils"
)

func GetBooks() (books []*model.Book, err error) {
	sqlQuery := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book *model.Book = &model.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return
}

func AddBook(b *model.Book) (err error) {
	sqlQuery := "insert into books (title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err = utils.Db.Exec(sqlQuery, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return
	}
	return nil
}

func DeleteBook(id int) (err error) {
	sqlQuery := "delete from books where id = ?"
	_, err = utils.Db.Exec(sqlQuery, id)
	if err != nil {
		return
	}
	return nil
}

func GetBookById(id int) (book *model.Book, err error) {
	sqlQuery := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	book = &model.Book{}
	row := utils.Db.QueryRow(sqlQuery, id)
	_ = row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return
}

func UpdateBook(b *model.Book) (err error) {
	sqlQuery := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err = utils.Db.Exec(sqlQuery, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Id)
	return
}

func GetPageBooks(pageNumber int) (page *model.Page, err error) {
	// get total records
	sqlQuery := "select count(*) from books"
	var totalRecords int
	row := utils.Db.QueryRow(sqlQuery)
	err = row.Scan(&totalRecords)
	// set up page size
	var pageSize int = 4
	// get pages
	var totalPages int = totalRecords / pageSize
	if totalRecords % pageSize != 0 {
		totalPages += 1
	}
	// get books in current page
	sqlGetBooks := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	rows, err := utils.Db.Query(sqlGetBooks, (pageNumber-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book = make([]*model.Book, 0)
	for rows.Next() {
		book := &model.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	// page
	page = &model.Page{
		Books:        books,
		PageNumber:   pageNumber,
		PageSize:     pageSize,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}
	return
}