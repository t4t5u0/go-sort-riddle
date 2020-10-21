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
	answer := getRandomWord()
	x := []rune(answer)
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	// sortRiddle := string([]rune(x))
	fmt.Printf("問題: %v\n", string([]rune(x)))

	for {
		var input string
		fmt.Print("答えを入力してください: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}
		y := []rune(input)
		if count := countCorrect(y, x); count == len(x) {
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
	// x := []rune(input)
	// y := []rune(answer)
	if len(input) != len(answer) {
		return -1
	}
	cnt := 0
	for i := 0; i < len(input); i++ {
		if input[i] == answer[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
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
	// fmt.Println(result)
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
