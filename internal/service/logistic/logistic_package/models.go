package logistic_package

import "fmt"

var allEntities = []*Package{
	{Title: "1", Sizes: Sizes{
		Width:  10,
		Height: 20,
		Long:   10,
	}},
	{Title: "2", Sizes: Sizes{
		Width:  10,
		Height: 20,
		Long:   10,
	}},
	{Title: "3", Sizes: Sizes{
		Width:  10,
		Height: 20,
		Long:   10,
	}},
	{Title: "4", Sizes: Sizes{
		Width:  10,
		Height: 20,
		Long:   10,
	}},
	{Title: "5", Sizes: Sizes{
		Width:  10,
		Height: 20,
		Long:   10,
	}},
}

type Package struct {
	Title string
	Sizes Sizes
}

type Sizes struct {
	Width  int
	Height int
	Long   int
}

func (p *Package) String() string {
	return fmt.Sprintf(`📦 %s (%dx%dx%d)`, p.Title, p.Sizes.Long, p.Sizes.Width, p.Sizes.Height)
}
