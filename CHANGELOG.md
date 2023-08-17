## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [1.3.0] - 2022-04-18
### Added
- The `Config` type which is used to configure the email service.

### Updated
- The `Service` type to include a `Config` field.

## [1.2.1] - 2022-04-17
### Changed
- The `SendMail` function needed a path to be declared, fixed to send to 
route `/send` of the email microservice. 

## [1.2.0] - 2022-04-11
### Added 
- `SendMail` function which communicates with the email service.
- Tests for `SendMail`.

## [1.1.0] - 2022-04-06
### Added 
- Added the type `Service` which is actually the `msp.Service` type from the
`github.com/johannesscr/micro` package. To be able to make it a mock-able
service.

## [1.0.1] - 2022-04-05
### Added 
- A method `Message.Validate` to validate that the basic requirements are
met to be able to send a message. With a comments on how the body should
preferably look as an HTML document.

## [1.0.0] - 2022-03-22
### Added
- Changed all the email address fields from `string` to `mail.Address` to better
align with built-in Go features.

## [Released]
## [0.0.0] - 2022-03-21
### Added
- basic types `Message` and `Header`