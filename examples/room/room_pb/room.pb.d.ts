import * as $protobuf from "protobufjs";
export = pb;

declare namespace pb {


    namespace stmp {

        namespace examples {

            namespace room {

                interface IUserModel {
                    name?: (string|null);
                    room?: (string|null);
                    status?: (stmp.examples.room.UserModel.Status|null);
                }

                class UserModel implements IUserModel {
                    constructor(properties?: stmp.examples.room.IUserModel);
                    public name: string;
                    public room: string;
                    public status: stmp.examples.room.UserModel.Status;
                    public static encode(message: stmp.examples.room.IUserModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.UserModel;
                }

                namespace UserModel {

                    enum Status {
                        Offline = 0,
                        Online = 1,
                        Chatting = 2,
                        ChattingOffline = 3
                    }
                }

                interface ILoginInput {
                    name?: (string|null);
                }

                class LoginInput implements ILoginInput {
                    constructor(properties?: stmp.examples.room.ILoginInput);
                    public name: string;
                    public static encode(message: stmp.examples.room.ILoginInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.LoginInput;
                }

                interface IListInput {
                    limit?: (number|Long|null);
                    offset?: (number|Long|null);
                }

                class ListInput implements IListInput {
                    constructor(properties?: stmp.examples.room.IListInput);
                    public limit: (number|Long);
                    public offset: (number|Long);
                    public static encode(message: stmp.examples.room.IListInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ListInput;
                }

                interface IListUserOutput {
                    total?: (number|Long|null);
                    users?: (stmp.examples.room.IUserModel[]|null);
                }

                class ListUserOutput implements IListUserOutput {
                    constructor(properties?: stmp.examples.room.IListUserOutput);
                    public total: (number|Long);
                    public users: stmp.examples.room.IUserModel[];
                    public static encode(message: stmp.examples.room.IListUserOutput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ListUserOutput;
                }

                class UserService extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public listUser(request: stmp.examples.room.IListInput, callback: stmp.examples.room.UserService.ListUserCallback): void;
                    public listUser(request: stmp.examples.room.IListInput): Promise<stmp.examples.room.ListUserOutput>;
                    public login(request: stmp.examples.room.ILoginInput, callback: stmp.examples.room.UserService.LoginCallback): void;
                    public login(request: stmp.examples.room.ILoginInput): Promise<stmp.examples.room.UserModel>;
                }

                namespace UserService {

                    type ListUserCallback = (error: (Error|null), response?: stmp.examples.room.ListUserOutput) => void;

                    type LoginCallback = (error: (Error|null), response?: stmp.examples.room.UserModel) => void;
                }

                class UserEvents extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public statusUpdated(request: stmp.examples.room.IUserModel, callback: stmp.examples.room.UserEvents.StatusUpdatedCallback): void;
                    public statusUpdated(request: stmp.examples.room.IUserModel): Promise<google.protobuf.Empty>;
                }

                namespace UserEvents {

                    type StatusUpdatedCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;
                }

                interface IChatMessageModel {
                    room?: (string|null);
                    user?: (string|null);
                    content?: (string|null);
                    createdAt?: (number|Long|null);
                }

                class ChatMessageModel implements IChatMessageModel {
                    constructor(properties?: stmp.examples.room.IChatMessageModel);
                    public room: string;
                    public user: string;
                    public content: string;
                    public createdAt: (number|Long);
                    public static encode(message: stmp.examples.room.IChatMessageModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ChatMessageModel;
                }

                interface IRoomModel {
                    name?: (string|null);
                    users?: ({ [k: string]: stmp.examples.room.IUserModel }|null);
                    messages?: (stmp.examples.room.IChatMessageModel[]|null);
                }

                class RoomModel implements IRoomModel {
                    constructor(properties?: stmp.examples.room.IRoomModel);
                    public name: string;
                    public users: { [k: string]: stmp.examples.room.IUserModel };
                    public messages: stmp.examples.room.IChatMessageModel[];
                    public static encode(message: stmp.examples.room.IRoomModel, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.RoomModel;
                }

                interface ICreateRoomInput {
                    name?: (string|null);
                }

                class CreateRoomInput implements ICreateRoomInput {
                    constructor(properties?: stmp.examples.room.ICreateRoomInput);
                    public name: string;
                    public static encode(message: stmp.examples.room.ICreateRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.CreateRoomInput;
                }

                interface IListRoomOutput {
                    total?: (number|Long|null);
                    rooms?: (stmp.examples.room.IRoomModel[]|null);
                }

                class ListRoomOutput implements IListRoomOutput {
                    constructor(properties?: stmp.examples.room.IListRoomOutput);
                    public total: (number|Long);
                    public rooms: stmp.examples.room.IRoomModel[];
                    public static encode(message: stmp.examples.room.IListRoomOutput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ListRoomOutput;
                }

                interface IJoinRoomInput {
                    room?: (string|null);
                }

