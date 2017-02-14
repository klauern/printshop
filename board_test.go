package printshop

import (
	"reflect"
	"testing"

	"github.com/VojtechVitek/go-trello"
	check "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type BoardSuite struct{}

var _ = check.Suite(&BoardSuite{})

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
