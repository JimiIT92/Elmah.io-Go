<p align="center"><img src="logo.png" width="250"></p>
<p align="center">
  <a href="https://goreportcard.com/report/github.com/jimiit92/elmah.io-go"><img src="https://goreportcard.com/badge/github.com/jimiit92/elmah.io-go"></a>
  <a href="https://godoc.org/github.com/jimiit92/elmah.io-go"><img src="/godoc-elmah.io.svg" alt="GoDoc"></a>
</p>

# Elmah.io-Go
Log errors on elmah.io from Go(lang) web applications.

-------------------------
- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Bugs / Feature Reporting](#bugs--feature-reporting)
- [Feedback](#feedback)
- [Continuous Integration](#continuous-integration)
- [License](#license)
-------------------------
## Introduction

This package <b>IS NOT OFFICIAL</b>, meaning that nor the Elmah.io team or any of their developers have worked on this.
This is just a user made package to start logging errors on Elmah.io, waiting for an official Go(lang) package.
This means that i'm not affiliated with Elmah.io in any means, and therefore any bug or missing feature on the package are not Elmah.io's fault.

Also, this package gives just the raw basics for logging errors on Elmah.io. For my personal use this is fine, however you can suggest changes or even contributing to them as described in the [Feedback](#feedback) section of this file.

-------------------------
## Requirements

- Go 1.10 or higher. Although this package has no dependencies, so it should work on all versions of Go, is always recommended to have the latest version of the language installed.
- An <a href="https://elmah.io/" target="_blank">Elmah.io account</a>
-------------------------
## Installation

The recommended way to get started using the Elmah.io Go package is by using `dep` to install the dependency in your project.

```bash
dep ensure -add "github.com/jimiit92/elmah.io-go@~1.0.0"
```
-------------------------
## Usage

To get started with the package, import the `elmahio` package and use the `Setup` function in your `main` function:
```go
import (
	elmahio "github.com/gimignanof/elmahio-go"
)

err := elmahio.Setup("Your-API-Key", "Elmah.io-Log-Id")
```

You can also setup the Application Version and the Application Source using the `SetupVersion` and `SetupSource` functions:
```go
elmahio.SetVersion(1.0)
elmahio.SetSource("Application Source")
```
To log the erorr on Elmah.io automatically from a web function, you need to declare the function like this:
```go
func handler(w http.ResponseWriter, r *http.Request) (*http.Response, error) {
// Do something and return a response,error pair
}
```
and then serve the function from the `ElmahHandler` wrapped like this (the example uses the <a href="https://github.com/gorilla/mux" target="_blank">gorilla/mux package</a>):
```go
router.Handle("/", elmahio.ElmahHandler(handler))
```
### Example application

<a href="/examples/example.go">Here </a>is listed and example application that uses the Elmah.io package and writes an error on Elmah.io. The project uses the <a href="https://github.com/gorilla/mux" target="_blank">gorilla/mux package</a> to create a rotuer and handle the route, however you are free to use any router package you want (or even none) as long as the handler function declaration remains the same and you can wrap it inside the  `ElmahHandler` middleware

-------------------------
## Bugs / Feature Reporting

New Features and bugs can be reported on the <a href="https://github.com/JimiIT92/Elmah.io-Go/issues">Issues tab</a>

-------------------------
## Feedback

The Elmah.io package is not feature complete, so any help is appreciated. You can clone and edit this repository to add as many features as you want. Keep in mind that each commit will be reviewed in order to mantain this package healty and to avoid malicious or un-necessary code. See our [contribution guidelines](CONTRIBUTING.md) for details.

-------------------------
## License

The Elmah.io package is licensed under the [MIT License](LICENSE).
