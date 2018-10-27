include "base.thrift"
const string uri = "internal/apps/test"

service Test {
    base.StatusCode getStatus(1: string arg),
}
