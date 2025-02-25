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


### Send mail using api

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


### Send email using rabbitmq

See env.example for rabbitmq configuration.

The service will take the message from the queue and send the email.

```
{
"from": "example.from@gmail.com",
"to": ["example.to@gmail.com"],
"subject": "example subject",
"body": "example body",
"isHtml": "true/false",

}
```

# Info


### This repository was made for personal use. 
### Feel free to use it and modify it as you wish.
### It should be used with a service account as sender.
### It is pretty straight forward and simple.

# Added Jwt Auth for external Service.
## Make sure you have the same JwtSecret. 
## Generate a jwt and send it in header as usual.