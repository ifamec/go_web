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