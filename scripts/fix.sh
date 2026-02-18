#!/bin/bash

set -e

HERE=$(dirname $0)
cd $HERE/..

go fix ./...
