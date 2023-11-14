# EnvNinja

This project is probably the dumbest but i needed something easier to manage my envs with different scopes so it is tailored to my specific needs only.

## Commands

1. [push](#push)
2. [pull](#pull)
3. [query](#query)
4. [list](#list)
5. [rename](#rename)
6. [remove](#remove)

> **NOTE:** Every command accepts -n and -s flag for project's name and scope. They are used to perform the actions on that particular project and scope. The default project the name of the current directory and the default scope is 'dev'.

### push

Push a new secret/s to the db from a env file or key/value pair. Auto creates project/scope if it doesn't already exists.

```
# Push secrets from .env file
envNinja push -p .env

# Push secret from key/value pair
envNinja push -k DB_NAME -v younime
```

### pull

Pull secrets into a file

```
# Pull secrets to .env
envNinja pull -p .env
```

### query

Get the value of secrets

```
# Get value of DB_NAME secret
envNinja query DB_NAME

# Get value of DB_NAME, APP_PORT secrets
envNinja query DB_NAME APP_PORT

# Get values of all secrets
envNinja query --all
```

### list

List projects and scopes

```
# List scopes within the current project
envNinja list

# List all scopes and projects
envNinja list --all
```

### rename

Rename secret's key or scope's name or project's name

```
# Rename secret's key
envNinja -k DB_NAME DBNAME

# Rename scope's name
envNinja -s dev prod

# Rename project's name
envNinja -n younime younime_client
```

### remove

Remove secrets or scope or project

```
# Remove DB_NAME secret
envNinja remove DB_NAME

# Remove DB_NAME, APP_PORT secret
envNinja remove DB_NAME APP_PORT

# Remove scope
envNinja remove -s prod

# Remove project
envNinja remove -n project
```

TODO:

- I don't intend to further improvise this project as it basically has every things that i'll prolly use
