# tibco-kafka
This activity provides your flogo application the ability to send a message to an Apache Kafka broker.


## Installation

```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/activity/kafka
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "token",
      "type": "string"
    },
    {
      "name": "message",
      "type": "string"
    }
  ],
  "outputs": []
}
```
## Settings
| Setting     | Description    |
|:------------|:---------------|
| topic | The Kafka Topic name |         
| message  | The message content  |

## Configuration Examples
### Simple
Configure a task in flow to send 'hello from flogo' to 'my-messages' Kafka Topic:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "tibco-kafka",
  "name": "Send Message",
  "attributes": [
    { "name": "topic", "value": "my-messages" },
    { "name": "message", "value": "hello from flogo" },
  ]
}
```

### Advanced
TODO
```
