# gowaka/resources

[![Build Status](https://travis-ci.org/gowaka/resources.svg?branch=master)](https://travis-ci.org/gowaka/resources)
[![GoDoc](https://godoc.org/github.com/gowaka/resources?status.svg)](http://godoc.org/github.com/gowaka/resources)

WakaTime API resources

## Example

Setup an [API client](https://github.com/gowaka/api) with your `api_key`.

    waka, err := api.NewClient("12345")
    if err != nil {
      // handle error
    }

Create a new CurrentUser resource and pass it to the API client.

    cu := resources.CurrentUser{}

    var u data.User
    err = waka.Get(cu, &u)
    if err != nil {
      // handle error
    }

## License

API
