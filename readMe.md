# DirWatcher

DirWatcher is a Go application that monitors a configured directory for file changes and counts occurrences of a specified magic string in the files. It provides a REST API for configuring the directory, interval, and magic string, as well as retrieving task run details.


## Installation

To install and run DirWatcher, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/dirwatcher.git
    ```

2. Navigate to the project directory:

    ```bash
    cd dirwatcher
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Build the application:

    ```bash
    go build -o dirwatcher ./cmd/dirWatcher
    ```

5. Run the application:

    ```bash
    ./dirwatcher
    ```

## Configuration

DirWatcher can be configured using environment variables or a configuration file. The following parameters can be configured:

- `dirwatcher`: The directory to monitor.
- `defaultInterval`: The interval (in seconds) for monitoring the directory.
- `defaultMagicString`: The magic string to search for in the files.
- `connectString`: The connection string is for postgres db connection.

## REST API

The application exposes the following REST API endpoints:

- `POST /start`: Starts the engine for monitoring the directory.
- `POST /stop`: Stops the engine.

## API Documentation

<iframe src="https://documenter.getpostman.com/view/9950111/2sA3JNaKxW" style="width: 100%; height: 500px;"></iframe>

View the API documentation in Postman [here](https://documenter.getpostman.com/view/9950111/2sA3JNaKxW).


