package store

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type FileStore struct {
	Path string
}

func (s FileStore) Add(item string) error {
	f, err := os.OpenFile(s.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s\n", item))
	if err != nil {
		return err
	}

	return nil
}

func (s FileStore) Next() (string, error) {
	f, err := os.OpenFile(s.Path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", nil
}

func (s FileStore) Pop() (string, error) {
	f, err := os.OpenFile(s.Path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(buf, f)
	if err != nil {
		return "", err
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	nw, err := io.Copy(f, buf)
	if err != nil {
		return "", err
	}

	err = f.Truncate(nw)
	if err != nil {
		return "", err
	}

	err = f.Sync()
	if err != nil {
		return "", err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	if len(line) == 0 {
		return "", nil
	}

	err = f.Close()
	if err != nil {
		return "", err
	}

	return string(line[:len(line)-1]), nil
}
