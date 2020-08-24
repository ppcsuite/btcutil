ppcutil
=======

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)

Package ppcutil provides peercoin-specific convenience functions and types.
A comprehensive suite of tests is provided to ensure proper functionality.  See
`test_coverage.txt` for the gocov coverage report.  Alternatively, if you are
running a POSIX OS, you can run the `cov_report.sh` script for a real-time
report.

This package was developed for ppcd, an alternative full-node implementation of
peercoin.  
Although it was primarily written for ppcd, this package has intentionally been designed so it
can be used as a standalone package for any projects needing the functionality
provided.

## Installation and Updating

```bash
$ go get -u github.com/ppcsuite/ppcutil
```


## License

Package ppcutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
