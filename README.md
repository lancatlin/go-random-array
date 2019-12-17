# go-random-array
A random array ganerater written in Go

## Feature
- custom array begin and end
- use crypto/rand for random seed

## Installation
Go to [Release](https://github.com/lancatlin/go-random-array/releases) and download the latest program.

## Using
``` shell
$ ./random-array -b $begin -e $end
```
Examble: from 1 to 32
``` shell
$ ./random-array -b 1 -e 32
Random seed: 3D3C9BE9B8AFC77F
[22 21 12 25 16 29 17 11 26 13 8 4 19 20 23 2 7 1 3 5 18 30 28 15 10 24 31 14 9 6 32 27]
```
P.S. Change the program name according operating system.
