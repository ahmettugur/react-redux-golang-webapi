package main

import (
	controller "./controllers"
	entity "./entities"
	business "./business"
	utilities "./utilities"
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
	var user entity.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	validateUser,_ := business.User{}.ValidateUser(user.Email,user.Password)

	if validateUser== nil {
		json.NewEncoder(w).Encode(utilities.Exception{Message: "User Not found"})
		return
	}
	//c:=ClaimUser{UserId:validateUser.UserId,FullName:validateUser.FullName,Password:validateUser.Email}
	roles,err := business.User{}.GetUserRoles(validateUser.UserId)
	if err != nil{
		fmt.Println(err)
	}
	validateUser.Roles = *roles
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":validateUser.UserId,
		"FullName":validateUser.FullName,
		"Password":validateUser.Password,
		"Email":validateUser.Email,
		"Roles":validateUser.Roles,
		"exp":time.Now().Add(time.Hour * 10).Unix(),
	})

	fmt.Println("validateUser: ",*validateUser)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("tokenString: ",tokenString )
	json.NewEncoder(w).Encode(utilities.JwtToken{Token: tokenString})
}

//func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
//	params := req.URL.Query()
//	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("There was an error")
//		}
//		return []byte("secret"), nil
//	})
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		var user entity.User
//		mapstructure.Decode(claims, &user)
//		json.NewEncoder(w).Encode(user)
//	} else {
//		json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
//	}
//}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(utilities.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(utilities.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(utilities.Exception{Message: "An authorization header is required"})
		}
	})
}

func TestEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user entity.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	mapstructure.Decode(decoded, &user)

	json.NewEncoder(w).Encode(user)
}

type CORSMiddleware struct {
	http.Handler
}

func (cm CORSMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}
	cm.Handler.ServeHTTP(w, r)
}

func main() {

	//p,err:=business.ProductWithCategory{}.GetAll();
	//if err != nil {
	//	fmt.Println("hatalı")
	//}
	//
	//fmt.Println(p)
	apiRoot := "/api"

	gorillaRoute:=mux.NewRouter()

	gorillaRoute.HandleFunc("/token", CreateTokenEndpoint).Methods("POST")
	gorillaRoute.HandleFunc(apiRoot+"/categories",controller.CategoriesController{}.GetAll).Methods("GET")
	gorillaRoute.HandleFunc(apiRoot+"/categories/{categoryId}",controller.CategoriesController{}.Get).Methods("GET")
	gorillaRoute.HandleFunc(apiRoot+"/products/{categoryId}/{page}",controller.ProductsController{}.GetAllProduct).Methods("GET")
	gorillaRoute.HandleFunc(apiRoot+"/products/{productId}",controller.ProductsController{}.GetProduct).Methods("GET")

	gorillaRoute.HandleFunc(apiRoot+"/admin/products/{page}",ValidateMiddleware(controller.ProductsController{}.GetAdminProduct)).Methods("GET")

	gorillaRoute.HandleFunc(apiRoot+"/admin/categories",controller.CategoriesController{}.Add).Methods("POST")
	gorillaRoute.HandleFunc(apiRoot+"/admin/categories",ValidateMiddleware(controller.CategoriesController{}.Update)).Methods("PUT")
	gorillaRoute.HandleFunc(apiRoot+"/admin/categories/{categoryId}",ValidateMiddleware(controller.CategoriesController{}.Delete)).Methods("DELETE")

	gorillaRoute.HandleFunc(apiRoot+"/admin/products",ValidateMiddleware(controller.ProductsController{}.Add)).Methods("POST")
	gorillaRoute.HandleFunc(apiRoot+"/admin/products",ValidateMiddleware(controller.ProductsController{}.Update)).Methods("PUT")
	gorillaRoute.HandleFunc(apiRoot+"/admin/products/{productId}",ValidateMiddleware(controller.ProductsController{}.Delete)).Methods("DELETE")



	var h http.Handler = CORSMiddleware{gorillaRoute}
	http.ListenAndServe(":3500",h)
	//http.Handle("/",gorillaRoute)
	//http.ListenAndServe(":3500",handlers.CORS()(gorillaRoute))

/*
	fmt.Println("Starting the application...")
	router.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")
	router.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
	router.HandleFunc("/test", ValidateMiddleware(TestEndpoint)).Methods("GET")
*/
}
