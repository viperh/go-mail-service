# Mail Service


## Description

Mail service to send mail using SMTP and go


## Installation

```bash
git clone https://github.com/qviperh/go-mail-service.git
```

```bash
cd go-mail-service
```

```bash
go mod download
```

```bash
go run cmd/main.go
```


## Recommended to use an app password for login

## Endpoints


### Send mail

<b>POST</b> /sendMail

```json
{
    "from": "example.from@gmail.com",
    "to": ["example.to@gmail.com"],
    "subject": "example subject",
    "body": "example body",
    "isHtml": "true/false"

}
```


# Info


### This repository was made for personal use. 
### Feel free to use it and modify it as you wish.
### It should be used with a service account as sender.
### It is pretty straight forward and simple.

