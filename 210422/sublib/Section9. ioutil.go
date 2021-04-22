package sublib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func IoutilExample() {
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content)) // Example text

	r := bytes.NewBuffer([]byte("abc"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2)) // abc
}
