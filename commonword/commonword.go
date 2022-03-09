package commonword

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type commonWordService struct {
	wordDictionary map[string]int
}

func New() *commonWordService {
	return new(commonWordService)
}

type CommonWordServiceResponse struct {
	commonWords []CommonWord `json:topTenWords`
}

type CommonWord struct {
	CommonWord string `json:commonWord`
	Occurrence int    `json:occurence`
}

func (s commonWordService) GetCommonWords(text string) (CommonWordServiceResponse, error) {
	log.Println(fmt.Sprintf("Start getting common words from: %s", text))
	var response CommonWordServiceResponse
	if text == "" {
		return response, errors.New("text is empty")
	}

	s.saveWordsInDictionary(text)
	commonWords := make([]CommonWord, 0)
	for word, occurence := range s.wordDictionary {
		commonWords = append(commonWords, CommonWord{CommonWord: word, Occurrence: occurence})
	}

	response = CommonWordServiceResponse{commonWords: commonWords}
	return response, nil
}

func (s commonWordService) saveWordsInDictionary(text string) {
	words := strings.Split(text, " ")
	for _, word := range words {
		s.wordDictionary[word] = s.wordDictionary[word] + 1
	}
}
