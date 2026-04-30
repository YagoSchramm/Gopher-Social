package domain

type PaginatedFeedQuery struct {
	Limit  int
	Offset int
	Sort   string
	Search string
	Tags   []string
}
