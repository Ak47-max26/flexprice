# FlexPrice Python SDK Scripts

This directory contains scripts used during the SDK generation process to add custom functionality.

## Scripts

### `add_python_async.py`

This script enhances the auto-generated Python SDK by adding asynchronous functionality for event submission. It is automatically run as a post-processing step during SDK generation.

## Features

The `add_python_async.py` script adds the following features to the Python SDK:

1. **Asynchronous Event Submission**: Fire-and-forget event posting that doesn't block your application
2. **Background Processing**: Events are processed in background threads
3. **Automatic Retries**: Failed events are automatically retried with exponential backoff
4. **Error Handling**: Optional callbacks for success/failure notification
5. **Graceful Shutdown**: Background threads are properly terminated on application exit

## Integration Points

The script is integrated into the SDK generation process:

1. It's executed as a post-processing step after the standard SDK generation
2. It creates an `async_utils.py` file with the async processing implementation
3. It adds an `events_post_async()` method to the `EventsApi` class

## Usage Example

```python
from flexprice import Configuration, ApiClient, EventsApi
from flexprice.models import DtoIngestEventRequest

# Configure API
configuration = Configuration(api_key={'ApiKeyAuth': 'YOUR_API_KEY'})
api_client = ApiClient(configuration)
events_api = EventsApi(api_client)

# Create event
event = DtoIngestEventRequest(
    external_customer_id="customer123",
    event_name="api_call",
    properties={"region": "us-west", "method": "GET"},
    timestamp="2023-01-01T12:00:00Z"
)

# Fire and forget submission
events_api.events_post_async(event)

# Submission with completion callback
def on_complete(result, error, success):
    if success:
        print(f"Event sent successfully: {result}")
    else:
        print(f"Event failed: {error}")

events_api.events_post_async(event, callback=on_complete)
```

## Implementation Details

The implementation uses a background thread pool to process events asynchronously:

1. Events are queued for processing when `events_post_async()` is called
2. Worker threads pick up events from the queue and process them
3. If an event fails, it is retried with exponential backoff and jitter
4. After maximum retries, it's considered permanently failed and the callback is notified

## Advanced Configuration

The async event processor can be configured by modifying the parameters in the `AsyncEventProcessor` class:

- `max_retries`: Maximum number of retry attempts (default: 3)
- `retry_delay`: Base delay between retries in seconds (default: 1.0)
- `max_queue_size`: Maximum size of the event queue (default: 1000)
- `workers`: Number of worker threads (default: 2)

## Troubleshooting

If you encounter issues with the async functionality:

1. Enable debug logging to see detailed information about event processing
2. Check for queue overflow if submitting events at a very high rate
3. Increase worker threads if processing throughput is insufficient

## Dependencies

The implementation uses only standard Python libraries:
- `threading`
- `queue`
- `logging`
- `time`
- `random` 