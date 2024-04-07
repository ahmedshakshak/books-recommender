package controllers

import "net/http"

type App struct {
	userController User
	bookController Book
	Handler        http.Handler
}

func NewApp(user User, book Book) App {
	app := App{
		userController: user,
		bookController: book,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/reading-interval", app.userController.SubmitReadingInterval)
	mux.HandleFunc("/recommend-books", app.bookController.RecommendBooks)

	app.Handler = mux
	return app
}
