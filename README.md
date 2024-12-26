# Go Gin Boilerplate with Laravel-Like Structure, Templ, Tailwind CSS, Preline and Air UI for Scalable Web Applications 

A boilerplate for building modern backend applications using the **Go Gin framework**, integrated with **TailwindCSS**, **Preline UI components**, and live reloading with **Air** for rapid development.

## Features

- Lightweight and efficient backend powered by [Gin](https://github.com/gin-gonic/gin).
- UI styling with [TailwindCSS](https://tailwindcss.com/) and components from [Preline](https://preline.io/).
- Live reload during development using [Air](https://github.com/cosmtrek/air).
- Well-organized project structure for scalability and maintainability.

---

## Getting Started

### Prerequisites

Ensure the following tools are installed on your system:

- [Go](https://golang.org/) (1.19 or newer recommended)
- [Node.js](https://nodejs.org/) (v16 or newer) with npm
- [Air](https://github.com/cosmtrek/air) for live reloading
- TailwindCSS and Preline libraries (included in the setup steps)

---

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/Forkan-dev/laragintempltail.git
   cd <your-repo-name>
   ```

2. **Set Up TailwindCSS**

   Navigate to the `resources` directory where TailwindCSS is configured.

   ```bash
   cd resources
   npm install
   ```

   This will install all required dependencies for TailwindCSS, including Preline components.

3. **Initialize Air**

   Ensure you have Air installed globally. If not, install it using the following command:

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

   Once installed, return to the root directory of the project and run:

   ```bash
   air init
   ```

   This will create an `.air.toml` configuration file. The `air` configuration is already included in the repository, so you should not need to modify it.

4. **Run the Application**

   Start the live reload server using Air:

   ```bash
   air
   ```

   Your application will be running at `http://localhost:8080` by default.

---

## Project Structure

```plaintext
├── main.go               # Entry point of the application
├── resources/            # Contains TailwindCSS configuration and assets
│   ├── tailwind.config.js  # TailwindCSS configuration file
│   ├── package.json        # NPM configuration file
│   └── ...                 # Other assets
├── templates/            # HTML templates for the application
├── static/               # Static files (CSS, JS, images)
├── routes/               # Gin route definitions
├── controllers/          # Handlers for routes
└── .air.toml             # Air configuration for live reloading
```

---

## Development

### TailwindCSS Customization

To customize the TailwindCSS configuration, modify the `tailwind.config.js` file located in the `resources` folder. After making changes, rebuild the CSS by running:

```bash
npm run build
```

Alternatively, use `npm run watch` for automatic rebuilds during development.

### Live Reload with Air

Air watches your Go files and automatically reloads the server upon detecting changes. To use it:

- Ensure `main.go` is correctly set up as the entry point.
- Start the server with `air` from the root directory.

---

## Deployment

1. **Build the Application**

   Use the Go build command to create a production-ready binary:

   ```bash
   go build -o app .
   ```

2. **Static Assets**

   Ensure all TailwindCSS and static assets are built before deployment:

   ```bash
   npm run build
   ```

3. **Run the Application**

   Start the compiled binary:

   ```bash
   ./app
   ```

---

## Contributing

Contributions are welcome! Please fork the repository, make changes, and submit a pull request for review.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

- [Gin Framework](https://github.com/gin-gonic/gin)
- [TailwindCSS](https://tailwindcss.com/)
- [Preline UI](https://preline.io/)
- [Air](https://github.com/cosmtrek/air)

---

Happy coding!

