// idl/hello.thrift
namespace go hello

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}

struct OtherReq {
    1: string Other (api.body="other");
}

struct OtherResp {
    1: string Resp;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
    OtherResp OtherMethod(1: OtherReq request) (api.post="/other");
}

service NewService {
    HelloResp NewMethod(1: HelloReq request) (api.get="/new");
}