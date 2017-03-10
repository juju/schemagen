# schemagen

Generates a JSON schema for the Juju API.

## Installation

To install, you will first need [go][] and [godeps][], then you can build
and install it with:

    go get github.com/juju/schemagen
    cd $GOPATH/src/github.com/juju/schemagen
    godeps -u dependencies.tsv
    go install

## Usage

Just run with no arguments, optionally redirecting the output to a file:

    schemagen > schemas.json

## Building for a New Juju Revision

Check out the specific revision of Juju that you wish to build agianst,
update the schemagen dependencies to match that, tag schemagen for that
version, then build and install the new version of schemagen:

    cd $GOPATH/src/github.com/juju/juju
    git checkout 2.1
    cd $GOPATH/src/github.com/juju/schemagen
    godeps > dependencies.tsv
    git ci -am 'Updated to 2.1'
    git tag 2.1
    git push --tags
    go install


[go]: https://golang.org/doc/install
[godeps]: https://github.com/rogpeppe/godeps
