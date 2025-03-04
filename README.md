#  Response to MUZZ Back-End Technical Test

### Golang gRPC Microservice

Repository: [GitHub - steve-westwood/explore-service](https://github.com/steve-westwood/explore-service)

### Overview

This repository contains a Golang-based gRPC microservice that interacts with a PostgreSQL database. Both the microservice and the database are containerized using Docker, and they are managed via docker-compose for easy deployment.

### Features

- gRPC Server to handle requests.
- PostgreSQL database for persistence.
- Containerized using Docker.
- docker-compose for build and start.

### Prerequisites

Before running the microservice, ensure you have the following installed:

- Docker
- Docker Compose
- Golang (only needed for local development)

### Getting Started

### Clone the repository

```
git clone https://github.com/steve-westwood/explore-service.git
```

### Build and Start Services

Run the following command to build and start the microservice and the database:

```
docker-compose up --build
```

This command will:

- Build the Golang gRPC microservice.

- Start the microservice container.

- Start the PostgreSQL database container.

- Create the database and the schema.

### Server Address

The server address is `localhost:9000`.

### Verify the Services

Once the services are up and running, you can verify the gRPC server is running using a gRPC client 

```
grpcurl -plaintext -d '{"recipient_id": "123e4567-e89b-12d3-a456-426614174000"}' localhost:9000 ExploreService/ListLikedYou
```

### Design

#### Microservice Design

The project follows a structured design pattern where main.go acts as the orchestrator for various services. This ensures separation of concerns and keeps the core business logic modular and maintainable.

main.go is responsible for initializing dependencies, setting up the gRPC server, and handling service orchestration.

Services are structured in a way that each component (e.g., database access, business logic) is handled separately to ensure modularity and ease of testing.

#### Database Design

The database uses a single table, `decisions`, to store user pass/like decisions. The simplicity of a single table enables efficient querying, with different queries for the required use cases.

```
CREATE TABLE Decisions (
    DecisionId UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    RecipientId UUID NOT NULL,
    ActorId UUID NOT NULL,
    Timestamp BIGINT NOT NULL,
    Liked BOOLEAN NOT NULL,
    UNIQUE (RecipientId,ActorId)
);
```

The primary method for inserting data into the table is through an `UPSERT` operation.
This design choice ensures that each interaction is unique and can be updated efficiently without introducing redundant records. In production a fuller auditable history would probably run beside this data. The other benefit of this structure is it's simplicity lends itself to easy sharding and scaling.

### Stopping the Service

To stop the running containers, use:

```
docker-compose down -v
```

This will remove volumes containing the database data.

### Running Unit Tests

To run the unit tests, use the following command:

```
go test ./...
```

This will execute all test cases within the project.

### Extending the Project

#### End-to-End Tests

To extend this microservice with end-to-end (E2E) tests, you could create a separate Go project that acts as a gRPC client and tests the containerized service locally. Although expensive, this test would be highly valuable for testing data integrity and performance. 

#### Logging and Metrics

I didn't have time to implement logging and metrics for the microservice. Again in production this would be a valuable addition.

