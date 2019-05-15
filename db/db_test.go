package db

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("data", "./data")
}

func teardown() {
	d := viper.GetString("data")
	os.Remove(d)
}

func TestSetGet(t *testing.T) {
	Initialize()
	defer teardown()

	b := []byte("foo")
	err := Set("1", b)
	if err != nil {
		t.Fatalf("%s", err)
	}

	bPrime, err := Get("1")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if string(b) != string(bPrime) {
		t.Fatalf("'%s' does not eq '%s'", b, bPrime)
	}
}

func TestMultilineString(t *testing.T) {
	Initialize()
	defer teardown()

	b := []byte("this is\na test")
	err := Set("1", b)
	if err != nil {
		t.Fatalf("%s", err)
	}

	bPrime, err := Get("1")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if string(b) != string(bPrime) {
		t.Fatalf("'%s' does not eq '%s'", b, bPrime)
	}
}

func TestBuildIndices(t *testing.T) {
	f, err := os.Create(viper.GetString("data"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := "Zm9vLGJhcgo=\nYmF6LHF1eAo=\n"

	_, err = f.WriteString(s)
	if err != nil {
		t.Fatal(err)
	}

	Initialize()
	defer teardown()

	foo, _ := Get("foo")
	if err != nil {
		t.Fatal(err)
	}

	baz, _ := Get("baz")
	if err != nil {
		t.Fatal(err)
	}

	if string(foo) != "bar" {
		t.Fatalf("'%s' does not eq 'bar'", foo)
	}

	if string(baz) != "qux" {
		t.Fatalf("'%s' does not eq 'qux'", baz)
	}
}
