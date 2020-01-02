/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-02 19:19:58
 */


import {Context, dialWebSocket, Server} from "stmp";
import pb from "../room_proto/room.pb";
import stmp from "../room_proto/room.stmp.example";
import ListUserInput = pb.stmp.examples.room.ListUserInput;
import ListUserOutput = pb.stmp.examples.room.ListUserOutput;
import UserServiceServer = stmp.stmp.examples.room.UserServiceServer;
import UserServiceClient = stmp.stmp.examples.room.UserServiceClient;
import UserServiceBuilder = stmp.stmp.examples.room.UserServiceBuilder;

class UserService implements UserServiceServer {
    async listUser(ctx: Context, input: ListUserInput, output: ListUserOutput) {
    }
}

export async function main() {
    const srv = new Server();
    const userService = new UserService();
    UserServiceServer.register(srv, userService);
    const conn = await dialWebSocket("ws://127.0.0.1:5001/ws");
    const nc = new UserServiceClient(conn);
    const users = await nc.listUser({limit: 20});
    srv.broadcast(UserServiceBuilder.listUser({limit: 20}));
}