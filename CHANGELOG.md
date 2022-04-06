## [Unreleased]
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