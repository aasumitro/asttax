package service

type ITransactionService interface{}

type transactionService struct{}

func NewTransactionService() ITransactionService {
	return &transactionService{}
}
