## GO ERP

**GO ERP** is a backend system written in Golang for Enterprise Resource Planning (ERP) functionalities.  This project allows you to manage authentication, items, customers, and orders.

### Features

* User authentication with secure password hashing
* Item management: Create, read, update, and delete items.
* Customer management: Create, read, update, and delete customer information.
* Order management: Create, read, update, and delete orders with associated items.

### Technologies

* Golang (programming language)
* **GORM** (ORM library for interacting with databases)
* **Fiber** (fast, flexible, and expressive web framework)
* **SQLite** (lightweight, serverless database)

### Prerequisites

* Golang installed (version 1.20 or later recommended)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/rodrigomedeirosbrazil/goerp
```

2. Navigate to the project directory:

```bash
cd goerp
```

3. Install dependencies:

```bash
go get -d ./...
```

4. Configure the database connection details in the `config.toml` file (details on configuration options will be added)

5. Run the application:

```bash
go run main.go
```

### Usage

**API Documentation**

The API documentation will be added shortly.  In the meantime, you can explore the codebase for insights on API endpoints and request/response structures.

**Additional Notes**

* This project is under development. Features and functionalities will be added over time.
* Contribution guidelines will be added soon.

### License

This project is licensed under the MIT License (see LICENSE file for details).

### Getting Help

For any questions or issues, feel free to open an issue on the GitHub repository.


**Please note:** This is a starting point for your README file.  You can customize it further by adding details on:

* Supported databases and configuration options (if using other databases besides SQLite)
* Specific API endpoints and usage examples
* Deployment instructions (if applicable)
* Contribution guidelines
* Screenshots or diagrams (if helpful)

I hope this helps!
