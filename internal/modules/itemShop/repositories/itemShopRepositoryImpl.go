package repositories

type itemShopRepositoryImpl struct{}

func NewItemShopRepositoryImpl() ItemShopRepository {
	return &itemShopRepositoryImpl{}
}
