package point

type Point struct {
	x float64
	y float64
}

func New() *Point {
	return &Point{}
}
func (p *Point) SetX(value float64) {
	p.x = value
}
func (p *Point) SetY(value float64) {
	p.y = value
}
func (p *Point) GetX() float64 {
	return p.x
}
func (p *Point) GetY() float64 {
	return p.y
}
