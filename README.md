
# template-go-cpp

A template to create a simple go project using cpp library. The structure in this repository is meant to be easily extensible and to allow quick-start of a go application with internal packages and configuration.

# Installation

- Clone the repo: `git clone git@github.com:Knoblauchpilze/template-go-cpp.git`.
- Install Go from [here](https://go.dev/doc/install). **NOTE**: this project expects Go 1.20 to be available on the system.
- Install Eigen from [here](http://eigen.tuxfamily.org/index.php?title=Main_Page#Download): e.g. just follow what's in the `INSTALL` file.
- Go to the project's directory `cd ~/path/to/the/repo`.
- Compile and install: `make`.
- Execute any application with `make run app_name`.

# General principle

This repository is divided into multiple sub-folders and follows the structure provided by [this](https://github.com/Knoblauchpilze/template-go) project where support for cpp libraries and apps have been added.

This is the list of the folders that we kept:
* [assets](https://github.com/Knoblauchpilze/template-go/tree/master/assets): holds all the resources for the project (images, logo, etc.).
* [bin](https://github.com/Knoblauchpilze/template-go/tree/master/bin): the executables.
* [cmd](https://github.com/Knoblauchpilze/template-go/tree/master/cmd): the list of applications defined by the project.
* [configs](https://github.com/Knoblauchpilze/template-go/tree/master/configs): configuration file templates or defaults.
* [internal](https://github.com/Knoblauchpilze/template-go/tree/master/internal): packages which we don't want to publish or make available to the public but are still needed by this project.
* [pkg](https://github.com/Knoblauchpilze/template-go/tree/master/pkg): the public packages which are defined by this project and can be reused.
* [scripts](https://github.com/Knoblauchpilze/template-go/tree/master/scripts): the list of scripts/tools used by this project.
* [test](https://github.com/Knoblauchpilze/template-go/tree/master/test): where tests are stored.

# What's already there

This project already defines some basic structure allowing applications and libraries defined in it to properly work together and be used in coordination.

## The root Makefile

The root [Makefile](Makefile) defines some targets to setup, build and clean the whole repository at once. This is very convenient in order to start from scratch easily and obtain a fully functional set of applications.

It also defines a `test` target, which helps to run the tests defined for this project.

Finally a convenience `run` target allows to execute an application defined in the repository in a controlled environment and with a setup of the necessary shared libraries search pathes needed. This can be accessed through `make run name-of-the-app-to-execute`. This is only available after a successful `make` command has been performed already.

## What to commit and not commit

The principle followed by this repository is that binaries and compiled libraries shouldn't be committed, along with the build resources: this is something that can be deduced from the source and which is best locally generated. This means that:
* the content of [bin](bin) is usually to be excluded from the versioning.
* the content of [cpp/bin](cpp/bin), [cpp/include](cpp/include) and [cpp/lib](cpp/lib) as well.

# How to extend this project

## Expand the Go projects

This is very similar to what is described in the original repository for a [template-go program](https://github.com/Knoblauchpilze/template-go#how-to-extend-this-project): the [internal](internal) folder is meant to receive new private package implementations, the [pkg](pkg) folder is meant to receive additional public packages and the [cmd](cmd) folder should receive the new applications.

## Expand the cpp framework

### Structure

The cpp files are located in the [cpp](cpp) folder. It contains the following sections:
* [bin](https://github.com/Knoblauchpilze/template-go-cpp/tree/master/cpp/bin): contains the executable of the cpp apps.
* [include](https://github.com/Knoblauchpilze/template-go-cpp/tree/master/cpp/include): this contains the 'published' public includes of the cpp library.
* [lib](https://github.com/Knoblauchpilze/template-go-cpp/tree/master/cpp/lib): contains the published version of the cpp libraries.
* [src](https://github.com/Knoblauchpilze/template-go-cpp/tree/master/cpp/src): contains the source code for cpp libraries and applications.

### Adding a library

An example library is provided in the [toy-lib](cpp/src/toy-lib) folder. Adding a new library can be done by copying this folder and renaming it as needed. The places where some renaming should happen are:
* the [Makefile](cpp/src/toy-lib/Makefile).
* the [CMakeLists.txt](cpp/src/toy-lib/CMakeLists.txt).
* the root cpp [Makefile](cpp/Makefile).

The existing `Makefile` already provides targets to build, clean and install the library in a folder which where they can be found by other cpp applications and libraries. The root `Makefile` can in turn package them automatically to be accessed by Go applications.

### Adding an application

An example application is provided in the [toy-app](cpp/src/toy-app) folder. Copying it to a new folder and renaming some properties is the easiest way to create a new application. The places where some renaming should happen are:
* the [Makefile](cpp/src/toy-app/Makefile).
* the [CMakeLists.txt](cpp/src/toy-app/CMakeLists.txt).
* the root cpp [Makefile](cpp/Makefile).

The existing `Makefile` already provides targets to build, clean and install the application using pre-existing cpp libraries to a place where the root `Makefile` can then find them and package them to be used by Go application can automatically access it if needed. It also already provides a neat environment so that it can be executed from the root of this project using the [make run app_name](#the-root-makefile) syntax.

### Using a cpp library in a Go application

When a cpp library is published (i.e. available in [cpp/lib](cpp/lib)), it is possible to use it in Go packages. This repository already provides an example in an [example](pkg/example/printer.go) package. It generally looks like this:

```go
// #include <stdlib.h>
// #cgo CFLAGS: -I../../cpp/include
// #include <toy-lib/lib_c_wrapper.h>
// #cgo LDFLAGS: -L../../cpp/lib -ltoy-lib
import "C"

func exampleFuncUsingCppLib() {
	C.lib_c_wrapper_say_hello()
}
```

The idea is that we need to include the header file with `<lib_name/wrapper_header_file.h>` and then set up the compiler/linker flags. Assuming that the library has been [published](#adding-a-library) it works out of the box.

## Tests

### Tests for Go

This section is very similar to the one provided in the [template-go](https://github.com/Knoblauchpilze/template-go#tests) testing section. This folder reuses it but should work properly to also link the shared cpp libraries to the testing framework so that it's possible to also test features which uses them.

In general just run the following command to run all tests:
```bash
make test
```