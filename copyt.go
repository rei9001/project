package main

import "fmt"

type Pos interface {
	Print()
	Copy() Pos
	Move(x, y, z int)
	Rename(string)
}

type pos struct {
	X, Y, Z int
	Name    string
}

func NewPos(name string, x, y, z int) Pos {
	p := new(pos)
	p.Name = name
	p.X = x
	p.Y = y
	p.Z = z
	return p
}

func (p *pos) Print() {
	fmt.Printf("%s (%d, %d, %d)\n", p.Name, p.X, p.Y, p.Z)
}

func (p *pos) Copy() Pos {
	copied := NewPos(p.Name, p.X, p.Y, p.Z)
	return copied
}

func (p *pos) Move(x, y, z int) {
	p.X = x
	p.Y = y
	p.Z = z
}

func (p *pos) Rename(name string) {
	p.Name = name
}

func main() {
	p1 := NewPos("wall", 0, 0, 0)
	p1.Print()

	p2 := p1.Copy()
	p2.Move(1, 0, 0)

	p1.Print()
	p2.Print()
}
