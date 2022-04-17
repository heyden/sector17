# Webhook Example

## Payload Specification

```
{
    Notification {
        Level     string
		Scope     string 
		Group     string
		Timestamp string
		Title     string
		Content   string
		Subject   object
    }
}
```