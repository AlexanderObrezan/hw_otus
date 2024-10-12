package main

// структура Book
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) SetID(i int) {
	b.id = i
}

func (b *Book) SetTitle(t string) {
	b.title = t
}

func (b *Book) SetAuthor(a string) {
	b.author = a
}

func (b *Book) SetYear(y int) {
	b.year = y
}

func (b *Book) SetSize(s int) {
	b.size = s
}

func (b *Book) SetRate(r float32) {
	b.rate = r
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float32 {
	return b.rate
}

type choice int

// enum для структуры сравнения книг
const (
	year choice = iota
	size
	rate
)

// структура сравнения книг
type CompareBook struct {
	compChoice choice
}

// метод сравнения книг
func (c CompareBook) Compare(firstBook *Book, secondBook *Book) bool {
	switch c.compChoice {
	case 0:
		return firstBook.year > secondBook.year
	case 1:
		return firstBook.size > secondBook.size
	case 2:
		return firstBook.rate > secondBook.rate
	}
	return false
}

func main() {
	// book1 := Book{
	// 	id:     1,
	// 	title:  "TestBook1",
	// 	author: "John Doe",
	// 	year:   1869,
	// 	size:   100,
	// 	rate:   10.0,
	// }

	// book2 := Book{
	// 	id:     3,
	// 	title:  "testBook2",
	// 	author: "Adam Smith",
	// 	year:   1923,
	// 	size:   15,
	// 	rate:   2.6,
	// }

	// fmt.Println(book.GetID(), book.GetTitle(), book.GetAuthor(), book.GetYear(), book.GetRate())

	// fmt.Println(book2.GetID(), book2.GetTitle(), book2.GetAuthor(), book2.GetYear(), book2.GetRate())

	// cb := CompareBook{3}

	// fmt.Println(cb.Compare(&book, &book2))
}
