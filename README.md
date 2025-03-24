# Easy-AI-Router

Server application with REST API that provides access to different providers or routers of AI models. Currently supported providers and routers: 
- OpenRouter.

## Setup

Make sure to install the dependencies:

```bash
go mod tidy
```

Environment variables should be put in .env file.

## Production

Build the application for production:

```bash
go build
```

## Environment variables

Application is using environment variables. You have to define:

- PORT on which the server will run locally.
- ACCESS_OPENROUTER_API_KEY to access application's REST API for OpenRouter.
- OPENROUTER_HOST is host URL for OpenRouter REST API.
- OPENROUTER_API_KEY is api key for OpenRouter REST API.

You can define all needed variables in .env file in root folder of this application.

## How to use Easy-AI-Router REST API

After deployment of this app on a server, you should have access to it's REST API. For example we will use: your-domain.com.

REST API hosts: 
- OpenRouter: your-domain.com/openrouter/api/v1. Api-Key to access this routes should be provided in headers["access-openrouter"], it is separate from OPENROUTER_API_KEY which is used to get access to OpenRouter's own REST API. Currently "easy-ai-router" uses deepseek/deepseek-r1:free model for chat completion text and google/gemma-3-12b-it:free for chat completion image description. We will expand the ability to choose which model to use in future updates.

### OpenRouter REST API v1

GET /ping

Response example:
```json
{
  "message": "openrouter ping"
}
```

GET /key

Response example:
```json
{
  "data": {
    "label": "<YOUR_OPENROUTER_API_KEY>",
    "usage": 0,
    "is_free_tier": true,
    "is_provisioning_key": false,
    "rate_limit": {
      "requests": 10,
      "interval": "10s"
    },
    "limit": 0,
    "limit_remaining": 0
  }
}
```

POST /chat

Body example:
```json
[
  {
    "role": "user",
    "content": "tell me a joke"
  }
]
```

Response example:
```json
{
  "id": "gen-1742842377-CFRRcQjjZD9gDui1uS8W",
  "choices": [
    {
      "message": {
        "role": "assistant",
        "content": "Sure! Here's a lighthearted joke for you:\n\nWhy don’t you ever see elephants hiding in trees?  \n*…Because they’re really good at it!* 🌳🐘  \n\n(Or if you’d prefer a groan-worthy pun):  \nWhat do you call a dog that can do magic?  \n*…A labracadabrador!* 🐕✨  \n\nLet me know if you need more laughs! 😄"
      }
    }
  ]
}
```

## Additional information

Easy-AI-Router is written in Go language (Go 1.24.1), uses: chi, godotenv. Please, before proceed be sure to check official documentation on corresponding technology.

## Known Issues

There are currently no known issues.

## Release Notes

### 0.1.0

Initial working version.

---

## For more information

* [GitHub](https://github.com/ikirja/easy-ai-router)
* [EasyOneWeb LLC](https://easyoneweb.ru)

# Copyright

EasyOneWeb LLC 2020 - 2025. All rights reserved. Code author: Kirill Makeev. See LICENSE.md for licensing and usage information.

**Enjoy!**