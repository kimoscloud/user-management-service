## USER MANAGEMENT SERVICE => Ex management service (Actually it's a monolith)

Service responsible for managing users and their roles.

### How to configure

1. Create a .env file in the root directory of the project
2. Request to nonils@tutanota.com a candidate .env file
3. Copy the content of the candidate .env file into the .env file
4. Fill the .env file with the correct values
5. Save the .env file
6. Run the service
7. Enjoy

### How to run

1. Install all go dependencies

```bash
go mod download
```

2. Run the service

```bash
  env ENV=dev go run cmd/main.go
```

# TODO:

- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Add e2e tests
- [ ] Add swagger documentation
- [ ] Add dockerfile that works

Note: This service is not finished yet.