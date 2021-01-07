# Confidentiality Plugin for Caddy

## Config

Add `order confidentiality first` to root level and add `confidentiality <classification>` to the servers you want to have the index on.

**Example:**

```
{
    order confidentiality first
}

http://localhost:2015 {
    confidentiality internal
    root * www/
    file_server
}
```

**Supported classifications:**

- internal
- restricted
- confidential

