namespace go api

struct Request {
	1: string message
}

struct Response {
	1: string message
}

service encrypt {
    Response encrypt(1: Request req)
}