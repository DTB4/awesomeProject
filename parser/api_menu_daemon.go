package parser

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DTB4/logger/v2"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func NewMenuParser(cfg *models.ParserConfig, logger *logger.Logger, supplierRepository *repository.SupplierRepository, productsRepo *repository.ProductsRepository) *MenuParser {
	return &MenuParser{
		cfg:          cfg,
		supplierRepo: supplierRepository,
		productsRepo: productsRepo,
		logger:       logger,
	}
}

type MenuParserI interface {
	TimedParsing()
	productCheckUpdateCreate(parsedProduct *models.ParserProduct, oldRestID int, newRestID int)
	restaurantCheckUpdateCreate(restaurant *models.ParsedSupplier)
	getAllRestaurants() ([]models.ParsedSupplier, error)
	getProductsFromRestByID(id int) (*[]models.ParserProduct, error)
	transformRestaurantModel(parsedRestaurant *models.ParsedSupplier) *models.Supplier
	transformProductModel(parsedProduct *models.ParserProduct, id int) *models.Product
	deleteNonUpdatedRestaurants()
	deleteNonUpdatedProducts()
}

type MenuParser struct {
	cfg          *models.ParserConfig
	supplierRepo *repository.SupplierRepository
	productsRepo *repository.ProductsRepository
	logger       *logger.Logger
}

func (m MenuParser) TimedParsing() {
	for {
		m.Parse()
		time.Sleep(time.Duration(m.cfg.ParsingDelaySeconds) * time.Second)
	}
}

func (m MenuParser) Parse() {

	m.logger.InfoLog("starting API parsing with timeout (s): ", m.cfg.ParsingDelaySeconds)
	ctx := context.Background()

	restaurants, err := func(ctx context.Context) (*[]models.ParsedSupplier, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Duration(m.cfg.ParsingDelaySeconds/10)*time.Second)
		defer cancel()
		return m.getAllRestaurants(ctx)
	}(ctx)

	if err != nil {
		m.logger.ErrorLog("Fail to get restaurants", err)
		return
	}

	var wg = sync.WaitGroup{}
	for i := range *restaurants {
		wg.Add(1)
		go func(i int) {
			m.supplierCheckUpdateCreate(&(*restaurants)[i])
			wg.Done()
		}(i)

	}
	wg.Wait()
	m.deleteNonUpdatedRestaurants()

}

func (m MenuParser) getAllRestaurants(ctx context.Context) (*[]models.ParsedSupplier, error) {
	url := m.cfg.URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	readBody, err := ioutil.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	var responseBodyRestaurants models.ResponseBodyRestaurants

	err = json.Unmarshal(readBody, &responseBodyRestaurants)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, errors.New("timeout in getAllRestaurants")
	default:
		return &responseBodyRestaurants.Suppliers, nil
	}
}

func (m MenuParser) getProductsFromRestByID(ctx context.Context, id int) (*[]models.ParserProduct, error) {
	url := fmt.Sprintf(m.cfg.FormatString, m.cfg.URL, id)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	readBody, err := ioutil.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	var products models.ResponseBodyMenu

	if err = json.Unmarshal(readBody, &products); err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		return nil, errors.New("timeout in getProductsFromRestByID")
	default:
		return &products.Menu, nil
	}
}

func (m MenuParser) transformSupplierModel(parsedRestaurant *models.ParsedSupplier) *models.Supplier {
	supplier := models.Supplier{
		ImgURL:  parsedRestaurant.Image,
		ID:      parsedRestaurant.ID,
		Name:    parsedRestaurant.Name,
		Type:    parsedRestaurant.Type,
		Opening: parsedRestaurant.WorkingHours.Opening,
		Closing: parsedRestaurant.WorkingHours.Closing,
	}
	return &supplier
}

func (m MenuParser) transformProductModel(parsedProduct *models.ParserProduct, id int) *models.Product {
	product := models.Product{
		ImgURL:      parsedProduct.Image,
		Name:        parsedProduct.Name,
		Type:        parsedProduct.Type,
		Price:       parsedProduct.Price,
		IDSupplier:  id,
		Ingredients: fmt.Sprint(parsedProduct.Ingredients),
	}
	return &product
}

