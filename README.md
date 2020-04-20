# slowcat

Slowcat is like "cat", but it's slow. This produces a visual effect you might
find pleasing.

## Installation

### With Go Environment

If you have a working `go` environment, simply run:

```
go get github.com/mreishus/slowcat
```

### Without Go Environment

Download the latest release and move the program inside your path. For
example, on 64 bit Linux with a ~/bin/ directory set up:

```
cd /tmp
wget https://github.com/mreishus/slowcat/releases/download/v0.1.0/slowcat_0.1.0_Linux_x86_64.tar.gz
mv slowcat ~/bin/
```

## Usage

Use `slowcat` just like you would use `cat`.

Read from STDIN:

```
fortune | slowcat
```

Read from file:

```
slowcat /tmp/file1 /tmp/file2
```
