package printshop

import "github.com/VojtechVitek/go-trello"
import "github.com/pkg/errors"
import "strings"
import "github.com/shurcooL/github_flavored_markdown"
import "sync"

var TrelloAPIError = errors.New("Error calling the Trello API")

const (
	MetaDataListName = "meta"
	TOKEN_ENV        = "TRELLO_TOKEN"
	KEY_ENV          = "TRELLO_APIKEY"
	MaxWorkers       = 5
)

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

type listID string

type BoardContainer struct {
	board *trello.Board
	lists []*trello.List
	cards map[listID][]trello.Card
	mux   sync.Mutex
}

func NewContainer(c *trello.Client, boardID string) (*BoardContainer, error) {
	board, err := c.Board(boardID)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to build Board container")
	}
	container := &BoardContainer{
		board: board,
	}
	lists, err := board.Lists()
	if err != nil {
		return container, errors.Wrapf(err, "Cannot get Lists for Board")
	}
	for _, list := range lists {
		container.lists = append(container.lists, &list)
	}

	return container, nil
}

func (c *BoardContainer) RetrieveCards() error {
	jobs := make(chan *trello.List, len(c.lists))
	results := make(chan []trello.Card, len(c.lists))
	// create the worker pool
	for i := 0; i < MaxWorkers; i++ {
		go c.ListWorker(i, jobs, results)
	}
	// Queue up the jobs
	for _, list := range c.lists {
		jobs <- list
	}
	close(jobs)

	for range c.lists {
		cardList := <-results
		id := listID(cardList[0].IdList)
		c.cards[id] = cardList
	}
	return nil
}

func (c *BoardContainer) ListWorker(id int, listJob <-chan *trello.List, results chan<- []trello.Card) {
	for j := range listJob {
		if cards, err := j.Cards(); err == nil {
			results <- cards
		}
	}
}

func NewEmail(b *trello.Board) *Email {
	lists, err := b.Lists()
	if err != nil {
		panic(err)
	}
	email := &Email{}
	for _, v := range lists {
		if v.Name == MetaDataListName {
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
