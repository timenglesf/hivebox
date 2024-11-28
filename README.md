# hivebox

[![Dynamic DevOps Roadmap](https://devopshive.net/badges/dynamic-devops-roadmap.svg)](https://github.com/DevOpsHiveHQ/dynamic-devops-roadmap)[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/{owner}/{repo}/badge)](https://scorecard.dev/viewer/?uri=github.com/{owner}/{repo})

Hivebox is an API that provides environmental data relating to bees. The enviornmental data is sources from [openSenseMap](https://opensensemap.org/)

## Getting Started

### Prerequisites

-   Python 3.12
-   Docker (optional, for containerized deployment)

### Installation

1. Clone the repository:

    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Create a virtual environment and activate it:

    ```sh
    python -m venv .venv
    source .venv/bin/activate  # On Windows use `.venv\Scripts\activate`
    ```

3. Install the dependencies:
    ```sh
    pip install -r requirements.txt
    ```

## Running the Application

To run the application locally, use the following command:

```sh
fastapi dev ./app/main.py
```

The application will be available at `http://127.0.0.1:8000`.

### Docker

To build and run the application using Docker:

1. Build the Docker image:

```sh
docker build -t hivebox .
```

2. Run the Docker container

```sh
docker run -p 8000:8000 hivebox
```

### Running Tests

To run tests, use the following command:

```sh
pytest
```

### CI/CD

This project uses GitHub Actions.

### Endpoints

-   `GET /`: Returns a greeting message.
-   `GET /version:` Returns the application version.
