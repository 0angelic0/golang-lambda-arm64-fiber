# Golang Lambda on Graviton2 ARM64

Ref1: https://aws.amazon.com/blogs/compute/migrating-aws-lambda-functions-to-arm-based-aws-graviton2-processors/
Ref2: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-custom.html

## Points

1. in `serverless.yml` use custom runtime `provided.al2` and architecture `arm64` and set packaging mode as `individually`

```
provider:
  runtime: provided.al2
  architecture: arm64

package:
  individually: true
```

2. in `services/{service}/serverless.yml` that declare a function for each service, the handler binary must be name `bootstrap` and set packaging artifact to the .zip file that will have `bootstrap` binary at root when unzipped.

```
{function_name}:
handler: bootstrap
  package:
    artifact: bin/{service}/bootstrap.zip
```

3. in `Makefile` use `go build` command with valid env `GOARCH=arm64` and tags `lambda.norpc`

```
GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -tags lambda.norpc -o {output_binary} {input_main.go}
```

4. Done! your go code don't need to change anything



# Golang generated go Fiber code from OpenAPI Spec

## Points

1. Develop OpenAPI Spec using split files and `$ref`

2. Bundle OpenAPI Spec using redocly-cli `bundle` command

3. Use `oapi-codegen` tools to generated code with custom template for Fiber (by override Echo template)

```
tools/oapi-codegen-templates/echo/echo-interface.tmpl
tools/oapi-codegen-templates/echo/echo-register.tmpl
tools/oapi-codegen-templates/echo/echo-wrappers.tmpl
```

4. in `Makefile` after generated code via `oapi-codegen`, use `sed` to replace import fiber to v2
 