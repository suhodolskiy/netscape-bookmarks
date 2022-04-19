package bookmarks

const attrLastModified = "last_modified"
const attrAddDate = "add_date"
const attrIcon = "icon"
const attrHref = "href"

type Folder struct {
	Name         string      `json:"name"`
	AddDate      string      `json:"add_date"`
	LastModified string      `json:"last_modified"`
	Children     Children    `json:"children"`
	Attributes   []Attribute `json:"attributes"`
}

type Bookmark struct {
	Name       string      `json:"name"`
	AddDate    string      `json:"add_date"`
	Icon       string      `json:"icon"`
	Href       string      `json:"href"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Children = []interface{}
