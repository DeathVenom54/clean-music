package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type RegexCleaner struct {
	pattern string
	replace string
}

var (
	regexpCleaners = []RegexCleaner{
		{pattern: "\\.mp3", replace: ""},
		{pattern: "\\(?\\d+ ?kbps\\)?", replace: ""},
		{pattern: " ?[-_] ?", replace: " "},
	}
)

func main() {
	fmt.Print("Directory: ")
	reader := bufio.NewReader(os.Stdin)
	directoryPath, inpErr := reader.ReadString('\n')
	if inpErr != nil {
		log.Fatal(inpErr)
	}
	directoryPath = strings.TrimSuffix(directoryPath, "\n")

	directory, dirErr := os.Stat(directoryPath)
	if (dirErr != nil && os.IsNotExist(dirErr)) || !directory.IsDir() {
		log.Fatalf("Directory %s doesn't exist or is not a valid directory\n", directoryPath)
	}

	files, readErr := os.ReadDir(directoryPath)
	if readErr != nil {
		log.Fatal(readErr)
	}

	client, authErr := authenticateClient()
	if authErr != nil {
		log.Fatalln(authErr)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".mp3") {
			continue
		}

		fullFileName := filepath.Join(directoryPath, file.Name())
		cleanFileName := cleanupName(file.Name())

		songs, searchErr := searchSongs(client, cleanFileName)
		if searchErr != nil {
			log.Fatalln(searchErr)
		}

		writeErr := WriteSongData(fullFileName, &songs[0])
		if writeErr != nil {
			log.Fatalln(writeErr)
		}

		fmt.Printf("Wrote %s\n%s - %s, album: %s, released: %s", fullFileName, songs[0].Title, strings.Join(songs[0].Artists, ", "), songs[0].AlbumName, songs[0].AlbumReleaseDate)
	}
}

func cleanupName(name string) string {
	for _, cleaner := range regexpCleaners {
		compiled := regexp.MustCompile(cleaner.pattern)
		name = compiled.ReplaceAllString(name, cleaner.replace)
	}
	return name
}
