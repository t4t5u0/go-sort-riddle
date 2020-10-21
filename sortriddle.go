package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

func main() {
	answer := []rune(getRandomWord())
	question := make([]rune, len(answer))
	copy(question, answer)
	sort.Slice(question, func(i, j int) bool {
		return question[i] < question[j]
	})
	fmt.Printf("問題: %v\n", string(question))

	for {
		var input string
		fmt.Print("答えを入力してください: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}
		inputRune := []rune(input)
		if count := countCorrect(inputRune, answer); count == len(answer) {
			fmt.Println("正解！")
			os.Exit(0)
		} else if count == -1 {
			fmt.Println("長さが違う")
		} else {
			fmt.Printf("%v 文字正解\n", count)
		}
	}

}

func countCorrect(input, answer []rune) int {
	if len(input) != len(answer) {
		return -1
	}
	cnt := 0
	for i := 0; i < len(input); i++ {
		if input[i] == answer[i] {
			cnt++
		}
	}
	return cnt
}

func getRandomWord() string {
	link := "https://ja.wikipedia.org/w/api.php?action=query&list=random&format=json&rnnamespace=0&rnlimit=1"
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(WikipediaObject)

	jsonBytes := ([]byte)(byteArray)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		panic(err)
	}
	result := data.Query.Random[0].Title
	return result
}

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
