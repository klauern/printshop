package printshop

import (
	"reflect"
	"sync"
	"testing"

	"os"

	"github.com/VojtechVitek/go-trello"
	"gopkg.in/check.v1"
)

type BoardSuite struct {
	client *trello.Client
}

func TestWithCheck(t *testing.T) {
	check.TestingT(t)
}

var _ = check.Suite(&BoardSuite{})

func (s *BoardSuite) SetUpSuite(c *check.C) {
	token := os.Getenv(TOKEN_ENV)
	key := os.Getenv(KEY_ENV)
	client, err := trello.NewAuthClient(key, &token)
	if err != nil {
		panic(err)
	}
	s.client = client
}

func (s *BoardSuite) TearDownSuite(c *check.C) {
}

func (s *BoardSuite) TestNewContainer(c *check.C) {
	type args struct {
		c       *trello.Client
		boardID string
	}
	tests := []struct {
		name    string
		args    args
		want    *BoardContainer
		wantErr bool
	}{
		{
			"generic",
			args{
				c:       s.client,
				boardID: "5893eac416340d068cb41a2b",
			},
			&BoardContainer{},
			false,
		},
	}
	for _, tt := range tests {
		got, err := NewContainer(tt.args.c, tt.args.boardID)
		c.Check(got, check.NotNil)
		c.Check(err, check.IsNil)
	}
}

func TestBoardContainer_RetrieveCards(t *testing.T) {
	type fields struct {
		board *trello.Board
		lists []*trello.List
		cards map[listID][]trello.Card
		mux   sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BoardContainer{
				board: tt.fields.board,
				lists: tt.fields.lists,
				cards: tt.fields.cards,
				mux:   tt.fields.mux,
			}
			if err := c.RetrieveCards(); (err != nil) != tt.wantErr {
				t.Errorf("BoardContainer.RetrieveCards() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoardContainer_ListWorker(t *testing.T) {
	type fields struct {
		board *trello.Board
		lists []*trello.List
		cards map[listID][]trello.Card
		mux   sync.Mutex
	}
	type args struct {
		id      int
		listJob <-chan *trello.List
		results chan<- []trello.Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &BoardContainer{
				board: tt.fields.board,
				lists: tt.fields.lists,
				cards: tt.fields.cards,
				mux:   tt.fields.mux,
			}
			c.ListWorker(tt.args.id, tt.args.listJob, tt.args.results)
		})
	}
}

func TestNewEmail(t *testing.T) {
	type args struct {
		b *trello.Board
	}
	tests := []struct {
		name string
		args args
		want *Email
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmail(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMetaData(t *testing.T) {
	type args struct {
		l *trello.List
	}
	tests := []struct {
		name    string
		args    args
		want    *MetaData
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMetaData(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMetaData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMetaData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSection(t *testing.T) {
	type args struct {
		l *trello.List
	}
	tests := []struct {
		name    string
		args    args
		want    *Section
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSection(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewArticle(t *testing.T) {
	type args struct {
		card *trello.Card
	}
	tests := []struct {
		name    string
		args    args
		want    *Article
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewArticle(tt.args.card)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
