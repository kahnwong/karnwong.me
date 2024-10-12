package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
)

func titleToSlug(title string) string {
	replacer := strings.NewReplacer("(", "", ")", "", "?", "", ",", "")
	sanitizedString := replacer.Replace(title)

	lowerString := strings.ToLower(sanitizedString)
	slug := strings.ReplaceAll(lowerString, " ", "-")

	return strings.ReplaceAll(slug, "'", "")
}

func getCurrentTime() (string, string) {
	rawCurrentTime := time.Now()

	currentTime := rawCurrentTime.Format("2006-01-02T15:04:05MST:00")
	currentYear := strconv.Itoa(rawCurrentTime.Year())

	return currentTime, currentYear
}

func formatTags(tags string) string {
	tagsSplit := strings.Split(tags, ",")

	// add `-` in front of each tag, to make it bullets
	prefixChar := "  - "
	for i := 0; i < len(tagsSplit); i++ {
		tagsSplit[i] = prefixChar + tagsSplit[i]
	}

	tagsFormatted := strings.Join(tagsSplit, "\n")

	return tagsFormatted
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

func main() {

	var (
		title    string
		tags     string
		hasImage bool
	)
	form := huh.NewForm(
		huh.NewGroup(
			// set title
			huh.NewInput().
				Title("Post title?").
				Value(&title).
				Validate(func(str string) error {
					if len(str) < 4 {
						return errors.New("Please enter a post title.")
					}
					return nil
				}),

			// set tags (comma separated)
			huh.NewInput().
				Title("Post tags?").
				Value(&tags).
				Validate(func(str string) error {
					if strings.HasSuffix(str, ",") {
						return errors.New("Please specify tags in comma separated format.")
					}
					return nil
				}),

			// (optional) set image
			huh.NewConfirm().
				Title("Post has images?").
				Value(&hasImage),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	//// create markdown file
	titleSlug := titleToSlug(title)
	currentTime, currentYear := getCurrentTime()
	tagsFormatted := formatTags(tags)

	// mkdir
	contentPath := fmt.Sprintf("content/posts/%s", currentYear)
	err = os.Mkdir(contentPath, os.ModePerm)

	if err != nil && !os.IsExist(err) {
		// Handle the error if it's not an "already exists" error
		fmt.Printf("Error creating directory: %v\n", err)
	}

	// set filepath
	var filePath string
	if hasImage {
		folder := fmt.Sprintf("%s/%s/images", contentPath, titleSlug)
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		filePath = fmt.Sprintf("%s/%s/index.md", contentPath, titleSlug)
	} else {
		filePath = fmt.Sprintf("%s/%s.md", contentPath, titleSlug)
	}

	// create file
	text := fmt.Sprintf(`---
title: %s
date: %s
draft: false
ShowToc: false
images:
tags:
%s
---`, title, currentTime, tagsFormatted)

	err = os.WriteFile(filePath, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Created %s", filePath)

	// add to git
	execCommand("git", "add", filePath)
}
