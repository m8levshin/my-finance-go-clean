package utils

import domainasset "github.com/mlevshin/my-finance-go-clean/internal/domain/asset"

func ResolveAssetTypeByName(name string) *domainasset.Type {
	for assetType, assetName := range domainasset.TypeNames {
		if name == assetName {
			return &assetType
		}
	}
	return nil
}
