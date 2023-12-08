package parse

import (
	"bufio"
	"os"
)

func GetLines(file string) ([]string, error) {
	return GetLinesAs[string](file, func(s string) (string, error) { return s, nil })
}

func GetLinesAsOne[T any](file string, conv func([]string) (T, error)) (T, error) {
	lines, err := GetLines(file)
	if err != nil {
		var empty T
		return empty, err
	}

	return conv(lines)
}

func GetLinesAs[T any](file string, conv func(string) (T, error)) ([]T, error) {
	fileIn, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	defer fileIn.Close()

	fileScanner := bufio.NewScanner(fileIn)
	fileScanner.Split(bufio.ScanLines)

	var res []T

	for fileScanner.Scan() {
		elm, err := conv(fileScanner.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, elm)
	}
	return res, nil
}
