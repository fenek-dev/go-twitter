package tweets

type Service struct {
	db Repository
}

func NewService(db Repository) *Service {
	return &Service{
		db: db,
	}
}
