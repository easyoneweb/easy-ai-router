# TODO:

## v0.3.0
- [x] internal package openrouter tests
- [x] refactor openrouter for SetConfig function, use internal global var config
- [x] refactor openrouter external request to check for response status code, which should be 200
- [x] setup services in their own function in main package

## v0.2.0
- [x] DB, MongoDB
- [x] Track openrouter limits with DB logs and env variable
- [x] Provide current limit via REST API endpoint
- [x] Block openrouter request if limit exceeded through middleware check
- [x] Use any Openrouter model