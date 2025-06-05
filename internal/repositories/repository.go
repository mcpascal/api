package repositories

type Repository struct {
}

func Paginator(page int, limit int) (int, int) {
	if page == 0 {
		page = 1
	}
	switch {
	case limit > 100:
		limit = 100
	case limit <= 0:
		limit = 10
	}
	offset := (page - 1) * limit
	return offset, limit
}
