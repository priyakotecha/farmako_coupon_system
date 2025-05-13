# farmako_coupon_system (Golang)

## Overview
This project is a backend coupon management system for a medicine ordering platform built in Golang using a clean architecture (Controller, Service, Repository layers).

## Features
- Admin can define coupon metadata
- Users can validate applicable coupons for their cart
- Constraints enforced: expiry, order total, category/medicine match
- RESTful APIs with Gin
- SQLite/PostgreSQL support
- Dockerized

## API Endpoints

### POST `/coupons/applicable`
**Request:**
```json
{
  "cart_items": [{ "id": "med_123", "category": "painkiller" }],
  "order_total": 700,
  "timestamp": "2025-05-05T15:00:00Z"
}
```

**Response:**
```json
{
  "applicable_coupons": [{ "coupon_code": "SAVE20", "discount_value": 20 }]
}
```

### POST `/coupons/validate`
**Request:**
```json
{
  "coupon_code": "SAVE20",
  "cart_items": [...],
  "order_total": 700,
  "timestamp": "2025-05-05T15:00:00Z"
}
```

**Response:**
```json
{
  "is_valid": true,
  "discount": {
    "items_discount": 50,
    "charges_discount": 20
  },
  "message": "coupon applied successfully"
}
```

## Setup
### Prerequisites
- Docker + Docker Compose

### Steps
```bash
git clone <repo-url>
cd coupon-system
docker-compose up --build
```