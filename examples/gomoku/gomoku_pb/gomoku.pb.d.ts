import * as $protobuf from "protobufjs";
export = pb;

declare namespace pb {


    namespace stmp {

        namespace examples {

            namespace gomoku {

                interface IEmpty {
                }

                class Empty implements IEmpty {
                    constructor(properties?: stmp.examples.gomoku.IEmpty);
                    public static encode(message: stmp.examples.gomoku.IEmpty, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.Empty;
                }

                interface IPlayerModel {
                    id?: (number|Long|null);
                    name?: (string|null);
                    status?: (stmp.examples.gomoku.PlayerModel.Status|null);
                    roomId?: (number|Long|null);
                    seat?: (number|null);
                    gameId?: (number|Long|null);
                    readyTimeout?: (number|Long|null);
                }

                class PlayerModel implements IPlayerModel {
                    constructor(properties?: stmp.examples.gomoku.IPlayerModel);
                    public id: (number|Long);
                    public name: string;
                    public status: stmp.examples.gomoku.PlayerModel.Status;
                    public roomId: (number|Long);
                    public seat: number;
                    public gameId: (number|Long);
                    public readyTimeout: (number|Long);
                    public static encode(message: stmp.examples.gomoku.IPlayerModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.PlayerModel;
                }

                namespace PlayerModel {

                    enum Status {
                        Reserved = 0,
                        Free = 1,
                        Standby = 2,
                        Unready = 3,
                        Ready = 4,
                        Playing = 5
                    }
                }

                interface IRoomModel {
                    id?: (number|Long|null);
                    players?: ({ [k: string]: (number|Long) }|null);
                    gameId?: (number|Long|null);
                    spectators?: ((number|Long)[]|null);
                }

                class RoomModel implements IRoomModel {
                    constructor(properties?: stmp.examples.gomoku.IRoomModel);
                    public id: (number|Long);
                    public players: { [k: string]: (number|Long) };
                    public gameId: (number|Long);
                    public spectators: (number|Long)[];
                    public static encode(message: stmp.examples.gomoku.IRoomModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.RoomModel;
                }

                namespace RoomModel {

                    enum Reasons {
                        Reserved = 0,
                        InRoomAlready = 1
                    }
                }

                interface IHandModel {
                    x?: (number|null);
                    y?: (number|null);
                    t?: (number|null);
                }

                class HandModel implements IHandModel {
                    constructor(properties?: stmp.examples.gomoku.IHandModel);
                    public x: number;
                    public y: number;
                    public t: number;
                    public static encode(message: stmp.examples.gomoku.IHandModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.HandModel;
                }

                interface IApplyModel {
                }

                class ApplyModel implements IApplyModel {
                    constructor(properties?: stmp.examples.gomoku.IApplyModel);
                    public static encode(message: stmp.examples.gomoku.IApplyModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.ApplyModel;
                }

                namespace ApplyModel {

                    enum Kind {
                        Reserved = 0,
                        GiveUp = 1,
                        Draw = 2
                    }
                }

                interface IGomokuModel {
                    id?: (number|Long|null);
                    roomId?: (number|Long|null);
                    playerBlack?: (number|Long|null);
                    playerWhite?: (number|Long|null);
                    seatBlack?: (number|null);
                    seatWhite?: (number|null);
                    history?: (stmp.examples.gomoku.IHandModel[]|null);
                    createdAt?: (number|Long|null);
                    result?: (stmp.examples.gomoku.GomokuModel.Result|null);
                    winner?: (number|Long|null);
                }

                class GomokuModel implements IGomokuModel {
                    constructor(properties?: stmp.examples.gomoku.IGomokuModel);
                    public id: (number|Long);
                    public roomId: (number|Long);
                    public playerBlack: (number|Long);
                    public playerWhite: (number|Long);
                    public seatBlack: number;
                    public seatWhite: number;
                    public history: stmp.examples.gomoku.IHandModel[];
                    public createdAt: (number|Long);
                    public result: stmp.examples.gomoku.GomokuModel.Result;
                    public winner: (number|Long);
                    public static encode(message: stmp.examples.gomoku.IGomokuModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.GomokuModel;
                }

                namespace GomokuModel {

