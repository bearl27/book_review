package repositories

type book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

type books []book

type BookRepository interface {
	FindBooks() ([]book, error)
	GetBookByID(id string) (book, error)
	CreateBook(book book) (book, error)
	UpdateBook(id string, book book) (book, error)
	DeleteBook(id string) error
}
