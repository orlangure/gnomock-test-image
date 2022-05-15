# gnomock-test-image

This repository includes a docker image used for
[Gnomock](https://github.com/orlangure/gnomock) tests.

It is a simple HTTP server that responds with "80" to `GET /` on port 80, and
"8080" to `GET /` on port 8080.

When started with environment variable `GNOMOCK_TEST_1`, it replies with its
value to `GET /env1` on port 80.

When started with environment variable `GNOMOCK_TEST_2`, it replies with its
value to `GET /env2` on port 8080.

When started with environment variable `GNOMOCK_REQUEST_TARGET`, it performs a
`GET` request to that URL (i.e `http://container-name`) and returns the same
status code it received from the target, every time it receives a new request
to `/request` endpoint.
