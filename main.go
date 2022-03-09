package main

import (
	"bufio"
	"log"
	"os"

	"github.com/keepitsimpol/topten/commonword"
	"github.com/sirupsen/logrus"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Enter text: ")
	text, _ := reader.ReadString('.')
	service := commonword.New()
	response, err := service.GetCommonWords(text)
	if err != nil {
		logrus.Errorf("Error in getting common words: %s", err.Error())
		return
	}

	logrus.Infof("Response is %+v", response)
}
