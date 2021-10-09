# This is just some quick Go code to get an AWS certificate 

## Build
`go build -o get-aws-cert main.go`

## Use
> NOTE: This application expects AWS configuration to exist in the default location. This location at the time of writing this is `~/.aws/config`.
https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html

```
$ ./get-aws-cert -arn <certificate's arn>
> Getting certificate from ARN: <certificate's arn>
-----BEGIN CERTIFICATE-----
j;lkjFLKJfRLK@#J$uF:LKj@$;l#lkFjEeElkjfP{:ofIjAKL@RTjh@fj@3yFj@h
[...]
-----END CERTIFICATE-----
```
