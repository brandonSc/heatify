package route

import "io/ioutil"

//
// an HTML page with a title (name of the file) and body content
//
type Page struct {
	Title string
	Body  []byte
	Data  string
}

//
// load an HTML page from the views/ directory
//
func LoadPage(title string) (*Page, error) {
	filename := "./views/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
