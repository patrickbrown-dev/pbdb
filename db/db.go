package db

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type writeData struct {
	offset int64
	length int
}

var hashIndex map[string]writeData
var dbFilePath string

func Initialize() {
	dbFilePath = viper.GetString("data")

	log.Printf("Using %s as data file\n", dbFilePath)

	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		log.Printf("%s does not exist; creating file\n", dbFilePath)
		f, err := os.Create(dbFilePath)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}

	log.Println("Initializing indices")

	BuildIndices()
}

// BuildIndices ...
func BuildIndices() {
	hashIndex = make(map[string]writeData)
	f, err := os.Open(dbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var o int64

	for scanner.Scan() {
		b := scanner.Bytes()
		k := strings.SplitN(string(b), ",", 2)[0]
		l := len(b) + len([]byte("\n"))
		hashIndex[k] = writeData{o, l}
		o += int64(l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Set ...
func Set(k string, v []byte) error {
	f, err := os.OpenFile(dbFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get offset of EOF
	o, err := f.Seek(0, 2)
	if err != nil {
		return err
	}

	// Write data to file
	s := fmt.Sprintf("%s,%s\n", k, v)
	l, err := f.WriteAt([]byte(s), o)
	if err != nil {
		return err
	}

	// Update hash index with offset
	hashIndex[k] = writeData{o, l}

	return nil
}

// Get ...
func Get(k string) ([]byte, error) {
	f, err := os.OpenFile(dbFilePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	wd := hashIndex[k]

	// Read bytes into buffer at offset
	b := make([]byte, wd.length)
	_, err = f.ReadAt(b, wd.offset)
	if err != nil {
		return nil, err
	}

	// Trim key, comma, and new line chars
	s := string(b)
	s = strings.TrimLeft(s, k)
	s = strings.TrimLeft(s, ",")
	s = strings.TrimRight(s, "\n")

	return []byte(s), nil
}
