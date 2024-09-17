# Fetch-Receipt-Processor
Coding Challenge Fetch Backend Developer

## Overview
The Receipt Processor Web Service is designed to process receipts, calculate points based on purchase details, and allow retrieval of receipt data via a RESTful API

# Run the application 
go run main.go

Server runs on port 8080
Access the api using http://localhost:8080


## API Endpoints

- **POST** `/receipts/process` - Submit a receipt for processing.
- **GET** `/receipts/{id}/points` - Retrieve points for a specific receipt by ID.
- **GET** `/receipts/{id}` - Retrieve receipt data for a specific receipt by ID.
