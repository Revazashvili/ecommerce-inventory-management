# E-Commerce Inventory Management

A simple implementation of an e-commerce inventory management system using Go.

## Project Overview
This repository demonstrates a basic inventory management system for e-commerce applications built with Go. It provides core functionality to track products, manage stock levels, and handle basic inventory operations.

## Features

* Product management
    * Getting products and products count with name filtering
    * Adding and updating products
        * [application listens to kafka topics](consumers/product_consumer.go) where ProductAdded and ProductUpdated events are published and syncs them in database, later this products are used to add stock or reserve them.
* Stock management
    * Gettting stocks with filtering(product id and date range)
    * Adding stock for products
    * Reserveing stocks
        * Reserving happens under order number and multiple products are reserved
    * Unreserving products


## Technologies
* Golang (1.23.3 version)
* Postgres
* Kafka
* Chi for router
* Sqlc for working with postgres
* Swaggo for generating swagger docs


## Usage

### Running

```bash
go build -o inventory-api main.go && ./inventory-api
```
