package internal

// Beer representation of beer into data struct
type Beer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Tagline     string `json:"tagline"`
	FirstBrewed string `json:"first_brewed"`
	Description string `json:"description"`
}

// BeerRepo definiton of methods to access a data beer
type BeerRepo interface {
	GetBeers() ([]Beer, error)
}

// NewBeer initialize struct beer
func NewBeer(ID int, name, tagLine, firstBrewed, description string) (b Beer) {
	b = Beer{
		ID:          ID,
		Name:        name,
		Tagline:     tagLine,
		FirstBrewed: firstBrewed,
		Description: description,
	}
	return
}
