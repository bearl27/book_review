package domain

import "fmt"

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

const (
	DefaultBookTitle   = "無題の本"
	DefaultBookSummary = "概要なし"
)

const (
	BookTitleMaxLength   = 100
	BookSummaryMaxLength = 1000
)

func (b Book) Validate() error {
	if len(b.ID) == 0 {
		return fmt.Errorf("id must not be empty")
	}
	if len(b.Title) == 0 || len(b.Title) > BookTitleMaxLength {
		return fmt.Errorf("title must be between 1 and %d characters", BookTitleMaxLength)
	}
	if len(b.Summary) > BookSummaryMaxLength {
		return fmt.Errorf("summary must be less than %d characters", BookSummaryMaxLength)
	}
	return nil
}

type BookList struct {
	Books []Book `json:"books"`
}
