package music_model

import "gateway/pkg/proto/music_service_proto"

type SongsGetModel struct {
	Songs []*SongGetModel `json:"songs"`
}

func (model *SongsGetModel) ToModel(resp *music_service_proto.GetSongsResp) {
	for _, song := range resp.Songs {
		songModel := &SongGetModel{}
		songModel.ToModel(song)
		model.Songs = append(model.Songs, songModel)
	}
}
