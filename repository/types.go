// This file contains types that are used in the repository layer.
package repository

type GetEstateByIdOutput struct {
	Id     string
	Length int
	Width  int
}

type InsertEstateInput struct {
	Length int
	Width  int
}

type InsertEstateObjectInput struct {
	Height int
	X      int
	Y      int
}

type GetEstateStatsOutput struct {
	Count  int
	Max    int
	Median int
	Min    int
}

type EstateObject struct {
	XLocation int
	YLocation int
	Height    int
}

type Estate struct {
	Width  int
	Length int
}
