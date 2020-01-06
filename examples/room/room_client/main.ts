/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-02 19:19:58
 */


import {Context, TCPClient} from "stmp";
import pb from "../room_pb/room.pb";
import stmp from "../room_pb/room.stmp";
import UserServiceClient = stmp.stmp.examples.room.UserServiceClient;
import UserEventsListener = stmp.stmp.examples.room.UserEventsListener;

class UserScene implements UserEventsListener {
    HandleStatusUpdated(ctx: Context, input: pb.stmp.examples.room.UserModel, output: pb.google.protobuf.Empty): void {
    }
}

export async function main() {
    const client = new TCPClient("ws://127.0.0.1:5001/ws");
    const usc = new UserServiceClient(client);
    await usc.ListUser({limit: 20});
    const userScene = new UserScene();
    UserEventsListener.register(client, userScene);
}
