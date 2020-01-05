// Code generated by protoc-gen-stmp. DO NOT EDIT.
// source: examples/room/room_pb/room.proto
import pb from "./room.pb";
import { CallOptions, Connection, ConnFilter, Context, Server, Client } from "stmp";

export default stmp;

declare namespace stmp {
  namespace stmp.examples.room {


    class UserServiceServer {
      static register(srv: Server, inst: UserServiceServer): void
      static unregister(srv: Server, inst: UserServiceServer): void
      ListUser(ctx: Context, input: pb.stmp.examples.room.ListInput, output: pb.stmp.examples.room.ListUserOutput): void | Promise<void>
      Login(ctx: Context, input: pb.stmp.examples.room.LoginInput, output: pb.stmp.examples.room.UserModel): void | Promise<void>
    }

    class UserServiceListener {
      static register(c: Client, inst: UserServiceListener): void
      static unregister(c: Client, inst: UserServiceListener): void
      HandleListUserOfUserService(ctx: Context, input: pb.stmp.examples.room.ListInput, output: pb.stmp.examples.room.ListUserOutput): void | Promise<void>
      HandleLoginOfUserService(ctx: Context, input: pb.stmp.examples.room.LoginInput, output: pb.stmp.examples.room.UserModel): void | Promise<void>
    }

