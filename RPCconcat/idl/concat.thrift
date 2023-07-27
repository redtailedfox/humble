namespace go api

struct Request {
	1: string message1
	2: string message2
}

struct Response {
	1: string message
}

service concat {
    Response concat(1: Request req)
}