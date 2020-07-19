#!/bin/bash

buildifier -r .
go fmt ./...
