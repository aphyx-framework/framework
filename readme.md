# Ryft Framework

Ryft Framework is a lightweight and efficient web development framework written in Go. It utilizes popular libraries
like GORM for ORM, FX for Dependency Injection, and Fiber for HTTP handling. Our goal is to provide a well-rounded and
easy for beginner framework that allows developers to quickly and easily create web apps in go easily.

> **Warning**
> This framework is still in alpha testing. We're frequently pushing updates and overhaul so your current version might
> be outdated. Please check back often for updates.

## Language

Go

## Features

- ORM via GORM
- Dependency Injection via FX
- HTTP Library via Fiber
- Model and Seeder Definition - all made from scratch
- Niceties: authentication via token, helper functions, routing
- CLI tools for migration, creating boilerplate, etc.

## Documentation

Work in progress

## Getting Started

1. Clone repository from https://github.com/rama-adi/RyFT-Framework
2. Rename the module name in the `go.mod` to be aligned with your project
3. Copy `.example.config.toml` and name it `config.toml` in the root folder
4. Change the app name, port
5. Create a 32-character long key with `go run main.go createkey`
6. Optionally enable features like database, auth, caching
7. Run the app via `go run entrypoint.go`

## Contributing

We welcome and appreciate any contributions to Ryft Framework! To contribute, please follow these steps:

1. Fork the repository
2. Create a new branch for your changes
3. Make your changes and commit them
4. Open a pull request. Please make sure to include a description of your changes and why you think they should be
   merged. If you're fixing a bug, please include steps to reproduce the bug.
5. Wait for your pull request to be reviewed and merged (currently only @rama-adi can merge pull requests)

Please note that we do not have a code of conduct yet, but we expect all contributors to be respectful and professional.

## License

Apache License 2.0