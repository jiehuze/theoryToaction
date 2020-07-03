namespace go rpc.thrift

service rpcService {        
    void RPCPush(1:string srcIP, 2:list<string> devIds, 3:i64 msgid, 4:string appid, 5:i8 pushType),
}
