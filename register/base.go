package register

import (
	reposiotires "github.com/ahmedshakshak/books-recommender/adapters/repositories"
	"github.com/ahmedshakshak/books-recommender/controllers"
	"github.com/ahmedshakshak/books-recommender/models"
	"github.com/ahmedshakshak/books-recommender/notifiers"
	"github.com/ahmedshakshak/books-recommender/usecases"
)

type Register struct {
	App controllers.App
}

func NewRegister() Register {
	db := models.GetDBConnection()
	bm := models.NewBookManager(db)

	n := notifiers.NewNotifier()
	ur := reposiotires.NewUser(bm, n)
	br := reposiotires.NewBook(bm)

	uc := usecases.NewUser(&ur)
	bc := usecases.NewBook(&br)

	user := controllers.NewUser(uc)
	book := controllers.NewBook(bc)
	app := controllers.NewApp(user, book)

	return Register{App: app}
}
