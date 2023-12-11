package structs

type Album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{Id: 1, Title: "Blue Train", Artist: "Jhon Coltrane", Price: 56.99},
	{Id: 2, Title: "Blue Train", Artist: "Jhon Coltrane", Price: 56.99},
	{Id: 3, Title: "Blue Train", Artist: "Jhon Coltrane", Price: 56.99},
}
