This repository has forked of the [filebrowser](https://github.com/filebrowser/filebrowser) project to impart file management functions for any other project. There is a git submodule which has client-side UI codes in the frontend directory. 

If you want to change something in UI, you can do it in the [forked repository](https://github.com/zenkronn/filemanager-ui). After that, you have to build a new binary in there to have updated UI.

An example of the fresh build;

    git clone https://github.com/zenkronn/filemanager.git
    cd filemanager
    git submodule update --init frontend
    cd frontend
    yarn install
    yarn build
    cd ..
    cd http
    rice embed-go
    cd ..
    go build -o filemanager
    ./filemanager config init
    ./filemanager users add admin admin --lockPassword=true --perm.admin=false --perm.share=false --perm.execute=false --viewMode=mosaic
    ./filemanager config set --port=9090 --root=/tmp
    ./filemanager config set --branding.disableExternal=true --branding.name='File Manager'
    ./filemanager config set --auth.method=platform --auth.endpoint=<PLATFORM_URL>
    ./filemanager
    open http://localhost:9090


---

<p align="center">
  <img src="https://raw.githubusercontent.com/filebrowser/logo/master/banner.png" width="550"/>
</p>

![Preview](https://user-images.githubusercontent.com/5447088/50716739-ebd26700-107a-11e9-9817-14230c53efd2.gif)

[![Travis](https://img.shields.io/travis/com/filebrowser/filebrowser.svg?style=flat-square)](https://travis-ci.com/filebrowser/filebrowser)
[![Go Report Card](https://goreportcard.com/badge/github.com/filebrowser/filebrowser?style=flat-square)](https://goreportcard.com/report/github.com/filebrowser/filebrowser)
[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/filebrowser/filebrowser)
[![Version](https://img.shields.io/github/release/filebrowser/filebrowser.svg?style=flat-square)](https://github.com/filebrowser/filebrowser/releases/latest)
[![Chat IRC](https://img.shields.io/badge/freenode-%23filebrowser-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23filebrowser)

> ℹ INFO: **This project is not under active development ATM. A small group of developers keeps the project alive, but due to lack of time, we can't continue adding new features or doing deep changes. Please read [#532](https://github.com/filebrowser/filebrowser/issues/532) for more info!**

filebrowser provides a file managing interface within a specified directory and it can be used to upload, delete, preview, rename and edit your files. It allows the creation of multiple users and each user can have its own directory. It can be used as a standalone app or as a middleware.

## Features

Please refer to our docs at [docs.filebrowser.xyz/features](https://docs.filebrowser.xyz/features)

## Install

Please refer to our docs at [docs.filebrowser.xyz](https://docs.filebrowser.xyz/).

## Usage

Please refer to our docs at [docs.filebrowser.xyz/usage](https://docs.filebrowser.xyz/usage).

## Contributing

Please refer to our docs at [docs.filebrowser.xyz/contributing](https://docs.filebrowser.xyz/contributing).
