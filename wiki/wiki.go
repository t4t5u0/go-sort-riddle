package main

import "fmt"

// WikipediaObject Wikipediaからランダムで持ってきたやつ
type WikipediaObject struct {
	Batchcomplete string `json:"batchcomplete"`
	Continue      struct {
		Rncontinue string `json:"rncontinue"`
		Continue   string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Random []struct {
			ID    int    `json:"id"`
			Ns    int    `json:"ns"`
			Title string `json:"title"`
		} `json:"random"`
	} `json:"query"`
}

// ReturnTitle です
func (arg WikipediaObject) ReturnTitle() {
	x := arg.Query.Random[0].Title
	fmt.Println(x)
}
