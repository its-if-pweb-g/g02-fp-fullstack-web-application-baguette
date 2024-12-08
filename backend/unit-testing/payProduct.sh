curl -X POST http://localhost:8000/api/user/pay \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjY3NTRkMzVhZTQ4NTI5OWRmYzIyZDE2YSIsImF1ZCI6ImJhZ3VldHRlIiwiZXhwIjoxNzMzOTAzOTAxLCJpYXQiOjE3MzM2NDQ3MDEsImlzcyI6ImJhZ3VldHRlIiwibmJmIjoxNzMzNjQ0NzAxLCJyb2xlIjoidXNlciJ9.2z6WEIooSOtTnDq0K8vk4FQO3ko52jyTHQhKHl563os" \
-d '{
  "product_id": "12345",
  "name": "Product Name",
  "price": 100000,
  "quantity": 2
}'

