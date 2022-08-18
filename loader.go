package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

type File struct {
	Path string
	Name string
}

func (self *File) Ask(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s %s? [y/n]: ", s, self.Name)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func (self *File) Download(destination string) error {
	response, err := http.Get(self.Path)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	out, err := os.Create(destination + self.Name)
	if err != nil {
		return err
	}
	defer out.Close()

	bar := progressbar.DefaultBytes(
		response.ContentLength,
		"Downloading",
	)
	_, err = io.Copy(io.MultiWriter(out, bar), response.Body)
	return err
}
