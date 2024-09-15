### Запускать проект с помощью 
```
docker-compose up --build
```
### Работает - ping 
```
curl http://localhost:8080/api/ping
```

### Работает - tender
```
curl -X POST "http://localhost:8080/api/tenders/new" \
-H "Content-Type: application/json" \
-d '{
      "name": "Test Tender",
      "description": "This is a test tender.",
      "serviceType": "Construction",
      "status": "Created",
      "organizationId": "org123",
      "creatorUsername": "user123"
    }'
```
```
curl -X GET "http://localhost:8080/api/tenders"
```

### Пока не работает - tender
```
curl -X GET "http://localhost:8080/api/tenders/{tenderId}"
```
```
curl -X PATCH "http://localhost:8080/api/tenders/{tenderId}/status" \
-H "Content-Type: application/json" \
-d '{
  "status": "Published"
}'
```

### Работает - bids
```
curl -X POST "http://localhost:8080/api/bids/new" \
-H "Content-Type: application/json" \
-d '{
  "name": "New Bid",
  "description": "This is a new bid.",
  "tenderId": "50137766-7f07-45ba-96e0-3ec252554a3e",
  "amount": 1500.00,
  "creatorUsername": "user123",
  "status": "Submitted"
}'
```



'новая реализация, с оглядкой на openapi.yml, работают api/ping, api/tenders/new, api/tenders, api/bids/new'