# apikeys

`apikeys` is a command line tool for working with Google's
[API Keys](https://cloud.google.com/service-infrastructure/docs/overview) API.

The `apikeys` tool uses [Sidecar](https://github.com/agentio/sidecar) and an
authenticating API proxy ([IO](https://agent.io/posts/io)) to make gRPC
requests to the API Keys API. As a result, this repo contains no secrets and
no code that ever possesses them!

## License

This software is released under the [Apache 2 license](/LICENSE).
