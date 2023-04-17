package memory_rw

import (
	"github.com/google/uuid"
	"github.com/mlevshin/my-finance-go-clean/internal/domain"
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"
	"sync"
)

type transactionRW struct {
	store *sync.Map
}

func (t transactionRW) GetTransactionsByAsset(assetId domain.Id) ([]*domainasset.Transaction, error) {
	value, _ := t.store.LoadOrStore(uuid.UUID(assetId), []domainasset.Transaction{})
	transactions := value.([]domainasset.Transaction)

	var result []*domainasset.Transaction
	for _, transaction := range transactions {
		result = append(result, &transaction)
	}

	return result, nil
}

func (t transactionRW) AddTransaction(assetId domain.Id, transaction domainasset.Transaction) error {
	assetUUID := uuid.UUID(assetId)
	value, _ := t.store.LoadOrStore(assetUUID, []domainasset.Transaction{})
	transactions := value.([]domainasset.Transaction)
	transactions = append(transactions, transaction)
	t.store.Swap(assetUUID, transactions)
	return nil
}

func (t transactionRW) SaveTransactionsByAsset(assetId domain.Id, transactions []domainasset.Transaction) error {
	assetUUID := uuid.UUID(assetId)
	t.store.Store(assetUUID, transactions)
	return nil
}
