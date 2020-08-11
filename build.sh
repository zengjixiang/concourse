#!/bin/bash

## From inside /concourse where your local code is

docker run -v $(pwd):/concourse -w /concourse --entrypoint /bin/bash --privileged -it concourse/dev

apt-get update
apt-get install wget
wget https://github.com/concourse/concourse/releases/download/v6.4.0/fly-6.4.0-linux-amd64.tgz
tar -xvf fly-6.4.0-linux-amd64.tgz
chmod +x fly && mv fly /usr/bin/fly

git clone https://github.com/concourse/concourse-docker.git
git clone https://github.com/concourse/ci.git

wget https://github.com/concourse/registry-image-resource/releases/download/v0.11.1/registry-image-resource-0.11.1-ubuntu.tgz
mkdir resource-types
mkdir resource-types/registry-image
tar -C resource-types/registry-image/ -xvf registry-image-resource-0.11.1-ubuntu.tgz

fly -t ci execute -c ci/tasks/concourse-build-linux.yml \
--input=concourse=concourse --include-ignored \
--input=resource-types=resource-types \
--inputs-from=concourse/build-concourse --output=concourse-tarball=concourse-docker/linux-rc

cd concourse-docker

### change second last line to say COPY entrypoint.sh ….

docker build -t concourse/[YOUR_IMAGE_NAME] .

docker push concourse/[YOUR_IMAGE_NAME]
