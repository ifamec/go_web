package model

type Page struct {
	Books        []*Book // books on current page
	PageNumber   int     // current page
	PageSize     int     // page size
	TotalPages   int     // total pages - calc
	TotalRecords int     // total records - from db
	PriceMin     float64
	PriceMax     float64
}

func (p *Page) HasPrev() bool {
	return p.PageNumber > 1
}
func (p *Page) HasNext() bool {
	return p.PageNumber < p.TotalPages
}
func (p *Page) GetPrev() int {
	if p.HasPrev() {
		return p.PageNumber - 1
	} else {
		return 1
	}
}
func (p *Page) GetNext() int {
	if p.HasNext() {
		return p.PageNumber + 1
	} else {
		return p.TotalPages
	}
}