package main

import (
	"fmt"
	"github.com/wangbjun/learnGo/Practice/ch3/github"
	"html/template"
	"log"
	"net/http"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/report", handlerReport)

	println("Server started at localhost:8000....")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handlerReport(w http.ResponseWriter, r *http.Request) {

	var keyword []string

	keyword = append(keyword, r.FormValue("keyword"))

	result, err := github.SearchIssues(keyword)

	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