                    enum Result {
                        Playing = 0,
                        Win = 1,
                        Draw = 2,
                        ApplyGiveUp = 3,
                        ApplyDraw = 4,
                        UserStepTimeout = 5,
                        UserTotalTimeout = 6
                    }
                }

                interface IFullRoomModel {
                    room?: (stmp.examples.gomoku.IRoomModel|null);
                    players?: ({ [k: string]: stmp.examples.gomoku.IPlayerModel }|null);
                    game?: (stmp.examples.gomoku.IGomokuModel|null);
                }

                class FullRoomModel implements IFullRoomModel {
                    constructor(properties?: stmp.examples.gomoku.IFullRoomModel);
                    public room?: (stmp.examples.gomoku.IRoomModel|null);
                    public players: { [k: string]: stmp.examples.gomoku.IPlayerModel };
                    public game?: (stmp.examples.gomoku.IGomokuModel|null);
                    public static encode(message: stmp.examples.gomoku.IFullRoomModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.FullRoomModel;
                }

                interface IListInput {
                    limit?: (number|Long|null);
                    offset?: (number|Long|null);
                }

                class ListInput implements IListInput {
                    constructor(properties?: stmp.examples.gomoku.IListInput);
                    public limit: (number|Long);
                    public offset: (number|Long);
                    public static encode(message: stmp.examples.gomoku.IListInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.ListInput;
                }

                interface IListRoomOutput {
                    total?: (number|Long|null);
                    rooms?: (stmp.examples.gomoku.IRoomModel[]|null);
                }

                class ListRoomOutput implements IListRoomOutput {
                    constructor(properties?: stmp.examples.gomoku.IListRoomOutput);
                    public total: (number|Long);
                    public rooms: stmp.examples.gomoku.IRoomModel[];
                    public static encode(message: stmp.examples.gomoku.IListRoomOutput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.ListRoomOutput;
                }

                interface ILookonRoomInput {
                    roomId?: (number|Long|null);
                }

                class LookonRoomInput implements ILookonRoomInput {
                    constructor(properties?: stmp.examples.gomoku.ILookonRoomInput);
                    public roomId: (number|Long);
                    public static encode(message: stmp.examples.gomoku.ILookonRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.LookonRoomInput;
                }

                interface IJoinRoomInput {
                    roomId?: (number|Long|null);
                    preferSeat?: (number|null);
                    seat?: (number|null);
                }

                class JoinRoomInput implements IJoinRoomInput {
                    constructor(properties?: stmp.examples.gomoku.IJoinRoomInput);
                    public roomId: (number|Long);
                    public preferSeat: number;
                    public seat: number;
                    public static encode(message: stmp.examples.gomoku.IJoinRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.JoinRoomInput;
                }

                class RoomService extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public matchRoom(request: stmp.examples.gomoku.IEmpty, callback: stmp.examples.gomoku.RoomService.MatchRoomCallback): void;
                    public matchRoom(request: stmp.examples.gomoku.IEmpty): Promise<stmp.examples.gomoku.FullRoomModel>;
                    public listRoom(request: stmp.examples.gomoku.IListInput, callback: stmp.examples.gomoku.RoomService.ListRoomCallback): void;
                    public listRoom(request: stmp.examples.gomoku.IListInput): Promise<stmp.examples.gomoku.ListRoomOutput>;
                    public lookonRoom(request: stmp.examples.gomoku.ILookonRoomInput, callback: stmp.examples.gomoku.RoomService.LookonRoomCallback): void;
                    public lookonRoom(request: stmp.examples.gomoku.ILookonRoomInput): Promise<stmp.examples.gomoku.FullRoomModel>;
                    public joinRoom(request: stmp.examples.gomoku.IJoinRoomInput, callback: stmp.examples.gomoku.RoomService.JoinRoomCallback): void;
                    public joinRoom(request: stmp.examples.gomoku.IJoinRoomInput): Promise<stmp.examples.gomoku.FullRoomModel>;
                    public ready(request: stmp.examples.gomoku.IEmpty, callback: stmp.examples.gomoku.RoomService.ReadyCallback): void;
                    public ready(request: stmp.examples.gomoku.IEmpty): Promise<stmp.examples.gomoku.Empty>;
                    public unready(request: stmp.examples.gomoku.IEmpty, callback: stmp.examples.gomoku.RoomService.UnreadyCallback): void;
                    public unready(request: stmp.examples.gomoku.IEmpty): Promise<stmp.examples.gomoku.Empty>;
                    public exitRoom(request: stmp.examples.gomoku.IEmpty, callback: stmp.examples.gomoku.RoomService.ExitRoomCallback): void;
                    public exitRoom(request: stmp.examples.gomoku.IEmpty): Promise<stmp.examples.gomoku.Empty>;
                }

