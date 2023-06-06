package models

type Book struct {
	Id      int
	Author  string
	Release int
	Name    string
}

var Books = []Book{
	{Id: 1, Author: "Author 1", Release: 2000, Name: "Book 1"},
	{Id: 2, Author: "Author 2", Release: 2011, Name: "Book 2"},
	{Id: 3, Author: "Author 3", Release: 2002, Name: "Book 3"},
}
