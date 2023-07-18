# Formatting with prettier

To format the battleship web client code, run the following:

```console
$ docker-compose exec web-client ./node_modules/.bin/prettier --single-quote --write "src/**/*.{js,jsx}"
```

This assumes that `docker-compose up` has been run from the repo root.
