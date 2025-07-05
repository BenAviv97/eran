# API Gateway

This service exposes a lightweight HTTP interface.

## Endpoints

| Method | Path       | Description                      |
| ------ | ---------- | -------------------------------- |
| GET    | /healthz   | Liveness probe returning `ok`.   |
| GET    | /channels  | Returns a JSON array of channel names. |
