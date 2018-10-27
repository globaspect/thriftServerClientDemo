const string uri = "/tmp/rpc"

enum StatusCode {
    OK,
    NOK
}

const string INVALID_RETURN_TYPE_MSG = "Invalid return type: expected \"%s\", got \"%s\""

enum InternalErrorCode {
    GENERIC_ERROR;
    INVALID_RETURN_TYPE;
}

exception ServiceCallUnavailable {
    1: optional string msg;
}

exception InternalError {
    1: InternalErrorCode error_code;
    2: string message;
}

typedef i32 timestamp