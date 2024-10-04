# aitokenizer - input tokenization microservice for AI workloads

Aitokenizer is a Nulladmin.com microservice which performs input
tokenization for AI workloads. It supports various encodings used by
OpenAI, Anthropic and others.

- Home: [https://nulladmin.com/aitokenizer/](https://nulladmin.com/aitokenizer/)
- Github: [https://github.com/nulladmin/aitokenizer](https://github.com/nulladmin/aitokenizer)
- OpenAPI Spec: [https://nulladmin.com/aitokenizer/openapi/](https://nulladmin.com/aitokenizer/openapi/)

The aitokenizer is primarily designed to be deployed as a stateless
Docker container in production Kubernetes clusters. It includes support
for structured JSON logging, graceful shutdown, health/readiness API,
informative error handling, and standard YAML/ENV configuration.

## Overview

aitokenizer is used to perform BPE input tokenization on text files.
It is fully compatible with the OpenAPI
[tiktoken](https://github.com/openai/tiktoken) encoder.

Support is available for the following BPE AI text encodings as used in various models:

- o200k_base - used by OpenAI gpt-4o-* models
- cl100k_base - used by OpenAI gpt-3-* and gpt-4-* models
- claude - used by Anthropic Claude models

The following are also supported encodings for older deprecated OpenAI models.

- p50k_base
- p50k_edit
- r50k_base
- gpt2

The encoding functions are provided through a well specified [OpenAPI REST
API](https://nulladmin.com/aitokenizer/openapi/)

## Additional features

Additional features can be made available on demand. Some possibilities:

- Various forms of authentication (JWT, Basic, Token, OAuth, Cert)
- Support for other encodings, e.g. the SentancePiece tokenizer used by Google Gemma models.
- OpenTracing support
- Support for Amazon S3 or a database as input and output sources.

Inquire at **contact@nulladmin.com**

## Building and Running the server

To build and run the server, follow these simple steps within the git repository.

While you can skip this step Nulladmin recommends that you always vendor your dependencies:
```
go mod vendor
```

To build and run the server with a default configuration:

```
make build
./aitokenizer
```
The server will be available on `http://localhost:8080`.

### Running in Docker

After building the server above you can generate a Docker container:

```
docker build --network=host -t aitokenizer .
```

Once the image is built you can run it on localhost as follows. 
Note that port 8080 is the default port for the server in the Docker container.

```
docker run --rm -p 8080:8080 -v /tmp:/tmp --hostname localhost -it aitokenizer
```

## Usage

```
$ aitokenizer --help

aitokenizer - a microservice to perform text input tokenization for AI workloads

Usage Examples:

  # Configuration from file
  aitokenizer -c config.yaml

Options:
  -c, --config               Path to the configuration file, yaml format
  -d, --debug                Turn on debug level logging/tracing
  -h, --help                 Print help
  -v, --version              Print version

All configuration values are specified either in the configuration file
or taken from environment variables.
```

## Example

Give a text file `test.txt` with the following content:

```
Sing, O goddess, the anger of Achilles son of Peleus, that brought
countless ills upon the Achaeans. Many a brave soul did it send hurrying
down to Hades, and many a hero did it yield a prey to dogs and vultures,
for so were the counsels of Jove fulfilled from the day on which the son of
Atreus, king of men, and great Achilles, first fell out with one another.
```

We can tokenize it using the `o200k_base` encoding used by OpenAI `gpt-4o` series 
of models:
```
$ curl -i -s -F "filename=@test.txt" http://example.com/v1/encode/o200k_base
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Oct 2024 20:08:01 GMT
Content-Length: 452

[54138,11,532,109209,11,290,32129,328,170018,2391,328,141304,385,11,484,11311,198,5420,2695,220,3678,7557,290,355,50763,616,13,11048,261,53768,16741,2242,480,4952,312,1057,20370,198,4653,316,487,4250,11,326,1991,261,20917,2242,480,14376,261,78166,316,16798,326,165872,1609,412,1938,813,1504,290,2692,199640,328,643,1048,61554,591,290,2163,402,1118,290,2391,328,198,2243,264,385,11,13793,328,1966,11,326,2212,170018,11,1577,18153,842,483,1001,3613,558]
```

## Configuration

The server has a small number of configuration variables which can be specified
either via environment variables or a configuration YAML file.

```
#
# Example Configuration file for aitokenizer showing default values
#

http_server:
  # Server bind address and port
  # address: "" is the default for serving on all interfaces
  address: ""
  port: 8080

  # HTTP read header timeout in seconds, the time from the connection is
  # accepted to when the API request headers are fully read.
  read_header_timeout: 20

  # HTTP read timeout in seconds, the time from when the connection is
  # accepted to when the API request body is fully read
  read_timeout: 60

  # HTTP write timeout in seconds, the time from the end of the request
  # header read to the end of the API response write
  write_timeout: 60

  # HTTP idle timeout in seconds
  idle_timeout: 120

  # Maximum size of request header in bytes, default is 1MB
  max_header_bytes: 1000000

# Enable verbose debug logging, true or false
debug: false
```

All configuration options can be set as environment variables on server startup, for example:

```
DEBUG=true HTTP_SERVER_PORT=8081 aitokenizer
```

## OpenAPI Specification

The full aitokenizer OpenAPI specification can be found here:

[https://nulladmin.com/aitokenizer/openapi/](https://nulladmin.com/aitokenizer/openapi/)

