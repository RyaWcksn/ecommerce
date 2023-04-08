# ecommerce

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white) [![codecov.io](https://codecov.io/github/RyaWcksn/ecommerce/coverage.svg?branch=master)](https://codecov.io/github/RyaWcksn/ecommerce?branch=master)

Ecommerce for seller and buyer

> This usecase will be as Gunpla store who sell gundam model kits


<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [ecommerce](#ecommerce)
    - [How to install](#how-to-install)
    - [Common Errors](#common-errors)
    - [Endpoints](#endpoints)
    - [Requests and Response](#requests-and-response)
        - [Login](#login)
        - [Seller Create Product](#seller-create-product)
        - [Seller Get Products](#seller-get-products)
        - [Seller Get Orders](#seller-get-orders)
        - [Seller Accept Order](#seller-accept-order)
        - [Buyer Create Order](#buyer-create-order)
        - [Buyer Get Products](#buyer-get-products)
        - [Buyer Get Orders](#buyer-get-orders)

<!-- markdown-toc end -->

## How to install

```bash
git clone https://github.com/RyaWcksn/ecommerce
cd ecommerce
docker-compose --build -d

./migrate.sh
```

- Buyer account
> user@mail.com
> password123

- Seller account
> seller@mail.com
> password123


## Common Errors

| Error Code | Error Message             | Description                                                                                                          |
|------------|---------------------------|----------------------------------------------------------------------------------------------------------------------|
| 401        | Unauthorized Access       | The user is not authorized to access the requested resource. This may be due to missing or invalid authentication.   |
| 403        | Forbidden                 | The user is authenticated but does not have sufficient permissions to access the requested resource.                |
| 404        | Not Found                 | The requested resource could not be found on the server. This may be due to a broken link or incorrect URL.           |
| 500        | Internal Server Error     | An unexpected error occurred on the server. This may be due to a bug in the application or a configuration issue.    |
| 503        | Service Unavailable       | The requested service is unavailable at the moment. This may be due to server maintenance or high traffic.           |

## Endpoints

| Endpoint                | Description                           |
|-------------------------|---------------------------------------|
| /api/v1/login           | Endpoint for user login               |
| /api/v1/seller/create   | Endpoint for seller to create a product|
| /api/v1/seller/products | Endpoint for seller to list their products|
| /api/v1/seller/orders   | Endpoint for seller to list their orders|
| /api/v1/seller/order/accept | Endpoint for seller to accept an order|
| /api/v1/order           | Endpoint for buyer to create an order  |
| /api/v1/orders          | Endpoint for buyer to list their orders|
| /api/v1/products        | Endpoint for buyer to list products    |


## Requests and Response

### Login

> Login endpoint have 2 roles, _Buyer_ and _Seller_, each role have different specs

- Header
Content-Type: Application/json
Accept: Application/json

- Request
```json
{
    "email": "user@mail.com",
    "password": "password123",
    "role": "buyer"
}
```

- Response
```json
{
    "code": 201,
    "message": "ok",
    "responseTime": "20230408102715",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InVzZXJAbWFpbC5jb20iLCJleHAiOjE2ODA5MjYyMzUsImlkIjoxLCJyb2xlIjoiYnV5ZXIifQ.0tXezwqgzCcU073i7e0xo5_02Lte65Z7kPMxtkjGvEM"
}
```

### Seller Create Product

- Header
Content-Type: Application/json
Accept: Application/json
Authorization: Bearer {JWT-TOKEN}

- Request
```json
{
    "name": "HG Zaku I Origin",
    "description": "HG Zaku I Origin from Kidou Senshi Gundam Origin",
    "price": "18000"
}
```

- Response
```json
{
    "code": 201,
    "message": "ok",
    "responseTime": "20230408103012",
    "product": {
        "name": "HG Zaku I Origin",
        "description": "HG Zaku I Origin from Kidou Senshi Gundam Origin",
        "price": "18000"
    }
}
```

### Seller Get Products

- Header
Content-Type: Application/json
Accept: Application/json
Authorization: Bearer {JWT-TOKEN}

- Response
```json
{
    "code": 201,
    "message": "ok",
    "responseTime": "20230408103214",
    "products": [
        {
            "id": 1,
            "productName": "HG Dynames Gundam",
            "description": "HG Dynames Gundam from Kidou Senshi Gundam 00",
            "price": "18000.00",
            "seller": 1
        },
        {
            "id": 2,
            "productName": "HG Exia Gundam",
            "description": "HG Exia Gundam from Kidou Senshi Gundam 00",
            "price": "18000.00",
            "seller": 1
        },
        {
            "id": 3,
            "productName": "HG Kyrios Gundam",
            "description": "HG Kyrios Gundam from Kidou Senshi Gundam 00",
            "price": "18000.00",
            "seller": 1
        },
        {
            "id": 4,
            "productName": "HG Virtue Gundam",
            "description": "HG Virtue Gundam from Kidou Senshi Gundam 00",
            "price": "18000.00",
            "seller": 1
        },
        {
            "id": 5,
            "productName": "HG Zaku I Origin",
            "description": "HG Zaku I Origin from Kidou Senshi Gundam Origin",
            "price": "18000.00",
            "seller": 1
        }
    ]
}
```

### Seller Get Orders

- Header
Content-Type: Application/json
Accept: Application/json
Authorization: Bearer {JWT-TOKEN}

- Response
```json
{
    "code": 200,
    "message": "ok",
    "responseTime": "20230408103403",
    "orders": [
        {
            "Id": 1,
            "Buyer": 1,
            "Seller": 1,
            "DeliverySource": "123 Main St",
            "DeliveryDestination": "123 Main St",
            "Items": ", HG Dynames Gundam",
            "Quantity": 1,
            "Price": "0",
            "TotalPrice": "0.00",
            "Status": {
                "message": "Seller is accepted the order",
                "status": "ACCEPTED"
            }
        }
    ]
}
```

### Seller Accept Order

- Header
Content-Type: Application/json
Accept: Application/json
Authorization: Bearer {JWT-TOKEN}

- Request
```json
{
    "orderId": 2
}
```

- Response
```json
{
    "code": 200,
    "message": "ok",
    "responseTime": "20230408103804",
    "order": {
        "Id": 2,
        "Buyer": 1,
        "Seller": 1,
        "DeliverySource": "123 Main St",
        "DeliveryDestination": "123 Main St",
        "Items": ", HG Kyrios Gundam",
        "Quantity": 1,
        "Price": "0",
        "TotalPrice": "0.00",
        "Status": {
            "message": "Seller is accepted the order",
            "status": "ACCEPTED"
        }
    }
}
```

### Buyer Create Order

- Header
Content-Type: Application/json
Accept: Application/json
Authorization: Bearer {JWT-TOKEN}

- Request
```json
{
    "productId": [3],
    "sellerId": 1
}
```

- Response
```json
{
    "code": 201,
    "message": "ok",
    "responseTime": "20230408103743",
    "order": {
        "message": "Waiting Seller to accept order",
        "status": "PENDING"
    }
}
```

### Buyer Get Products

- Header
Content-Type: Application/json
Accept: Application/json
Authorization: Bearer {JWT-TOKEN}

- Response
```json
{
    "code": 200,
    "message": "ok",
    "responseTime": "20230408111515",
    "products": [
        {
            "id": 1,
            "productName": "HG Dynames Gundam",
            "description": "HG Dynames Gundam from Kidou Senshi Gundam 00",
            "price": 18000,
            "seller": 1
        },
        {
            "id": 2,
            "productName": "HG Exia Gundam",
            "description": "HG Exia Gundam from Kidou Senshi Gundam 00",
            "price": 18000,
            "seller": 1
        },
        {
            "id": 3,
            "productName": "HG Kyrios Gundam",
            "description": "HG Kyrios Gundam from Kidou Senshi Gundam 00",
            "price": 18000,
            "seller": 1
        },
        {
            "id": 4,
            "productName": "HG Virtue Gundam",
            "description": "HG Virtue Gundam from Kidou Senshi Gundam 00",
            "price": 18000,
            "seller": 1
        },
        {
            "id": 5,
            "productName": "HG Zaku I Origin",
            "description": "HG Zaku I Origin from Kidou Senshi Gundam Origin",
            "price": 18000,
            "seller": 1
        },
        {
            "id": 6,
            "productName": "HG Zaku II Origin",
            "description": "HG Zaku II Origin from Kidou Senshi Gundam Origin",
            "price": 180000,
            "seller": 1
        }
    ]
}
```


### Buyer Get Orders

- Header
Content-Type: Application/json
Accept: Application/json
Authorization: Bearer {JWT-TOKEN}

- Response
```json
{
    "code": 200,
    "message": "ok",
    "responseTime": "20230408111414",
    "orders": [
        {
            "Id": 1,
            "Buyer": 1,
            "Seller": 1,
            "DeliverySource": "123 Main St",
            "DeliveryDestination": "123 Main St",
            "Items": ", HG Kyrios Gundam",
            "Quantity": 1,
            "Price": 18000,
            "TotalPrice": 18000,
            "Status": {
                "message": "Waiting Seller to accept order",
                "status": "PENDING"
            }
        },
        {
            "Id": 2,
            "Buyer": 1,
            "Seller": 1,
            "DeliverySource": "123 Main St",
            "DeliveryDestination": "123 Main St",
            "Items": ", HG Zaku II Origin",
            "Quantity": 1,
            "Price": 180000,
            "TotalPrice": 180000,
            "Status": {
                "message": "Waiting Seller to accept order",
                "status": "PENDING"
            }
        },
        {
            "Id": 3,
            "Buyer": 1,
            "Seller": 1,
            "DeliverySource": "123 Main St",
            "DeliveryDestination": "123 Main St",
            "Items": ", HG Dynames Gundam, HG Exia Gundam, HG Kyrios Gundam, HG Virtue Gundam, HG Zaku I Origin, HG Zaku II Origin",
            "Quantity": 6,
            "Price": 270000,
            "TotalPrice": 270000,
            "Status": {
                "message": "Waiting Seller to accept order",
                "status": "PENDING"
            }
        }
    ]
}
```
