package main

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json::"color, omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bpgart", "Ingrid Bergman"}},
}
