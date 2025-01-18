# API Testing CLI Tool

Welcome to the **API Testing CLI Tool**! 🚀

This command-line application empowers you to test and analyze API performance with advanced metrics and customizable parameters.

---

## Features

- **Comprehensive API Testing**:
  - Measure response times.
  - Assess resilience under load.
  - Identify failure rates.
- **Support for All HTTP Methods**:
  - GET, POST, PUT, DELETE, PATCH, and more.
- **Dynamic Command Parsing**:
  - Flexible input handling with support for flags and interactive mode.
- **Stress and Load Testing**:
  - Simulate multiple concurrent requests.
- **Detailed Reports**:
  - Metrics include average, minimum, and maximum response times, as well as error rates.
- **Batch Testing**:
  - Test multiple endpoints in a single session.
- **Extensibility**:
  - Easily add custom tests or metrics.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/api-testing-cli.git
   cd api-testing-cli
   ```

2. Build the application:
   ```bash
   go build -o mycli ./cmd
   ```

3. Run the CLI tool:
   ```bash
   ./mycli
   ```

---

## Usage

### General Command Structure

```bash
<command> [--url <URL>] [--method <METHOD>] [--body <BODY>] [--requests <NUM>] [--concurrency=<NUM>]
```

### Available Commands

- **test**: Perform API testing with specified parameters.
  ```bash
  test --url="https://api.example.com" --method=GET --requests=10 --concurrency=2
  ```
  Parameters:
  - `--url`: API endpoint to test (required).
  - `--method`: HTTP method (default: GET).
  - `--body`: Request body (for POST/PUT requests).
  - `--requests`: Number of requests to send (default: 1).
  - `--concurrency`: Number of concurrent requests (default: 1).

- **help**: Display help information.
  ```bash
  help
  ```

### Interactive Mode

Launch the application without commands to enter interactive mode. Example:
```bash
Введите команду: test --url=https://api.example.com --method=POST --requests=10 --concurrency=2
```

---

---

## Example Output

```
later...
```

---

## Future Enhancements

- **Real-Time Graphs**: Visualize metrics during test execution.
- **Authentication Support**: Add headers for tokens and API keys.
- **Custom Plugins**: Allow users to integrate their own test logic.

---

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
