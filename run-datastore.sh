#!/usr/bin/env bash
export DATASTORE_EMULATOR_HOST=localhost:8081 && gcloud beta emulators datastore start --no-store-on-disk