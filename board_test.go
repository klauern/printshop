package printshop

import (
	"os"
	"reflect"
	"testing"

	"github.com/VojtechVitek/go-trello"
	check "github.com/go-check/check"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type BoardSuite struct {
	client *trello.Client
}

var _ = check.Suite(&BoardSuite{})

func (b *BoardSuite) NewTrelloBoard() *trello.Board {
	return &trello.Board{}
}

func (b *BoardSuite) NewTrelloCard(id string) *trello.Card {
	const DefaultCardID = ""
	if id == "" {
		card, err := b.client.Card(DefaultCardID)
		if err != nil {
			panic(err)
		}
		return card
	}
	card, err := b.client.Card(id)
	if err != nil {
		panic(err)
	}
	return card
}

func (b *BoardSuite) NewTrelloList() *trello.List {
	return &trello.List{}
}

func (b *BoardSuite) SetUpSuite(c *check.C) {
	key := os.Getenv(KEY_ENV)
	token := os.Getenv(TOKEN_ENV)
	client, err := trello.NewAuthClient(key, &token)
	if err != nil {
		panic(err)
	}
	b.client = client
}

var testSuite *BoardSuite

func init() {
	testSuite = &BoardSuite{}
}

func (b *BoardSuite) TestNewEmail(t *check.C) {
	type args struct {
		b *trello.Board
	}
	tests := []struct {
		name string
		args args
		want *Email
	}{
		{
			"generic",
			args{
				b: b.NewTrelloBoard(),
			},
			&Email{},
		},
	}
	for _, tt := range tests {
		t.Check(NewEmail(tt.args.b), check.DeepEquals, tt.want)
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
		{
			"nil, empty",
			args{
				&trello.List{},
			},
			&MetaData{},
			false,
		},
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

func (b *BoardSuite) TestNewSection(t *check.C) {
	type args struct {
		l *trello.List
	}
	tests := []struct {
		name    string
		args    args
		want    *Section
		wantErr bool
	}{
		{
			"generic",
			args{
				b.NewTrelloList(),
			},
			&Section{},
			false,
		},
	}
	for _, tt := range tests {
		got, err := NewSection(tt.args.l)
		t.Check(got, check.DeepEquals, tt.want)
		t.Check(err, check.IsNil)
	}
}

func (b *BoardSuite) TestNewArticle(t *check.C) {
	type args struct {
		card *trello.Card
	}
	tests := []struct {
		name    string
		args    args
		want    *Article
		wantErr bool
	}{
		{
			"generic",
			args{
				card: b.NewTrelloCard(""),
			},
			&Article{},
			false,
		},
	}
	for _, tt := range tests {
		got, err := NewArticle(tt.args.card)
		t.Check(err, check.IsNil)
		t.Check(got, check.DeepEquals, tt.want)
	}
}