                namespace RoomService {

                    type MatchRoomCallback = (error: (Error|null), response?: stmp.examples.gomoku.FullRoomModel) => void;

                    type ListRoomCallback = (error: (Error|null), response?: stmp.examples.gomoku.ListRoomOutput) => void;

                    type LookonRoomCallback = (error: (Error|null), response?: stmp.examples.gomoku.FullRoomModel) => void;

                    type JoinRoomCallback = (error: (Error|null), response?: stmp.examples.gomoku.FullRoomModel) => void;

                    type ReadyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UnreadyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type ExitRoomCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;
                }

                interface IUserJoinEvent {
                    userId?: (number|Long|null);
                    seat?: (number|null);
                    readyTimeout?: (number|null);
                }

                class UserJoinEvent implements IUserJoinEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserJoinEvent);
                    public userId: (number|Long);
                    public seat: number;
                    public readyTimeout: number;
                    public static encode(message: stmp.examples.gomoku.IUserJoinEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserJoinEvent;
                }

                interface IUserReadyEvent {
                    userId?: (number|Long|null);
                }

                class UserReadyEvent implements IUserReadyEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserReadyEvent);
                    public userId: (number|Long);
                    public static encode(message: stmp.examples.gomoku.IUserReadyEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserReadyEvent;
                }

                interface IUserUnreadyEvent {
                    userId?: (number|Long|null);
                    readyTimeout?: (number|null);
                }

                class UserUnreadyEvent implements IUserUnreadyEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserUnreadyEvent);
                    public userId: (number|Long);
                    public readyTimeout: number;
                    public static encode(message: stmp.examples.gomoku.IUserUnreadyEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserUnreadyEvent;
                }

                interface IUserLookonEvent {
                    userId?: (number|Long|null);
                }

                class UserLookonEvent implements IUserLookonEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserLookonEvent);
                    public userId: (number|Long);
                    public static encode(message: stmp.examples.gomoku.IUserLookonEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserLookonEvent;
                }

                interface IUserExitEvent {
                    userId?: (number|Long|null);
                }

                class UserExitEvent implements IUserExitEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserExitEvent);
                    public userId: (number|Long);
                    public static encode(message: stmp.examples.gomoku.IUserExitEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserExitEvent;
                }

