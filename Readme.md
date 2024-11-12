
# 🐳 Docker & API Steps Guide

## Step 1: Build the Docker Image

Run the following command to build the Docker image for the app:

```bash
docker build --no-cache -t gin-app .
```

## Step 2: Run the Docker Container

Start the container by running:

```bash
docker run -p 8080:8080 gin-app
```

## Step 3: Send a POST Request to Process Receipt

Use the `curl` command to send a POST request to the API and process a receipt:

```bash
curl -X POST http://localhost:8080/receipts/process   -H "Content-Type: application/json"   -d '{
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
  }'
```

This will return an **ID**.

## Step 4: Use the ID to Get Points

Once you have the ID from step 3, use it to execute this command and get the points associated with the receipt:

```bash
curl -X GET http://localhost:8080/receipts/{06f8fef0-58de-406b-964c-645d6d77af40}/points
```

Replace `{06f8fef0-58de-406b-964c-645d6d77af40}` with the actual ID returned from step 3.

Enjoy! 🎉
