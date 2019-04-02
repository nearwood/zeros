# zeros

Counts number of zero bytes (0x0) present. Reports percentage of the file size that are zeros.

Useful for determining orphaned downloads were close to completion or not. Depends on the files being initialized to zero, of course.

## Usage

`./zeros file.mp4`

## Roadmap

* v0.5
  - [x] Count zeros
* v1.0
  - [ ] flag to skip non-contiguous zeros
  - [ ] flag to report filename
  - [ ] flag to skip (no output) files with less than X percent zero


