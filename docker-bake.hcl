group "default" {
  targets = [
    "token_service",
    "user_service",
    "post_service",
    "playlist_service",
    "player_service",
    "photo_service",
    "music_service",
    "gateway"
  ]
}

target "token_service" {
  context    = "token_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/token_service"]
}

target "user_service" {
  context    = "user_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/user_service"]
}

target "post_service" {
  context = "post_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/post_service"]
}

target "playlist_service" {
  context = "playlist_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/playlist_service"]
}

target "player_service" {
  context = "player_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/player_service"]
}

target "photo_service" {
  context = "photo_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/photo_service"]
}

target "music_service" {
  context = "music_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/music_service"]
}

target "gateway" {
  context = "gateway"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/gateway"]
}

target "collection_service" {
  context = "collection_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/collection_service"]
}

target "comment_service" {
  context = "comment_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/comment_service"]
}

target "recommendation_service" {
  context = "recommendation_service"
  dockerfile = "Dockerfile"
  tags       = ["docker.io/regretto/recommendation_service"]
}
