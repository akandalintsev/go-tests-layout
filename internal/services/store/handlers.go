package delivery

type Store interface {
	Stocks(ids []string) ([]ProductStock, error)
}

type StoreService struct {
	config struct {
		baseUrl string
	}
}

func (s *StoreService) Stocks(ids []string) ([]ProductStock, error) {
	return nil, nil
}
