# OOM

A Lambda which will run out of memory.

```bash
aws lambda invoke \
--invocation-type RequestResponse \
--function-name oom-dev-memory \
--region eu-west-2 \
--log-type Tail \
--payload "{}" \
output.txt
```

```json
{
    "LogResult": "U1RBUlQgUmVxdWVzdElkOiA4YTg2NmU4Zi0yNmZjLTExZTgtOWQwNi1iZjZjYmRkYTE3ZjIgVmVyc2lvbjogJExBVEVTVApJJ20gYWJvdXQgdG8gdXNlIHVwIGEgbG90IG9mIFJBTS4uLgoxTUIgY29uc3VtZWQKMk1CIGNvbnN1bWVkCjNNQiBjb25zdW1lZAo0TUIgY29uc3VtZWQKNU1CIGNvbnN1bWVkCjZNQiBjb25zdW1lZAo3TUIgY29uc3VtZWQKOE1CIGNvbnN1bWVkCjlNQiBjb25zdW1lZAoxME1CIGNvbnN1bWVkCjExTUIgY29uc3VtZWQKMTJNQiBjb25zdW1lZAoxM01CIGNvbnN1bWVkCjE0TUIgY29uc3VtZWQKMTVNQiBjb25zdW1lZAoxNk1CIGNvbnN1bWVkCjE3TUIgY29uc3VtZWQKMThNQiBjb25zdW1lZAoxOU1CIGNvbnN1bWVkCjIwTUIgY29uc3VtZWQKMjFNQiBjb25zdW1lZAoyMk1CIGNvbnN1bWVkCjIzTUIgY29uc3VtZWQKMjRNQiBjb25zdW1lZAoyNU1CIGNvbnN1bWVkCjI2TUIgY29uc3VtZWQKMjdNQiBjb25zdW1lZAoyOE1CIGNvbnN1bWVkCjI5TUIgY29uc3VtZWQKMzBNQiBjb25zdW1lZAozMU1CIGNvbnN1bWVkCjMyTUIgY29uc3VtZWQKMzNNQiBjb25zdW1lZAozNE1CIGNvbnN1bWVkCjM1TUIgY29uc3VtZWQKRU5EIFJlcXVlc3RJZDogOGE4NjZlOGYtMjZmYy0xMWU4LTlkMDYtYmY2Y2JkZGExN2YyClJFUE9SVCBSZXF1ZXN0SWQ6IDhhODY2ZThmLTI2ZmMtMTFlOC05ZDA2LWJmNmNiZGRhMTdmMglEdXJhdGlvbjogMjgwNzAuOTYgbXMJQmlsbGVkIER1cmF0aW9uOiAyODEwMCBtcyAJTWVtb3J5IFNpemU6IDEyOCBNQglNYXggTWVtb3J5IFVzZWQ6IDEyOSBNQgkKUmVxdWVzdElkOiA4YTg2NmU4Zi0yNmZjLTExZTgtOWQwNi1iZjZjYmRkYTE3ZjIgUHJvY2VzcyBleGl0ZWQgYmVmb3JlIGNvbXBsZXRpbmcgcmVxdWVzdAoK",
    "FunctionError": "Unhandled",
    "StatusCode": 200
}
```