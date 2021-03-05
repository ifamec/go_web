package model

type Page struct {
	Books        []*Book // books on current page
	PageNumber   int     // current page
	PageSize     int     // page size
	TotalPages   int     // total pages - calc
	TotalRecords int     // total records - from db
}
