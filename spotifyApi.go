package main

import (
	"context"
	"errors"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type SongData struct {
	Title            string
	AlbumName        string
	AlbumReleaseDate string
	Artists          []string
	Duration         int
}

const (
	clientID     = "6e174ceb365a4f33abec6223e337247b"
	clientSecret = "fcbecfea304143b99dde61c19d932116"
	searchLimit  = 3
)

func authenticateClient() (*spotify.Client, error) {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		return nil, err
	}

	client := spotify.Authenticator{}.NewClient(token)
	return &client, nil
}

func searchSongs(client *spotify.Client, name string) ([]SongData, error) {
	var limit = searchLimit
	result, err := client.SearchOpt(name, spotify.SearchTypeTrack, &spotify.Options{Limit: &limit})
	if err != nil {
		return []SongData{}, err
	}

	if result.Tracks == nil {
		return []SongData{}, errors.New("song not found")
	}

	var songs = []SongData{}

	for _, track := range result.Tracks.Tracks {
		data := SongData{
			Title:            track.Name,
			AlbumName:        track.Album.Name,
			AlbumReleaseDate: track.Album.ReleaseDate,
			Duration:         track.Duration,
			Artists:          []string{},
		}
		for _, artist := range track.Artists {
			data.Artists = append(data.Artists, artist.Name)
		}
		songs = append(songs, data)
	}

	return songs, nil
}
