# landscape-go-api-client

## Generating from the OpenAPI spec

This project uses [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) to generate the core API client from the [Landscape API OpenAPI spec](https://github.com/jansdhillon/landscape-openapi-spec). Update the generated code with:

```sh
cd client && go generate ./...
```

## Using the CLI tool

> [!CAUTION]
> I mainly created the CLI tool for manually testing the API client and it's missing a lot. Ideally this could be generated as well.

This repository also contains a CLI wrapper around the generated code. First, build the CLI tool:

```sh
go build ./cmd/landscape-api
```

You must provide the base URL of a Landscape Server instance as well as both an API access key and API secret key for it as flags or as environment variables. For example:

```sh
export LANDSCAPE_BASE_URL="https://landscape.canonical.com"
export LANDSCAPE_ACCESS_KEY="XXXXX"
export LANDSCAPE_SECRET_KEY="YYYYY"
```

Alternatively, you can provide an email/password combination (and, optionally, an account to log into) if the Landscape instance accepts password-based authentication:

```sh
export LANDSCAPE_BASE_URL="https://landscape.example.com"
export LANDSCAPE_EMAIL="jan@example.com"
export LANDSCAPE_PASSWORD="5uper5ecurep@ssword!"
export LANDSCAPE_ACCOUNT="example-org"
```

If set, these values will be used to attempt to log into Landscape, instead of the access key/secret key pair.

> [!TIP]
> See the help text for the CLI by passing `-h` to any of the commands. For example:
>
> ````sh
> ./landscape-api -h
> ````

Now, you can use the CLI tool to call the Landscape API. For example, to create a new V2 (versioned, with status) script:

```sh
./landscape-api script create -title coolscript -code $'#!/bin/bash\nB)' -script-type V2
```

...

```json
{
  "id": 21433,
  "title": "coolscript",
  "version_number": 1,
  "created_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "created_at": "2025-11-10T02:56:25.366742",
  "last_edited_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "last_edited_at": "2025-11-10T02:56:25.366742",
  "script_profiles": [],
  "status": "ACTIVE",
  "attachments": [],
  "code": "B)",
  "interpreter": "/bin/bash",
  "access_group": "global",
  "time_limit": 300,
  "username": ""
}
```

Edit it:

```sh
./landscape-api script edit 21433 -c $'#!/bin/bash\nBo)' -t coolerscript
```

...

```json
{
  "id": 21433,
  "title": "coolerscript",
  "version_number": 2,
  "created_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "created_at": "2025-11-10T02:56:25.366742",
  "last_edited_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "last_edited_at": "2025-11-10T02:57:57.842792",
  "script_profiles": [],
  "status": "ACTIVE",
  "attachments": [],
  "code": "Bo)",
  "interpreter": "/bin/bash",
  "access_group": "global",
  "time_limit": 300,
  "username": ""
}
```

Create an attachment for it:

```sh
code=$(echo "attachment" | base64)
./landscape-api script attachment create -i 21433 -f "attachment.txt\$\$$code"
```

...

```text
"attachment.txt"
```

View the script:

```sh
./landscape-api script get 21433
```

...

```json
{
  "id": 21433,
  "title": "coolerscript",
  "version_number": 2,
  "created_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "created_at": "2025-11-10T02:56:25.366742",
  "last_edited_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "last_edited_at": "2025-11-10T02:57:57.842792",
  "script_profiles": [],
  "status": "ACTIVE",
  "attachments": [
    {
      "filename": "attachment.txt",
      "id": 50544
    }
  ],
  "code": "Bo)",
  "interpreter": "/bin/bash",
  "access_group": "global",
  "time_limit": 300,
  "username": "",
  "is_redactable": true,
  "is_editable": true,
  "is_executable": true
}
```

### V1 (legacy) scripts

You can also create and manage V1 scripts (i.e., those shown in the legacy UI) by omitting the `-script-type`:

```sh
./landscape-api script create -title legacy-script -code $'#!/bin/bash\nlegacy script'
```

...

```json
{
  "id": 21434,
  "access_group": "global",
  "creator": {
    "name": "Jan-Yaeger Dhillon",
    "email": "jan.dhillon@canonical.com",
    "id": 127649
  },
  "title": "legacy-script",
  "time_limit": 300,
  "username": "",
  "attachments": [],
  "status": "V1"
}
```

Create an attachment for it:

```sh
code=$(echo "legacy attachment" | base64)
./landscape-api script attachment create -i 21434 -f "legacy-attachment.txt\$\$$code"
```

...

```text
"legacy-attachment.txt"
```

View the script:

```sh
./landscape-api script get 21434
```

...

```json
{
  "id": 21434,
  "title": "legacy-script",
  "creator": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon",
    "email": "jan.dhillon@canonical.com"
  },
  "attachments": [
    "attachment.txt",
    "legacy-attachment.txt"
  ],
  "access_group": "global",
  "time_limit": 300,
  "username": "",
  "status": "V1"
}
```