    class UserServiceBroadcaster {
      constructor()
     static ListUser(input: pb.stmp.examples.room.IListInput, conn: Connection, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.ListUserOutput>
     static ListUserToSet(input: pb.stmp.examples.room.IListInput, conns: Set<Connection>, excludes?: Connection[]): void
     static ListUserToAll(input: pb.stmp.examples.room.IListInput, srv: Server, filter?: ConnFilter): void
     static Login(input: pb.stmp.examples.room.ILoginInput, conn: Connection, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.UserModel>
     static LoginToSet(input: pb.stmp.examples.room.ILoginInput, conns: Set<Connection>, excludes?: Connection[]): void
     static LoginToAll(input: pb.stmp.examples.room.ILoginInput, srv: Server, filter?: ConnFilter): void
    }

    class UserServiceClient {
      private client: Client;
      constructor(client: Client)
      ListUser(data: pb.stmp.examples.room.IListInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.ListUserOutput>
      Login(data: pb.stmp.examples.room.ILoginInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.UserModel>
    }

    class UserEventsServer {
      static register(srv: Server, inst: UserEventsServer): void
      static unregister(srv: Server, inst: UserEventsServer): void
      StatusUpdated(ctx: Context, input: pb.stmp.examples.room.UserModel, output: pb.google.protobuf.Empty): void | Promise<void>
    }

    class UserEventsListener {
      static register(c: Client, inst: UserEventsListener): void
      static unregister(c: Client, inst: UserEventsListener): void
      HandleStatusUpdatedOfUserEvents(ctx: Context, input: pb.stmp.examples.room.UserModel, output: pb.google.protobuf.Empty): void | Promise<void>
    }

    class UserEventsBroadcaster {
      constructor()
     static StatusUpdated(input: pb.stmp.examples.room.IUserModel, conn: Connection, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
     static StatusUpdatedToSet(input: pb.stmp.examples.room.IUserModel, conns: Set<Connection>, excludes?: Connection[]): void
     static StatusUpdatedToAll(input: pb.stmp.examples.room.IUserModel, srv: Server, filter?: ConnFilter): void
    }

    class UserEventsClient {
      private client: Client;
      constructor(client: Client)
      StatusUpdated(data: pb.stmp.examples.room.IUserModel, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
    }

    class RoomServiceServer {
      static register(srv: Server, inst: RoomServiceServer): void
      static unregister(srv: Server, inst: RoomServiceServer): void
      CreateRoom(ctx: Context, input: pb.stmp.examples.room.CreateRoomInput, output: pb.stmp.examples.room.RoomModel): void | Promise<void>
      ListRoom(ctx: Context, input: pb.stmp.examples.room.ListInput, output: pb.stmp.examples.room.ListRoomOutput): void | Promise<void>
      JoinRoom(ctx: Context, input: pb.stmp.examples.room.JoinRoomInput, output: pb.stmp.examples.room.RoomModel): void | Promise<void>
      ExitRoom(ctx: Context, input: pb.stmp.examples.room.ExitRoomInput, output: pb.google.protobuf.Empty): void | Promise<void>
      SendMessage(ctx: Context, input: pb.stmp.examples.room.SendMessageInput, output: pb.google.protobuf.Empty): void | Promise<void>
    }

    class RoomServiceListener {
      static register(c: Client, inst: RoomServiceListener): void
      static unregister(c: Client, inst: RoomServiceListener): void
      HandleCreateRoomOfRoomService(ctx: Context, input: pb.stmp.examples.room.CreateRoomInput, output: pb.stmp.examples.room.RoomModel): void | Promise<void>
      HandleListRoomOfRoomService(ctx: Context, input: pb.stmp.examples.room.ListInput, output: pb.stmp.examples.room.ListRoomOutput): void | Promise<void>
      HandleJoinRoomOfRoomService(ctx: Context, input: pb.stmp.examples.room.JoinRoomInput, output: pb.stmp.examples.room.RoomModel): void | Promise<void>
      HandleExitRoomOfRoomService(ctx: Context, input: pb.stmp.examples.room.ExitRoomInput, output: pb.google.protobuf.Empty): void | Promise<void>
      HandleSendMessageOfRoomService(ctx: Context, input: pb.stmp.examples.room.SendMessageInput, output: pb.google.protobuf.Empty): void | Promise<void>
    }

    class RoomServiceBroadcaster {
      constructor()
     static CreateRoom(input: pb.stmp.examples.room.ICreateRoomInput, conn: Connection, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.RoomModel>
     static CreateRoomToSet(input: pb.stmp.examples.room.ICreateRoomInput, conns: Set<Connection>, excludes?: Connection[]): void
     static CreateRoomToAll(input: pb.stmp.examples.room.ICreateRoomInput, srv: Server, filter?: ConnFilter): void
     static ListRoom(input: pb.stmp.examples.room.IListInput, conn: Connection, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.ListRoomOutput>
     static ListRoomToSet(input: pb.stmp.examples.room.IListInput, conns: Set<Connection>, excludes?: Connection[]): void
     static ListRoomToAll(input: pb.stmp.examples.room.IListInput, srv: Server, filter?: ConnFilter): void
     static JoinRoom(input: pb.stmp.examples.room.IJoinRoomInput, conn: Connection, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.RoomModel>
     static JoinRoomToSet(input: pb.stmp.examples.room.IJoinRoomInput, conns: Set<Connection>, excludes?: Connection[]): void
     static JoinRoomToAll(input: pb.stmp.examples.room.IJoinRoomInput, srv: Server, filter?: ConnFilter): void
     static ExitRoom(input: pb.stmp.examples.room.IExitRoomInput, conn: Connection, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
     static ExitRoomToSet(input: pb.stmp.examples.room.IExitRoomInput, conns: Set<Connection>, excludes?: Connection[]): void
     static ExitRoomToAll(input: pb.stmp.examples.room.IExitRoomInput, srv: Server, filter?: ConnFilter): void
     static SendMessage(input: pb.stmp.examples.room.ISendMessageInput, conn: Connection, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
     static SendMessageToSet(input: pb.stmp.examples.room.ISendMessageInput, conns: Set<Connection>, excludes?: Connection[]): void
     static SendMessageToAll(input: pb.stmp.examples.room.ISendMessageInput, srv: Server, filter?: ConnFilter): void
    }

    class RoomServiceClient {
      private client: Client;
      constructor(client: Client)
      CreateRoom(data: pb.stmp.examples.room.ICreateRoomInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.RoomModel>
      ListRoom(data: pb.stmp.examples.room.IListInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.ListRoomOutput>
      JoinRoom(data: pb.stmp.examples.room.IJoinRoomInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.room.RoomModel>
      ExitRoom(data: pb.stmp.examples.room.IExitRoomInput, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
      SendMessage(data: pb.stmp.examples.room.ISendMessageInput, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
    }

    class RoomEventsServer {
      static register(srv: Server, inst: RoomEventsServer): void
      static unregister(srv: Server, inst: RoomEventsServer): void
      UserEnter(ctx: Context, input: pb.stmp.examples.room.UserEnterEvent, output: pb.google.protobuf.Empty): void | Promise<void>
      UserExit(ctx: Context, input: pb.stmp.examples.room.UserExitEvent, output: pb.google.protobuf.Empty): void | Promise<void>
      NewMessage(ctx: Context, input: pb.stmp.examples.room.ChatMessageModel, output: pb.google.protobuf.Empty): void | Promise<void>
    }

    class RoomEventsListener {
      static register(c: Client, inst: RoomEventsListener): void
      static unregister(c: Client, inst: RoomEventsListener): void
      HandleUserEnterOfRoomEvents(ctx: Context, input: pb.stmp.examples.room.UserEnterEvent, output: pb.google.protobuf.Empty): void | Promise<void>
      HandleUserExitOfRoomEvents(ctx: Context, input: pb.stmp.examples.room.UserExitEvent, output: pb.google.protobuf.Empty): void | Promise<void>
      HandleNewMessageOfRoomEvents(ctx: Context, input: pb.stmp.examples.room.ChatMessageModel, output: pb.google.protobuf.Empty): void | Promise<void>
    }

    class RoomEventsBroadcaster {
      constructor()
     static UserEnter(input: pb.stmp.examples.room.IUserEnterEvent, conn: Connection, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
     static UserEnterToSet(input: pb.stmp.examples.room.IUserEnterEvent, conns: Set<Connection>, excludes?: Connection[]): void
     static UserEnterToAll(input: pb.stmp.examples.room.IUserEnterEvent, srv: Server, filter?: ConnFilter): void
     static UserExit(input: pb.stmp.examples.room.IUserExitEvent, conn: Connection, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
     static UserExitToSet(input: pb.stmp.examples.room.IUserExitEvent, conns: Set<Connection>, excludes?: Connection[]): void
     static UserExitToAll(input: pb.stmp.examples.room.IUserExitEvent, srv: Server, filter?: ConnFilter): void
     static NewMessage(input: pb.stmp.examples.room.IChatMessageModel, conn: Connection, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
     static NewMessageToSet(input: pb.stmp.examples.room.IChatMessageModel, conns: Set<Connection>, excludes?: Connection[]): void
     static NewMessageToAll(input: pb.stmp.examples.room.IChatMessageModel, srv: Server, filter?: ConnFilter): void
    }

    class RoomEventsClient {
      private client: Client;
      constructor(client: Client)
      UserEnter(data: pb.stmp.examples.room.IUserEnterEvent, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
      UserExit(data: pb.stmp.examples.room.IUserExitEvent, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
      NewMessage(data: pb.stmp.examples.room.IChatMessageModel, options?: Partial<CallOptions>): Promise<pb.google.protobuf.Empty>
    }
  }
}
