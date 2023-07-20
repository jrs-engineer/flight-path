# flight-path
This is the simple microservice to detemine the flight path from the history.

## How to run

```bash
# run the service
make run

curl -X POST -H "Content-Type: application/json" -d "[[\"A\",\"F\"]]" http://localhost:8080/calculate
```

```bash
# run the tests
make test

make lint
```

