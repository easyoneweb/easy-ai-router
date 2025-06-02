# Change Log

All notable changes to the "easy-ai-router" application will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - ----------

### Added

- Integration with MongoDB.
- Log model to save request's meta data to models providers.
- New env variables: DB_URI and DB_NAME for database connection, OPENROUTER_LIMIT to limit how many requests per day can be made to Openrouter's API.
- Use any Openrouter model with chat and chat with image requests, defaults to deepseek/deepseek-r1:free if model name wasn't provided.

### Changed

- Openrouter chat and chat with image functions now require request identity string to save log data after successful requests to openrouter api.
- Openrouter REST API v1 /chat and /chat/image now use different body for request, refer to README.md. 

## [0.1.0] - 2025-03-14

### Added

- Initial working version.

### Fixed

- Initial working version.

### Changed

- Initial working version.

### Removed

- Initial working version.