package akuma

import (
	"io/fs"
	"log"
	"os"
)

// retrive all the manga items in the directory
func GetMangaList() []fs.DirEntry {
	files, err := os.ReadDir("manga")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// look up manga chapters from their hash
func GetMangaChapters(mangaName string) []fs.DirEntry {
	entries, err := os.ReadDir("manga/" + mangaName)
	if err != nil {
		log.Println("Error: GetMangaChapters(mangaName string): ", err)
	}
	return entries
}

// takes next, prev parameters and checks if next/prev chapter exists
// returns -1 if next/prev chapter does not exists
// else returns vol+1 or vol-1 if next/prev chapters exists
func DoesNextChapterExists(vol int, entries []fs.DirEntry) (int, int) {
	if vol == 0 {
		return vol + 1, -1
	} else if vol+1 >= len(entries) {
		return -1, vol - 1
	}
	return vol + 1, vol - 1
}
