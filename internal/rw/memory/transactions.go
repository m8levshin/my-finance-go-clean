package memory

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	"github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"github.com/mlevshin/my-finance-go-clean/internal/uc/rw"
	"sync"
)

type transactionRW struct {
	store *sync.Map
}

func NewMemoryTransactionRW() rw.TransactionRW {
	return transactionRW{
		store: &sync.Map{},
	}
}

func (t transactionRW) GetTransactionsByAsset(assetId domain.Id) ([]*asset.Transaction, error) {
	value, _ := t.store.LoadOrStore(uuid.UUID(assetId), []asset.Transaction{})
	transactions := value.([]asset.Transaction)

	result := make([]*asset.Transaction, 0, len(transactions))
	for _, transaction := range transactions {
		result = append(result, &transaction)
	}

	return result, nil
}

func (t transactionRW) AddTransaction(assetId domain.Id, transaction asset.Transaction) error {
	assetUUID := uuid.UUID(assetId)
	value, _ := t.store.LoadOrStore(assetUUID, []asset.Transaction{})
	transactions := value.([]asset.Transaction)
	transactions = append(transactions, transaction)
	t.store.Swap(assetUUID, transactions)
	return nil
}

func (t transactionRW) SaveTransactionsByAsset(assetId domain.Id, transactions []asset.Transaction) error {
	assetUUID := uuid.UUID(assetId)
	t.store.Store(assetUUID, transactions)
	return nil
}
