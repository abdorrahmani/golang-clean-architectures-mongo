# Golang GIN MongoDB Clean Architectures

![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)
![Gin](https://img.shields.io/badge/Framework-Gin-green?logo=go)
![MongoDB](https://img.shields.io/badge/Database-MongoDB-brightgreen?logo=mongodb)
![License](https://img.shields.io/github/license/abdorrahmani/golang-clean-architectures-mongo)
[![Website](https://img.shields.io/badge/Website-Online-brightgreen?style=flat&logo=google-chrome)](https://anophel.com)
![Stars](https://img.shields.io/github/stars/abdorrahmani/golang-clean-architectures-mongo?style=social)
---

## Overview

This repository is a collection of **Clean Architecture** examples implemented in **Golang** with **Gin** framework and **MongoDB** as the database. It demonstrates different architectural patterns to build scalable, maintainable, and testable applications.

### Included Architectures:

- **Feature-Based Architecture**
- **Hexagonal Architecture**
- **Domain-Driven Design (DDD)**
- Other modern architecture patterns

---

## Features

- Clean and modular code structure
- Separation of concerns with layers
- Example implementations for each architecture
- Dependency injection and inversion
- Testable components
- Optimized for scalability and maintainability

---

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Architectures](#architectures)
  - [Feature-Based](#feature-based-architecture)
  - [Hexagonal](#hexagonal-architecture)
  - [DDD](#domain-driven-design-ddd-architecture)
- [Contributing](#contributing)
- [License](#license)

---

## Installation

```bash
git clone https://github.com/abdorrahmani/golang-clean-architectures-mongo.git
cd golang-clean-architectures-mongo
go mod tidy
```

---

## Usage

1. Choose the architecture you want to explore from the `examples` folder.
2. Navigate to the respective folder.
3. Run the application:

```bash
go run main.go
```

---

## Architectures

### Feature-Based Architecture

![Feature-Based](https://img.shields.io/badge/Feature--Based-Architecture-informational)

A simple yet effective way to organize code by features. Each feature has its own folder containing its controllers, services, and repository.

### Hexagonal Architecture

![Hexagonal](https://img.shields.io/badge/Hexagonal-Architecture-important)

Also known as "Ports and Adapters," this architecture emphasizes decoupling the core application logic from external dependencies.

### Domain-Driven Design (DDD) Architecture

![DDD](https://img.shields.io/badge/DDD-Architecture-critical)

A more advanced architectural style focused on modeling the business domain and its logic in a way that closely aligns with real-world processes.

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature.
3. Commit your changes and create a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Connect

![GitHub followers](https://img.shields.io/github/followers/abdorrahmani?style=social)
