package tetris

type Vector struct {
	x, y int
}

func (v1 Vector) plus(v2 Vector) Vector {
	return Vector{v1.x + v2.x, v1.y + v2.y}
}

func (v1 Vector) equals(v2 Vector) bool {
	return v1.x == v2.x && v1.y == v2.y
}
