package commonword

import (
	"testing"
)

func TestGetCommonWords(t *testing.T) {
	scenarios := []struct {
		testCaseName        string
		text                string
		maxTopWords         int
		expectedNumOfWords  int
		expectedCommonWords []CommonWord
		hasError            bool
		errorMessage        string
	}{
		{
			testCaseName:       "Text has size below 10",
			text:               "A quick brown fox jumps over the lazy dog",
			maxTopWords:        10,
			expectedNumOfWords: 9,
			expectedCommonWords: []CommonWord{
				{word: "A", occurence: 1},
				{word: "QUICK", occurence: 1},
				{word: "BROWN", occurence: 1},
				{word: "FOX", occurence: 1},
				{word: "JUMPS", occurence: 1},
				{word: "OVER", occurence: 1},
				{word: "THE", occurence: 1},
				{word: "LAZY", occurence: 1},
				{word: "DOG", occurence: 1},
			},
		},
		{
			testCaseName:       "Text has size ten",
			text:               "A quick brown fox jumps over the lazy dog again",
			maxTopWords:        10,
			expectedNumOfWords: 10,
			expectedCommonWords: []CommonWord{
				{word: "A", occurence: 1},
				{word: "QUICK", occurence: 1},
				{word: "BROWN", occurence: 1},
				{word: "FOX", occurence: 1},
				{word: "JUMPS", occurence: 1},
				{word: "OVER", occurence: 1},
				{word: "THE", occurence: 1},
				{word: "LAZY", occurence: 1},
				{word: "DOG", occurence: 1},
				{word: "AGAIN", occurence: 1},
			},
		},
		{
			testCaseName:       "Text size exceeds ten",
			text:               "A quick brown fox jumps over the lazy dog again and again",
			maxTopWords:        10,
			expectedNumOfWords: 10,
			expectedCommonWords: []CommonWord{
				{word: "A", occurence: 1},
				{word: "QUICK", occurence: 1},
				{word: "BROWN", occurence: 1},
				{word: "FOX", occurence: 1},
				{word: "JUMPS", occurence: 1},
				{word: "OVER", occurence: 1},
				{word: "LAZY", occurence: 1},
				{word: "DOG", occurence: 1},
				{word: "AGAIN", occurence: 2},
			},
		},
		{
			testCaseName:       "Text size exceeds ten and has many top ten candidates",
			text:               "A quick A brown quick fox jumps over over over over the lazy dog gh gh gh again too too too many many many and many many many again",
			maxTopWords:        10,
			expectedNumOfWords: 10,
			expectedCommonWords: []CommonWord{
				{word: "A", occurence: 2},
				{word: "BROWN", occurence: 1},
				{word: "FOX", occurence: 1},
				{word: "JUMPS", occurence: 1},
				{word: "LAZY", occurence: 1},
				{word: "DOG", occurence: 1},
				{word: "AGAIN", occurence: 2},
			},
		},
		{
			testCaseName: "Text is empty",
			text:         "",
			maxTopWords:  10,
			hasError:     true,
			errorMessage: "text is empty",
		},
		{
			testCaseName: "Text is a tab",
			text: "	",
			maxTopWords:  10,
			hasError:     true,
			errorMessage: "text is empty",
		},
		{
			testCaseName: "Text has extra space in between",
			text: "A  quick  brown fox     jumps over   the 	lazy		 dog",
			maxTopWords:        10,
			expectedNumOfWords: 9,
			expectedCommonWords: []CommonWord{
				{word: "A", occurence: 1},
				{word: "QUICK", occurence: 1},
				{word: "BROWN", occurence: 1},
				{word: "FOX", occurence: 1},
				{word: "JUMPS", occurence: 1},
				{word: "OVER", occurence: 1},
				{word: "THE", occurence: 1},
				{word: "LAZY", occurence: 1},
				{word: "DOG", occurence: 1},
			},
		},
		{
			testCaseName: "Text has spaces before and after",
			text: " A quick brown fox jumps over the lazy dog	",
			maxTopWords:        10,
			expectedNumOfWords: 9,
			expectedCommonWords: []CommonWord{
				{word: "A", occurence: 1},
				{word: "QUICK", occurence: 1},
				{word: "BROWN", occurence: 1},
				{word: "FOX", occurence: 1},
				{word: "JUMPS", occurence: 1},
				{word: "OVER", occurence: 1},
				{word: "THE", occurence: 1},
				{word: "LAZY", occurence: 1},
				{word: "DOG", occurence: 1},
			},
		},
		{
			testCaseName:       "Text has numbers and symbols",
			text:               "A quick 1 brown ! fox jumps over the lazy dog dog123",
			maxTopWords:        10,
			expectedNumOfWords: 9,
			expectedCommonWords: []CommonWord{
				{word: "A", occurence: 1},
				{word: "QUICK", occurence: 1},
				{word: "BROWN", occurence: 1},
				{word: "FOX", occurence: 1},
				{word: "JUMPS", occurence: 1},
				{word: "OVER", occurence: 1},
				{word: "THE", occurence: 1},
				{word: "LAZY", occurence: 1},
				{word: "DOG", occurence: 1},
			},
		},
	}

	for _, tc := range scenarios {
		t.Run(tc.testCaseName, func(t *testing.T) {
			service := New(tc.maxTopWords)
			response, err := service.GetCommonWords(tc.text)

			if (err != nil) != tc.hasError {
				t.Errorf("Expected to error: %t but was: %t", tc.hasError, err != nil)
			}

			if len(response.commonWords) != tc.expectedNumOfWords {
				t.Errorf("Expected num of words: %d but was %d", tc.expectedNumOfWords, len(response.commonWords))
			}

			if err != nil && err.Error() != tc.errorMessage {
				t.Errorf("Expected error message: %s but was: %s", tc.errorMessage, err.Error())
			}

			responseMap := make(map[string]int)
			for _, commonWord := range response.commonWords {
				responseMap[commonWord.word] = commonWord.occurence
			}

			for _, commonword := range tc.expectedCommonWords {
				wordOccurence, exist := responseMap[commonword.word]
				if !exist {
					t.Errorf("Expected word: %s but was not found", commonword.word)
				} else if wordOccurence != commonword.occurence {
					t.Errorf("Expected occurence for word %s: %d but was %d", commonword.word, commonword.occurence, wordOccurence)
				}
			}
		})
	}
}
