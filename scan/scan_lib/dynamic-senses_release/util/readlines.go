package util

import (
    "strings"
    "io/ioutil"
)


func Read_lines(infile string) ([]string,  error) {
    f, err := ioutil.ReadFile(infile)
    return strings.Split(string(f), "\n"), err
}
