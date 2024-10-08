# atlas-maps

Mushroom game maps Service

## Overview

A RESTful resource which provides maps services.

## Environment

- JAEGER_HOST - Jaeger [host]:[port]
- LOG_LEVEL - Logging level - Panic / Fatal / Error / Warn / Info / Debug / Trace
- BOOTSTRAP_SERVERS - Kafka [host]:[port]
- GAME_DATA_SERVICE_URL - [scheme]://[host]:[port]/api/gis/ 
- MONSTER_SERVICE_URL - [scheme]://[host]:[port]/api/mos/
- EVENT_TOPIC_CHARACTER_STATUS - Kafka Topic for transmitting character status events
- EVENT_TOPIC_MAP_STATUS - Kafka Topic for transmitting map status events

## API

### Header

All RESTful requests require the supplied header information to identify the server instance.

```
TENANT_ID:083839c6-c47c-42a6-9585-76492795d123
REGION:GMS
MAJOR_VERSION:83
MINOR_VERSION:1
```

### Requests

Requests are documented via Bruno collection.