package transaction

type Service interface {
	CreateData(input Transaction) (Transaction, error)
	UpdateData(inputID int, input FormUpdateDataInput) (Transaction, error)
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

func (s *service) UpdateData(inputID int, input FormUpdateDataInput) (Transaction, error) {
	transaction, err := s.repository.FindByID(inputID)
	if err != nil {
		return transaction, err
	}

	transaction.Nama = input.Nama
	transaction.TanggalTransaksi = input.TanggalTransaksi
	transaction.Keterangan = input.Keterangan
	transaction.MingguKe = input.MingguKe
	transaction.JumlahMasuk = input.JumlahMasuk

	result, err := s.repository.Update(transaction)
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
