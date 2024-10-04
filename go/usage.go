package aitokenizer

const Usage = `
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

`
