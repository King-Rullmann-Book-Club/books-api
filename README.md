# Books-API

An api for managing books, book clubs, and their members.

## Start the API

Use the following command to start the API:

```bash
go run ./cmd/main.go
```

The API should be running on port `8080`.

Or use the start script:

```bash
chmod +x ./start.sh
./start.sh
```

Use the build script to build a static binary:

```bash
./build.sh
```

## Using Docker

You can build, and then run the container with Docker. Docker compose is used to run the 
docker container with the contents of `.db-data` mounted as a volume, and listening on 
port `8000`. `--build` is used to ensure it rebuilds the container with an up-to-date binary.

```bash
sudo docker compose up -d --build

# use down to shut it down again
sudo docker compose down
```

## Using NixOS

There is a flake setup for this repo. To use it run the following command:

```bash
nix develop
```


