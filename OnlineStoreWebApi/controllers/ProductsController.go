package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	business "onlinestorewebapi/business"
	entity "onlinestorewebapi/entities"
	model "onlinestorewebapi/models"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductsController struct {
}

func (p ProductsController) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	page, errPage := strconv.Atoi(urlParams["page"])
	if errPage != nil {
		page = 1
	}
	categoryId, errCategory := strconv.Atoi(urlParams["categoryId"])
	if errCategory != nil {
		categoryId = 0
	}
	pageSize := 12
	fmt.Println("Page: ", page)
	fmt.Println("Category: ", categoryId)
	products, err := business.Product{}.GatAll(categoryId, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	var productResponse model.ProductResponse
	productResponse.Products = *paginateProduct(*products, (page-1)*pageSize, pageSize)
	productResponse.CurrentCategory = categoryId
	productResponse.CurrentPage = page
	productResponse.PageSize = pageSize
	productResponse.PageCount = int(math.Ceil(float64(len(*products)) / float64(pageSize)))

	output, errJson := json.Marshal(productResponse)
	if errJson != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(errJson)
	}
	fmt.Println("products")
	fmt.Println(string(output))
	fmt.Println("products")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(output))
}

func (p ProductsController) GetProduct(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	productId, _ := strconv.Atoi(urlParams["productId"])
	product, err := business.Product{}.Get(productId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	output, errJson := json.Marshal(product)
	if errJson != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(errJson)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(output))
}

func (p ProductsController) GetAdminProduct(w http.ResponseWriter, r *http.Request) {
	pageSize := 12

	urlParams := mux.Vars(r)

	page, _ := strconv.Atoi(urlParams["page"])

	products, _ := business.ProductWithCategory{}.GetAll()
	//fmt.Println(products)
	var productWithCategoryResponse model.ProductWithCategoryResponse
	productWithCategoryResponse.Products = *paginateProductWithCategory(*products, (page-1)*pageSize, pageSize)
	productWithCategoryResponse.PageSize = pageSize
	productWithCategoryResponse.PageCount = int(math.Ceil(float64(len(*products)) / float64(pageSize)))

	//fmt.Println(products)

	output, errJson := json.Marshal(productWithCategoryResponse)
	if errJson != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(errJson)
	}
	fmt.Println("products")
	fmt.Println(string(output))
	fmt.Println("products")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(output))
}

func (p ProductsController) Add(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	addedProduct, _ := business.Product{}.Add(&product)

	output, outputErr := json.Marshal(addedProduct)
	if outputErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(outputErr)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(output))
}

func (p ProductsController) Update(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	updatedProduct, _ := business.Product{}.Update(&product)

	output, outputErr := json.Marshal(updatedProduct)

	if outputErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(outputErr)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(output))
}

func (p ProductsController) Delete(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	productId, _ := strconv.Atoi(urlParams["productId"])

	product := entity.Product{Id: productId}

	err := business.Product{}.Delete(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusNoContent)
}

func paginateProduct(x []entity.Product, skip int, size int) *[]entity.Product {
	limit := func() int {
		if skip+size > len(x) {
			return len(x)
		} else {
			return skip + size
		}

	}

	start := func() int {
		if skip > len(x) {
			return len(x)
		} else {
			return skip
		}

	}
	data := x[start():limit()]
	return &data
}

func paginateProductWithCategory(x []entity.ProductWithCategory, skip int, size int) *[]entity.ProductWithCategory {
	limit := func() int {
		if skip+size > len(x) {
			return len(x)
		} else {
			return skip + size
		}

	}

	start := func() int {
		if skip > len(x) {
			return len(x)
		} else {
			return skip
		}

	}
	data := x[start():limit()]
	return &data
}
