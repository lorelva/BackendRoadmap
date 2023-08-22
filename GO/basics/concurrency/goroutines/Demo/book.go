package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)

}

var books = []Book{
	{
		ID:            1,
		Title:         "A Soul of Ash and Blood",
		Author:        "Jennifer L. Armentrout",
		YearPublished: 2019,
	},

	{
		ID:            2,
		Title:         "The Perks of Being a Wallflower",
		Author:        "Stephen Chbosky",
		YearPublished: 2012,
	},
	{
		ID:            3,
		Title:         "Paper Towns",
		Author:        "John Green",
		YearPublished: 2015,
	},
	{
		ID:            4,
		Title:         "The Trial",
		Author:        "Franz Kafka",
		YearPublished: 1925,
	},
	{
		ID:            5,
		Title:         "Atomic Habits",
		Author:        "James Clear",
		YearPublished: 2018,
	},
}
