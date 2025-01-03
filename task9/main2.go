
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

func parseFilesMetadata(filePath string) []fileMetadata {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)
	isFile := true
	fileIndex := 0
	filePosition := 0
	fileMetadataSlice := make([]fileMetadata, 0)
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
			filePosition += length
		}
		isFile = !isFile
	}
	fmt.Println("len of metadataSlice", len(fileMetadataSlice))
	fmt.Println("last structure", fileMetadataSlice[len(fileMetadataSlice)-1])
	return fileMetadataSlice
}

func moveFile(spacePosition, spaceLength int, metadata *fileMetadata, newMetadataSlice *[]fileMetadata) (int, int) {
	// call function that will take last file and space, it manipulates in place new slice and returns remaining space
	if spaceLength >= metadata.length {
		(*newMetadataSlice) = append((*newMetadataSlice), fileMetadata{startPosition: spacePosition, length: metadata.length, fileIndex: metadata.fileIndex})
		spaceLength = spaceLength - metadata.length
		spacePosition = spacePosition + metadata.length
		metadata.length = 0
		return spacePosition, spaceLength
	}
	(*newMetadataSlice) = append((*newMetadataSlice), fileMetadata{startPosition: spacePosition, length: spaceLength, fileIndex: metadata.fileIndex})
	metadata.length = metadata.length - spaceLength
	return spacePosition + spaceLength, 0

}
func moveFiles(metadataSlice []fileMetadata) []fileMetadata {
	index := 0
	newMetadataSlice := make([]fileMetadata, 0)
	for index < len(metadataSlice)-1 {
		spacePosition := metadataSlice[index].startPosition + metadataSlice[index].length
		spaceLength := metadataSlice[index+1].startPosition - spacePosition
		for spaceLength > 0 && index+1 < len(metadataSlice) {
			spacePosition, spaceLength = moveFile(spacePosition, spaceLength, &(metadataSlice[len(metadataSlice)-1]), &newMetadataSlice)
			if metadataSlice[len(metadataSlice)-1].length <= 0 {
				metadataSlice = metadataSlice[:len(metadataSlice)-1]
			}
		}
		index += 1
	}
	// connects 2 slices and return them
	return append(newMetadataSlice, metadataSlice...)
}

func countResult(metadataSlice []fileMetadata) int {
	result := 0
	length := 0
	for _, metadata := range metadataSlice {
		for index := range metadata.length {
			result += (metadata.startPosition + index) * metadata.fileIndex
		}
		if length != metadata.startPosition {
			panic("wrong starting position")
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
	output := parseFilesMetadata("input.txt")
	result := moveFiles(output)
	writeFieldToFile(result, "output.txt")
	fmt.Println(countResult(output))
	//fmt.Println(countResult(result))
	fmt.Println(countResult(result))
}
