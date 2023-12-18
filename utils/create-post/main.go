package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
)

func titleToSlug(title string) string {
	re := regexp.MustCompile("[a-zA-Z ]+")
	sanitizedString := re.FindString(title)
	lowerString := strings.ToLower(sanitizedString)
	slug := strings.ReplaceAll(lowerString, " ", "-")

	return slug
}

func getCurrentTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02T15:04:05MST:00")

	return formattedTime
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
	currentTime := getCurrentTime()
	tagsFormatted := formatTags(tags)

	// set filepath
	var filePath string
	if hasImage {
		folder := fmt.Sprintf("content/posts/%s/images", titleSlug)
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}

		filePath = fmt.Sprintf("content/posts/%s/index.md", titleSlug)
	} else {
		filePath = fmt.Sprintf("content/posts/%s.md", titleSlug)
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
}
