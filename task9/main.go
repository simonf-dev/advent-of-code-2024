package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type fileMetadata struct {
	fileIndex     int
	startPosition int
	length        int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFilesMetadata(filePath string) ([]fileMetadata, []fileMetadata) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	isFile := true
	fileIndex := 0
	filePosition := 0
	fileMetadataSlice := make([]fileMetadata, 0)
	spaces := make([]fileMetadata, 0)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		} else if r == '\n' {
			continue
		}
		if isFile {
			length, err := strconv.Atoi(string(r))
			check(err)
			fileMetadataSlice = append(fileMetadataSlice, fileMetadata{fileIndex: fileIndex, startPosition: filePosition, length: length})
			fileIndex += 1
			filePosition += length

		} else {
			length, err := strconv.Atoi(string(r))
			check(err)
			spaces = append(spaces, fileMetadata{fileIndex: -1, startPosition: filePosition, length: length})
			filePosition += length
		}
		isFile = !isFile
	}
	fmt.Println("len of metadataSlice", len(fileMetadataSlice))
	fmt.Println("last structure", fileMetadataSlice[len(fileMetadataSlice)-1])
	return fileMetadataSlice, spaces
}

func moveFile(metadata *fileMetadata, spaces *[]fileMetadata) {
	// call function that will take last file and space, it manipulates in place new slice and returns remaining space
	for spaceIndex := range len(*spaces) {
		if (*spaces)[spaceIndex].startPosition > metadata.startPosition {
			return
		}
		if (*spaces)[spaceIndex].length >= metadata.length {
			metadata.startPosition = (*spaces)[spaceIndex].startPosition
			(*spaces)[spaceIndex].startPosition += metadata.length
			(*spaces)[spaceIndex].length -= metadata.length
			return
		}
	}
}
func moveFiles(files, spaces []fileMetadata) []fileMetadata {
	index := 1
	for index < len(files) {
		moveFile(&files[len(files)-index], &spaces)
		index += 1
	}
	return files
}

func countResult(metadataSlice []fileMetadata) int {
	result := 0
	length := 0
	for _, metadata := range metadataSlice {
		for index := range metadata.length {
			result += (metadata.startPosition + index) * metadata.fileIndex
		}
		length += metadata.length

	}
	return result
}

func writeFieldToFile(metadataSlice []fileMetadata, filePath string) error {
	// Sort the slice by startPosition
	sort.Slice(metadataSlice, func(i, j int) bool {
		return metadataSlice[i].startPosition < metadataSlice[j].startPosition
	})

	// Open the file for writing
	file, err := os.Create(filePath) // Creates the file if it doesn't exist
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // Ensure file is closed after the function exits

	// Write the data to the file
	for _, value := range metadataSlice {
		for i := 0; i < value.length; i++ { // Write for 'length' times
			_, err := fmt.Fprintf(file, "%d,", value.fileIndex)
			if err != nil {
				return fmt.Errorf("failed to write to file: %w", err)
			}
		}
	}

	return nil
}
func main() {
	files, spaces := parseFilesMetadata("input.txt")
	fmt.Println(files)
	fmt.Println(spaces)
	result := moveFiles(files, spaces)
	writeFieldToFile(result, "output.txt")
	//fmt.Println(countResult(output))
	fmt.Println(result)
	fmt.Println(spaces)
	fmt.Println(countResult(result))
}
