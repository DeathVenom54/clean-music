package main

import (
	"github.com/bogem/id3v2"
	"strings"
)

func WriteSongData(fileName string, data *SongData) error {
	tag, openErr := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if openErr != nil {
		return openErr
	}

	tag.SetTitle(data.Title)
	tag.SetAlbum(data.AlbumName)
	tag.SetArtist(strings.Join(data.Artists, ", "))

	if data.AlbumReleaseDate != "" {
		tag.SetYear(data.AlbumReleaseDate)
	}

	saveErr := tag.Save()
	if saveErr != nil {
		return saveErr
	}
	return nil
}
