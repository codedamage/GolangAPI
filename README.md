# Hackathon 2020

#### Configuration

This app using .env file to store settings for DB connection, etc.
To make this app work you need to rename .env.example to .env and 
configure connection to your DB in "DB_DSN" variable.

Example: `user_name:password@tcp(host:3306)/db_name?charset=utf8&parseTime=True&loc=Local`

---

#### How to use
**TODO: Add dockerfile or docker-compose**

0. Edit .env file change DB connection variables, create DB to work with
1. Run project as a module in a GoLand
2. Build app
3. Choose csv-to-database import option in terminal (First run - choose "5", example files already present in a repository)
4. Go to your favorite API testing tool (app tested with Postman)
5. Generate a token, by sending POST request with login/pass(form{username:admin, password:pass}) arguments to [http://localhost:8081/api/login](http://localhost:8081/api/login) Example: [http://localhost:8081/api/login?username=user&password=pass](http://localhost:8081/api/login?username=user&password=pass)
6. Use generated token to get or set info from/into DB
7. Get info about the product(GET request): [http://localhost:8081/api/v1/{product_asin}?token={generated_token}&per_page=1&reviews_page=1](http://localhost:8081/api/v1/?token={generated_token}&per_page=1&reviews_page=1)
8. Add review (POST request, {asin: stirng, title: string, content: string, token: generated_token }): [http://localhost:8081/api/put?asin={asin}&title={title}&content={content}&token={generated_token}](http://localhost:8081/api/put?asin={asin}&title={title}&content={content}&token={generated_token})
---

Pagination arguments not required. Default read parameters for pagination: 2 reviews per page, first page to display.
You can pass any login/pass arguments on a login, token will be generating regardless.
ASIN for test: B07N9BJT4R.

---
#### You could download CSV files using this links:
Products
https://docs.google.com/spreadsheets/d/1roypo_8amDEIYc-RFCQrb3WyubMErd3bxNCJroX-HVE/edit#gid=0

Reviews
https://docs.google.com/spreadsheets/d/1iSR0bR0TO5C3CfNv-k1bxrKLD5SuYt_2HXhI2yq15Kg/edit#gid=0 

---

#### TODO:
1. **DONE** API endpoint (POST) - add entities to DB 
2. **DONE** Add manual csv-to-db console program, do not launch import every time.
3. **DONE** Add limits and pagination to reviews
4. Cache GET endpoints
5. Pack into docker container
6. **DONE** Add placeholder outputs, like "auth is failed/wrong url/access forbidden" or something, on wrong urls in json
7. **DONE** Prettify code, maybe make operations of reading/writing data like isolated functions(e.g "get_reviews(asin) or get_book_data(asin)")
8. **DONE** Improve JWT
9. Add an answer for put operation