package service_address_config

import (
	"flag"
	"fmt"
)

var (
	UserServicePort                     = flag.Int("usport", defaultUserServicePort, "port of user_service service")
	defaultUserServiceAddress           = fmt.Sprintf("%s:%d", "localhost", *UserServicePort)
	UserServiceAddress                  = flag.String("usaddr", defaultUserServiceAddress, "address of user_service service")
	Port                                = flag.Int("port", defaultPort, "port for gateway")
	defaultAddress                      = fmt.Sprintf("%s:%d", "", *Port)
	Address                             = flag.String("addr", defaultAddress, "address for gateway")
	TokenServicePort                    = flag.Int("tsport", defaultTokenServicePort, "port of token service")
	defaultTokenServiceAddress          = fmt.Sprintf("%s:%d", "localhost", *TokenServicePort)
	TokenServiceAddress                 = flag.String("tsaddr", defaultTokenServiceAddress, "address of token service")
	PhotoServicePort                    = flag.Int("psport", defaultPhotoServicePort, "port of photo service")
	defaultPhotoServiceAddress          = fmt.Sprintf("%s:%d", "localhost", *PhotoServicePort)
	PhotoServiceAddress                 = flag.String("psaddr", defaultPhotoServiceAddress, "address of photo service")
	MusicServicePort                    = flag.Int("msport", defaultMusicServicePort, "port of music service")
	defaultMusicServiceAddress          = fmt.Sprintf("%s:%d", "localhost", *MusicServicePort)
	MusicServiceAddress                 = flag.String("msaddr", defaultMusicServiceAddress, "address of music service")
	PlaylistServicePort                 = flag.Int("plsport", defaultPlaylistServicePort, "port of playlist service")
	defaultPlaylistServiceAddress       = fmt.Sprintf("%s:%d", "localhost", *PlaylistServicePort)
	PlaylistServiceAddress              = flag.String("pladdr", defaultPlaylistServiceAddress, "address of playlist service")
	CollectionServicePort               = flag.Int("csport", defaultCollectionServicePort, "port of collection service")
	defaultCollectionServiceAddress     = fmt.Sprintf("%s:%d", "localhost", *CollectionServicePort)
	CollectionServiceAddress            = flag.String("csaddr", defaultCollectionServiceAddress, "address of collection service")
	PostServicePort                     = flag.Int("posport", defaultPostServicePort, "port of grpc_post service")
	defaultPostServiceAddress           = fmt.Sprintf("%s:%d", "localhost", *PostServicePort)
	PostServiceAddress                  = flag.String("posaddr", defaultPostServiceAddress, "address of grpc_post service")
	defaultCommentServiceAddress        = fmt.Sprintf("%s:%d", "localhost", defaultCommentServicePort)
	CommentServiceAddress               = flag.String("commaddr", defaultCommentServiceAddress, "address of comment service")
	defaultRecommendationServiceAddress = fmt.Sprintf("%s:%d", "localhost", defaultRecommendationServicePort)
	RecommendationServiceAddress        = flag.String("recaddr", defaultRecommendationServiceAddress, "address of recommendation service")
)

const (
	defaultPort                      = 9999
	defaultUserServicePort           = 8080
	defaultTokenServicePort          = 8794
	defaultPhotoServicePort          = 7354
	defaultMusicServicePort          = 7070
	defaultPlaylistServicePort       = 6061
	defaultCollectionServicePort     = 9876
	defaultPostServicePort           = 5555
	defaultCommentServicePort        = 1000
	defaultRecommendationServicePort = 4561
)
