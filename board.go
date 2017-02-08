package printshop

import "github.com/VojtechVitek/go-trello"
import "github.com/pkg/errors"
import "strings"
import "github.com/shurcooL/github_flavored_markdown"

var TrelloAPIError = errors.New("Error Calling the Trello API")

const META = "meta"

type Article struct {
	Title    string
	BodyHTML []byte
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
			meta, err := NewMetaData(&v)
			if err == nil {
				email.meta = *meta
			}
		} else {
			section, err := NewSection(&v)
			if err == nil {
				email.sections = append(email.sections, *section)
			}
		}
	}
	return email
}

func NewMetaData(l *trello.List) (*MetaData, error) {
	cards, err := l.Cards()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to retrieve cards for list %s", l.Name)
	}
	for _, card := range cards {
		if strings.ToLower(card.Name) == "subject" {
			return &MetaData{card.Desc}, nil
		}
	}
	return nil, errors.New("No card with 'subject' found in name")
}

func NewSection(l *trello.List) (*Section, error) {
	section := &Section{
		Title:    l.Name,
		Articles: make([]Article, 0),
	}
	cards, err := l.Cards()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to retrieve cards for list %s", l.Name)
	}
	for _, card := range cards {
		article, err := NewArticle(&card)
		if err == nil {
			section.Articles = append(section.Articles, *article)
		}
	}
	return section, nil
}

func NewArticle(card *trello.Card) (*Article, error) {
	return &Article{
		Title:    card.Name,
		BodyHTML: github_flavored_markdown.Markdown([]byte(card.Desc)),
	}, nil
}
