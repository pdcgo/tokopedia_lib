package main

import (
	"log"
	"testing"
)

func TestPurnaSessionWalk(t *testing.T) {
	WalkPurnaSession("../../asset_sampah/session/", func(file string, sess *PurnaSession) error {
		log.Println(file)
		return nil
	})
}
