Single page basic Online Store app; React and Redux with ASP.NET Web Api.

# Web Site
- http://localhost:3000/web

# Admin Panel
- http://localhost:3000/admin

# Web API
- Generic Repostory Pattern
- Aspect Oriented Programming with Postsharp (4.2.17)
  - AuthorizationAspects
  - CacheAsepcts
  - ExceptionAspects
  - LogAspects
  - ValidationAspects
- FluentValidation
- log4net with logging
- Bearer Token Authentication
- Ninject IOC
- AutoMapper

# React Component
  - redux-from
  - material UI
  - axios
  - react-router
  - bootpage
  - bootbox

# Prerequisites
  
  - .NetFramework 4.7 (VS 2015/2017)
  - Postsharp (4.2.17)
  - node.js 8 >

### Database and Postsharp
```sh
Download Postsahrp (4.2.17). Install your compoter
Open SQL Server Management Studio > File > Open > File  select Store.sql and execute
Change OnlineStore.WebApi > Web.config file connection string Data Source your server name
Change OnlineStore.WebApi > log4net.config file connection string Data Source your server name
```

### Installation Node Module

Open command prompt

```sh
cd OnlineStoreReact folder location
npm install 
npm start
```
