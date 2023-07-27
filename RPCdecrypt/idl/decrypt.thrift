namespace go api

struct Request {
	1: string message
}

struct Response {
	1: string message
}

service decrypt {
    Response decrypt(1: Request req)
}