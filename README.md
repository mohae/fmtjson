fmtjson
=======

fmtjson format's json files.

I got tired of dealing with consistently formatting JSON data when hand creating format files, e.g. creating [Packer]*(https://packer.io) templates.  This is a tool to consistently format JSON files.

## Usage

    $ git clone https://github.com/mohae/fmtjson

    $ cd $GOPATH/src/github.com/mohae/fmtjson

    $ go build -O $GOPATH/bin/fmtjson

    $ fmtjson /path/to/json/file

fmtjson can format multiple files, each file to be formatted should be separated by a space:

    $ fmtjson file1.json file2.json file3.json

## Options

Flag | Short | Type | Default | Info  
:-- | :-- | :-- | :-- | :--  
-help | -h | bool | false | fmtjson Help  
-spaces | -s | int |  | number of spaces to use for indents; defaults to /; must be > 0  
