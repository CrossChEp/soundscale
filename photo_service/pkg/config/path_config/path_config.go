package path_config

import "flag"

var (
	//
	StaticDir          = flag.String("sdir", "./", "path to static directory")
	PFPPath            = *StaticDir + "photos/pfps"
	LogsPath           = *StaticDir + "logs/logs.log"
	SongCoversPath     = *StaticDir + "photos/songs/covers/"
	PlaylistCoversPath = *StaticDir + "photos/playlists/covers"
	AlbumCoverPath     = *StaticDir + "photos/albums/covers"
	PostPath           = *StaticDir + "photos/posts"
)
