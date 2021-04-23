package sub

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpExample() {
	// resp, _ := http.Get("https://nextjs-api-harry.herokuapp.com/api/detail-post/1/")
	// defer resp.Body.Close()

	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	fmt.Println(base) // http://example.com
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint) // http://example.com/test?a=1&b=2

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyzzy`)
	q := req.URL.Query()
	fmt.Println(q) // map[a:[1] b:[2]]
	q.Add("c", "3")
	fmt.Println(q, q.Encode()) // map[a:[1] b:[2] c:[3]], a=1&b=2&c=3

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
