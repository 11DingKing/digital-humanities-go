#!/bin/sh
set -eu
docker build -f benzhi.Dockerfile -t digital-humanities-go:latest .
