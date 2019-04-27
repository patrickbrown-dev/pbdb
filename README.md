# pbdb

[![GoDoc](https://godoc.org/github.com/kineticdial/pbdb?status.svg)](https://godoc.org/github.com/kineticdial/pbdb)

## Description

`pbdb` is a key/value database inspired by chapter 3 of Designing Data-Intensive Applications by Martin Kleppmann. 

> The underlying storage format is very simple: a text file where each line contains a key-value pair, separated by a comma (roughly like a CSV file, ignoring escaping issues). Every call to db_set appends to the end of the file, so if you update a key several times, the old versions of the value are not overwritten — you need to look at the last occurrence of a key in a file to find the latest value [...] (Kleppmann 2017)

It also uses a hash index (like Kleppmann later suggests) to speed up reads.

## Usage

`pbdb` is an experiment and not meant for production by any sense of the word. However, if you wish to run it here's how:

```
pbdb run --data DATABASE_FILE --port PORT
```

This will start the database and bind it to `PORT` (default: `1728`). The database will survive restarts, but will need a moment to reindex its in-memory hash index. You can interact with the database using http.

```
$ curl localhost:1728/set/foo -d bar
$ curl localhost:1728/get/foo
bar$
```

## Caveats

You'll have a bad time if you put commas in your keys.

You'll also have a bad time if your values have new line characters.

Finally, _there's no authentication_. You have been warned.

## Works Cited

Kleppmann, Martin. Designing Data-Intensive Applications (Kindle Locations 1934-1937). O'Reilly Media. Kindle Edition. 

