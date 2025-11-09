# landscape-go-client

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

> [!CAUTION]
> See the help text for the CLI by passing `-h` to any of the commands. For example:
>
> ````sh
> ./landscape-api -h
> ````

Now, you can use the CLI tool to call the Landscape API. For example, to create a new V1 (legacy) script:

```sh
./landscape-api script create -title coolscript -code $'#!/bin/bash\nB)\n'
```

...

```json
{
  "id": 21428,
  "access_group": "global",
  "creator": {
    "name": "Jan-Yaeger Dhillon",
    "email": "jan.dhillon@canonical.com",
    "id": 127649
  },
  "title": "coolscript",
  "time_limit": 300,
  "username": "",
  "attachments": [],
  "status": "V1"
}
```

Edit it:

```sh
./landscape-api script edit 21428 -c $'#!/bin/bash\nedited\n' -t newtitle
```

...

```json
{
  "id": 21428,
  "access_group": "global",
  "creator": {
    "name": "Jan-Yaeger Dhillon",
    "email": "jan.dhillon@canonical.com",
    "id": 127649
  },
  "title": "newtitle",
  "time_limit": 300,
  "username": "",
  "attachments": [],
  "status": "V1"
}
```

Create an attachment for it:

```sh
code=$(echo "attachment" | base64)
./landscape-api script attachment create -i 21428 -f "note.txt\$\$$code"
```

...

```text
"note.txt"
```

View the script:

```sh
./landscape-api script get 21428
```

...

```json
{
  "id": 21428,
  "title": "newtitle",
  "creator": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon",
    "email": "jan.dhillon@canonical.com"
  },
  "attachments": [
    "note.txt"
  ],
  "access_group": "global",
  "time_limit": 300,
  "username": "",
  "status": "V1"
}
```

### V2 scripts

You can also create and manage V2 scripts by setting the `-script-type` flag to `V2`:

```sh
./landscape-api script create -title v2script -code $'#!/bin/bash\nv2\n' -script-type V2
```

...

```json
{
  "id": 21431,
  "title": "v2script",
  "version_number": 1,
  "created_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "created_at": "2025-11-09T22:07:37.989440",
  "last_edited_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "last_edited_at": "2025-11-09T22:07:37.989440",
  "script_profiles": [],
  "status": "ACTIVE",
  "attachments": [],
  "code": "v2\n",
  "interpreter": "/bin/bash",
  "access_group": "global",
  "time_limit": 300,
  "username": ""
}
```

Create an attachment for it:

```sh
code=$(echo "v2_attachment" | base64)
./landscape-api script attachment create -i 21431 -f "note.txt\$\$$code"
```

View the script:

```sh
./landscape-api script get 21431
```

...

```json
{
  "id": 21431,
  "title": "v2script",
  "version_number": 1,
  "created_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "created_at": "2025-11-09T22:07:37.989440",
  "last_edited_by": {
    "id": 127649,
    "name": "Jan-Yaeger Dhillon"
  },
  "last_edited_at": "2025-11-09T22:07:37.989440",
  "script_profiles": [],
  "status": "ACTIVE",
  "attachments": [
    {
      "filename": "note.txt",
      "id": 50543
    }
  ],
  "code": "v2\n",
  "interpreter": "/bin/bash",
  "access_group": "global",
  "time_limit": 300,
  "username": "",
  "is_redactable": true,
  "is_editable": true,
  "is_executable": true
}
```
