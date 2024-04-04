package store

import (
	"bufio"
	"os"
)

type Storer interface {
	Write(line string) error
	Read() ([]string, error)
}

type FileStorer struct {
	fpath string
}

func New(fpath string) *FileStorer {
	return &FileStorer{fpath}
}

func (fs *FileStorer) Write(line string) error {
	file, err := os.OpenFile(fs.fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(line + "\n")
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStorer) Read() ([]string, error) {
	file, err := os.Open(fs.fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
