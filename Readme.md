Step 1: docker build --no-cache -t gin-app .
Step 2:  docker run -p 8080:8080 gin-app 
Step 3 : curl -X POST http://localhost:8080/receipts/process \
  -H "Content-Type: application/json" \
  -d '{
    "retailer": "M&M Corner Market",
    "purchaseDate": "2022-03-20",
    "purchaseTime": "14:33",
    "items": [
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      },
      {
        "shortDescription": "Gatorade",
        "price": "2.25"
      }
    ],
    "total": "9"
  }

  This will return an ID

  Step 4:
  Use the id returned by step 3 to execute this command.
   
  curl -X GET http://localhost:8080/receipts/{06f8fef0-58de-406b-964c-645d6d77af40}/points
  