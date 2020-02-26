// Code generated by protoc-gen-stmp. DO NOT EDIT.
// source: examples/gomoku/gomoku_pb/gomoku.proto
import pb from "./gomoku.pb";
import { CallOptions, Connection, ConnFilter, Context, Server, Client } from "stmp";

export default stmp;

declare namespace stmp {
  
  namespace stmp.examples.gomoku {
    
    
    class RoomServiceServer {
      static register(srv: Server, inst: RoomServiceServer): void
      static unregister(srv: Server, inst: RoomServiceServer): void
      
      MatchRoom(ctx: Context, input: pb.stmp.examples.gomoku.Empty, output: pb.stmp.examples.gomoku.FullRoomModel): void | Promise<void>
      
      ListRoom(ctx: Context, input: pb.stmp.examples.gomoku.ListInput, output: pb.stmp.examples.gomoku.ListRoomOutput): void | Promise<void>
      
      LookonRoom(ctx: Context, input: pb.stmp.examples.gomoku.LookonRoomInput, output: pb.stmp.examples.gomoku.FullRoomModel): void | Promise<void>
      
      JoinRoom(ctx: Context, input: pb.stmp.examples.gomoku.JoinRoomInput, output: pb.stmp.examples.gomoku.FullRoomModel): void | Promise<void>
      
      Ready(ctx: Context, input: pb.stmp.examples.gomoku.Empty, output: pb.stmp.examples.gomoku.Empty): void | Promise<void>
      
      Unready(ctx: Context, input: pb.stmp.examples.gomoku.Empty, output: pb.stmp.examples.gomoku.Empty): void | Promise<void>
      
      ExitRoom(ctx: Context, input: pb.stmp.examples.gomoku.Empty, output: pb.stmp.examples.gomoku.Empty): void | Promise<void>
      
    }

    class RoomServiceClient {
      private client: Client;
      constructor(client: Client)
      
