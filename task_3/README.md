# Central Bank of the Russian Federation Exchange Rate Analyzer

A Go program for obtaining daily exchange rates from the API of the Central Bank of the Russian Federation for the last 90 calendar days, analyzing them and outputting the maximum, minimum and average values of the ruble exchange rate for the entire period for all currencies.

The program uses the "Worker Pool" pattern for efficient and competitive data acquisition, and also includes retry logic for resilience to temporary network or API problems.

## Requirements

To build and run the project, you must have the following installed:

* **Go**: Versions **1.16 or higher**.
* **Make**: A utility for automating assembly (usually pre-installed on most Unix-like systems, available for installation on Windows).

## Installation and Launch

1. Clone the repository or make sure that you have all the project files in the `task_3` directory.
2. Open the terminal and navigate to the root directory of the `task_3` project:
    ```bash
    cd path/to/your/task_3
    ```
3. Initialize the Go module and download the necessary dependencies. If you run `make build` or `make run`, these steps will be performed automatically, but it's useful to know about them.:
    ```bash
    go mod tidy
    # or go mod download
    ```
4. Assemble and run the program using the `make` utility:
    ```bash
    make run
    ```
    The `make run` command will first assemble the executable file (`parse_cbr` or `parse_cbr.exe `), and then it will launch it.

## Available `make` commands

The `Makefile` defines standard goals for ease of development:

```bash
make help
```

This will list the available commands.:

* `make all`: Builds an executable file (default action);
* `make build`: Compiles the Go application (`parse_cbr` or `parse_cbr.exe `);
* `make run`: Builds and launches the app;
* `make vet`: Performs static code analysis using `go vet`;
* `make fmt`: Formats the Go source code using `go fmt`;
* `make clean`: Deletes the compiled executable file.

## Expected Output

The program will output the results of currency exchange rate analysis in the following format to the standard output stream:


```bash
Max valute:
  Value: <value of the maximum exchange rate per unit>
  Name: <name of the currency with the maximum exchange rate>
  Date: <date when the maximum exchange rate was fixed>
Min value:
  Value: <value of the minimum exchange rate per unit>
  Name: <name of the currency with the minimum exchange rate>
  Date: <date when the minimum exchange rate was fixed>
Avg value: <the average value of the ruble exchange rate for the entire period for all currencies>
```
Error messages may also be displayed during execution (for example, in case of network problems, parsing, or data conversion) indicating the date for which they occurred.