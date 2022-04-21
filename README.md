# JSON Converter

*This is a proof of concept. Use at your own risk.*

It is common for workloads initiated by events or messages to receive JSON 
payloads, which are then parsed and used to call other services/APIs.

Convert a JSON input to a format defined by a template.

### Usage

```shell
$ ./json-transform-poc --help
Usage of ./json-transform-poc:
  -json-path string
        json file path (default "./sample.json")
  -template-path string
        template file path (default "./default.template")
```

`--json-path` is the path to the source JSON file.

`--template-path` is the path to the template file.

It is recommended that you save multiple template files, 
and vary the one used when deploying.