      MatchRoom(data: pb.stmp.examples.gomoku.IEmpty, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.FullRoomModel>
      
      ListRoom(data: pb.stmp.examples.gomoku.IListInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.ListRoomOutput>
      
      LookonRoom(data: pb.stmp.examples.gomoku.ILookonRoomInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.FullRoomModel>
      
      JoinRoom(data: pb.stmp.examples.gomoku.IJoinRoomInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.FullRoomModel>
      
      Ready(data: pb.stmp.examples.gomoku.IEmpty, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.Empty>
      
      Unready(data: pb.stmp.examples.gomoku.IEmpty, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.Empty>
      
      ExitRoom(data: pb.stmp.examples.gomoku.IEmpty, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.Empty>
      
    }
    
    
    
    
    
    class RoomEventsListener {
      static register(c: Client, inst: RoomEventsListener): void
      static unregister(c: Client, inst: RoomEventsListener): void
      
      HandleUserJoin(ctx: Context, input: pb.stmp.examples.gomoku.UserJoinEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserReady(ctx: Context, input: pb.stmp.examples.gomoku.UserReadyEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserUnready(ctx: Context, input: pb.stmp.examples.gomoku.UserUnreadyEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserLookon(ctx: Context, input: pb.stmp.examples.gomoku.UserLookonEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserExit(ctx: Context, input: pb.stmp.examples.gomoku.UserExitEvent, output: pb.stmp.examples.gomoku.Empty): void
      
    }

    class RoomEventsBroadcaster {
      
      static UserJoin(input: pb.stmp.examples.gomoku.IUserJoinEvent, conn: Connection): void
      static UserJoinToSet(input: pb.stmp.examples.gomoku.IUserJoinEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserJoinToAll(input: pb.stmp.examples.gomoku.IUserJoinEvent, srv: Server, filter?: ConnFilter): void
      
      static UserReady(input: pb.stmp.examples.gomoku.IUserReadyEvent, conn: Connection): void
      static UserReadyToSet(input: pb.stmp.examples.gomoku.IUserReadyEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserReadyToAll(input: pb.stmp.examples.gomoku.IUserReadyEvent, srv: Server, filter?: ConnFilter): void
      
      static UserUnready(input: pb.stmp.examples.gomoku.IUserUnreadyEvent, conn: Connection): void
      static UserUnreadyToSet(input: pb.stmp.examples.gomoku.IUserUnreadyEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserUnreadyToAll(input: pb.stmp.examples.gomoku.IUserUnreadyEvent, srv: Server, filter?: ConnFilter): void
      
      static UserLookon(input: pb.stmp.examples.gomoku.IUserLookonEvent, conn: Connection): void
      static UserLookonToSet(input: pb.stmp.examples.gomoku.IUserLookonEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserLookonToAll(input: pb.stmp.examples.gomoku.IUserLookonEvent, srv: Server, filter?: ConnFilter): void
      
      static UserExit(input: pb.stmp.examples.gomoku.IUserExitEvent, conn: Connection): void
      static UserExitToSet(input: pb.stmp.examples.gomoku.IUserExitEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserExitToAll(input: pb.stmp.examples.gomoku.IUserExitEvent, srv: Server, filter?: ConnFilter): void
      
      constructor()
    }
    
    
    
    class GomokuServiceServer {
      static register(srv: Server, inst: GomokuServiceServer): void
      static unregister(srv: Server, inst: GomokuServiceServer): void
      
      Play(ctx: Context, input: pb.stmp.examples.gomoku.HandModel, output: pb.stmp.examples.gomoku.Empty): void | Promise<void>
      
      Apply(ctx: Context, input: pb.stmp.examples.gomoku.ApplyInput, output: pb.stmp.examples.gomoku.Empty): void | Promise<void>
      
      Reply(ctx: Context, input: pb.stmp.examples.gomoku.ReplyInput, output: pb.stmp.examples.gomoku.Empty): void | Promise<void>
      
    }

    class GomokuServiceClient {
      private client: Client;
      constructor(client: Client)
      
      Play(data: pb.stmp.examples.gomoku.IHandModel, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.Empty>
      
      Apply(data: pb.stmp.examples.gomoku.IApplyInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.Empty>
      
      Reply(data: pb.stmp.examples.gomoku.IReplyInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.Empty>
      
    }
    
    
    
    
    
    class GomokuEventsListener {
      static register(c: Client, inst: GomokuEventsListener): void
      static unregister(c: Client, inst: GomokuEventsListener): void
      
      HandleGameStart(ctx: Context, input: pb.stmp.examples.gomoku.GomokuModel, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserPlay(ctx: Context, input: pb.stmp.examples.gomoku.UserPlayEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserApply(ctx: Context, input: pb.stmp.examples.gomoku.UserApplyEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserReply(ctx: Context, input: pb.stmp.examples.gomoku.UserReplyEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserDisconnected(ctx: Context, input: pb.stmp.examples.gomoku.UserDisconnectedEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleUserConnected(ctx: Context, input: pb.stmp.examples.gomoku.UserConnectedEvent, output: pb.stmp.examples.gomoku.Empty): void
      
      HandleGameOver(ctx: Context, input: pb.stmp.examples.gomoku.GomokuModel, output: pb.stmp.examples.gomoku.Empty): void
      
    }

    class GomokuEventsBroadcaster {
      
      static GameStart(input: pb.stmp.examples.gomoku.IGomokuModel, conn: Connection): void
      static GameStartToSet(input: pb.stmp.examples.gomoku.IGomokuModel, conns: Set<Connection>, excludes?: Connection[]): void
      static GameStartToAll(input: pb.stmp.examples.gomoku.IGomokuModel, srv: Server, filter?: ConnFilter): void
      
      static UserPlay(input: pb.stmp.examples.gomoku.IUserPlayEvent, conn: Connection): void
      static UserPlayToSet(input: pb.stmp.examples.gomoku.IUserPlayEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserPlayToAll(input: pb.stmp.examples.gomoku.IUserPlayEvent, srv: Server, filter?: ConnFilter): void
      
      static UserApply(input: pb.stmp.examples.gomoku.IUserApplyEvent, conn: Connection): void
      static UserApplyToSet(input: pb.stmp.examples.gomoku.IUserApplyEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserApplyToAll(input: pb.stmp.examples.gomoku.IUserApplyEvent, srv: Server, filter?: ConnFilter): void
      
      static UserReply(input: pb.stmp.examples.gomoku.IUserReplyEvent, conn: Connection): void
      static UserReplyToSet(input: pb.stmp.examples.gomoku.IUserReplyEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserReplyToAll(input: pb.stmp.examples.gomoku.IUserReplyEvent, srv: Server, filter?: ConnFilter): void
      
      static UserDisconnected(input: pb.stmp.examples.gomoku.IUserDisconnectedEvent, conn: Connection): void
      static UserDisconnectedToSet(input: pb.stmp.examples.gomoku.IUserDisconnectedEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserDisconnectedToAll(input: pb.stmp.examples.gomoku.IUserDisconnectedEvent, srv: Server, filter?: ConnFilter): void
      
      static UserConnected(input: pb.stmp.examples.gomoku.IUserConnectedEvent, conn: Connection): void
      static UserConnectedToSet(input: pb.stmp.examples.gomoku.IUserConnectedEvent, conns: Set<Connection>, excludes?: Connection[]): void
      static UserConnectedToAll(input: pb.stmp.examples.gomoku.IUserConnectedEvent, srv: Server, filter?: ConnFilter): void
      
      static GameOver(input: pb.stmp.examples.gomoku.IGomokuModel, conn: Connection): void
      static GameOverToSet(input: pb.stmp.examples.gomoku.IGomokuModel, conns: Set<Connection>, excludes?: Connection[]): void
      static GameOverToAll(input: pb.stmp.examples.gomoku.IGomokuModel, srv: Server, filter?: ConnFilter): void
      
      constructor()
    }
    
    
    
    class PlayerServiceServer {
      static register(srv: Server, inst: PlayerServiceServer): void
      static unregister(srv: Server, inst: PlayerServiceServer): void
      
      Login(ctx: Context, input: pb.stmp.examples.gomoku.LoginInput, output: pb.stmp.examples.gomoku.PlayerModel): void | Promise<void>
      
      ListUser(ctx: Context, input: pb.stmp.examples.gomoku.ListPlayerInput, output: pb.stmp.examples.gomoku.ListPlayerOutput): void | Promise<void>
      
    }

    class PlayerServiceClient {
      private client: Client;
      constructor(client: Client)
      
      Login(data: pb.stmp.examples.gomoku.ILoginInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.PlayerModel>
      
      ListUser(data: pb.stmp.examples.gomoku.IListPlayerInput, options?: Partial<CallOptions>): Promise<pb.stmp.examples.gomoku.ListPlayerOutput>
      
    }
    
    
    
    
    
    class PlayerEventsListener {
      static register(c: Client, inst: PlayerEventsListener): void
      static unregister(c: Client, inst: PlayerEventsListener): void
      
      HandleStatusUpdated(ctx: Context, input: pb.stmp.examples.gomoku.PlayerModel, output: pb.stmp.examples.gomoku.Empty): void
      
    }

    class PlayerEventsBroadcaster {
      
      static StatusUpdated(input: pb.stmp.examples.gomoku.IPlayerModel, conn: Connection): void
      static StatusUpdatedToSet(input: pb.stmp.examples.gomoku.IPlayerModel, conns: Set<Connection>, excludes?: Connection[]): void
      static StatusUpdatedToAll(input: pb.stmp.examples.gomoku.IPlayerModel, srv: Server, filter?: ConnFilter): void
      
      constructor()
    }
    
    
  }
  
}