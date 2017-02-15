package printshop

import (
	"os"
	"reflect"
	"testing"

	"github.com/VojtechVitek/go-trello"
	check "gopkg.in/check.v1"
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

func (b *BoardSuite) NewTrelloCard() *trello.Card {
	return &trello.Card{}
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

var suite *BoardSuite

func init() {
	suite = &BoardSuite{}
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
		{
			"generic",
			args{
				&trello.Card{},
			},
			&Article{},
			false,
		},
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
