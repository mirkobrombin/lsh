package core

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sync"
	"time"
)

// Bookmark represents a saved connection bookmark
type Bookmark struct {
	Name       string `json:"name"`
	Connection string `json:"connection"`
}

// bookmarksFile is the path to the bookmarks JSON file
var bookmarksFile string

// bookmarksMu protects access to the bookmarks slice
var bookmarksMu sync.Mutex

// bookmarks is a slice to store loaded bookmarks
var bookmarks []Bookmark

func init() {
	initBookmarksFilePath()
	loadBookmarks()
}

// initBookmarksFilePath initializes the bookmarks file path
func initBookmarksFilePath() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting current user: %v\n", err)
		os.Exit(1)
	}

	bookmarksDir := filepath.Join(currentUser.HomeDir, ".local", "share", "LSH")
	err = os.MkdirAll(bookmarksDir, 0700)
	if err != nil {
		fmt.Printf("Error creating bookmarks directory: %v\n", err)
		os.Exit(1)
	}

	bookmarksFile = filepath.Join(bookmarksDir, "bookmarks.json")
}

// loadBookmarks loads existing bookmarks from the file
func loadBookmarks() {
	data, err := os.ReadFile(bookmarksFile)
	if err != nil {
		if os.IsNotExist(err) {
			saveBookmarks()
			return
		}

		fmt.Printf("Error reading bookmarks file: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &bookmarks)
	if err != nil {
		fmt.Printf("Error unmarshalling bookmarks: %v\n", err)
		os.Exit(1)
	}
}

// saveBookmarks saves the current bookmarks to the file
func saveBookmarks() {
	data, err := json.MarshalIndent(bookmarks, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling bookmarks: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(bookmarksFile, data, 0600)
	if err != nil {
		fmt.Printf("Error writing bookmarks file: %v\n", err)
		os.Exit(1)
	}
}

// LoadBookmark loads a bookmark by name
func LoadBookmark(name string) (Bookmark, error) {
	bookmarksMu.Lock()
	defer bookmarksMu.Unlock()

	for _, bookmark := range bookmarks {
		if bookmark.Name == name {
			return bookmark, nil
		}
	}

	return Bookmark{}, fmt.Errorf("bookmark not found: %s", name)
}

// SaveBookmark saves a new bookmark
func SaveBookmark(bookmark Bookmark) error {
	bookmarksMu.Lock()
	defer bookmarksMu.Unlock()

	for _, existingBookmark := range bookmarks {
		if existingBookmark.Name == bookmark.Name {
			return fmt.Errorf("bookmark with name %s already exists", bookmark.Name)
		}
	}

	bookmarks = append(bookmarks, bookmark)
	saveBookmarks()

	return nil
}

// Connect connects to a bookmark using the ConnectSSH method
func (b Bookmark) Connect() error {
	return ConnectSSH(b.Connection)
}

// ListBookmarks returns the list of saved bookmarks
func ListBookmarks() ([]Bookmark, error) {
	bookmarksMu.Lock()
	defer bookmarksMu.Unlock()

	return bookmarks, nil
}

// ExportBookmarks exports the current bookmarks to a JSON file
func ExportBookmarks(filePath string) (string, error) {
	bookmarksMu.Lock()
	defer bookmarksMu.Unlock()

	fileInfo, err := os.Stat(filePath)
	if err == nil && fileInfo.IsDir() {
		timestamp := time.Now().Format("20060102_150405")
		filePath = filepath.Join(filePath, "bookmarks_"+timestamp+".json")
	}

	data, err := json.MarshalIndent(bookmarks, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling bookmarks: %v", err)
	}

	err = os.WriteFile(filePath, data, 0600)
	if err != nil {
		return "", fmt.Errorf("error writing bookmarks file: %v", err)
	}

	return filePath, nil
}

// ImportBookmarks imports bookmarks from a JSON file and replaces the
// existing bookmarks
func ImportBookmarks(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading bookmarks file: %v", err)
	}

	var importedBookmarks []Bookmark
	err = json.Unmarshal(data, &importedBookmarks)
	if err != nil {
		return fmt.Errorf("error unmarshalling bookmarks: %v", err)
	}

	bookmarksMu.Lock()
	defer bookmarksMu.Unlock()

	bookmarks = importedBookmarks
	saveBookmarks()

	return nil
}

// RemoveBookmark removes a bookmark by name
func RemoveBookmark(name string) error {
	bookmarksMu.Lock()
	defer bookmarksMu.Unlock()

	var index int
	for i, bookmark := range bookmarks {
		if bookmark.Name == name {
			index = i
			break
		}
	}

	bookmarks = append(bookmarks[:index], bookmarks[index+1:]...)
	saveBookmarks()

	return nil
}