                class RoomEvents extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public userJoin(request: stmp.examples.gomoku.IUserJoinEvent, callback: stmp.examples.gomoku.RoomEvents.UserJoinCallback): void;
                    public userJoin(request: stmp.examples.gomoku.IUserJoinEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userReady(request: stmp.examples.gomoku.IUserReadyEvent, callback: stmp.examples.gomoku.RoomEvents.UserReadyCallback): void;
                    public userReady(request: stmp.examples.gomoku.IUserReadyEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userUnready(request: stmp.examples.gomoku.IUserUnreadyEvent, callback: stmp.examples.gomoku.RoomEvents.UserUnreadyCallback): void;
                    public userUnready(request: stmp.examples.gomoku.IUserUnreadyEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userLookon(request: stmp.examples.gomoku.IUserLookonEvent, callback: stmp.examples.gomoku.RoomEvents.UserLookonCallback): void;
                    public userLookon(request: stmp.examples.gomoku.IUserLookonEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userExit(request: stmp.examples.gomoku.IUserExitEvent, callback: stmp.examples.gomoku.RoomEvents.UserExitCallback): void;
                    public userExit(request: stmp.examples.gomoku.IUserExitEvent): Promise<stmp.examples.gomoku.Empty>;
                }

                namespace RoomEvents {

                    type UserJoinCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserReadyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserUnreadyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserLookonCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserExitCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;
                }

                interface IApplyInput {
                    kind?: (stmp.examples.gomoku.ApplyModel.Kind|null);
                }

                class ApplyInput implements IApplyInput {
                    constructor(properties?: stmp.examples.gomoku.IApplyInput);
                    public kind: stmp.examples.gomoku.ApplyModel.Kind;
                    public static encode(message: stmp.examples.gomoku.IApplyInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.ApplyInput;
                }

                interface IReplyInput {
                    kind?: (stmp.examples.gomoku.ApplyModel.Kind|null);
                    accept?: (boolean|null);
                }

                class ReplyInput implements IReplyInput {
                    constructor(properties?: stmp.examples.gomoku.IReplyInput);
                    public kind: stmp.examples.gomoku.ApplyModel.Kind;
                    public accept: boolean;
                    public static encode(message: stmp.examples.gomoku.IReplyInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.ReplyInput;
                }

                class GomokuService extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public play(request: stmp.examples.gomoku.IHandModel, callback: stmp.examples.gomoku.GomokuService.PlayCallback): void;
                    public play(request: stmp.examples.gomoku.IHandModel): Promise<stmp.examples.gomoku.Empty>;
                    public apply(request: stmp.examples.gomoku.IApplyInput, callback: stmp.examples.gomoku.GomokuService.ApplyCallback): void;
                    public apply(request: stmp.examples.gomoku.IApplyInput): Promise<stmp.examples.gomoku.Empty>;
                    public reply(request: stmp.examples.gomoku.IReplyInput, callback: stmp.examples.gomoku.GomokuService.ReplyCallback): void;
                    public reply(request: stmp.examples.gomoku.IReplyInput): Promise<stmp.examples.gomoku.Empty>;
                }

                namespace GomokuService {

                    type PlayCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type ApplyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type ReplyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;
                }

                interface IUserPlayEvent {
                    userId?: (number|Long|null);
                    hand?: (stmp.examples.gomoku.IHandModel|null);
                }

                class UserPlayEvent implements IUserPlayEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserPlayEvent);
                    public userId: (number|Long);
                    public hand?: (stmp.examples.gomoku.IHandModel|null);
                    public static encode(message: stmp.examples.gomoku.IUserPlayEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserPlayEvent;
                }

                interface IUserApplyEvent {
                    kind?: (stmp.examples.gomoku.ApplyModel.Kind|null);
                }

                class UserApplyEvent implements IUserApplyEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserApplyEvent);
                    public kind: stmp.examples.gomoku.ApplyModel.Kind;
                    public static encode(message: stmp.examples.gomoku.IUserApplyEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserApplyEvent;
                }

                interface IUserReplyEvent {
                    kind?: (stmp.examples.gomoku.ApplyModel.Kind|null);
                    accepted?: (boolean|null);
                }

                class UserReplyEvent implements IUserReplyEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserReplyEvent);
                    public kind: stmp.examples.gomoku.ApplyModel.Kind;
                    public accepted: boolean;
                    public static encode(message: stmp.examples.gomoku.IUserReplyEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserReplyEvent;
                }

                interface IUserDisconnectedEvent {
                    userId?: (number|Long|null);
                    waitTimeout?: (number|null);
                }

                class UserDisconnectedEvent implements IUserDisconnectedEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserDisconnectedEvent);
                    public userId: (number|Long);
                    public waitTimeout: number;
                    public static encode(message: stmp.examples.gomoku.IUserDisconnectedEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserDisconnectedEvent;
                }

                interface IUserConnectedEvent {
                    userId?: (number|Long|null);
                }

                class UserConnectedEvent implements IUserConnectedEvent {
                    constructor(properties?: stmp.examples.gomoku.IUserConnectedEvent);
                    public userId: (number|Long);
                    public static encode(message: stmp.examples.gomoku.IUserConnectedEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.UserConnectedEvent;
                }

