
---

# Number Classifier API

This is a simple API built using the **Go (Golang)** programming language and **Gin** framework that classifies a number based on its properties, checks whether it's a prime number, an Armstrong number, and returns a fun fact about the number.

## Features
- Classifies numbers as **Armstrong**, **Odd**, or **Even**.
- Checks whether a number is **prime**.
- Returns the **sum of its digits**.
- Provides a **fun fact** about the number.
- Returns results in **JSON format**.

## API Endpoint

### `GET /api/classify-number?number={number}`

### Parameters
- `number`: The number you want to classify. It should be a valid integer.

### Example Requests

1. **Valid Request with Number 371**
   ```sh
   GET https://number-classifier-api-k8cw.onrender.com/api/classify-number?number=371
   ```

   **Response:**
   ```json
   {
     "number": 371,
     "is_prime": false,
     "is_perfect": false,
     "properties": ["armstrong", "odd"],
     "digit_sum": 11,
     "fun_fact": "371 is a narcissistic number."
   }
   ```

2. **Invalid Request with Non-Numeric Input**
   ```sh
   GET https://number-classifier-api-k8cw.onrender.com/api/classify-number?number=hello
   ```

   **Response:**
   ```json
   {
     "number": "hello",
     "error": true
   }
   ```

---

## Installation

### Prerequisites

Ensure you have the following tools installed:
- Go 1.23.4 or later
- Git
- Optional: Postman or cURL for testing the API locally

### Local Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/rowjay007/number-classifier-api.git
   ```

2. Navigate to the project folder:
   ```sh
   cd number-classifier-api
   ```

3. Install dependencies (using `go mod`):
   ```sh
   go mod tidy
   ```

4. Build the Go application:
   ```sh
   go build -tags netgo -ldflags '-s -w' -o app
   ```

5. Run the application:
   ```sh
   ./app
   ```

The server will start running at `http://localhost:8080`.

---

## Deployment

The API is deployed to **Render** and can be accessed publicly at:

[https://number-classifier-api-k8cw.onrender.com](https://number-classifier-api-k8cw.onrender.com)

---

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/your-feature-name`).
5. Create a new Pull Request.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---