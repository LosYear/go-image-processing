package model

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type VectorPolygon struct {
	Transparent bool    `json:"transparent"`
	Points      []Point `json:"points"`
}

type Dimensions struct {
	width, height float64
	minX, minY    float64
}

type VectorModel struct {
	Polygons   []VectorPolygon
	Dimensions Dimensions
}
