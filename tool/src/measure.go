package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func readLines(path string) []string {
    srcBytes, err := ioutil.ReadFile(path)
    check(err)
    return strings.Split(string(srcBytes), "\n")
}

var todoPat = regexp.MustCompile("\\/\\/ todo: ")

func main() {
    sourcePaths, err := filepath.Glob("./src/0*/*")
    check(err)
    foundLongFile := false
    for _, sourcePath := range sourcePaths {
        foundLongLine := false
        lines := readLines(sourcePath)
        for _, line := range lines {
            if !foundLongLine && !todoPat.MatchString(line) && (len(line) > 58) {
                fmt.Println(sourcePath)
                foundLongLine = true
                foundLongFile = true
            }
        }
    }
    if foundLongFile {
        os.Exit(1)
    }
}