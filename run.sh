#! /usr/bin/env bash

# Get to the project root directory
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE"
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
cd -P "$( dirname "$SOURCE" )"

# Constants
ROOT_DIR=$(pwd)
DBAPP_DIR="$ROOT_DIR/dbapp"
LOG_FILE="$ROOT_DIR/test.log"

# Check if dbapp dir exists
if [ ! -d "$DBAPP_DIR" ] ; then
    echo "Project's dbapp directory not found! Exiting."
    exit 1
fi

# Run db commands
cd "$DBAPP_DIR"
make run
if [ $? -ne 0 ]; then
    echo "Make command failed! Exiting."
    exit 1
fi

# Run all test cases
cd "$ROOT_DIR"
go clean -modcache 2>&1 | tee "$LOG_FILE"
go clean -testcache 2>&1 | tee -a "$LOG_FILE"
go test -v ./... 2>&1 | tee -a "$LOG_FILE"

# Clean exit
exit 0
