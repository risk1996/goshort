# Goshort

A simple Go-based service to manage shortened links with support for create, access, edit, enable, and disable functionalities.

## Features
- Create short links for long URLs.
- Access and redirect to the original URLs.
- Edit or update the destination URL of an existing link.
- Enable or disable links.
- Idempotency checks for state change actions.

## Requirements
- [Go](https://go.dev) v1.23.4 installed on your system.
- [Hurl](https://hurl.dev/) for integration testing.

## Default Configuration
- **Port**: `8080`
- **Base URL**: `http://localhost:8080`

## Usage

### Start the Service
To run the service, use the following command:
```bash
go run main.go
```
The service will start on port 8080 by default. You can access it via http://localhost:8080.

### Integration Testing
To test the service using the provided Hurl scripts:
```
hurl --variables-file ./integration/.local.properties --test .
```
The `./integration/.local.properties` file should contain your test-specific variables, such as:
```
base_url=http://localhost:8080

target_url=https://www.bing.com/
target_url_normalized=https://www.bing.com

edit_url=https://www.google.com/
edit_url_normalized=https://www.google.com

wrong_data=123
```
