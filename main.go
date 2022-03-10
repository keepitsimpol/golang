package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/keepitsimpol/topten/commonword"
	"github.com/sirupsen/logrus"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	logrus.Infoln("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	text = strings.Replace(text, "\r", "", 1)

	service := commonword.New()
	response, err := service.GetCommonWords(text)
	if err != nil {
		logrus.Errorf("Error in getting common words: %s", err.Error())
		return
	}

	logrus.Infof("Response is %+v", response)
}
