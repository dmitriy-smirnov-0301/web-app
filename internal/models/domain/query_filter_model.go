package domain

type QueryFilter struct {
	SortBy string
	Order  string
	Limit  int
	Offset int
}
