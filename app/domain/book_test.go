package domain

import (
	"fmt"
	"strings"
	"testing"
)

func TestBookTitle_Validate(t *testing.T) {
	maxTitle := strings.Repeat("a", BookTitleMaxLength)
	overMaxTitle := strings.Repeat("a", BookTitleMaxLength+1)
	maxSummary := strings.Repeat("b", BookSummaryMaxLength)
	overMaxSummary := strings.Repeat("b", BookSummaryMaxLength+1)

	type fields struct {
		ID      string
		Title   string
		Summary string
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "valid book",
			fields: fields{
				ID:      "1",
				Title:   "Valid Title",
				Summary: "Valid Summary",
			},
			want: nil,
		},
		{
			name: "empty id",
			fields: fields{
				ID:      "",
				Title:   "Valid Title",
				Summary: "Valid Summary",
			},
			want: fmt.Errorf("id must not be empty"),
		},
		{
			name: "empty title",
			fields: fields{
				ID:      "1",
				Title:   "",
				Summary: "Valid Summary",
			},
			want: fmt.Errorf("title must be between 1 and %d characters", BookTitleMaxLength),
		},
		{
			name: "title max length",
			fields: fields{
				ID:      "1",
				Title:   maxTitle,
				Summary: "Valid Summary",
			},
			want: nil,
		},
		{
			name: "title over max length",
			fields: fields{
				ID:      "1",
				Title:   overMaxTitle,
				Summary: "Valid Summary",
			},
			want: fmt.Errorf("title must be between 1 and %d characters", BookTitleMaxLength),
		},
		{
			name: "summary max length",
			fields: fields{
				ID:      "1",
				Title:   "Valid Title",
				Summary: maxSummary,
			},
			want: nil,
		},
		{
			name: "summary over max length",
			fields: fields{
				ID:      "1",
				Title:   "Valid Title",
				Summary: overMaxSummary,
			},
			want: fmt.Errorf("summary must be less than %d characters", BookSummaryMaxLength),
		},
		{
			name: "valid book with empty summary",
			fields: fields{
				ID:      "1",
				Title:   "Valid Title",
				Summary: "",
			},
			want: nil,
		},
		{
			name: "valid book with default title and summary",
			fields: fields{
				ID:      "1",
				Title:   DefaultBookTitle,
				Summary: DefaultBookSummary,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Book{
				ID:      tt.fields.ID,
				Title:   tt.fields.Title,
				Summary: tt.fields.Summary,
			}
			got := b.Validate()
			if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) {
				t.Errorf("Book.Validate() = %v, want %v", got, tt.want)
				return
			}
			if got != nil && tt.want != nil && got.Error() != tt.want.Error() {
				t.Errorf("Book.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
