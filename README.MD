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
- [Swag](https://github.com/swaggo/swag) for generating Swagger documentation. Optionally, add the swag executables into your `PATH` variable.
- [Hurl](https://hurl.dev/) for integration testing.
- [Lefthook](https://github.com/evilmartians/lefthook) for Git hooks.
- [Koji](https://github.com/cococonscious/koji) for Conventional Commits helper.

## Documentation
Open `https://editor.swagger.io` and select `File > Import URL` from the top menu and put this URL:
```
https://raw.githubusercontent.com/risk1996/goshort/refs/heads/main/pkg/docs/swagger.yaml
```

### Setup
Run this to install Lefthook as configured:
```
lefthook install
```

## Usage

### Start the Service
To run the service for debugging, use the following command:
```bash
go run ./main.go
```
The service will start on port 8080 by default. You can access it via http://localhost:8080.

#### Flags
| Name   | Default | Description |
| ------ | ------- | ----------- |
| `port` | `8080`  | Server port |

#### Local Documentation
Local documentation can be viewed when running the application and opening `http://localhost:8080/swagger/index.html`.

### Deployment
You need to build and run the application in release mode, use the following command:
```bash
go build ./main.go
GIN_MODE=release ./main
```

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
