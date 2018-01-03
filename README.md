# emailApi
[![CircleCI](https://circleci.com/gh/mike001005/emailApi.svg?style=svg)](https://circleci.com/gh/mike001005/emailApi)
[![codecov](https://codecov.io/gh/mike001005/emailApi/branch/master/graph/badge.svg)](https://codecov.io/gh/mike001005/emailApi)


## Installation
HTTPS:
```bash
cd go/src/github.com/...
git clone https://github.com/mike001005/emailApi.git 
```
SSH:
```bash
cd go/src/github.com/...
git@github.com:mike001005/emailApi.git
```

## Edit jsonFile
vim:
```bash
vim emails.json
```
or
nano:
```bash
nano email.json
```
Edit Json file:
```json
{
  "senders": [
    "example@example.com",
    "example@example.com",
    "example@example.com",
    "example@example.com"
  ],
  "receivers": [
    "example@example.com",
    "example@example.com",
    "example@example.com",
    "example@example.com"
  ]
}
```

## Set Env Vars
Option 1: Export Vars in terminal
```bash
export MG_API_KEY=your_MG_API_KEY
export MG_DOMAIN=your_Mg_DOMAIN
export MG_PUBLIC_API_KEY=your_PUBLIC_KEY
export SENDGRID_API_KEY=your_API_KEY
```
Option 2: set vars in docker with -e --env
```bash
docker run --publish 6060:8081 --name email --rm email -e MG_API_KEY="xxxxx" -e MG_DOMAIN="xxxxx" -e MG_PUBLIC_API_KEY="xxxxx" -e SENDGRID_API_KEY="xxxxx"
```
## Docker
Option 1:
```bash
docker build -t email .
docker run --publish 6060:8081 --name email --rm email
```
Option 2: Vars not set
```bash
docker run --publish 6060:8081 --name email --rm email -e MG_API_KEY="xxxxx" -e MG_DOMAIN="xxxxx" -e MG_PUBLIC_API_KEY="xxxxx" -e SENDGRID_API_KEY="xxxxx"
```

## Call Api
Open new terminal
```bash
curl -i -d @emails.json http://localhost:6060/mailgun
curl -i -d @emails.json http://localhost:6060/sendgird
```

