package bookmarks_test

import (
	"os"
	"testing"

	"github.com/go-test/deep"
	bookmarks "github.com/suhodolskiy/netscape-bookmarks"
	"gotest.tools/v3/assert"
)

var want = bookmarks.Children{
	bookmarks.Folder{
		Name:         "Bookmarks Bar",
		AddDate:      "1611420722",
		LastModified: "1650397116",
		Children: bookmarks.Children{
			bookmarks.Bookmark{
				Name:       "Google",
				AddDate:    "1650397103",
				Href:       "https://www.google.com/",
				Attributes: []bookmarks.Attribute{},
			},
			bookmarks.Folder{
				Name:         "Github",
				AddDate:      "1650397124",
				LastModified: "1650397155",
				Children: bookmarks.Children{
					bookmarks.Bookmark{
						Name:       "The React documentation website",
						AddDate:    "1650397152",
						Href:       "https://github.com/reactjs/reactjs.org",
						Attributes: []bookmarks.Attribute{},
					},
				},
				Attributes: []bookmarks.Attribute{},
			},
		},
		Attributes: []bookmarks.Attribute{},
	},
	bookmarks.Bookmark{
		Name:    "Golang Net html",
		AddDate: "1650055527",
		Href:    "https://zetcode.com/golang/net-html/",
		Attributes: []bookmarks.Attribute{
			{
				Key:   "custom_attr",
				Value: "TRUE",
			},
		},
	},
}

func TestBookmarks(t *testing.T) {
	file, err := os.OpenFile("./example.html", os.O_RDONLY, 0644)
	assert.NilError(t, err)

	defer file.Close()

	result, err := bookmarks.Parse(file)
	assert.NilError(t, err)

	diff := deep.Equal(result, want)
	if diff != nil {
		t.Errorf("compare failed: %v", diff)
	}
}
