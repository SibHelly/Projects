package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"products/pkg/models"
	"products/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var NewProduct models.Product

func GetProduct(w http.ResponseWriter, r *http.Request) {
	newProducts := models.GetAllProducts()
	res, _ := json.Marshal(newProducts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	productDetails, _ := models.GetProductById(ID)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	CreateProduct := &models.Product{}
	utils.ParseBody(r, CreateProduct)
	b := CreateProduct.CreateProduct()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	product := models.DeleteProduct(ID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct = &models.Product{}
	utils.ParseBody(r, updateProduct)
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	productDetails, db := models.GetProductById(ID)
	if updateProduct.Name != "" {
		productDetails.Name = updateProduct.Name
	}
	if updateProduct.Category != "" {
		productDetails.Category = updateProduct.Category
	}
	if updateProduct.Count != 0 {
		productDetails.Count = updateProduct.Count
	}
	if updateProduct.Price != 0 {
		productDetails.Price = updateProduct.Price
	}
	db.Save(&productDetails)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
