
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

### Adding an application

### Using a library in a Go application

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

### Tests for cpp

TODO: Handle this.