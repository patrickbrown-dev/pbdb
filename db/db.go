// Package db provides low-level database operations with a simple
// "Get" and "Set" API. It is important to note that the database
// must be initialized (with the method Initialize) before use.
package db

import (
	"bufio"
	"encoding/base64"
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

// Initialize creates the data file if it does not exist, and then
// builds an in-memory hash index from the data file.
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

	buildIndices()
}

func buildIndices() {
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
		bytes, err := base64.StdEncoding.DecodeString(string(b))
		if err != nil {
			log.Fatal(err)
		}
		k := strings.SplitN(string(bytes), ",", 2)[0]
		l := len(b) + len([]byte("\n"))
		hashIndex[k] = writeData{o, l}
		o += int64(l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Set stores a value at a given key. Under the hood, Set is simply
// appending to the datafile and never actually writes over any old
// data.
func Set(k string, v []byte) error {
	err := validateKey(k)
	if err != nil {
		return err
	}

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

	s := fmt.Sprintf("%s,%s", k, v)
	s = base64.StdEncoding.EncodeToString([]byte(s)) + "\n"

	// Write (i.e. append) data to file
	l, err := f.Write([]byte(s))
	if err != nil {
		return err
	}

	// Update hash index with offset
	hashIndex[k] = writeData{o, l}

	return nil
}

// Get retrieves a value by a given key, taking advantage of the hash
// index to provide constant time lookups.
func Get(k string) ([]byte, error) {
	err := validateKey(k)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(dbFilePath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	wd := hashIndex[k]

	// Read bytes into buffer at offset
	b := make([]byte, wd.length)
	_, err = f.ReadAt(b, wd.offset)
	if err != nil {
		return nil, err
	}

	// Trim key, comma, and new line chars
	bytes, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return nil, err
	}
	s := string(bytes)
	s = strings.TrimPrefix(s, k+",")
	s = strings.TrimSuffix(s, "\n")

	return []byte(s), nil
}
