package printshop

import "github.com/VojtechVitek/go-trello"
import "errors"

var TrelloAPIError = errors.New("Error Calling the Trello API")

const META = "meta"

type Article struct {
	Title    string
	BodyHTML string
}

type Section struct {
	Title    string
	Articles []Article
}

type MetaData struct {
	subject string
}

type Email struct {
	meta     MetaData
	sections []Section
}

func NewEmail(b *trello.Board) *Email {
	lists, err := b.Lists()
	if err != nil {
		panic(err)
	}
	email := &Email{}
	for _, v := range lists {
		if v.Name == META {
			email.meta = *NewMetaData(&v)
		} else {
			email.sections = append(email.sections, *NewSection(&v))
		}
	}
	return email
}

func NewMetaData(l *trello.List) *MetaData {
	cards, err := l.Cards()

}

func NewSection(l *trello.List) *Section {

}
