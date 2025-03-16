package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/pelletier/go-toml/v2"
)

type Taxonomies struct {
	Categories []string `toml:"categories"`
	Tags       []string `toml:"tags"`
}
type Frontmatter struct {
	Title      string     `toml:"title"`
	Date       string     `toml:"date"`
	Path       string     `toml:"path"`
	Taxonomies Taxonomies `toml:"taxonomies"`
}

func main() {
	// get user input
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

	// parse user input
	titleSlug := titleToSlug(title)
	currentDate, currentYear, currentMonth := getCurrentDate()
	tagsSlice := tagsToSlice(tags)

	// create post dir & file
	filePath := createFiles(currentDate, titleSlug, hasImage)

	// create file
	frontmatter := Frontmatter{
		Title: title,
		Date:  currentDate,
		Path:  fmt.Sprintf("/posts/%s/%s/%s", currentYear, currentMonth, titleSlug),
		Taxonomies: Taxonomies{
			Categories: []string{"api",
				"book-highlights",
				"career",
				"ci-cd",
				"data-analysis",
				"data-engineering",
				"data-science",
				"devops",
				"frontend",
				"homelab",
				"infrastructure",
				"kubernetes",
				"languages",
				"life",
				"llm",
				"migration",
				"mlops",
				"platform-engineering",
				"programming",
				"security",
				"software-engineering",
				"sre",
				"wasm",
				"web-hosting",
			}, // pre-filled
			Tags: tagsSlice,
		},
	}

	//// marshal to toml
	frontmatterBytes, err := toml.Marshal(frontmatter)
	if err != nil {
		panic(err)
	}

	// create file
	text := fmt.Sprintf(`+++
%s+++`, string(frontmatterBytes))

	err = os.WriteFile(filePath, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Created %s", filePath)

	// add to git
	execCommand("git", "add", filePath)
}
