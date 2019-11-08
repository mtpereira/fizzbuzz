# FizzBuzz microservice

## Local deploy for testing

- Start minikube: `minikube start`
- Setup ingress on it: `minikube addons enable ingress`
- Setup Helm on it: `helm init`
- Install the microservice: `helm install chart/fizzbuzz`
- Add the result from `minikube ip` to the local `/etc/hosts/` file, in the form of: `192.168.99.102 minikube.local`
- Access it: `curl -POST -d '{"n": 1}' minikube.local/fizzbuzz/v1/single`
