package day3

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// This code that be copied from google, but LGTM!

// In ra danh sách folder và files dưới dạng tree
func printListing(entry string, depth int) {
	indent := strings.Repeat("|   ", depth)
	fmt.Printf("%s|-- %s\n", indent, entry)
}

// Tìm kiếm danh sách các files và folder
// nếu gặp folder sẽ tiếp tục dùng đệ quy để tiếp tục duyệt theo depth level đã được truyền vào
func printDirectory(path string, depth int) error {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	printListing(path, depth)
	for _, entry := range entries {
		if (entry.Mode() & os.ModeSymlink) == os.ModeSymlink {
			full_path, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				return err
			} else {
				printListing(entry.Name()+" -> "+full_path, depth+1)
			}
		} else if entry.IsDir() {
			printDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {
			printListing(entry.Name(), depth+1)
		}
	}
	return nil
}
