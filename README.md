# Shopping Cart Api

This api is designed as an example of a e-commerce basket api

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of Go.

## Installing Shopping Cart Api

To clone Shopping Cart Api to your repo follow these steps:

```
git clone https://github.com/Property-Finder-Golang-Bootcamp/bitirme-projesi-aydinsercan.git
```

## Repository Overview

```bash
├── README.md
├── main.go
├── config.yaml
├── go.mod
├── go.sum
├── controller
│   ├── carts
│   │   ├── v1/routes.go
│   │   ├── cart.go
│   ├── products
│   │   ├── v1/routes.go
│   │   └── product.go
│   └── routes.go
├── helper
│   └── helper.go
├── middleware
│   ├── auth.go
│   └── request.go
├── model
│   ├── cart
│   │   ├── cart.go
│   ├── product
│   │   └── product.go
│   ├── migrate.go
│   └── db.go

```

## Getting Started

Before running the project you must set up your database credentials in the config.yaml file. 

## Using Shopping Cart Api

The app provides the following endpoints:

#### `GET /api/v1/products` : List all the products 

#### `POST /api/v1/carts` : Add products to cart

#### `GET /api/v1/carts` : List the products user have added to their cart, total price and VAT of the cart. 

#### `DELETE /api/v1/carts/:id` : Delete product from cart

#### `GET /api/v1/carts/orders` : Show order information


## Tool set

- Go
- Echo
- Mysql
- Postman
- Swagger

## Contributors

Thanks to the following people who have contributed to this project:

- [@aydinsercan](https://github.com/aydinsercan) 📖

## License

This project uses the following license: [MIT](https://opensource.org/licenses/MIT).
