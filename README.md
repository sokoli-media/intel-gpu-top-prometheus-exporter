# intel_gpu_top prometheus exporter

This project is meant to be run as a docker container on an UnRaid server and expose metrics from Intel GPU to 
Prometheus.

## Why would I want to use it?

If you use this tool, you'll be able to export metrics from your Intel GPU to the Prometheus. If you want, you can
add alerts and get notified when something happens to those values. If you just want to monitor them, and you have
e.g. Grafana already installed, you can create a dashboard and stare at them for absolutely no reason!

## Does it work with my UnRaid machine?

The Docker image is prebuilt for linux/amd64 machines.

## How to install it on my UnRaid machine?

1. Go to your UnRaid instance
2. Go to `Docker` in the top menu
3. Click `Add container` button on the bottom of the page
4. Fill in the data:
   1. Use `intel-gpu-top-prometheus-exporter` as a name (or anything else you wish)
   2. Use `docker.io/maciejplonski/intel-gpu-top-prometheus-exporter:latest` as a repository
   3. Check `Privileged` (it's required so that we can access GPU's metrics)
   4. Click on `Add another Path, Port, Variable, Label or Device`
      1. Choose `Device` as a `Config Type`
      2. Use `/dev/dri` as a `Name` and `Value`
         - make sure that this directory exists first
      3. Click `Add`
   5. Click on `Add another Path, Port, Variable, Label or Device`
      1. Choose `Port` as a `Config Type`
      2. Use `http` as a `Name`
      3. Use `9000` as a `Container Port`
      4. Click `Add`
   6. Use any port number you want for the `http` port added in the previous point
       * remember to make sure this port is not used by any other service
   7. Click `Apply`
5. Your container should be started right now!
    - You can go to `http://YOUR-UNRAID-IP:9876` to double check if it works (change `9876` to the port number you chose)
6. Once you make sure it works, add it as a another targer to your Prometheus.
7. Once it's added, try finding metrics that start with `intel_gpu_top_` and use them as you wish :)

## Why it's not an app in UnRaid?

It's a simple project I created for myself. If there'll be more people trying to use it,
I'll try to add it to the UnRaid community apps directory.

## I found a bug! / I want to contribute!

Feel free to add a GitHub issue if you found something not to work as expected.

Also, feel free to submit a pull request with a change if you can fix it yourself :)
