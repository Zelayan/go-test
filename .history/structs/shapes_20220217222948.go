package structs

type Rectangle struct {
    Width float64
    Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(igth float64, height float64) float64 {
	return wigth * height
}