func (m MenuParser) supplierCheckUpdateCreate(supplier *models.ParsedSupplier) {
	var wg = sync.WaitGroup{}
	var lastSupplierID int
	dbSupplier, err := m.supplierRepo.GetByName(supplier.Name)
	if err != nil {
		m.logger.ErrorLog("fail to search supplier", err)
		return
	}
	if dbSupplier.ID != 0 {

		rowsAffected, err := m.supplierRepo.Update(dbSupplier)
		if err != nil {
			m.logger.ErrorLog("fail to edit supplier", err)
			return
		}
		m.logger.InfoLog("rows in supplier renewed", rowsAffected)
		lastSupplierID = dbSupplier.ID
	} else {

		lastRestID, err := m.supplierRepo.Create(m.transformSupplierModel(supplier))
		if err != nil {
			m.logger.ErrorLog("fail to create new supplier", err)
			return
		}
		m.logger.InfoLog("Saved supplier with ID ", lastRestID)
		lastSupplierID = lastRestID
	}
	ctx := context.Background()
	products, err := func(ctx context.Context, id int) (*[]models.ParserProduct, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Duration(m.cfg.ParsingDelaySeconds/10)*time.Second)
		defer cancel()
		return m.getProductsFromRestByID(ctx, supplier.ID)
	}(ctx, supplier.ID)
	if err != nil {
		m.logger.ErrorLog("fail to get product by rest ID", err)
		return
	}
	for i := range *products {
		wg.Add(1)
		go func(i int) {
			m.productCheckUpdateCreate(&(*products)[i], lastSupplierID)
			wg.Done()
		}(i)

	}
	wg.Wait()
	m.deleteNonUpdatedProducts(lastSupplierID)
}

func (m MenuParser) productCheckUpdateCreate(parsedProduct *models.ParserProduct, supplierID int) {

	oldProductID, err := m.productsRepo.SearchBySupIDAndName(supplierID, parsedProduct.Name)
	product := m.transformProductModel(parsedProduct, supplierID)

	if err != nil {
		m.logger.ErrorLog("fail to search existed product", err)
		return
	}
	if oldProductID != 0 {
		product.ID = oldProductID
		rowsAffected, err := m.productsRepo.Update(product)
		if err != nil {
			m.logger.ErrorLog("fail to edit existed product", err)
			return
		}
		m.logger.InfoLog(" rows in product renewed", rowsAffected)
	} else {
		lastProdID, err := m.productsRepo.Create(product)
		if err != nil {
			m.logger.ErrorLog("fail to create new product", err)
			return
		}
		m.logger.InfoLog("Product with ID, saved to DB", lastProdID)
	}
}

func (m MenuParser) deleteNonUpdatedRestaurants() {
	rowsAffected, err := m.supplierRepo.SoftDeleteNotUpdated(m.cfg.ParsingDelaySeconds)
	if err != nil {
		m.logger.ErrorLog("Failed to delete not updated restaurants", err)
	}
	if rowsAffected != 0 {
		m.logger.InfoLog("Rows was deleted from restaurants due to old update date", rowsAffected)
	}
	if rowsAffected == 0 {
		m.logger.InfoLog("All suppliers is up to date", rowsAffected, " rows was changed")
	}
}

func (m MenuParser) deleteNonUpdatedProducts(supplierID int) {
	rowsAffected, err := m.productsRepo.SoftDeleteNotUpdated(m.cfg.ParsingDelaySeconds, supplierID)
	if err != nil {
		m.logger.ErrorLog("Failed to delete not updated products", err)
	}
	if rowsAffected != 0 {
		m.logger.InfoLog("Rows was deleted in products due to old update date", rowsAffected)
	}
	if rowsAffected == 0 {
		m.logger.InfoLog("All products is up to date in supplier ID", supplierID)

	}
}
