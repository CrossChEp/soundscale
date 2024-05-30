package constants

const (
	MongoAddrEnvName = "mongoaddr"
	TypePost         = "post"
	TypeSong         = "song"
	TypePlaylist     = "playlist"
	TypeAlbum        = "album"
)

var (
	EntityTypes = []string{TypePost, TypeAlbum, TypePlaylist, TypeSong}
)
