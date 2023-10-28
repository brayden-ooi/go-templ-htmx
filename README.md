## Golang-Templ-HTMX Todo App
This is an exploratory project creating a simple CRUD todo app using `Golang`, `Templ`, and `HTMX`.

1. Go's standard library is used for the back-end to handle requests
2. In this project there is only one exposed route other than the root: `/todo`
   1. `/todo` accepts the 'GET', 'POST', 'PUT', 'DELETE' requests and will respond to them correspondingly
   2. In addition, the client can also modify the request header `'Accept'` to either `text/html` or `application/json` to receive payload in different format
   3. The requests to accept `text/html` and `application/json` are both demonstrated in the client code
3. The client code is created using `templ` and `HTMX`. Requests triggered by `HTMX` will get back `HTML` responses
4. This is an SSR solution that requires minimal JS/TS code and centralises business logic to the back-end. However, JS/TS code is still necessary for more complex front-end use cases (eg. browser APi manipulation).

## How to start
1. `git clone` the repository
2. run `make`
3. run `make run`