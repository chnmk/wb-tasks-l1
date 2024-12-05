package point

type Point struct {
	x int
	y int
}

func (p *Point) SetPoints(new_x int, new_y int) {
	p.x = new_x
	p.y = new_y
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}
