Single page basic Online Store app; React and Redux with Golang Web Api.

# Web Site
- http://localhost:3000/web

# Admin Panel
- http://localhost:3000/admin

# Web API
- JWT Token Authentication
- Golang Gorm (mssql)

# React Component
  - redux-from
  - material UI
  - axios
  - react-router
  - bootpage
  - bootbox

# Prerequisites
  
  - Go 1.9.1
  - node.js 8 >
  - VS COde or Jetbrains GoLand or atom editor

### Database
```sh
Open SQL Server Management Studio > File > Open > File  select Store.sql and execute
Change OnlineStoreWebApi>database>database.go>InitDB function connString variable change Data Source your server name
```

### Installation Go Package
```sh
go get "github.com/jinzhu/gorm"
go get "github.com/gorilla/mux"
go get "github.com/gorilla/context"
go get "github.com/mitchellh/mapstructure"
go get "github.com/dgrijalva/jwt-go"

```

### Installation Node Module

Open command prompt

```sh
cd OnlineStoreReact folder location
npm install 
npm start
```
