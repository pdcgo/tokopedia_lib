package autochat

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdcgo/common_conf/pdc_application"
)

type AutochatMessage struct {
	Patterns []string
	Messages []string
}

const (
	PatternDir = "message_patterns"
	ReplyDir   = "message_replies"
)

func fileTxtWalk(dir string, handler func(file *os.File) error) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {

		if info == nil {
			return nil
		}

		isTxt := filepath.Ext(info.Name()) == ".txt"
		if !info.IsDir() && isTxt {

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()
			handler(file)
		}

		return nil
	})
}

func fileLineSplit(file *os.File) (lines []string, err error) {

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func NewAutochatMessage(base pdc_application.BaseApplication) (*AutochatMessage, error) {

	// create dir first
	for _, dir := range []string{PatternDir, ReplyDir} {
		dir = base.Path(dir)
		if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}

	automessage := AutochatMessage{}

	// load patterns
	err := fileTxtWalk(base.Path(PatternDir), func(file *os.File) error {
		lines, err := fileLineSplit(file)
		automessage.Patterns = lines
		return err
	})
	if err != nil {
		return &automessage, err
	}
	if len(automessage.Patterns) == 0 {
		log.Println("[ warning ] message patterns kosong")
	}

	// load messages
	err = fileTxtWalk(base.Path(ReplyDir), func(file *os.File) error {
		lines, err := fileLineSplit(file)
		automessage.Messages = append(automessage.Messages, lines...)
		return err
	})
	if err != nil {
		return &automessage, err
	}

	if len(automessage.Messages) == 0 {
		log.Println("[ warning ] message replies kosong")
	}
	return &automessage, nil
}

func (p *AutochatMessage) InPattern(msg string) bool {

	msg = strings.ToLower(msg)
	for _, pat := range p.Patterns {

		pat = strings.ToLower(pat)
		inPattern := strings.Contains(msg, pat)
		if inPattern {
			return inPattern
		}
	}

	return false
}

func (p *AutochatMessage) GetMessage() string {

	msglen := len(p.Messages)
	if msglen > 0 {
		msgind := rand.Intn(msglen - 1)
		return p.Messages[msgind]
	}

	return ""
}
