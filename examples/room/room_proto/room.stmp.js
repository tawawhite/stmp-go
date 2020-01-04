// Code generated by protoc-gen-stmp. DO NOT EDIT.
// source: examples/room/room_proto/room.proto
const pb = require("./room.pb");
const { PayloadMap, registerMethodAction, notifyOptions } = require("stmp");

const stmp = Object.create(null);

module.exports = stmp;
module.exports.default = module.exports;

function initNamespace(root, ns, factory) {
    for (const item of ns.split(".")) {
        root = (root[item] = root[item] || Object.create(null))
    }
	factory(root)
}

initNamespace(stmp, "stmp.examples.room", (ns) => {


  registerMethodAction("stmp.examples.room.UserService.ListUser", "1001", pb.stmp.examples.room.ListInput, pb.stmp.examples.room.ListUserOutput);
  registerMethodAction("stmp.examples.room.UserService.Login", "1002", pb.stmp.examples.room.LoginInput, pb.stmp.examples.room.UserModel);

  ns.UserServiceServer = class UserServiceServer {
    static register(srv, inst) {
      srv.register(inst, "stmp.examples.room.UserService.ListUser", inst.ListUser);
      srv.register(inst, "stmp.examples.room.UserService.Login", inst.Login);
    }

    static unregister(srv, inst) {
      srv.unregister(inst, "stmp.examples.room.UserService.ListUser");
      srv.unregister(inst, "stmp.examples.room.UserService.Login");
    }

    ListUser(ctx, input, output) { throw new Error("not implemented") }
    Login(ctx, input, output) { throw new Error("not implemented") }
  };
  
  ns.UserServiceBroadcaster = class UserServiceBroadcaster {
    ListUserToOne(input, conn, options) { return conn.invoke("stmp.examples.room.UserService.ListUser", input, options) }
    ListUserToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.UserService.ListUser", pm.get(conn), notifyOptions) }
    ListUserToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.UserService.ListUser", input, filter) }
    LoginToOne(input, conn, options) { return conn.invoke("stmp.examples.room.UserService.Login", input, options) }
    LoginToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.UserService.Login", pm.get(conn), notifyOptions) }
    LoginToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.UserService.Login", input, filter) }
  };
  
  ns.UserServiceClient = class UserServiceClient {
    constructor(conn) { this.conn = conn }
    ListUser(input, options) { return this.conn.invoke("stmp.examples.room.UserService.ListUser", input, options) }
    Login(input, options) { return this.conn.invoke("stmp.examples.room.UserService.Login", input, options) }
  };

  registerMethodAction("stmp.examples.room.UserEvents.StatusUpdated", "1101", pb.stmp.examples.room.UserModel, pb.google.protobuf.Empty);

  ns.UserEventsServer = class UserEventsServer {
    static register(srv, inst) {
      srv.register(inst, "stmp.examples.room.UserEvents.StatusUpdated", inst.StatusUpdated);
    }

    static unregister(srv, inst) {
      srv.unregister(inst, "stmp.examples.room.UserEvents.StatusUpdated");
    }

    StatusUpdated(ctx, input, output) { throw new Error("not implemented") }
  };
  
  ns.UserEventsBroadcaster = class UserEventsBroadcaster {
    StatusUpdatedToOne(input, conn, options) { return conn.invoke("stmp.examples.room.UserEvents.StatusUpdated", input, options) }
    StatusUpdatedToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.UserEvents.StatusUpdated", pm.get(conn), notifyOptions) }
    StatusUpdatedToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.UserEvents.StatusUpdated", input, filter) }
  };
  
  ns.UserEventsClient = class UserEventsClient {
    constructor(conn) { this.conn = conn }
    StatusUpdated(input, options) { return this.conn.invoke("stmp.examples.room.UserEvents.StatusUpdated", input, options) }
  };

  registerMethodAction("stmp.examples.room.RoomService.CreateRoom", "1201", pb.stmp.examples.room.CreateRoomInput, pb.stmp.examples.room.RoomModel);
  registerMethodAction("stmp.examples.room.RoomService.ListRoom", "1202", pb.stmp.examples.room.ListInput, pb.stmp.examples.room.ListRoomOutput);
  registerMethodAction("stmp.examples.room.RoomService.JoinRoom", "1203", pb.stmp.examples.room.JoinRoomInput, pb.stmp.examples.room.RoomModel);
  registerMethodAction("stmp.examples.room.RoomService.ExitRoom", "1204", pb.stmp.examples.room.ExitRoomInput, pb.google.protobuf.Empty);
  registerMethodAction("stmp.examples.room.RoomService.SendMessage", "1205", pb.stmp.examples.room.SendMessageInput, pb.google.protobuf.Empty);

  ns.RoomServiceServer = class RoomServiceServer {
    static register(srv, inst) {
      srv.register(inst, "stmp.examples.room.RoomService.CreateRoom", inst.CreateRoom);
      srv.register(inst, "stmp.examples.room.RoomService.ListRoom", inst.ListRoom);
      srv.register(inst, "stmp.examples.room.RoomService.JoinRoom", inst.JoinRoom);
      srv.register(inst, "stmp.examples.room.RoomService.ExitRoom", inst.ExitRoom);
      srv.register(inst, "stmp.examples.room.RoomService.SendMessage", inst.SendMessage);
    }

    static unregister(srv, inst) {
      srv.unregister(inst, "stmp.examples.room.RoomService.CreateRoom");
      srv.unregister(inst, "stmp.examples.room.RoomService.ListRoom");
      srv.unregister(inst, "stmp.examples.room.RoomService.JoinRoom");
      srv.unregister(inst, "stmp.examples.room.RoomService.ExitRoom");
      srv.unregister(inst, "stmp.examples.room.RoomService.SendMessage");
    }

    CreateRoom(ctx, input, output) { throw new Error("not implemented") }
    ListRoom(ctx, input, output) { throw new Error("not implemented") }
    JoinRoom(ctx, input, output) { throw new Error("not implemented") }
    ExitRoom(ctx, input, output) { throw new Error("not implemented") }
    SendMessage(ctx, input, output) { throw new Error("not implemented") }
  };
  
  ns.RoomServiceBroadcaster = class RoomServiceBroadcaster {
    CreateRoomToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomService.CreateRoom", input, options) }
    CreateRoomToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomService.CreateRoom", pm.get(conn), notifyOptions) }
    CreateRoomToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomService.CreateRoom", input, filter) }
    ListRoomToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomService.ListRoom", input, options) }
    ListRoomToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomService.ListRoom", pm.get(conn), notifyOptions) }
    ListRoomToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomService.ListRoom", input, filter) }
    JoinRoomToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomService.JoinRoom", input, options) }
    JoinRoomToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomService.JoinRoom", pm.get(conn), notifyOptions) }
    JoinRoomToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomService.JoinRoom", input, filter) }
    ExitRoomToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomService.ExitRoom", input, options) }
    ExitRoomToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomService.ExitRoom", pm.get(conn), notifyOptions) }
    ExitRoomToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomService.ExitRoom", input, filter) }
    SendMessageToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomService.SendMessage", input, options) }
    SendMessageToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomService.SendMessage", pm.get(conn), notifyOptions) }
    SendMessageToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomService.SendMessage", input, filter) }
  };
  
  ns.RoomServiceClient = class RoomServiceClient {
    constructor(conn) { this.conn = conn }
    CreateRoom(input, options) { return this.conn.invoke("stmp.examples.room.RoomService.CreateRoom", input, options) }
    ListRoom(input, options) { return this.conn.invoke("stmp.examples.room.RoomService.ListRoom", input, options) }
    JoinRoom(input, options) { return this.conn.invoke("stmp.examples.room.RoomService.JoinRoom", input, options) }
    ExitRoom(input, options) { return this.conn.invoke("stmp.examples.room.RoomService.ExitRoom", input, options) }
    SendMessage(input, options) { return this.conn.invoke("stmp.examples.room.RoomService.SendMessage", input, options) }
  };

  registerMethodAction("stmp.examples.room.RoomEvents.UserEnter", "1301", pb.stmp.examples.room.UserEnterEvent, pb.google.protobuf.Empty);
  registerMethodAction("stmp.examples.room.RoomEvents.UserExit", "1302", pb.stmp.examples.room.UserExitEvent, pb.google.protobuf.Empty);
  registerMethodAction("stmp.examples.room.RoomEvents.NewMessage", "1303", pb.stmp.examples.room.ChatMessageModel, pb.google.protobuf.Empty);

  ns.RoomEventsServer = class RoomEventsServer {
    static register(srv, inst) {
      srv.register(inst, "stmp.examples.room.RoomEvents.UserEnter", inst.UserEnter);
      srv.register(inst, "stmp.examples.room.RoomEvents.UserExit", inst.UserExit);
      srv.register(inst, "stmp.examples.room.RoomEvents.NewMessage", inst.NewMessage);
    }

    static unregister(srv, inst) {
      srv.unregister(inst, "stmp.examples.room.RoomEvents.UserEnter");
      srv.unregister(inst, "stmp.examples.room.RoomEvents.UserExit");
      srv.unregister(inst, "stmp.examples.room.RoomEvents.NewMessage");
    }

    UserEnter(ctx, input, output) { throw new Error("not implemented") }
    UserExit(ctx, input, output) { throw new Error("not implemented") }
    NewMessage(ctx, input, output) { throw new Error("not implemented") }
  };
  
  ns.RoomEventsBroadcaster = class RoomEventsBroadcaster {
    UserEnterToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomEvents.UserEnter", input, options) }
    UserEnterToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomEvents.UserEnter", pm.get(conn), notifyOptions) }
    UserEnterToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomEvents.UserEnter", input, filter) }
    UserExitToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomEvents.UserExit", input, options) }
    UserExitToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomEvents.UserExit", pm.get(conn), notifyOptions) }
    UserExitToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomEvents.UserExit", input, filter) }
    NewMessageToOne(input, conn, options) { return conn.invoke("stmp.examples.room.RoomEvents.NewMessage", input, options) }
    NewMessageToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("stmp.examples.room.RoomEvents.NewMessage", pm.get(conn), notifyOptions) }
    NewMessageToAll(input, srv, filter) { return srv.broadcast("stmp.examples.room.RoomEvents.NewMessage", input, filter) }
  };
  
  ns.RoomEventsClient = class RoomEventsClient {
    constructor(conn) { this.conn = conn }
    UserEnter(input, options) { return this.conn.invoke("stmp.examples.room.RoomEvents.UserEnter", input, options) }
    UserExit(input, options) { return this.conn.invoke("stmp.examples.room.RoomEvents.UserExit", input, options) }
    NewMessage(input, options) { return this.conn.invoke("stmp.examples.room.RoomEvents.NewMessage", input, options) }
  };
});
