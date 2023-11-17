package main

import (
	"fmt"
	"testing"
)

func TestAuthenticateClient(t *testing.T) {
	client, err := authenticateClient()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", client)
}

func TestSearchSong(t *testing.T) {
	client, err := authenticateClient()
	if err != nil {
		t.Error(err)
	}
	songs, searchErr := searchSongs(client, "SnapInsta.io - Twenty One Pilots - Chlorine (Lyrics) (128 kbps).mp3")
	if searchErr != nil {
		t.Error(err)
	}
	for _, song := range songs {
		fmt.Printf("%#v\n", song)
	}
}