                class GomokuEvents extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public gameStart(request: stmp.examples.gomoku.IGomokuModel, callback: stmp.examples.gomoku.GomokuEvents.GameStartCallback): void;
                    public gameStart(request: stmp.examples.gomoku.IGomokuModel): Promise<stmp.examples.gomoku.Empty>;
                    public userPlay(request: stmp.examples.gomoku.IUserPlayEvent, callback: stmp.examples.gomoku.GomokuEvents.UserPlayCallback): void;
                    public userPlay(request: stmp.examples.gomoku.IUserPlayEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userApply(request: stmp.examples.gomoku.IUserApplyEvent, callback: stmp.examples.gomoku.GomokuEvents.UserApplyCallback): void;
                    public userApply(request: stmp.examples.gomoku.IUserApplyEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userReply(request: stmp.examples.gomoku.IUserReplyEvent, callback: stmp.examples.gomoku.GomokuEvents.UserReplyCallback): void;
                    public userReply(request: stmp.examples.gomoku.IUserReplyEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userDisconnected(request: stmp.examples.gomoku.IUserDisconnectedEvent, callback: stmp.examples.gomoku.GomokuEvents.UserDisconnectedCallback): void;
                    public userDisconnected(request: stmp.examples.gomoku.IUserDisconnectedEvent): Promise<stmp.examples.gomoku.Empty>;
                    public userConnected(request: stmp.examples.gomoku.IUserConnectedEvent, callback: stmp.examples.gomoku.GomokuEvents.UserConnectedCallback): void;
                    public userConnected(request: stmp.examples.gomoku.IUserConnectedEvent): Promise<stmp.examples.gomoku.Empty>;
                    public gameOver(request: stmp.examples.gomoku.IGomokuModel, callback: stmp.examples.gomoku.GomokuEvents.GameOverCallback): void;
                    public gameOver(request: stmp.examples.gomoku.IGomokuModel): Promise<stmp.examples.gomoku.Empty>;
                }

                namespace GomokuEvents {

                    type GameStartCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserPlayCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserApplyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserReplyCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserDisconnectedCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type UserConnectedCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;

                    type GameOverCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;
                }

                interface ILoginInput {
                    name?: (string|null);
                }

                class LoginInput implements ILoginInput {
                    constructor(properties?: stmp.examples.gomoku.ILoginInput);
                    public name: string;
                    public static encode(message: stmp.examples.gomoku.ILoginInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.LoginInput;
                }

                interface IListPlayerInput {
                    limit?: (number|Long|null);
                    offset?: (number|Long|null);
                    ids?: ((number|Long)[]|null);
                }

                class ListPlayerInput implements IListPlayerInput {
                    constructor(properties?: stmp.examples.gomoku.IListPlayerInput);
                    public limit: (number|Long);
                    public offset: (number|Long);
                    public ids: (number|Long)[];
                    public static encode(message: stmp.examples.gomoku.IListPlayerInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.ListPlayerInput;
                }

                interface IListPlayerOutput {
                    total?: (number|Long|null);
                    players?: (stmp.examples.gomoku.IPlayerModel[]|null);
                }

                class ListPlayerOutput implements IListPlayerOutput {
                    constructor(properties?: stmp.examples.gomoku.IListPlayerOutput);
                    public total: (number|Long);
                    public players: stmp.examples.gomoku.IPlayerModel[];
                    public static encode(message: stmp.examples.gomoku.IListPlayerOutput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.gomoku.ListPlayerOutput;
                }

                class PlayerService extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public login(request: stmp.examples.gomoku.ILoginInput, callback: stmp.examples.gomoku.PlayerService.LoginCallback): void;
                    public login(request: stmp.examples.gomoku.ILoginInput): Promise<stmp.examples.gomoku.PlayerModel>;
                    public listUser(request: stmp.examples.gomoku.IListPlayerInput, callback: stmp.examples.gomoku.PlayerService.ListUserCallback): void;
                    public listUser(request: stmp.examples.gomoku.IListPlayerInput): Promise<stmp.examples.gomoku.ListPlayerOutput>;
                }

                namespace PlayerService {

                    type LoginCallback = (error: (Error|null), response?: stmp.examples.gomoku.PlayerModel) => void;

                    type ListUserCallback = (error: (Error|null), response?: stmp.examples.gomoku.ListPlayerOutput) => void;
                }

                class PlayerEvents extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public statusUpdated(request: stmp.examples.gomoku.IPlayerModel, callback: stmp.examples.gomoku.PlayerEvents.StatusUpdatedCallback): void;
                    public statusUpdated(request: stmp.examples.gomoku.IPlayerModel): Promise<stmp.examples.gomoku.Empty>;
                }

                namespace PlayerEvents {

                    type StatusUpdatedCallback = (error: (Error|null), response?: stmp.examples.gomoku.Empty) => void;
                }
            }
        }
    }
}
