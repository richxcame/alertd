# alertd

The tool is designed to monitor the server status using the iLO API and send alerts to the system administrators using the SMS API.
Tested with the following servers:

-   iLO 4 ProLiant DL380 Gen9 Intel(R) Xeon(R) CPU E5-2683 v4 @ 2.10GHz

## Preqrequisites:

-   Golang (1.16+)
-   [go-sms](https://github.com/richxcame/go-sms)

## Installation:

1. Create `.env`

```bash
cp .env.example .env
```

2. Server ilo list as shown in `servers.example.json` (which needs to be monitored)

```bash
cp servers.example.json servers.json
```

3. System administrators list as shown in `admins.example.json` (Which will be used to send alerts)

```bash
cp admins.example.json admins.json
```

## Run the application in development mode:

```bash
go run main.go
```

## Buid the application in production mode:

```bash
$ go build -o alertd
$ ./alertd
```
