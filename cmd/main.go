package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

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
	var ex io.Reader

	ex = strings.NewReader(examp1)
	l := link.Link{}
	links, err := l.Parse(ex)
	if err != nil {
		fmt.Println(err)
	}
	_ = links
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
