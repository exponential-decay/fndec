# fndec

Created for Artefactual Systems Inc.

Guess filename encodings... based on @richardlehane's [characterize][char]
golang package.

## Get up and running

Install and update dependencies:  
`go get -u`  

Run release script (if you are creating releases for multiple platforms):  
`bash fndec.sh`  

Otherwise run `go build` or `go install` to create a single binary on your
machine. The binary will be created in the current folder or `$GOBIN`.

Run command (or create alias):  
`path/to/release/your-release-version -file path/to/filename`  

## Binaries

A handful of binaries for different operating systems are created on release 
and are available here: https://github.com/exponential-decay/fndec/releases/latest

## Usage
```
    ross-spencer@artefactual:~:$ fndec -h
    Usage of fndec:
      -file string
    	    File to read the filename from.
      -version
    	    Return version information.

```
## Example output
![alt text][example-one]

[char]: https://github.com/richardlehane/characterize "Characterize Package"
[example-one]: docs/images/example-output.png "Example encoding detection"
