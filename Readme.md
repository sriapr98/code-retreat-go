## Requirements

The project requires Go

### Build the project

Compiles the project and then creates an executable file

```console
$ go build
```

Run the application using executable file produced by the `build` task. The application will be
listening to port `8080`.

```Run the application
$ ./code-retreat-go
```

### Run the application

Run the application which will be listening on port `8080`.

## API

Below is a list of API endpoints with their respective input and output. Please note that the application needs to be
running for the following endpoints to work. For more information about how to run the application, please refer
to [run the application](#run-the-application) section above.

### Get employees

```text
GET /employees
```

### Get terminations

```text
GET /terminations
```