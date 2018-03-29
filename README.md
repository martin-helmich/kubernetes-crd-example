# Kubernetes Custom Resource Definition (CRD) Example

This repository contains the example files for my [blog article on defining and accessing Kubernetes Custom Resources](https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html) with the [Kubernetes client-go package](https://github.com/kubernetes/client-go).

## Setup

Building the example:

    $ go get github.com/martin-helmich/kubernetes-crd-example

Setting up a custom resource definition (CRD) with an example object:

    $ kubectl apply -f https://raw.githubusercontent.com/martin-helmich/kubernetes-crd-client/master/kubernetes/crd.yaml
    $ kubectl apply -f https://raw.githubusercontent.com/martin-helmich/kubernetes-crd-client/master/kubernetes/project.yaml