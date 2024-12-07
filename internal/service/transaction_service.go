package service

type transactionService struct{}

func NewTransactionService() ITransactionService {
	return &transactionService{}
}
