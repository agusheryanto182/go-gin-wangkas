package transaction

type Service interface {
	CreateData(input Transaction) (Transaction, error)
	UpdateData(input Transaction) (Transaction, error)
	DeleteData(ID int) error
	GetAllDataTransactions() ([]Transaction, error)
	GetByID(ID int) (Transaction, error)
	GetByWeekID(weekID int) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateData(input Transaction) (Transaction, error) {
	result, err := s.repository.Save(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *service) UpdateData(input Transaction) (Transaction, error) {
	result, err := s.repository.Update(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *service) DeleteData(ID int) error {
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetByID(ID int) (Transaction, error) {
	result, err := s.repository.FindByID(ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *service) GetByWeekID(weekID int) ([]Transaction, error) {
	result, err := s.repository.FindAllByWeekID(weekID)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *service) GetAllDataTransactions() ([]Transaction, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return result, err
	}
	return result, nil
}
