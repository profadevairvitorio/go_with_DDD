# go_with_DDD
Simple GO Project using DDD (Domain-Driven Design)
## Repository Structure
`domain`: The heart of the software, representing business logic and rules.
- `entities`: Fundamental objects within our system, like `Product` and `Seller`.
- `application`: Contains use-case specific operations that interact with the domain layer.
- `infrastructure`: Supports the higher layers with technical capabilities like database access.
    - `db`: Database access and models.
    - `mappers`: Converters between domain entities and database models.
    - `repositories`: Concrete implementations of our storage needs.
- `interface`: The external layer which interacts with the outside world, like API endpoints.
    - `api/rest`: Handlers or controllers for managing HTTP requests and responses.


