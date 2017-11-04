package controllers

import (
	business "../business"
	entity "../entities"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
)

type CategoriesController struct {}

func (c CategoriesController) GetAll(w http.ResponseWriter,r *http.Request) {
	categories,err:=business.Category{}.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	output,errJson:=json.Marshal(categories)
	if errJson != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(errJson)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,string(output))

}

func (c CategoriesController) Get(w http.ResponseWriter,r *http.Request)  {

	urlParams:=mux.Vars(r)
	categoryId,_:=strconv.Atoi(urlParams["categoryId"])
	category,err:=business.Category{}.Get(categoryId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	output,outputErr:=json.Marshal(category)
	if outputErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(outputErr)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,string(output))
}

func (c CategoriesController) Add(w http.ResponseWriter,r *http.Request)  {
	var category entity.Category
	_ = json.NewDecoder(r.Body).Decode(&category)

	addedCategory,err:=business.Category{}.Add(&category);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	output,errOutput:=json.Marshal(addedCategory);
	if errOutput != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(errOutput)
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,string(output))
}

func (c CategoriesController) Update(w http.ResponseWriter,r *http.Request)  {
	var category entity.Category

	_ =json.NewDecoder(r.Body).Decode(&category)
	updatedCategory,err:=business.Category{}.Update(&category)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	fmt.Println("updatedCategory")
	fmt.Println(updatedCategory)
	fmt.Println("updatedCategory")
	output,errOutput :=json.Marshal(updatedCategory);
	if errOutput != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(errOutput)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,string(output))
}
func (c CategoriesController) Delete(w http.ResponseWriter,r *http.Request)  {

	urlParams :=mux.Vars(r)
	categoryId,_:=strconv.Atoi(urlParams["categoryId"])

	category:=entity.Category{Id:categoryId}
	err:=business.Category{}.Delete(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	w.WriteHeader(http.StatusNoContent)
	//fmt.Fprintf(w,"")
}