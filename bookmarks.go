// Golang library to work with bookmark export files in Netscape bookmark format
package bookmarks

import (
	"io"

	"golang.org/x/net/html"
)

// Parse html data and return a list of bookmarks and folders
// 	import (
// 		bookmarks "github.com/suhodolskiy/netscape-bookmarks"
// 	)

// 	func main() {
// 		input, _ := os.OpenFile("./example.html", os.O_RDONLY, 0644)
// 		data, _ := bookmarks.Parse(input)
// 		fmt.Println("Result", data)
// 	}
func Parse(r io.Reader) (Children, error) {
	var c Children

	n, err := html.Parse(r)

	if err != nil {
		return c, err
	}

	initialNode := findInitialNode(n)

	if initialNode != nil && initialNode.FirstChild != nil {
		parse(initialNode.FirstChild, &c)
	}

	return c, nil
}

// Recursive find for the first <DL> node
func findInitialNode(n *html.Node) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "dl" {
			return c
		}

		if n := findInitialNode(c); n != nil {
			return n
		}
	}

	return nil
}

func findNextDlNode(n *html.Node) *html.Node {
	for c := n; c != nil; c = c.NextSibling {
		if c.Data == "dl" {
			return c
		}
	}

	return nil
}

func parse(n *html.Node, result *Children) {
	for c := n; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			continue
		}

		if c.Data == "dt" && c.FirstChild.Data == "h3" {
			f := Folder{
				Name:       c.FirstChild.FirstChild.Data,
				Attributes: []Attribute{},
			}

			for _, a := range c.FirstChild.Attr {
				switch a.Key {
				case attrAddDate:
					{
						f.AddDate = a.Val
					}
				case attrLastModified:
					{
						f.LastModified = a.Val
					}
				default:
					f.Attributes = append(f.Attributes, Attribute{Key: a.Key, Value: a.Val})
				}
			}

			if c.FirstChild != nil {
				dl := findNextDlNode(c.FirstChild)

				if dl != nil && dl.FirstChild != nil && dl.FirstChild.Data == "p" {
					parse(dl.FirstChild, &f.Children)
				}
			}

			*result = append(*result, f)
		}

		if c.Data == "dt" && c.FirstChild.Data == "a" {
			b := Bookmark{
				Name:       c.FirstChild.FirstChild.Data,
				Attributes: []Attribute{},
			}

			for _, a := range c.FirstChild.Attr {
				switch a.Key {
				case attrAddDate:
					{
						b.AddDate = a.Val
					}
				case attrIcon:
					{
						b.Icon = a.Val
					}
				case attrHref:
					{
						b.Href = a.Val
					}
				default:
					b.Attributes = append(b.Attributes, Attribute{Key: a.Key, Value: a.Val})
				}
			}

			*result = append(*result, b)
		}
	}
}
