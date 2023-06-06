package utils

import (
	domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/finance/model"
	"strings"
)

func ResolveAssetTypeByName(name string) *domainasset.Type {
	for assetType, assetName := range domainasset.TypeNames {
		if strings.EqualFold(assetName, name) {
			return &assetType
		}
	}
	return nil
}
