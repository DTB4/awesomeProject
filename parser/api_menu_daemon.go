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

func NewMenuParser(url, format string, logger *logger.Logger, restaurantRepo *repository.SupplierRepository, productsRepo *repository.ProductsRepository, frequencySeconds int) *MenuParser {
	return &MenuParser{
		restaurantRepo:   restaurantRepo,
		productsRepo:     productsRepo,
		logger:           logger,
		url:              url,
		format:           format,
		frequencySeconds: frequencySeconds,
	}
}

type MenuParserI interface {
	TimedParsing()
	productCheckUpdateCreate(parsedProduct *models.ParserProduct, oldRestID int, newRestID int)
	restaurantCheckUpdateCreate(restaurant *models.ParserRestaurant)
	getAllRestaurants() ([]models.ParserRestaurant, error)
	getProductsFromRestByID(id int) (*[]models.ParserProduct, error)
	transformRestaurantModel(parsedRestaurant *models.ParserRestaurant) *models.Supplier
	transformProductModel(parsedProduct *models.ParserProduct, id int) *models.Product
	deleteNonUpdatedRestaurants()
	deleteNonUpdatedProducts()
}

type MenuParser struct {
	restaurantRepo   *repository.SupplierRepository
	productsRepo     *repository.ProductsRepository
	logger           *logger.Logger
	url              string
	format           string
	frequencySeconds int
}

func (m MenuParser) TimedParsing() {
	for {
		m.Parse()
		time.Sleep(time.Duration(m.frequencySeconds) * time.Second)
	}
}

func (m MenuParser) Parse() {

	m.logger.InfoLog("starting API parsing with timeout (s): ", m.frequencySeconds)
	ctx := context.Background()

	restaurants, err := func(ctx context.Context) (*[]models.ParserRestaurant, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Duration(m.frequencySeconds/10)*time.Second)
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
		go m.restaurantCheckUpdateCreate(&(*restaurants)[i])
		wg.Done()
	}
	wg.Wait()
	m.deleteNonUpdatedRestaurants()

}

func (m MenuParser) getAllRestaurants(ctx context.Context) (*[]models.ParserRestaurant, error) {
	url := m.url
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
		return &responseBodyRestaurants.Restaurants, nil
	}
}

func (m MenuParser) getProductsFromRestByID(ctx context.Context, id int) (*[]models.ParserProduct, error) {
	url := fmt.Sprintf(m.format, m.url, id)

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

func (m MenuParser) transformRestaurantModel(parsedRestaurant *models.ParserRestaurant) *models.Supplier {
	supplier := models.Supplier{
		ImgURL: parsedRestaurant.Image,
		ID:     parsedRestaurant.ID,
		Name:   parsedRestaurant.Name,
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

func (m MenuParser) restaurantCheckUpdateCreate(restaurant *models.ParserRestaurant) {
	var wg = sync.WaitGroup{}
	dbSupplier, err := m.restaurantRepo.GetByName(restaurant.Name)
	if err != nil {
		m.logger.ErrorLog("fail to search supplier", err)
		return
	}
	if dbSupplier != nil {
		result, err := m.restaurantRepo.SoftDelete(dbSupplier.ID)
		if err != nil {
			m.logger.ErrorLog("fail to edit supplier", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			m.logger.ErrorLog("fail to get rowsAffected from result", err)
			return
		}
		m.logger.InfoLog("rows in restaurant renewed", rowsAffected)

	}
	result, err := m.restaurantRepo.Create(m.transformRestaurantModel(restaurant))
	if err != nil {
		m.logger.ErrorLog("fail to create new supplier", err)
		return
	}
	lastRestID, err := result.LastInsertId()
	if err != nil {
		m.logger.ErrorLog("fail to get lastInsertID from result", err)
		return
	}
	m.logger.InfoLog("Saved restaurant with ID ", lastRestID)
	ctx := context.Background()
	products, err := func(ctx context.Context, id int) (*[]models.ParserProduct, error) {
		ctx, cancel := context.WithTimeout(ctx, time.Duration(m.frequencySeconds/10)*time.Second)
		defer cancel()
		return m.getProductsFromRestByID(ctx, restaurant.ID)
	}(ctx, restaurant.ID)
	if err != nil {
		m.logger.ErrorLog("fail to get product by rest ID", err)
		return
	}
	for i := range *products {
		wg.Add(1)
		go m.productCheckUpdateCreate(&(*products)[i], dbSupplier.ID, int(lastRestID))
		wg.Done()
	}
	wg.Wait()
	m.deleteNonUpdatedProducts()
}

func (m MenuParser) productCheckUpdateCreate(parsedProduct *models.ParserProduct, oldRestID int, newRestID int) {

	product := m.transformProductModel(parsedProduct, newRestID)
	oldProductID, err := m.productsRepo.SearchBySupIDAndName(oldRestID, parsedProduct.Name)
	if err != nil {
		m.logger.ErrorLog("fail to search existed product", err)
		return
	}
	if oldProductID != 0 {
		result, err := m.productsRepo.SoftDelete(oldProductID)
		if err != nil {
			m.logger.ErrorLog("fail to edit existed product", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			m.logger.ErrorLog("fail to get rowsAffected from result", err)
			return
		}
		m.logger.InfoLog(" rows in product renewed", rowsAffected)
	}
	result, err := m.productsRepo.Create(product)
	if err != nil {
		m.logger.ErrorLog("fail to create new product", err)
		return
	}
	lastProdID, err := result.LastInsertId()
	if err != nil {
		m.logger.ErrorLog("fail to get lastInsertID from result", err)
		return
	}
	m.logger.InfoLog("Product with ID, saved to DB", lastProdID)

}

func (m MenuParser) deleteNonUpdatedRestaurants() {
	result, err := m.restaurantRepo.SoftDeleteNotUpdated(m.frequencySeconds)
	if err != nil {
		m.logger.ErrorLog("Failed to delete not updated restaurants", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		m.logger.ErrorLog("failed to get rows affected by soft delete not updated restaurants", err)
	}
	if rowsAffected != 0 {
		m.logger.InfoLog("Rows was deleted from restaurants due to old update date", rowsAffected)

	}
}

func (m MenuParser) deleteNonUpdatedProducts() {
	result, err := m.productsRepo.SoftDeleteNotUpdated(m.frequencySeconds)
	if err != nil {
		m.logger.ErrorLog("Failed to delete not updated products", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		m.logger.ErrorLog("failed to get rows affected by soft delete not updated products", err)
	}
	if rowsAffected != 0 {
		m.logger.InfoLog("Rows was deleted in products due to old update date", rowsAffected)

	}
}
