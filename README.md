# gorocksdb, a Go wrapper for RocksDB

> **Note**  
> Segment has paused maintenance on this project, but may return it to an active status in the future. Issues and pull requests from external contributors are not being considered, although internal contributions may appear from time to time. The project remains available under its open source license for anyone to use.

[![Build Status](https://travis-ci.org/tecbot/gorocksdb.png)](https://travis-ci.org/tecbot/gorocksdb) [![GoDoc](https://godoc.org/github.com/tecbot/gorocksdb?status.png)](http://godoc.org/github.com/tecbot/gorocksdb)

## Install

You'll need to build [RocksDB](https://github.com/facebook/rocksdb) v5.5+ on your machine.

After that, you can install gorocksdb using the following command:

    CGO_CFLAGS="-I/path/to/rocksdb/include" \
    CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
      go get github.com/tecbot/gorocksdb

Please note that this package might upgrade the required RocksDB version at any moment.
Vendoring is thus highly recommended if you require high stability.

*The [embedded CockroachDB RocksDB](https://github.com/cockroachdb/c-rocksdb) is no longer supported in gorocksdb.*
