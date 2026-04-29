#!/bin/bash

export GOOGLE_APPLICATION_CREDENTIALS=~/Downloads/key.json
export PROJECT=$(gcloud config get-value project)

go run main.go