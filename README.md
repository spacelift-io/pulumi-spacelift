# Spacelift Pulumi Provider

WARNING: This provider is a work in progress.

## Installing

This package is available in many languages in the standard packaging formats.

In order to run a project locally using this package, you have to install the Spacelift provider binary. This is done automatically on Spacelift workers.
```
pulumi plugin install --reinstall --server https://downloads.spacelift.io/pulumi-plugins resource spacelift 0.0.0
```

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```
npm install @spacelift-io/pulumi-spacelift
```

or `yarn`:

```
yarn add @spacelift-io/pulumi-spacelift
```

However, since we're hosting the packages on GitHub packages, you'll need to add an `.npmrc` file (in your project or user home directory) with the following contents:
```
@spacelift-io:registry=https://npm.pkg.github.com
//npm.pkg.github.com/:_authToken=${GITHUB_TOKEN}
```

where ${GITHUB_TOKEN} is a GitHub access token with the Read capability on GitHub packages.

### Go

To use from Go, use `go get` to grab the latest version of the library:

```
go get -u github.com/spacelift-io/pulumi-spacelift
```

The import paths will then start with `github.com/spacelift-io/pulumi-spacelift/sdk/go/`.

### Python

To use from Python, install using `pip`:

```
pip install "git+https://github.com/spacelift-io/pulumi-spacelift#egg=pulumi_spacelift&subdirectory=sdk/python/bin"
```

or:

```
pip install "git+https://github.com/spacelift-io/pulumi-spacelift@v1.2.0#egg=pulumi_spacelift&subdirectory=sdk/python/bin"
```

to install a specific version, v1.2.0 in this case.

### .NET Core (C#/F#)

To use from .NET Core, install using
```
dotnet add *.csproj package Pulumi.Spacelift
```

However, since we're hosting the packages on GitHub packages, you'll need to add a `nuget.config` file with the following contents:
```xml
<configuration>
    <packageSources>
        <add key="github" value="https://nuget.pkg.github.com/spacelift-io/index.json" />
    </packageSources>
    <packageSourceCredentials>
        <github>
            <add key="Username" value="${GITHUB_USER}" />
            <add key="ClearTextPassword" value="${GITHUB_TOKEN}" />
        </github>
    </packageSourceCredentials>
</configuration>
```

where ${GITHUB_USER} and ${GITHUB_TOKEN} is respectively a GitHub username and access token with the Read capability on GitHub packages.

## Configuration

The following configuration points are available for the `spacelift` provider:

- `spacelift:apiKeyEndpoint` (environment: `SPACELIFT_API_KEY_ENDPOINT`) - Endpoint to use when authenticating with an API key outside of Spacelift
- `spacelift:apiKeyId` (environment: `SPACELIFT_API_KEY_ID`) - ID of the API key to use when executing outside of Spacelift
- `spacelift:apiKeySecret` (environment: `SPACELIFT_API_KEY_SECRET`) - API key secret to use when executing outside of Spacelift
- `spacelift:apiToken` (environment: `SPACELIFT_API_TOKEN`) - Spacelift token generated by a run, only useful from within Spacelift
