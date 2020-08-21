# get-ip
Simple go program that queries https://www.ipify.org/ to get your external IP address.  

## Install
`go get github.com/MEhrn00/get-ip`

This will install the binary in the `$GOPATH/bin` directory.

Either add `$GOPATH/bin` to your `$PATH` or move the binary to a directory in your path to use it.

You can also download the binary directly from https://github.com/MEhrn00/get-ip/releases

## Usage
Very simple to use:

This will get your external IPv4 address:
```bash
./get-ip
```

To specify IPv6 just pass it with the `-protocol` flag:
```bash
./get-ip -protocol 6
```
