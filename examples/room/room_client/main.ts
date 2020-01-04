/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-02 19:19:58
 */


import {Context, Server, TCPClient} from "stmp";
import pb from "../room_proto/room.pb";
import ListUserInput = pb.stmp.examples.room.ListUserInput;
import ListUserOutput = pb.stmp.examples.room.ListUserOutput;
import stmp from "../room_proto/room.stmp";
import UserServiceServer = stmp.stmp.examples.room.UserServiceServer;
import UserServiceClient = stmp.stmp.examples.room.UserServiceClient;
import UserServiceBroadcaster = stmp.stmp.examples.room.UserServiceBroadcaster;

class UserService implements UserServiceServer {
    async ListUser(ctx: Context, input: ListUserInput, output: ListUserOutput) {
    }
}

export async function main() {
    const srv = new Server();
    const userService = new UserService();
    UserServiceServer.register(srv, userService);
    const conn = new TCPClient("ws://127.0.0.1:5001/ws");
    const usc = new UserServiceClient(conn);
    const usb = new UserServiceBroadcaster();
    const users = await usc.ListUser({limit: 20});
    usb.ListUserToAll({limit: 20}, srv);
    const users2 = await usb.ListUserToOne({limit: 20}, conn);
    console.log(users.total == users2.total);
}