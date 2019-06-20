package controllers

import (
	"testing"

	"github.com/mike-webster/golang-validation/models"
)

func TestPostAlbum(t *testing.T) {
	t.Run("AlbumTests", func(t *testing.T) {
		tests := []testCase{
			testCase{
				Name:        "artists-not-provided",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist"},
				ExpMessages: []string{"Artist is required"},
				Body: models.AlbumExample{
					Name: "dude ranch",
				},
			},
			testCase{
				Name:        "artists-empty",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist"},
				ExpMessages: []string{"Artist must contain at least 1 entry"},
				Body: models.AlbumExample{
					Artist: []string{},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "artists-too-large",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist"},
				ExpMessages: []string{"Artist must contain no more than 5 entries"},
				Body: models.AlbumExample{
					Artist: []string{"fda", "fda", "fdas", "fdas", "fdas", "fda"},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "artist-entry-too-short",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist[0]"},
				ExpMessages: []string{"Artist [ 0 ] must contain at least 2 characters"},
				Body: models.AlbumExample{
					Artist: []string{"f"},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "artist-entry-too-long",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Artist[0]"},
				ExpMessages: []string{"Artist [ 0 ] must contain no more than 50 characters"},
				Body: models.AlbumExample{
					Artist: []string{"asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasd"},
					Name:   "dude ranch",
				},
			},
			testCase{
				Name:        "name-not-provided",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name is required"},
				Body: models.AlbumExample{
					Artist: []string{"blink 182"},
				},
			},
			testCase{
				Name:        "name-too-short",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name must contain at least 2 characters"},
				Body: models.AlbumExample{
					Artist: []string{"blink 182"},
					Name:   "a",
				},
			},
			testCase{
				Name:        "name-too-long",
				Path:        "/album",
				ExpCode:     400,
				ExpFields:   []string{"Name"},
				ExpMessages: []string{"Name must contain no more than 50 characters"},
				Body: models.AlbumExample{
					Artist: []string{"blink 182"},
					Name:   "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasd",
				},
			},
		}
		runTests(t, tests, GetRouter())
	})
}
