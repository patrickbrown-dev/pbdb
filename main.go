// pbdb is a key/value database inspired by chapter 3 of Designing
// Data-Intensive Applications by Martin Kleppmann.
//
// “The underlying storage format is very simple: a text file where
// each line contains a key-value pair, separated by a comma
// (roughly like a CSV file, ignoring escaping issues). Every call
// to db_set appends to the end of the file, so if you update a key
// several times, the old versions of the value are not overwritten
// — you need to look at the last occurrence of a key in a file to
// find the latest value [...]” (Kleppmann 2017)
//
// It also uses a hash index (like Kleppmann later suggests) to speed
// up reads.
package main

import (
	"github.com/patrickbrown-dev/pbdb/cmd"
)

func main() {
	cmd.Execute()
}
