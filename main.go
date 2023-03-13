package main

import (
	"flag"
	"fmt"
	"strings"
	"wordleWizard/config"
	"wordleWizard/query"
	simplechecker "wordleWizard/word-checker/simple-checker"
	memory_repo "wordleWizard/word_repo/memory-repo"
)

func main() {

	var eStr, iStr, pStr string

	flag.StringVar(&eStr, "e", "", "Characters excluded")
	flag.StringVar(&iStr, "i", "", "Characters included")
	flag.StringVar(&pStr, "p", "", "Characters placed")
	flag.Parse()

	cfg := config.Config{LetterCount: 5}

	q, err := query.CreateQuery(eStr, iStr, pStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = query.SanitizeQuery(cfg, &q)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(q.String())

	w, err := memory_repo.NewMemoryRepo()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	chkr := simplechecker.NewSimpleChecker(cfg, q)
	words, err := w.FindWord(chkr)
	if err != nil {
		fmt.Println("Error occurred finding words: ", err.Error())
		return
	}

	if len(words) == 0 {
		fmt.Println("Couldn't think of any words sorry :(")
		return
	}

	var sb = strings.Builder{}

	if len(words) > 1 {
		sb.WriteString("These ")
		sb.WriteString("words ")
	} else {
		sb.WriteString("This ")
		sb.WriteString("word ")
	}
	sb.WriteString("may work ")
	sb.WriteString(fmt.Sprintf("(%d possibilities):", len(words)))

	fmt.Println(sb.String())
	fmt.Println("=========================================")
	for i, word := range words {
		fmt.Printf("%d) %s\n", i+1, word)
	}
}
