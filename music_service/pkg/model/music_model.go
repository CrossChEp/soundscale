package model

import (
	"time"
)

type Song struct {
	Id            string    `bson:"_id"`
	Name          string    `bson:"name"`
	Genre         string    `bson:"genre"`
	AuthorId      string    `bson:"author"`
	Collaborators []string  `bson:"collaborators"`
	ReleaseDate   time.Time `bson:"release_date"`
	Listened      []string  `bson:"listened"`
	Played        []string  `bson:"played"`
	Liked         []string  `bson:"liked"`
	Disliked      []string  `bson:"disliked"`
	Exclusive     string    `bson:"exclusive"`
}
