package commonword

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

type commonWordService struct {
	maxTopCommonWords int
}

func New(maxTopCommonWords int) *commonWordService {
	service := new(commonWordService)
	service.maxTopCommonWords = maxTopCommonWords
	return service
}

type CommonWordServiceResponse struct {
	commonWords []CommonWord
}

type CommonWord struct {
	word      string
	occurence int
}

func (s commonWordService) GetCommonWords(text string) (CommonWordServiceResponse, error) {
	log.Println(fmt.Sprintf("Start getting common words from: %s", text))
	var response CommonWordServiceResponse
	wordDictionary := make(map[string]int)

	text = strings.TrimSpace(text)
	if text == "" {
		return response, errors.New("text is empty")
	}

	s.saveWordsInDictionary(text, wordDictionary)
	commonWords := s.determineTopCommonWords(wordDictionary)
	sort.SliceStable(commonWords, func(i, j int) bool {
		return commonWords[i].occurence > commonWords[j].occurence
	})

	response = CommonWordServiceResponse{commonWords: commonWords}
	return response, nil
}

func (s commonWordService) saveWordsInDictionary(text string, dictionary map[string]int) {
	logrus.Infoln("Saving words in our dictionary")
	spaceRemover := regexp.MustCompile(`\s+`)
	text = spaceRemover.ReplaceAllString(text, " ")
	text = strings.ToUpper(text)

	words := strings.Split(text, " ")
	logrus.Infof("Words to be saved in dictionary: %v with size: %d", words, len(words))
	r := regexp.MustCompile(`[^a-zA-Z]`)

	for _, word := range words {
		logrus.Infof("Current word to save: %s", word)

		notAWord := r.MatchString(word)
		if notAWord {
			logrus.Warnf("%s is not a word. Discarding this entry", word)
		} else {
			dictionary[word] = dictionary[word] + 1
		}
	}

	logrus.Infof("Dictionary completed: %v", dictionary)
}

func (s commonWordService) determineTopCommonWords(wordDictionary map[string]int) []CommonWord {
	logrus.Infoln("Start determining top ten words")
	keys := make([]string, 0, len(wordDictionary))
	for k := range wordDictionary {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	commonWords := make([]CommonWord, 0)
	for _, key := range keys {
		if len(commonWords) >= s.maxTopCommonWords {
			logrus.Infoln("Common words are more than ten. Updating the top ten list.")
			keyToDelete := ""
			for i := 0; i < len(commonWords); i++ {
				if commonWords[i].occurence < wordDictionary[key] {
					keyToDelete = key
				}
			}

			if keyToDelete != "" {
				logrus.Infof("Word to be removed from the top-ten list: %s", keyToDelete)
				delete(wordDictionary, keyToDelete)
			} else {
				logrus.Infof("Not adding word: %s", key)
			}
		}
		if len(commonWords) < s.maxTopCommonWords {
			logrus.Infof("Adding word: %s in our top ten", key)
			commonWords = append(commonWords, CommonWord{word: key, occurence: wordDictionary[key]})
		}
	}
	return commonWords
}
