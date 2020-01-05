/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-02 19:19:58
 */


import {Context, Server, TCPClient} from "stmp";
import pb from "../room_pb/room.pb";
import stmp from "../room_pb/room.stmp";
import UserServiceServer = stmp.stmp.examples.room.UserServiceServer;
import UserServiceClient = stmp.stmp.examples.room.UserServiceClient;
import UserServiceBroadcaster = stmp.stmp.examples.room.UserServiceBroadcaster;
import UserEventsListener = stmp.stmp.examples.room.UserEventsListener;

class UserService implements UserServiceServer {
    ListUser(ctx: Context, input: pb.stmp.examples.room.ListInput, output: pb.stmp.examples.room.ListUserOutput): void | Promise<void> {
    }

    Login(ctx: Context, input: pb.stmp.examples.room.LoginInput, output: pb.stmp.examples.room.UserModel): void | Promise<void> {
    }
}

class UserScene implements UserEventsListener {
    HandleStatusUpdatedOfUserEvents(ctx: Context, input: pb.stmp.examples.room.UserModel, output: pb.google.protobuf.Empty): void | Promise<void> {
    }
}

export async function main() {
    const srv = new Server();
    const userService = new UserService();
    UserServiceServer.register(srv, userService);
    const client = new TCPClient("ws://127.0.0.1:5001/ws");
    const usc = new UserServiceClient(client);
    const users = await usc.ListUser({limit: 20});
    UserServiceBroadcaster.ListUserToAll({limit: 20}, srv);
    const userScene = new UserScene();
    UserEventsListener.register(client, userScene);
    const users2 = await UserServiceBroadcaster.ListUser({limit: 20}, client);
    console.log(users.total == users2.total);
}