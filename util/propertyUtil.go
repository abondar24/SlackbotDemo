package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const PropsFilename = "../slack.properties"

type SlackProps map[string]string

func ReadProperties() (SlackProps, error) {

	props := SlackProps{}

	if len(PropsFilename) == 0 {
		return props, nil
	}

	file, err := os.Open(PropsFilename)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				props[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return props, nil
}
