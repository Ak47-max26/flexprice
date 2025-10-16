### Implementing the ClickHouse store

Building a ClickHouse store for the events. The store should be able to:

- Insert events
- Get usage of an event, aggregated in different ways (count, sum, max, unique count, latest)


### Inserting events
Sample payload:


```json
// Basic event with minimal fields
{
	"customer_id": "123",
	"event_name": "purchase",
}
```

```json
// API call event with numeric properties
{
    "id": "evt_abc123",
    "customer_id": "cust_123",
    "event_name": "api_call",
    "timestamp": "2024-03-20T10:15:00Z",
    "properties": {
        "duration_ms": 150,
        "status_code": 200,
        "bytes_transferred": 1024
    }
}
```

```json
// Purchase event with monetary values
{
    "customer_id": "cust_123",
    "event_name": "purchase",
    "timestamp": "2024-03-20T11:00:00Z",
    "properties": {
        "amount": 99.99,
        "currency": "USD",
        "items": 3,
        "product_id": "prod_456"
    }
}
```

### Getting usage

```
# Count all page views
curl "http://localhost:8080/usage?customer_id=cust_123&event_name=page_view&aggregation_type=count"

# Sum of purchase amounts
curl "http://localhost:8080/usage?customer_id=cust_123&event_name=purchase&property_name=amount&aggregation_type=sum"


```


