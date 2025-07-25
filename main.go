package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"slices"

	"go.senan.xyz/taglib"
)

func main() {

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("can't read dir: %s", err.Error())
	}

	fmt.Printf("Found %d total files in the folder\n", len(files))

	// Remove non-mp3 files from folder
	files = slices.DeleteFunc(files, func(s os.DirEntry) bool {
		return path.Ext(s.Name()) != ".mp3"
	})

	fmt.Printf("Found %d .mp3 files in the folder\n", len(files))

	// Shuffle "files" slice
	rand.Shuffle(len(files), func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	for index, file := range files {

		// Retrieve current track number
		tags, err := taglib.ReadTags(file.Name())
		if err != nil {
			log.Fatalf("failed while reading tags of file '%s' with error: '%s'", file.Name(), err.Error())
		}
		trackNumberBefore := tags[taglib.TrackNumber]

		// Set track number to current index
		err = taglib.WriteTags(file.Name(), map[string][]string{
			taglib.TrackNumber: {fmt.Sprintf("%d", index)},
		}, 0)
		if err != nil {
			log.Fatalf("failed while setting tags of file '%s' with error: '%s'", file.Name(), err.Error())
		}

		fmt.Printf("Updated trackNumber of file '%s' from '%s' to '%d'\n", file.Name(), trackNumberBefore, index)
	}
}
