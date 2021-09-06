package parser

import (
	"awesomeProject/models"
	"awesomeProject/repository"
	"encoding/json"
	"fmt"
	"github.com/DTB4/logger/v2"
	"io/ioutil"
	"net/http"
	"time"
)

func NewMenuParser(url, format string, logger *logger.Logger, restaurantRepo *repository.SupplierRepository, productsRepo *repository.ProductsRepository) *MenuParser {
	return &MenuParser{
		restaurantRepo: restaurantRepo,
		productsRepo:   productsRepo,
		logger:         logger,
		url:            url,
		format:         format,
	}
}

type MenuParserI interface {
	TimedParsing(frequencySeconds int)
	productWork(products *[]models.ParserProduct, restID int)
	restaurantCheckUpdateCreate(restaurants *[]models.ParserRestaurant)
}

type MenuParser struct {
	restaurantRepo *repository.SupplierRepository
	productsRepo   *repository.ProductsRepository
	logger         *logger.Logger
	url            string
	format         string
}

func (m MenuParser) TimedParsing(frequencySeconds int) {
	for {
		time.Sleep(time.Duration(frequencySeconds) * time.Second)
		restaurants, err := m.getAllRestaurants()
		if err != nil {
			m.logger.ErrorLog("Fail to get restaurants", err)
		}

		for i := range restaurants {
			go m.restaurantCheckUpdateCreate(&restaurants[i])
		}
	}
}

func (m MenuParser) getAllRestaurants() ([]models.ParserRestaurant, error) {
	url := m.url

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	readedBody, err := ioutil.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	var responseBodyRestaurants models.ResponseBodyRestaurants

	err = json.Unmarshal(readedBody, &responseBodyRestaurants)
	if err != nil {
		return nil, err
	}

	return responseBodyRestaurants.Restaurants, nil
}

func (m MenuParser) getProductsFromRestByID(id int) (*[]models.ParserProduct, error) {
	url := fmt.Sprintf(m.format, m.url, id)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	readedBody, err := ioutil.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		return nil, err
	}

	var products models.ResponseBodyMenu

	if err := json.Unmarshal(readedBody, &products); err != nil {
		return nil, err
	}

	return &products.Menu, nil
}

func transformRestaurantModel(parsedRestaurant *models.ParserRestaurant) *models.Supplier {
	supplier := models.Supplier{
		ImgURL: parsedRestaurant.Image,
		ID:     parsedRestaurant.ID,
		Name:   parsedRestaurant.Name,
	}
	return &supplier
}

func transformProductModel(parsedProduct *models.ParserProduct, id int) *models.Product {
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

	dbSupplier, err := m.restaurantRepo.SearchByID(restaurant.ID)
	if err != nil {
		m.logger.ErrorLog("fail to search supplier", err)
		return
	}
	if dbSupplier {
		result, err := m.restaurantRepo.Update(transformRestaurantModel(restaurant))
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

	} else {
		result, err := m.restaurantRepo.Create(transformRestaurantModel(restaurant))
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
	}
	products, err := m.getProductsFromRestByID(restaurant.ID)
	for i := range *products {
		go m.productCheckUpdateCreate(&(*products)[i], restaurant.ID)
	}

}

func (m MenuParser) productCheckUpdateCreate(parsedproduct *models.ParserProduct, restID int) {

	product := transformProductModel(parsedproduct, restID)
	productID, err := m.productsRepo.SearchBySupIDAndName(product.IDSupplier, product.Name)
	if err != nil {
		m.logger.ErrorLog("fail to search existed product", err)
		return
	}
	if productID != 0 {
		product.ID = productID
		result, err := m.productsRepo.Update(product)
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
	} else {
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

}