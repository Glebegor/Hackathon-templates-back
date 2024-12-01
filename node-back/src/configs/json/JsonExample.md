# Example of Json

## Description

Change and create json config which describes your project. Application need to have 3 config:

- config.dev.json
- config.prod.json
- config.test.json

## Example

```config
{
    "SERVER": {
        "PORT": 3000,
        "HOST": "localhost"
    },
    "DB": {
        "HOST": "localhost",
        "PORT": 5436,
        "USER": "postgres",
        "SSLMODE": "disable",
        "DATABASE": "postgres"
    }
}
```

