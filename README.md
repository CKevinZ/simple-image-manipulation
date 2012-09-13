# Simple image manipulation

/!\ THIS PROJECT IS ABSOLUTELY NOT MATURE /!\

An image manipulation binary, written in go.
I use it in a personal ruby on rails app.

# Installation

```bash
$ git clone git://github.com/CKevinZ/simple-image-manipulation.git
$ cd simple-image-manipulation
$ go build
```

# Usage

```bash
$ ./img -info picture.jpg # Prints jsonp formatted information about picture.jpg
$ ./img -resize picture.jpg -geo '300 200' -out picture.out.jpg
$ ./img -crop picture.jpg -geo '100 200 150 250' -out picture.out.jpg
```
