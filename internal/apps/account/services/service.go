package services

type AccountService interface {
}

type accountService struct {
}

func NewAccountService() AccountService {
	return &accountService{}
}
