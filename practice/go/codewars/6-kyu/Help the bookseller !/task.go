package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Book struct {
	code string
}

func (b Book) Category() string {
	return strings.ToUpper(string(b.code[0]))
}
func (b Book) Count() int {
	re := regexp.MustCompile("[0-9]+")
	digits := re.FindAllString(b.code, -1)
	v := strings.Join(digits, "")
	count, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return count
}

func mapStringBooksToBooks(strBooks []string) (books []Book) {
	for _, strBook := range strBooks {
		books = append(books, Book{strBook})
	}
	return books
}

type Stock struct {
	s map[string]int
}

func (s Stock) CountBooksCategories(ourBooks []Book, categoriesRequested []string) {
	for _, book := range ourBooks {
		category := book.Category()
		count := book.Count()
		for _, cat := range categoriesRequested {
			if _, ok := s.s[cat]; !ok {
				s.s[cat] = 0
			}
			if cat == category {
				s.s[category] += count
				break
			}
		}
	}
}

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	stock := Stock{make(map[string]int)}
	stock.CountBooksCategories(mapStringBooksToBooks(listArt), listCat)
	stockList := make([]string, 0)
	for category, count := range stock.s {
		stockList = append(stockList, fmt.Sprintf("(%s : %d)", category, count))
	}
	return strings.Join(stockList, " - ")
}

/*


A bookseller has lots of books classified in 26 categories labeled A, B, ... Z. Each book has a code c of 3, 4, 5 or more characters. The 1st character of a code is a capital letter which defines the book category.

In the bookseller's stocklist each code c is followed by a space and by a positive integer n (int n >= 0) which indicates the quantity of books of this code in stock.

For example an extract of a stocklist could be:

L = {"ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"}.
or
L = ["ABART 20", "CDXEF 50", "BKWRK 25", "BTSQZ 89", "DRTYM 60"] or ....
You will be given a stocklist (e.g. : L) and a list of categories in capital letters e.g :

M = {"A", "B", "C", "W"}
or
M = ["A", "B", "C", "W"] or ...
and your task is to find all the books of L with codes belonging to each category of M and to sum their quantity according to each category.

For the lists L and M of example you have to return the string (in Haskell/Clojure/Racket/Prolog a list of pairs):

(A : 20) - (B : 114) - (C : 50) - (W : 0)
where A, B, C, W are the categories, 20 is the sum of the unique book of category A, 114 the sum corresponding to "BKWRK" and "BTSQZ", 50 corresponding to "CDXEF" and 0 to category 'W' since there are no code beginning with W.

If L or M are empty return string is "" (Clojure/Racket/Prolog should return an empty array/list instead).

Notes:
In the result codes and their values are in the same order as in M.
See "Samples Tests" for the return.

*/
