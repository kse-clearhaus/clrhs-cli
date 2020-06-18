# Clearhaus CLI

## TODO

- Credits
- Support ARes, PARes
- Print errors
- Add a SQLite database?
- ...


## Installation

Make sure `$GOPATH` is defined, e.g. in `~/.bashrc`:

```
export GOPATH="${HOME}/.go
```

```
go install github.com/kse-clearhaus/clrhs-cli/cmd/clrhs
```

## Usage

First read `clrhs -h`.

Export `$APIKEY`, `$SIGNINGKEY` and `$PRIVKEY` or pass the as arguments.

### Authorization

Create an authorization:
``` bash
clrhs auth -a <JSON_FILE> --rreq <RREQ_JSON_FILE>
```

Where `<JSON_FILE>` could contain e.g.:
```
{
  "amount": "100",
  "currency": "DKK",
  "text_on_statement": "Test",
  "card[pan]": "4111111111111111",
  "card[expire_month]": "01",
  "card[expire_year]": "2024",
  "card[csc]": "123"
}
```

Example response:
```json
{
  "3dsecure": {
    "version": "2.1.0"
  },
  "_links": {
    "captures": {
      "href": "/authorizations/6165ac5e-bd92-4d8c-bc5e-d8107a646ce6/captures"
    },
    "refunds": {
      "href": "/authorizations/6165ac5e-bd92-4d8c-bc5e-d8107a646ce6/refunds"
    },
    "voids": {
      "href": "/authorizations/6165ac5e-bd92-4d8c-bc5e-d8107a646ce6/voids"
    }
  },
  "amount": 100,
  "csc": {
    "matches": true,
    "present": true
  },
  "currency": "DKK",
  "id": "6165ac5e-bd92-4d8c-bc5e-d8107a646ce6",
  "processed_at": "2020-06-18T11:02:13+00:00",
  "recurring": false,
  "status": {
    "code": 20000
  },
  "text_on_statement": "Test",
  "threed_secure": true
}
```

### Void

Void an authorization:

```bash
clrhs void <Transaction ID>
```

Example response:

```json
{
  "id": "6165ac5e-bd92-4d8c-bc5e-d8107a646ce6",
  "processed_at": "2020-06-18T11:56:14+00:00",
  "status": {
    "code": 20000
  }
}
```
