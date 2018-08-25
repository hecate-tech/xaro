# Xaro

A massively multiplayer online role playing game.

## Getting Started

These instructions are meant for if you'd like to build the game on your local machine for development purposes.

### Prerequisites

You'll need to install [engo](https://engo.io/) and [mage](https://github.com/magefile/mage) in order to build this project. Engo wasn't included into the magefile because of dependency issues with glfw3 so you'll have to install it manually.

### Installing

Install the repository directly into your gopath using this command.

```
go get -v github.com/damienfamed75/engo-xaro
```

Once you have the two dependencies installed then you can direct your terminal to the root directory of the project and with your gopath bin folder set as an environment variable for `PATH` you'll be able to use the `mage` command.

```
mage
```

The mage command will run through the magefile and install the application and place the executable in the project's `/bin` folder. Then all you have to do is run the bash script to run the file.

```
./runXaro.sh
```

If you want to run a server then you'll have to use the `runServer.sh` bash script instead and that will build the server if it's not built yet as well.

```
./runServer.sh
```

If you're interested in changing the `.proto` files and rebuilding them then you have to have [protobuf](https://github.com/golang/protobuf) installed on your machine by Google.

Afterwards you'll be able to edit the `.proto` files and rebuild them by running the shell script located in the root directory.

```
./generateProto.sh
```

## Authors

* **Damien Stamates** - Lead Developer/Graphic Designer

## License

This project is licensed under the [Apache 2.0 License](LICENSE)

## Acknowledgments

* [u2i superstellar](https://github.com/u2i/superstellar)
    * I took inspiration and code snippets from this open-source multiplayer game written in golang.
* [SolarLune](https://github.com/SolarLune)
    * Taking inspiration from SolarLune's creativity and videos on game development and design.
