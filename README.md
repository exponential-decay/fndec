# fndec

Created for Artefactual Systems Inc.

Guess filename encodings... based on @richardlehane's [characterize][char]
golang package.

## Get up and running

Install dependency:  
`go get github.com/richardlehane/characterize`  

Run release script:  
`bash fndec.sh`  

Run command (or create alias):  
`path/to/release/your-release-version -file path/to/filename`  


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
