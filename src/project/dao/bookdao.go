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