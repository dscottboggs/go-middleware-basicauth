# BasicAuth middleware for Go APIs
A high-level middleware module for authenticated with basic auth. Plug-and-play in your existing API.

## SSL - IMPORTANT
HTTP basic authentication sends the username and password in clear text. DO NOT
use this in a secure context without SSL as well.

### Usage
see this [gist](https://gist.githubusercontent.com/dscottboggs/e55b1add1fede8cfa515ea288bd51c7e/raw/52d2da55638ab9d9590b60c828c3449156d525b2/middleware.go)
