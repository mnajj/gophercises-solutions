package main

import (
	"fmt"
	"log"
	"strings"

	"ex4/link"
)

const htmlEx1 string = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(htmlEx1)
	links, err := link.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", links)
}
