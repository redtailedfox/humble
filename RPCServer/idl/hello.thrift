namespace go api

struct Request {
	1: string message
}

struct Response {
	1: string message
}

service call {
    Response call(1: Request req)
}
