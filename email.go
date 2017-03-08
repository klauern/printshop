package printshop

import (
	markdown "github.com/shurcooL/github_flavored_markdown"
)

type Renderer interface {
	Render(md string) ([]byte, error)
}

func (m *MailClient) Render(md string) ([]byte, error) {
	return markdown.Markdown([]byte(md)), nil
}
