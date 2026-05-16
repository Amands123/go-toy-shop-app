# Happy Kids Toy Shop Application

<img width="1362" height="627" alt="image" src="https://github.com/user-attachments/assets/d2d2097b-a4dc-474c-af01-064ad0320eb1" />



A Toy Shop web application built using Go (Golang) where users can browse toy categories, view products, add products to cart, and place orders.

---

# Features

- Product Categories
  - Teddy
  - Robot
  - Car
  - Lego

- View products by category
- Add products to cart
- Cart management
- Place order functionality
- Order success page
- Responsive UI using HTML and CSS
- Unit test support

---

# Tech Stack

- Go (Golang)
- HTML5
- CSS3
- Go Templates
- Net/HTTP Package

---

# Project Structure

```bash
go-app/
│
├── main.go
├── main_test.go
├── go.mod
│
├── handlers/
├── templates/
└── static/
```

---

# Run Application

Install dependencies:

```bash
go mod tidy
```

Run application:

```bash
go run main.go
```

Application URL:

```text
http://localhost:8080
```

---

# Build Application

Windows:

```bash
go build -o main.exe .
```

Linux/Mac:

```bash
go build -o main .
```

---

# Run Test Cases

```bash
go test -v
```

Run coverage:

```bash
go test -cover
```

---

# Application Flow

```text
Home Page
   ↓
Select Category
   ↓
View Products
   ↓
Add To Cart
   ↓
Cart Page
   ↓
Place Order
```

---

# Future Enhancements

- Database Integration
- User Authentication
- Payment Gateway
- Docker Support
- Kubernetes Deployment
- CI/CD Pipeline

---

# Author

Aman Agarwal
