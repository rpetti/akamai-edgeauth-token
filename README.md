# akamai-edgeauth-token

This is a partial implementation of Akamai's HMAC token auth ([Akamai Docs](https://learn.akamai.com/en-us/webhelp/adaptive-media-delivery/adaptive-media-delivery-implementation-guide/GUID-041AEFDE-7E25-4AD8-B6C4-73F1B7200F02.html))

It's currently untested, and only supports URL tokens.

## Example

```go
tokenGenerator := EdgeAuthToken {
    Key: "adf765d7854adfdf", //Shared Key as a hex string (required)
    WindowSeconds: 300, //Number of seconds the token should be valid (default 300)
    ClientIP: "127.127.127.127", //IP address of the client that will be using the token (optional)
}

token, err := tokenGenerator.GenerateURLToken("/path/to/my/asset")
//token can then be added to the final URL to be sent to the client
// eg: "https://akamaidomain/path/to/my/asset?__token__=<token>"
// where <token> is the value obtained above
```
