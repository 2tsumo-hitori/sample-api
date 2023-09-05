package model

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = make(map[string]Album)

func init() {
	albums["1"] = Album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}
	albums["2"] = Album{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99}
	albums["3"] = Album{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99}
}

func Albums() map[string]Album {
	return albums
}

func SetAlbum(album Album) {
	albums[album.ID] = album
}
