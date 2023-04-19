# Go Fiber URL Shortener

This is a simple URL Shortener REST API project that I am using to learn Go and the Fiber web framework.

## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/rileygrotenhuis/go-fiber-url-shortener.git
   ```

2. Install dependencies

   ```bash
   cd go-fiber-url-shortener
   go mod tidy
   ```

3. Copy the `.env.example` file as `.env.local` and `.env` then update the given variables

   ```bash
   cp .env.example .env
   ```

4. Start the development server

   ```
   go run main.go
   ```

5. Your application will now begin running at http://localhost:8080

## Contributing

All contributions, big or small, are welcome and I appreciate any effort to improve this repository!

View our [Contributing Guide](CONTRIBUTING.md) for more details.

## License

[MIT License](LICENSE.txt)