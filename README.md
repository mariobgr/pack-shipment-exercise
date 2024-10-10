# pack-shipment-api

## Purpose
This is a solution for an exercise that aims to showcase skills in Golang and algorithms by providing an API
that calculates the smallest number of packages of configurable size for a requested number of items

## Installation
In order to install and run, you will need docker and docker-compose (and a text editor to edit .env file)

Make sure your port 8080 is not in use and simply type `docker-compose up -d --build` in your Terminal

## Usage
In order to make requests to http://localhost:8080 you need to make sure to add header `X-Api-Key: Ap1K3y` to authenticate

The API exposes two endpoints:
```
GET / - list available pack sizes

POST /calculate - calculates number of packs required for requested number of items
```
Example requests can be found in [requests.http](requests.http) file

If you want to configure tha packs sizes, edit [.env](.env) and change value for  `CONFIG_PACK_SIZES`

The app has the functionality to reload the env values on a set interval, configurable via `UPDATE_INTERVAL_SECONDS` **on startup!**, defaults to 60 seconds

## Some thoughts on implementation
The application is written using a simplified version of the [hexagonal architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)), to which I'm a big fan.

Ideally the configuration would live in a database repository that provides methods for updating. Initially I was considering using sqlite, however, given the simplistic nature of the task I decided it would be too much.

Instead, I opted for self-updating env-based config, which I believe is a good example for some other Golang techniques.