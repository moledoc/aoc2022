package read

import (
	"bufio"
	"os"
)

func File(filename string) ([]string, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return []string{}, err
	}
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		input = append(input, val)
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return input, nil
}
