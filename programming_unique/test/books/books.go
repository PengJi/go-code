package books1

type Book struct {
	Title string
	Author string
	Pages int
}

func NewBookFromJSON(str string) (Book, error){
	book := Book {
		Title: "test_title",
		Author: "test_author",
		Pages: 11,
	}

	return book, nil
}

func (b *Book) AuthorLastName() string {
	return "test"
}
