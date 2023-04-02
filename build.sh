#!/bin/bash

IMAGE_NAME="limeapi"
IMAGE_TAG="latest"

docker build -t $IMAGE_NAME:$IMAGE_TAG .
