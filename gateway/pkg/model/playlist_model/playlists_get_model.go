package playlist_model

import "gateway/pkg/proto/playlist_service_proto"

type PlaylistsGetModel struct {
	Playlists []PlaylistGetModel `json:"playlists"`
}

func (model *PlaylistsGetModel) ToModel(resp *playlist_service_proto.GetPlaylistsResp) {
	var playlists []PlaylistGetModel
	for _, playlist := range resp.Playlists {
		playlistModel := PlaylistGetModel{}
		playlistModel.ToModel(playlist)
		playlists = append(playlists, playlistModel)
	}
	model.Playlists = playlists
}
