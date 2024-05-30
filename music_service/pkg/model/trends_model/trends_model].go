package trends_model

import "music_service/pkg/model"

type TrendsModel struct {
	Songs []model.Song `json:"songs"`
}