                class JoinRoomInput implements IJoinRoomInput {
                    constructor(properties?: stmp.examples.room.IJoinRoomInput);
                    public room: string;
                    public static encode(message: stmp.examples.room.IJoinRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.JoinRoomInput;
                }

                interface IExitRoomInput {
                    room?: (string|null);
                }

                class ExitRoomInput implements IExitRoomInput {
                    constructor(properties?: stmp.examples.room.IExitRoomInput);
                    public room: string;
                    public static encode(message: stmp.examples.room.IExitRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ExitRoomInput;
                }

                interface ISendMessageInput {
                    room?: (string|null);
                    content?: (string|null);
                }

                class SendMessageInput implements ISendMessageInput {
                    constructor(properties?: stmp.examples.room.ISendMessageInput);
                    public room: string;
                    public content: string;
                    public static encode(message: stmp.examples.room.ISendMessageInput, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.SendMessageInput;
                }

                class RoomService extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public createRoom(request: stmp.examples.room.ICreateRoomInput, callback: stmp.examples.room.RoomService.CreateRoomCallback): void;
                    public createRoom(request: stmp.examples.room.ICreateRoomInput): Promise<stmp.examples.room.RoomModel>;
                    public listRoom(request: stmp.examples.room.IListInput, callback: stmp.examples.room.RoomService.ListRoomCallback): void;
                    public listRoom(request: stmp.examples.room.IListInput): Promise<stmp.examples.room.ListRoomOutput>;
                    public joinRoom(request: stmp.examples.room.IJoinRoomInput, callback: stmp.examples.room.RoomService.JoinRoomCallback): void;
                    public joinRoom(request: stmp.examples.room.IJoinRoomInput): Promise<stmp.examples.room.RoomModel>;
                    public exitRoom(request: stmp.examples.room.IExitRoomInput, callback: stmp.examples.room.RoomService.ExitRoomCallback): void;
                    public exitRoom(request: stmp.examples.room.IExitRoomInput): Promise<google.protobuf.Empty>;
                    public sendMessage(request: stmp.examples.room.ISendMessageInput, callback: stmp.examples.room.RoomService.SendMessageCallback): void;
                    public sendMessage(request: stmp.examples.room.ISendMessageInput): Promise<google.protobuf.Empty>;
                }

                namespace RoomService {

                    type CreateRoomCallback = (error: (Error|null), response?: stmp.examples.room.RoomModel) => void;

                    type ListRoomCallback = (error: (Error|null), response?: stmp.examples.room.ListRoomOutput) => void;

                    type JoinRoomCallback = (error: (Error|null), response?: stmp.examples.room.RoomModel) => void;

                    type ExitRoomCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;

                    type SendMessageCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;
                }

                interface IUserEnterEvent {
                    room?: (string|null);
                    user?: (stmp.examples.room.IUserModel|null);
                }

                class UserEnterEvent implements IUserEnterEvent {
                    constructor(properties?: stmp.examples.room.IUserEnterEvent);
                    public room: string;
                    public user?: (stmp.examples.room.IUserModel|null);
                    public static encode(message: stmp.examples.room.IUserEnterEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.UserEnterEvent;
                }

                interface IUserExitEvent {
                    room?: (string|null);
                }

                class UserExitEvent implements IUserExitEvent {
                    constructor(properties?: stmp.examples.room.IUserExitEvent);
                    public room: string;
                    public static encode(message: stmp.examples.room.IUserExitEvent, writer?: $protobuf.Writer): $protobuf.Writer;
                    public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.UserExitEvent;
                }

                class RoomEvents extends $protobuf.rpc.Service {
                    constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                    public userEnter(request: stmp.examples.room.IUserEnterEvent, callback: stmp.examples.room.RoomEvents.UserEnterCallback): void;
                    public userEnter(request: stmp.examples.room.IUserEnterEvent): Promise<google.protobuf.Empty>;
                    public userExit(request: stmp.examples.room.IUserExitEvent, callback: stmp.examples.room.RoomEvents.UserExitCallback): void;
                    public userExit(request: stmp.examples.room.IUserExitEvent): Promise<google.protobuf.Empty>;
                    public newMessage(request: stmp.examples.room.IChatMessageModel, callback: stmp.examples.room.RoomEvents.NewMessageCallback): void;
                    public newMessage(request: stmp.examples.room.IChatMessageModel): Promise<google.protobuf.Empty>;
                }

                namespace RoomEvents {

                    type UserEnterCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;

                    type UserExitCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;

                    type NewMessageCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;
                }
            }
        }
    }

    namespace google {

        namespace protobuf {

            interface IEmpty {
            }

            class Empty implements IEmpty {
                constructor(properties?: google.protobuf.IEmpty);
                public static encode(message: google.protobuf.IEmpty, writer?: $protobuf.Writer): $protobuf.Writer;
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.Empty;
            }
        }
    }
}
