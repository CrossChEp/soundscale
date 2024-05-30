package response

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"music_service/pkg/model"
	"music_service/pkg/proto/music_service_proto"
)

func CreateSimpleGetSongsResp(songs []model.Song) *music_service_proto.GetSongsResp {
	var songResps []*music_service_proto.GetResp
	for _, song := range songs {
		songResps = append(songResps, createSimpleGetResp(song))
	}
	return &music_service_proto.GetSongsResp{Songs: songResps}
}

func CreateGetSongsResp(songs []*model.Song) *music_service_proto.GetSongsResp {
	var songResps []*music_service_proto.GetResp
	for _, song := range songs {
		songResps = append(songResps, CreateGetResp(song))
	}
	return &music_service_proto.GetSongsResp{Songs: songResps}
}

func CreateGetErrorSongsResp(message string) *music_service_proto.GetSongsResp {
	return &music_service_proto.GetSongsResp{Error: message}
}

func CreateGetResp(song *model.Song) *music_service_proto.GetResp {
	return &music_service_proto.GetResp{
		Id:            song.Id,
		SongName:      song.Name,
		AuthorId:      song.AuthorId,
		Collaborators: song.Collaborators,
		Genre:         song.Genre,
		ReleaseDate:   timestamppb.New(song.ReleaseDate),
		Listened:      song.Listened,
		Played:        song.Played,
		Liked:         song.Liked,
		Disliked:      song.Disliked,
		Exclusive:     song.Exclusive,
	}
}

func createSimpleGetResp(song model.Song) *music_service_proto.GetResp {
	return &music_service_proto.GetResp{
		Id:            song.Id,
		SongName:      song.Name,
		AuthorId:      song.AuthorId,
		Collaborators: song.Collaborators,
		Genre:         song.Genre,
		ReleaseDate:   timestamppb.New(song.ReleaseDate),
		Listened:      song.Listened,
		Played:        song.Played,
		Liked:         song.Liked,
		Disliked:      song.Disliked,
		Exclusive:     song.Exclusive,
	}
}

func CreateErrorResp(message string) *music_service_proto.GetResp {
	return &music_service_proto.GetResp{Error: message}
}
