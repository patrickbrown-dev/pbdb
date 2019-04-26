package db_test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/kineticdial/pbdb/db"
)

func TestPutGet(t *testing.T) {
	b := []byte("foo")
	err := db.Put("1", b)
	if err != nil {
		t.Fatalf("%s", err)
	}

	bPrime, err := db.Get("1")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if string(b) != string(bPrime) {
		t.Fatalf("%s does not eq %s", b, bPrime)
	}

	os.Remove(viper.GetString("dbFilePath"))
}

func TestBuildIndices(t *testing.T) {
	f, err := os.Create(viper.GetString("dbFilePath"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := "foo,bar\nbaz,qux\n"

	_, err = f.WriteString(s)
	if err != nil {
		t.Fatal(err)
	}

	db.BuildIndices()

	foo, _ := db.Get("foo")
	if err != nil {
		t.Fatal(err)
	}

	baz, _ := db.Get("baz")
	if err != nil {
		t.Fatal(err)
	}

	if string(foo) != "bar" {
		t.Fatalf("%s does not eq bar", foo)
	}

	if string(baz) != "qux" {
		t.Fatalf("%s does not eq qux", baz)
	}

	os.Remove(viper.GetString("dbFilePath"))
}
