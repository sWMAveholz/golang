package main
import "strings"
import "strconv"
//
type Page[]string

//
type Book[]Page

//
type Index map[string][]int

func MakeIndex(book Book) Index{
	idx := make(Index)
	for i,page := range book{
		for _, word := range page{
			pages := idx[word]
		idx[word] = append(pages,i)
		}
	}
	return idx
}

//Stringer

func (b Book)String()string{
	stringBook := []string{}
	for _,p := range b{
		stringBook = append(stringBook, p.String())
	}
	return strings.Join(stringBook," ")
	
}

func(p Page)String() string{
	return strings.Join(p," ")
}

func (ind Index) String()string{
	stringIndex := []string{}
	for key, value := range ind{
		stringRow := []string{}
		for _,intvalue := range value{
			stringRow = append(stringRow, strconv.Itoa(intvalue))
		}
		stringIndex = append(stringIndex, key + " + " + strings.Join(stringRow, " ")+ " ")

	}
	return strings.Join(stringIndex, " ")
}