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
nano emails.json
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
Option 1: Set Var in 
```bash
vim Dockerfile
```
or 
```bash
nano Dockerfile
```
Edit Env Var in DockerFile:
```Dockerfile
MG_API_KEY=your_MG_API_KEY
MG_DOMAIN=your_Mg_DOMAIN
MG_PUBLIC_API_KEY=your_PUBLIC_KEY
SENDGRID_API_KEY=your_API_KEY
```
## Docker
```bash
docker build -t email .
docker run --publish 6060:8081 --name email --rm email
```
## Call Api
Open new terminal
```bash
curl -i -d @emails.json http://localhost:6060/mailgun
curl -i -d @emails.json http://localhost:6060/sendgird
```

