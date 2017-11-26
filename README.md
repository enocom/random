# random

Random is a web application which collects random Wikipedia pages and displays
links to those pages as swatches. It is deployed at [random.ecom.com](http://random.ecom.com)
and is running within a Kubernetes cluster on Google Cloud Platform.

## Development

To run the application locally, use:

```
make run
```

Then, visit `localhost:8080` in a browser. Note, random will poll Wikipedia
every 10 seconds for a random link.

## Docker

Random is also available as a Docker container. To run that container, use:

```
make docker-run
```

For more details, see the Makefile.
