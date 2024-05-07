package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type InventoryRepository interface {
	AddProduct(product models.AddInventories, url string) (models.ProductsResponse, error)
	AddInventory(inventory models.AddInventories, url string) (models.ProductsResponse, error)
	ListProducts() ([]models.ProductsResponseDisp, error)
	EditInventory(domain.Inventories, int) (domain.Inventories, error)
	DeleteInventory(id string) error
	CheckInventory(pid int) (bool, error)
	CheckInventoryByCatAndName(cat int, prdct string) (bool, error)
	UpdateInventory(pid int, stock int) (models.ProductsResponse, error)
	ShowIndividualProducts(id string) (models.ProductsResponse, error)
	CheckStock(inventory_id int) (int, error)
	FetchProductDetails(productId uint) (models.Inventories, error)
	GetInventory(prefix string) ([]models.ProductsResponse, error)
	FilterByCategory(CategoryIdInt int) ([]models.ProductsResponse, error)
	ImageUploader(inventoryID int, url string) error
}
