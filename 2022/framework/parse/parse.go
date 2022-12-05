package parse

import (
    "bufio"
    "os"
)

func GetLines(file string) ([]string, error) {
    fileIn, err := os.Open(file)

    if err != nil {
        return nil, err
    }

    defer fileIn.Close()

    fileScanner := bufio.NewScanner(fileIn)
    fileScanner.Split(bufio.ScanLines)

    var res []string

    for fileScanner.Scan() {
        res = append(res, fileScanner.Text())
    }
    return res, nil
}
