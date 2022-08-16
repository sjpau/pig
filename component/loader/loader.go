package loader

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sjpau/pig/component/grep"
)

type File struct {
	Path string
	Name string
}

func New(url string) *File {
	return &File{
		Path: url,
		Name: grep.FileNameFromPath(url),
	}
}

func (self *File) Ask(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s %s? [y/n]: ", s, self.Path)
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

func (self *File) Download(url string) error {
	response, err := http.Get(url + self.Path)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if self.Ask("Download") {
		out, err := os.Create(self.Name)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, response.Body)
		return err
	} else {
		return err
	}
}
