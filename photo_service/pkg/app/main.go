package main

import (
	"flag"
	"fmt"
	"os"
	"photo_service/pkg/config/logger_config"
	"photo_service/pkg/config/path_config"
	"photo_service/pkg/service/setup"
)

func main() {
	flag.Parse()
	fmt.Println(*path_config.StaticDir)
	if exists(*path_config.StaticDir) {
		path_config.PFPPath = *path_config.StaticDir + "photos/pfps"
		if err := os.MkdirAll(path_config.PFPPath, 0777); err != nil {
			logger_config.InfoConsoleLogger.Println(fmt.Sprintf("%s %v", path_config.PFPPath, err))
		}
		path_config.PostPath = *path_config.StaticDir + "photos/posts"
		os.Mkdir(path_config.PostPath, os.FileMode(0777))
		os.Mkdir(*path_config.StaticDir+"photos/songs", os.FileMode(0777))
		path_config.SongCoversPath = *path_config.StaticDir + "photos/songs/covers/"
		os.Mkdir(path_config.SongCoversPath, os.FileMode(0522))
		os.Mkdir(*path_config.StaticDir+"photos/playlists", os.FileMode(0522))
		path_config.PlaylistCoversPath = *path_config.StaticDir + "photos/playlists/covers"
		os.Mkdir(path_config.PlaylistCoversPath, os.FileMode(0522))
		os.Mkdir(*path_config.StaticDir+"photos/albums", os.FileMode(0522))
		path_config.AlbumCoverPath = *path_config.StaticDir + "photos/albums/covers"
		os.Mkdir(path_config.AlbumCoverPath, os.FileMode(0522))
	} else {
		logger_config.InfoConsoleLogger.Println(fmt.Sprintf("%s", *path_config.StaticDir))
		path_config.PFPPath = *path_config.StaticDir + "photos/pfps"
		path_config.PostPath = *path_config.StaticDir + "photos/posts"
		path_config.SongCoversPath = *path_config.StaticDir + "photos/songs/covers/"
		path_config.PlaylistCoversPath = *path_config.StaticDir + "photos/playlists/covers"
		path_config.AlbumCoverPath = *path_config.StaticDir + "photos/albums/covers"
	}
	setup.ServiceSetup()
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}
	return true
}
