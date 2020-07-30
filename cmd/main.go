package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	link "github.com/oucema001/link"
)

var examp1 = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>

`

func main() {

	urlp := "https://www.calhoun.io"
	content, err := http.Get(urlp)
	if err != nil {
		log.Println(err)
	}
	defer content.Body.Close()

	l, err := link.Parse(content.Body)
	if err != nil {
		log.Println(err)
	}
	for _, v := range l {
		fmt.Printf("%s : %s \n", v.Href, v.Text)
	}

	//	fmt.Printf("%+v", links)
}

func openFile(fileName string) ([]byte, error) {
	f, err := ioutil.ReadFile(fileName)

	/*
		if err != nil {
			log.Println(err)
			return nil, err
		}
	*/
	return f, err
}
