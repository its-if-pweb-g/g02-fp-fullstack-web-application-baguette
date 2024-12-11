curl -X POST "http://localhost:8000/api/user/cart/products/inc/6756cc53d43baa467b1bac40" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6IjY3NTRjZTkwYjc0ZGJhY2JhN2RlYTdiMSIsImF1ZCI6ImJhZ3VldHRlIiwiZXhwIjoxNzM0MTYzNDA1LCJpYXQiOjE3MzM5MDQyMDUsImlzcyI6ImJhZ3VldHRlIiwibmJmIjoxNzMzOTA0MjA1LCJyb2xlIjoiYWRtaW4ifQ.YK3QcV0PRM023nBbZ_fEWALVe0QiCRkXtJQIoWA_3HM" \
-H "ngrok-skip-browser-warning: true" \
-d '{"quantity": 10}'

