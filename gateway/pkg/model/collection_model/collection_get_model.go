package collection_model

import (
	"gateway/pkg/model/music_model"
	"gateway/pkg/model/playlist_model"
	"gateway/pkg/proto/collection_service_proto"
	"gateway/pkg/service/album_service"
	grpc_service "gateway/pkg/service/grpc_service/music"
	"gateway/pkg/service/playlist_service"
)

type CollectionGetModel struct {
	Id            string                             `json:"id"`
	UserId        string                             `json:"user_id"`
	Songs         []*music_model.SongGetModel        `json:"songs"`
	Albums        []*playlist_model.PlaylistGetModel `json:"albums"`
	Playlists     []*playlist_model.PlaylistGetModel `json:"playlists"`
	Genres        []string                           `json:"genres"`
	CreatedGenres []string                           `json:"created_genres"`
	Followed      []string                           `json:"followed"`
	Subscribed    []string                           `json:"subscribed"`
}

func (model *CollectionGetModel) ToModel(resp *collection_service_proto.GetFavouritesResp) {
	model.Id = resp.Id
	model.UserId = resp.UserId
	model.addSongs(resp.Songs)
	model.addPlaylists(resp.Playlists)
	model.addAlbums(resp.Albums)
	model.Genres = resp.Genres
	model.CreatedGenres = resp.CreatedGenres
	model.Subscribed = resp.Subscribed
	model.Followed = resp.Followed
}

func (model *CollectionGetModel) addSongs(songIds []string) {
	var songs []*music_model.SongGetModel
	for _, songId := range songIds {
		generateSongModel(songId, &songs)
	}
	model.Songs = songs
}

func generateSongModel(songId string, songs *[]*music_model.SongGetModel) {
	resp, err := grpc_service.GetSongById(songId)
	if err != nil || resp.Error != "" {
		*songs = append(*songs, &music_model.SongGetModel{SongId: songId})
		return
	}
	song := &music_model.SongGetModel{}
	song.ToModel(resp)
	*songs = append(*songs, song)
}

func (model *CollectionGetModel) addPlaylists(playlistIds []string) {
	var playlists []*playlist_model.PlaylistGetModel
	for _, playlistId := range playlistIds {
		generatePlaylistModel(playlistId, &playlists)
	}
	model.Playlists = playlists
}

func generatePlaylistModel(playlistId string, playlists *[]*playlist_model.PlaylistGetModel) {
	playlist, err := playlist_service.GetPlaylist(playlistId)
	if err != nil {
		*playlists = append(*playlists, &playlist_model.PlaylistGetModel{Id: playlistId})
	}
	*playlists = append(*playlists, playlist)
}

func (model *CollectionGetModel) addAlbums(albumIds []string) {
	var albums []*playlist_model.PlaylistGetModel
	for _, albumId := range albumIds {
		generateAlbumModel(albumId, &albums)
	}
	model.Albums = albums
}

func generateAlbumModel(albumId string, albums *[]*playlist_model.PlaylistGetModel) {
	album, err := album_service.GetAlbum(albumId)
	if err != nil {
		*albums = append(*albums, &playlist_model.PlaylistGetModel{Id: albumId})
	}
	*albums = append(*albums, album)
}
