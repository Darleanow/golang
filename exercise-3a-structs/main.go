// Command exercise-3a-structs covers structs and their associated methods
// through a few geometric shapes (Point, Rectangle, Circle).
package main

import (
	"fmt"
	"math"
)

// Point represents a coordinate in a 2D plane.
type Point struct {
	X float64
	Y float64
}

// DistanceTo returns the euclidean distance between the receiver point and q.
// Value receiver: it only reads the point, no mutation needed.
func (p Point) DistanceTo(q Point) float64 {
	return math.Sqrt(math.Pow(q.X-p.X, 2) + math.Pow(q.Y-p.Y, 2))
}

// Rectangle is defined by its lower-left (Min) and upper-right (Max) corners.
type Rectangle struct {
	Min Point
	Max Point
}

// Width returns the horizontal size of the rectangle.
func (r Rectangle) Width() float64 {
	return r.Max.X - r.Min.X
}

// Height returns the vertical size of the rectangle.
func (r Rectangle) Height() float64 {
	return r.Max.Y - r.Min.Y
}

// Area returns the surface of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width() * r.Height()
}

// Perimeter returns the perimeter of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width() + r.Height())
}

// Move shifts the rectangle by (dx, dy).
// Pointer receiver: it mutates the original rectangle, so we must operate on
// the instance itself and not on a copy.
func (r *Rectangle) Move(dx, dy float64) {
	r.Min.X += dx
	r.Min.Y += dy
	r.Max.X += dx
	r.Max.Y += dy
}

// String makes Rectangle satisfy fmt.Stringer, so it is printed nicely by fmt.
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle Min(%.1f, %.1f) Max(%.1f, %.1f) - largeur %.1f, hauteur %.1f",
		r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, r.Width(), r.Height())
}

// Circle is defined by its center and radius.
type Circle struct {
	Center Point
	Radius float64
}

// Area returns the surface of the circle.
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Circumference returns the circumference of the circle.
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// Scale multiplies the radius by factor.
// Pointer receiver: it mutates the original circle's radius.
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

// String makes Circle satisfy fmt.Stringer.
func (c Circle) String() string {
	return fmt.Sprintf("Cercle de centre (%.1f, %.1f) et rayon %.1f", c.Center.X, c.Center.Y, c.Radius)
}

// NewRectangle validates the dimensions before building a Rectangle, returning
// an error when the corners would yield a negative width or height.
func NewRectangle(min, max Point) (Rectangle, error) {
	if max.X < min.X || max.Y < min.Y {
		return Rectangle{}, fmt.Errorf("dimensions invalides : Max %v doit être >= Min %v", max, min)
	}
	return Rectangle{Min: min, Max: max}, nil
}

// NewCircle validates the radius before building a Circle.
func NewCircle(center Point, radius float64) (Circle, error) {
	if radius < 0 {
		return Circle{}, fmt.Errorf("rayon invalide : %.1f doit être >= 0", radius)
	}
	return Circle{Center: center, Radius: radius}, nil
}

func exercise1Rectangle() {
	fmt.Println("Exercice 1 - Point et Rectangle")

	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 4, Y: 6}
	fmt.Printf("Distance entre p1 et p2 : %.2f\n", p1.DistanceTo(p2))

	rect := Rectangle{Min: Point{X: 0, Y: 0}, Max: Point{X: 5, Y: 3}}
	fmt.Println("Largeur   :", rect.Width())
	fmt.Println("Hauteur   :", rect.Height())
	fmt.Println("Surface   :", rect.Area())
	fmt.Println("Périmètre :", rect.Perimeter())

	rect.Move(2, 1)
	fmt.Println("Après déplacement (2, 1) :")
	fmt.Printf("  Min(%.1f, %.1f) Max(%.1f, %.1f)\n", rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
}

func exercise2Circle() {
	fmt.Println("\nExercice 2 - Cercle")

	c := Circle{Center: Point{X: 1, Y: 1}, Radius: 2}
	fmt.Printf("Surface       : %.2f\n", c.Area())
	fmt.Printf("Circonférence : %.2f\n", c.Circumference())

	c.Scale(1.5)
	fmt.Println("Après Scale(1.5) :")
	fmt.Printf("  Rayon         : %.1f\n", c.Radius)
	fmt.Printf("  Surface       : %.2f\n", c.Area())
	fmt.Printf("  Circonférence : %.2f\n", c.Circumference())
}

func exercise3Improvements() {
	fmt.Println("\nExercice 3 - Améliorations")

	rect := Rectangle{Min: Point{X: 0, Y: 0}, Max: Point{X: 5, Y: 3}}
	circle := Circle{Center: Point{X: 1, Y: 1}, Radius: 2}
	fmt.Println(rect)
	fmt.Println(circle)

	if _, err := NewRectangle(Point{X: 5, Y: 5}, Point{X: 0, Y: 0}); err != nil {
		fmt.Println("Rectangle rejeté :", err)
	}
	if _, err := NewCircle(Point{X: 0, Y: 0}, -3); err != nil {
		fmt.Println("Cercle rejeté    :", err)
	}

	if r, err := NewRectangle(Point{X: 0, Y: 0}, Point{X: 4, Y: 2}); err == nil {
		fmt.Println("Rectangle valide :", r)
	}
}

func main() {
	exercise1Rectangle()
	exercise2Circle()
	exercise3Improvements()
}
