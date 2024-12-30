package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func titleToSlug(title string) string {
	replacer := strings.NewReplacer("(", "", ")", "", "?", "", ",", "")
	sanitizedString := replacer.Replace(title)

	lowerString := strings.ToLower(sanitizedString)
	slug := strings.ReplaceAll(lowerString, " ", "-")

	return strings.ReplaceAll(slug, "'", "")
}

func getCurrentDate() (string, string, string) {
	currentTime := time.Now()

	currentDate := currentTime.Format("2006-01-02")
	currentYear := strconv.Itoa(currentTime.Year())
	currentMonth := fmt.Sprintf("%02d", int(currentTime.Month()))

	return currentDate, currentYear, currentMonth
}

func tagsToSlice(tags string) []string {
	return strings.Split(tags, ",")
}

func execCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(stdout))
	}
}

func createFiles(currentDate string, titleSlug string, hasImage bool) string {
	contentPath := fmt.Sprintf("content/posts/%s-%s", currentDate, titleSlug)
	var filePath string
	if hasImage {
		imageFolder := fmt.Sprintf("%s/images", contentPath)
		err := os.MkdirAll(imageFolder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		filePath = fmt.Sprintf("%s/index.md", contentPath)
	} else {
		filePath = fmt.Sprintf("%s.md", contentPath)
	}

	return filePath
}
