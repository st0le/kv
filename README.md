# kv

Dumb Program to store simple Key-Value notes on persistent storage. Integrates with your command line pipeline so you can pull/store values at any time.

## Get the code

`go get -u github.com/st0le/kv`

## Build 

`go build`

## Instructions

You can invoke `kv` in 4 modes. Ideally copy the binary into a location that is in your PATH.

1) list all keys

    `kv`

2) get a single key

    `kv <key>`

3) set a value 

    `kv <key> <value>`

4) pipe contents to a key

    `echo test | kv <key>`

The values are stored in files named by the key in `$HOME\.kv`. Of course, keys will be unique. Files are not encrypted, I wouldn't use it for anything sensitive like your password, apikeys, etc 

## Tip for Windows Powershell Users

Use 

`New-Alias -Name kv -Value "kv.exe"`

To avoid typing ".\kv.exe" each time.


