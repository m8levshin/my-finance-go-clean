package gorm

func mapList[F any, T any](from []*F, mapFunc func(*F) *T) []*T {
	result := make([]*T, 0, len(from))
	for _, itemToConvert := range from {
		result = append(result, mapFunc(itemToConvert))
	}
	return result
}
