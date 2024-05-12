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

## TaskRun Table Schema

The `taskrun` table stores information about each task run in the DirWatcher application.

## Schema Design

| Column       | Type          | Description                                  |
|--------------|---------------|----------------------------------------------|
| ID           | INTEGER       | Unique identifier for the task run.           |
| StartTime    | TIMESTAMP     | Start time of the task run.                  |
| EndTime      | TIMESTAMP     | End time of the task run.                    |
| Runtime      | INTERVAL      | Duration of the task run.                    |
| Status       | VARCHAR(20)   | Status of the task run (e.g., Success, Failed). |
| FilesAdded   | TEXT ARRAY    | List of files added during the task run.     |
| FilesDeleted | TEXT ARRAY    | List of files deleted during the task run.   |
| Occurrences  | INTEGER       | Total occurrences of the magic string.       |
| Directory    | VARCHAR(255)  | Directory being monitored during the task run. |
| MagicString  | VARCHAR(255)  | Magic string being searched for.             |



## REST API

The application exposes the following REST API endpoints:

### Endpoints

- `POST /config`: Updates the configuration for the directory monitoring task.

    **Request Body:**

    ```bash
    curl --location --request POST 'http://localhost:8080/config' \
    --data '{
        "dir": "../Users/test",
        "interval": 10,
        "magic_string": "go"
    }'
    ```

    **Response:**

    ```bash
    HTTP/1.1 200 OK
    ```

- `GET /getTaskRuns`: Retrieves information about recent task runs.

    **Request:**

    ```bash
    curl --location 'http://localhost:8080/getTaskRuns'
    ```

    **Response:**

    ```json
    [
        {
            "ID": 107,
            "StartTime": "2024-05-12T21:16:27.143611Z",
            "EndTime": "2024-05-12T21:16:27.145639Z",
            "Runtime": "00:00:00.002028",
            "Status": "Success",
            "FilesAdded": null,
            "FilesDeleted": null,
            "Occurrences": 1,
            "Directory": "/Users/khalith/projects/DirWatcher/.vscode",
            "MagicString": "cmd"
        },
        {
            "ID": 106,
            "StartTime": "2024-05-12T21:16:07.144481Z",
            "EndTime": "2024-05-12T21:16:07.146514Z",
            "Runtime": "00:00:00.002033",
            "Status": "Success",
            "FilesAdded": [
                "launch.json"
            ],
            "FilesDeleted": null,
            "Occurrences": 1,
            "Directory": "/Users/khalith/projects/DirWatcher/.vscode",
            "MagicString": "cmd"
        },
        ...
    ]
    ```

- `POST /stopMonitoringTask`: Stops the monitoring task.

    **Request:**

    ```bash
    curl --location --request POST 'http://localhost:8080/stopMonitoringTask'
    ```

    **Response:**

    ```json
    {
        "message": "Task stopped successfully"
    }
    ```

- `POST /startMonitoringTask`: Starts the monitoring task.

    **Request:**

    ```bash
    curl --location --request POST 'http://localhost:8080/startMonitoringTask'
    ```

    **Response:**

    ```json
    {
        "message": "Task already running"
    }
    ```

- `POST /toggleMonitoringTask`: Toggles the monitoring task (start if stopped, stop if running).

    **Request:**

    ```bash
    curl --location --request POST 'http://localhost:8080/toggleMonitoringTask'
    ```

    **Response:**

    ```json
    {
        "message": "Task started successfully"
    }
    ```


## API Documentation

View the API documentation in Postman [here](https://documenter.getpostman.com/view/9950111/2sA3JNaKxW).


