package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID	string	`json:"id"` //xác định tên trường và chuyển đổi thành Json
	Title	string	`json:"title"`
	Author	string	`json:"author"`
	Quantity	int	`json:"quantity"`
}

var books = []book {
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2}, 
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5}, 
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) { //GET method
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book 
	if err := c.BindJSON(&newBook); err != nil {	//POST method
													//BindJSON để gán dữ liệu Json bên ngoài vào struct
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook) //IndentedJSON căn lề và xuống dòng, không thì dùng JSON
}

func getBookById(c *gin.Context) {
	id := c.Param("id")

	for i, b := range books{
		if id == b.ID{
			c.IndentedJSON(http.StatusOK, books[i]) 
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Not found book with id: " + id}) // dùng gin.H là 1 map[string]any để hiển thị lỗi, + để nối chuỗi
}

func main() {
	router := gin.Default() //tạo router để xử lý các request 
	router.GET("/book/getAll",getBooks)
	router.POST("/book/createBook", createBook)
	router.GET("book/getById/:id", getBookById)
	router.Run("localhost:8080")
}
