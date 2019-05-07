# zeros

Counts number of zero bytes (0x0) present. Reports percentage of the file size that are zeros.

Useful for determining orphaned downloads were close to completion or not. Depends on the files being initialized to zero, of course.

## Usage

```
Usage: zeros [-bytes] [-print] [-skipnc] [-threshold FLOAT] -file FILE

  -bytes
    	print byte counts
  -file string
    	file to parse
  -print
    	print filename in output (default true)
  -skipnc
    	skip non-contiguous bytes (default true)
  -threshold float
    	threshold (0-100%) to print results
```

## Roadmap

* v0.5
  - [x] Count zeros
* v1.0
  - [x] flag to skip non-contiguous zeros
  - [x] flag to report filename
  - [x] flag to skip (no output) files with less than X percent zero


