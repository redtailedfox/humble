namespace go api

struct Request {
	1: string message (api.query="msg");
}

struct Response {
	1: string message
}

struct concatreq {
    1: string message1 (api.query = "msg1")
    2: string message2 (api.query = "msg2")
}

service thriftCall {
    Response call(1: Request req) (api.post="/post")
}

service concat {
    Response concat(1: concatreq req) (api.post = "/concat")
}
