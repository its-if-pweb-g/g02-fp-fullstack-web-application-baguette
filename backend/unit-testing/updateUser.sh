curl -X PUT http://localhost:8000/api/user \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjY3NTkxN2E0Y2JmNTBiZjZkMmU4ODE4MCIsImF1ZCI6ImJhZ3VldHRlIiwiZXhwIjoxNzM0MTUyNTM4LCJpYXQiOjE3MzM4OTMzMzgsImlzcyI6ImJhZ3VldHRlIiwibmJmIjoxNzMzODkzMzM4LCJyb2xlIjoidXNlciJ9.0Lx2XSY17poupZluQlP68w7wdzjsUJn9Iv9QAICPcN4" \
-d '{
  "email": "testing@gmail.com",
  "name": "testing",
  "address": "ja kamboja",
  "phone": "09832142"
}'
