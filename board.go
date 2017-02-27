package printshop

import "github.com/VojtechVitek/go-trello"
import "github.com/pkg/errors"
import "github.com/shurcooL/github_flavored_markdown"
import "strings"
import "bytes"

// ErrTrelloAPI represents an error with the underlying Trello API calls.
var ErrTrelloAPI = errors.New("Error calling the Trello API")

const (
	// EnvTrelloToken represents the environment variables that holds a particular authorization token for Trello.
	EnvTrelloToken = "TRELLO_TOKEN"
	// EnvTrelloKey represents the environment variable that holds the Trello API Key.
	EnvTrelloKey = "TRELLO_APIKEY"
	// MetaDataListName is the list used to pull in metadata.
	MetaDataListName = "meta"
	// MaxWorkers is the number of worker goroutines that will make parallel calls to Trello's API.
	MaxWorkers = 5
)

// Article represents the link and descriptive text for a part of the email.
type Article struct {
	Title    string
	BodyHTML []byte
}

// Section includes specific Articles and a Title for a particular piece of the email.
type Section struct {
	Title    string
	Articles []Article
}

// MetaData represents the data about the email that will be sent.
type MetaData struct {
	subject string
	from    string
}

// Email is the overarching struct for sending an email.
type Email struct {
	meta     MetaData
	sections []Section
}

type listID string

// BoardContainer will hold the underlying Trello boards, lists, and cards.
type BoardContainer struct {
	board *trello.Board
	lists []*trello.List
	cards map[listID][]trello.Card
}

// NewContainer will create a new BoardContainer with a given Trello Client and
// a specific boardID of the board to build out.
//
// This will build out and retrieve all of the data on a board, including all of the
// cards, lists, metadata, etc., about a board.
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

// RetrieveCards will retrieve all of the cards for a particular BoardContainer
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

// ListWorker is the worker that will retrieve the cards for a particular Trello list.
func (c *BoardContainer) ListWorker(id int, listJob <-chan *trello.List, results chan<- []trello.Card) {
	for j := range listJob {
		if cards, err := j.Cards(); err == nil {
			results <- cards
		}
	}
}

func (e *Email) RenderBody() (string, error) {
	var body bytes.Buffer
	for _, s := range e.sections {
		s.
	}
}

// NewEmail will create a newsletter Email from a given Trello Board
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

// NewMetaData will create the MetaData type from a given Trello List.
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

// NewSection will create a Section out of a Trello List.
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

// NewArticle will create a new Article from a Trello Card.
func NewArticle(card *trello.Card) (*Article, error) {
	return &Article{
		Title:    card.Name,
		BodyHTML: github_flavored_markdown.Markdown([]byte(card.Desc)),
	}, nil
}
