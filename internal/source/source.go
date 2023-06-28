package source

type Entry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Source []Entry
