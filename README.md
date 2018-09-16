# Xaro

A space metroidvania game.

## Getting Started

These instructions are meant for if you'd like to build the game on your local machine for development purposes.

### Prerequisites

You'll need to install [engo](https://engo.io/), [mage](https://github.com/magefile/mage), and [dep](https://github.com/golang/dep) in order to build this project. Engo wasn't included into the magefile because of dependency issues with glfw3 so you'll have to install it manually.

* **engo**
    * `go get -u engo.io/engo`
* **magefile**
    * `go get -u -d github.com/magefile/mage`
* **dep**
    * MacOS
        * `brew install dep`
        * `brew upgrade dep`
    * Other platforms
        * `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`

### Installing

Install the repository directly into your gopath using this command.

```shell
go get -v github.com/hecategames/xaro
```

Once you have the three prerequisites installed then you can direct your terminal to the root directory of the project and with your gopath bin folder set as an environment variable for `PATH` you'll be able to use the `mage` command.

The mage command will run the `magefile.go` file which is what this project uses instead of a `Makefile`. Magefiles were chosen so it was easy to build the project cross platform.

```shell
mage
```

The mage command will clean the repo, rebuild dep, install dependencies, vet, fmt, lint, build, install the application and place the executable in the project's `/bin` folder. Then all you have to do is run the bash script to run the file.

```shell
# windows
./runXaro.sh

# others
sh runXaro.sh
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
