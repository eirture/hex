
# hex

A tiny tool to encode or decode hex data.


## Usage

The hex command only read data from stdin.

```sh
$ hex --help
Usage of hex:
  -d    Decode incoming Hex stream into binary data.

$ echo -n 9d4c8fe7d9ead6125efbc892655f61ba818793c6 | hex -d | base64
nUyP59nq1hJe+8iSZV9huoGHk8Y=

$ base64 -d <<<nUyP59nq1hJe+8iSZV9huoGHk8Y= | hex
9d4c8fe7d9ead6125efbc892655f61ba818793c6
```
