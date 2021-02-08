# Spacelift Pulumi Provider

WARNING: This provider is a work in progress.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/spacelift

or `yarn`:

    $ yarn add @pulumi/spacelift

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_spacelift

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/spacelift-io/pulumi-spacelift/sdk/go/...

## Configuration

The following configuration points are available for the `spacelift` provider:

- `spacelift:apiKeyEndpoint` (environment: `SPACELIFT_API_KEY_ENDPOINT`) - Endpoint to use when authenticating with an API key outside of Spacelift
- `spacelift:apiKeyId` (environment: `SPACELIFT_API_KEY_ID`) - ID of the API key to use when executing outside of Spacelift
- `spacelift:apiKeySecret` (environment: `SPACELIFT_API_KEY_SECRET`) - API key secret to use when executing outside of Spacelift
- `spacelift:apiToken` (environment: `SPACELIFT_API_TOKEN`) - Spacelift token generated by a run, only useful from within Spacelift
