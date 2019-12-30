import * as $protobuf from "protobufjs";
/** Namespace stmp. */
export namespace stmp {

    /** Namespace examples. */
    namespace examples {

        /** Namespace room. */
        namespace room {

            /** Properties of a UserModel. */
            interface IUserModel {

                /** UserModel id */
                id?: (number|Long|null);

                /** UserModel nickname */
                nickname?: (string|null);

                /** UserModel avatar */
                avatar?: (string|null);
            }

            /** Represents a UserModel. */
            class UserModel implements IUserModel {

                /**
                 * Constructs a new UserModel.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IUserModel);

                /** UserModel id. */
                public id: (number|Long);

                /** UserModel nickname. */
                public nickname: string;

                /** UserModel avatar. */
                public avatar: string;

                /**
                 * Creates a new UserModel instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns UserModel instance
                 */
                public static create(properties?: stmp.examples.room.IUserModel): stmp.examples.room.UserModel;

                /**
                 * Encodes the specified UserModel message. Does not implicitly {@link stmp.examples.room.UserModel.verify|verify} messages.
                 * @param message UserModel message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IUserModel, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified UserModel message, length delimited. Does not implicitly {@link stmp.examples.room.UserModel.verify|verify} messages.
                 * @param message UserModel message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IUserModel, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a UserModel message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns UserModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.UserModel;

                /**
                 * Decodes a UserModel message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns UserModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.UserModel;

                /**
                 * Verifies a UserModel message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a UserModel message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns UserModel
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.UserModel;

                /**
                 * Creates a plain object from a UserModel message. Also converts values to other types if specified.
                 * @param message UserModel
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.UserModel, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this UserModel to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Properties of a RoomModel. */
            interface IRoomModel {

                /** RoomModel id */
                id?: (number|Long|null);

                /** RoomModel name */
                name?: (string|null);

                /** RoomModel users */
                users?: ({ [k: string]: stmp.examples.room.IUserModel }|null);
            }

            /** Represents a RoomModel. */
            class RoomModel implements IRoomModel {

                /**
                 * Constructs a new RoomModel.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IRoomModel);

                /** RoomModel id. */
                public id: (number|Long);

                /** RoomModel name. */
                public name: string;

                /** RoomModel users. */
                public users: { [k: string]: stmp.examples.room.IUserModel };

                /**
                 * Creates a new RoomModel instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns RoomModel instance
                 */
                public static create(properties?: stmp.examples.room.IRoomModel): stmp.examples.room.RoomModel;

                /**
                 * Encodes the specified RoomModel message. Does not implicitly {@link stmp.examples.room.RoomModel.verify|verify} messages.
                 * @param message RoomModel message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IRoomModel, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified RoomModel message, length delimited. Does not implicitly {@link stmp.examples.room.RoomModel.verify|verify} messages.
                 * @param message RoomModel message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IRoomModel, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a RoomModel message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns RoomModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.RoomModel;

                /**
                 * Decodes a RoomModel message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns RoomModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.RoomModel;

                /**
                 * Verifies a RoomModel message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a RoomModel message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns RoomModel
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.RoomModel;

                /**
                 * Creates a plain object from a RoomModel message. Also converts values to other types if specified.
                 * @param message RoomModel
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.RoomModel, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this RoomModel to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Properties of a CreateRoomInput. */
            interface ICreateRoomInput {

                /** CreateRoomInput name */
                name?: (string|null);
            }

            /** Represents a CreateRoomInput. */
            class CreateRoomInput implements ICreateRoomInput {

                /**
                 * Constructs a new CreateRoomInput.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.ICreateRoomInput);

                /** CreateRoomInput name. */
                public name: string;

                /**
                 * Creates a new CreateRoomInput instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns CreateRoomInput instance
                 */
                public static create(properties?: stmp.examples.room.ICreateRoomInput): stmp.examples.room.CreateRoomInput;

                /**
                 * Encodes the specified CreateRoomInput message. Does not implicitly {@link stmp.examples.room.CreateRoomInput.verify|verify} messages.
                 * @param message CreateRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.ICreateRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified CreateRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.CreateRoomInput.verify|verify} messages.
                 * @param message CreateRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.ICreateRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a CreateRoomInput message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns CreateRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.CreateRoomInput;

                /**
                 * Decodes a CreateRoomInput message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns CreateRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.CreateRoomInput;

                /**
                 * Verifies a CreateRoomInput message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a CreateRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns CreateRoomInput
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.CreateRoomInput;

                /**
                 * Creates a plain object from a CreateRoomInput message. Also converts values to other types if specified.
                 * @param message CreateRoomInput
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.CreateRoomInput, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this CreateRoomInput to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Properties of a ListRoomInput. */
            interface IListRoomInput {

                /** ListRoomInput limit */
                limit?: (number|Long|null);

                /** ListRoomInput offset */
                offset?: (number|Long|null);
            }

            /** Represents a ListRoomInput. */
            class ListRoomInput implements IListRoomInput {

                /**
                 * Constructs a new ListRoomInput.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IListRoomInput);

                /** ListRoomInput limit. */
                public limit: (number|Long);

                /** ListRoomInput offset. */
                public offset: (number|Long);

                /**
                 * Creates a new ListRoomInput instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns ListRoomInput instance
                 */
                public static create(properties?: stmp.examples.room.IListRoomInput): stmp.examples.room.ListRoomInput;

                /**
                 * Encodes the specified ListRoomInput message. Does not implicitly {@link stmp.examples.room.ListRoomInput.verify|verify} messages.
                 * @param message ListRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IListRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified ListRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.ListRoomInput.verify|verify} messages.
                 * @param message ListRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IListRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a ListRoomInput message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns ListRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ListRoomInput;

                /**
                 * Decodes a ListRoomInput message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns ListRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.ListRoomInput;

                /**
                 * Verifies a ListRoomInput message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a ListRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns ListRoomInput
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.ListRoomInput;

                /**
                 * Creates a plain object from a ListRoomInput message. Also converts values to other types if specified.
                 * @param message ListRoomInput
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.ListRoomInput, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this ListRoomInput to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Properties of a ListRoomOutput. */
            interface IListRoomOutput {

                /** ListRoomOutput total */
                total?: (number|Long|null);

                /** ListRoomOutput rooms */
                rooms?: (stmp.examples.room.IRoomModel[]|null);
            }

            /** Represents a ListRoomOutput. */
            class ListRoomOutput implements IListRoomOutput {

                /**
                 * Constructs a new ListRoomOutput.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IListRoomOutput);

                /** ListRoomOutput total. */
                public total: (number|Long);

                /** ListRoomOutput rooms. */
                public rooms: stmp.examples.room.IRoomModel[];

                /**
                 * Creates a new ListRoomOutput instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns ListRoomOutput instance
                 */
                public static create(properties?: stmp.examples.room.IListRoomOutput): stmp.examples.room.ListRoomOutput;

                /**
                 * Encodes the specified ListRoomOutput message. Does not implicitly {@link stmp.examples.room.ListRoomOutput.verify|verify} messages.
                 * @param message ListRoomOutput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IListRoomOutput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified ListRoomOutput message, length delimited. Does not implicitly {@link stmp.examples.room.ListRoomOutput.verify|verify} messages.
                 * @param message ListRoomOutput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IListRoomOutput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a ListRoomOutput message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns ListRoomOutput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ListRoomOutput;

                /**
                 * Decodes a ListRoomOutput message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns ListRoomOutput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.ListRoomOutput;

                /**
                 * Verifies a ListRoomOutput message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a ListRoomOutput message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns ListRoomOutput
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.ListRoomOutput;

                /**
                 * Creates a plain object from a ListRoomOutput message. Also converts values to other types if specified.
                 * @param message ListRoomOutput
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.ListRoomOutput, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this ListRoomOutput to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Properties of a JoinRoomInput. */
            interface IJoinRoomInput {

                /** JoinRoomInput id */
                id?: (number|Long|null);
            }

            /** Represents a JoinRoomInput. */
            class JoinRoomInput implements IJoinRoomInput {

                /**
                 * Constructs a new JoinRoomInput.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IJoinRoomInput);

                /** JoinRoomInput id. */
                public id: (number|Long);

                /**
                 * Creates a new JoinRoomInput instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns JoinRoomInput instance
                 */
                public static create(properties?: stmp.examples.room.IJoinRoomInput): stmp.examples.room.JoinRoomInput;

                /**
                 * Encodes the specified JoinRoomInput message. Does not implicitly {@link stmp.examples.room.JoinRoomInput.verify|verify} messages.
                 * @param message JoinRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IJoinRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified JoinRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.JoinRoomInput.verify|verify} messages.
                 * @param message JoinRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IJoinRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a JoinRoomInput message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns JoinRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.JoinRoomInput;

                /**
                 * Decodes a JoinRoomInput message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns JoinRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.JoinRoomInput;

                /**
                 * Verifies a JoinRoomInput message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a JoinRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns JoinRoomInput
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.JoinRoomInput;

                /**
                 * Creates a plain object from a JoinRoomInput message. Also converts values to other types if specified.
                 * @param message JoinRoomInput
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.JoinRoomInput, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this JoinRoomInput to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Properties of an ExitRoomInput. */
            interface IExitRoomInput {

                /** ExitRoomInput id */
                id?: (number|Long|null);
            }

            /** Represents an ExitRoomInput. */
            class ExitRoomInput implements IExitRoomInput {

                /**
                 * Constructs a new ExitRoomInput.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IExitRoomInput);

                /** ExitRoomInput id. */
                public id: (number|Long);

                /**
                 * Creates a new ExitRoomInput instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns ExitRoomInput instance
                 */
                public static create(properties?: stmp.examples.room.IExitRoomInput): stmp.examples.room.ExitRoomInput;

                /**
                 * Encodes the specified ExitRoomInput message. Does not implicitly {@link stmp.examples.room.ExitRoomInput.verify|verify} messages.
                 * @param message ExitRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IExitRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified ExitRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.ExitRoomInput.verify|verify} messages.
                 * @param message ExitRoomInput message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IExitRoomInput, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes an ExitRoomInput message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns ExitRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.ExitRoomInput;

                /**
                 * Decodes an ExitRoomInput message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns ExitRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.ExitRoomInput;

                /**
                 * Verifies an ExitRoomInput message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates an ExitRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns ExitRoomInput
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.ExitRoomInput;

                /**
                 * Creates a plain object from an ExitRoomInput message. Also converts values to other types if specified.
                 * @param message ExitRoomInput
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.ExitRoomInput, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this ExitRoomInput to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Represents a RoomService */
            class RoomService extends $protobuf.rpc.Service {

                /**
                 * Constructs a new RoomService service.
                 * @param rpcImpl RPC implementation
                 * @param [requestDelimited=false] Whether requests are length-delimited
                 * @param [responseDelimited=false] Whether responses are length-delimited
                 */
                constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

                /**
                 * Creates new RoomService service using the specified rpc implementation.
                 * @param rpcImpl RPC implementation
                 * @param [requestDelimited=false] Whether requests are length-delimited
                 * @param [responseDelimited=false] Whether responses are length-delimited
                 * @returns RPC service. Useful where requests and/or responses are streamed.
                 */
                public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): RoomService;

                /**
                 * Calls CreateRoom.
                 * @param request CreateRoomInput message or plain object
                 * @param callback Node-style callback called with the error, if any, and RoomModel
                 */
                public createRoom(request: stmp.examples.room.ICreateRoomInput, callback: stmp.examples.room.RoomService.CreateRoomCallback): void;

                /**
                 * Calls CreateRoom.
                 * @param request CreateRoomInput message or plain object
                 * @returns Promise
                 */
                public createRoom(request: stmp.examples.room.ICreateRoomInput): Promise<stmp.examples.room.RoomModel>;

                /**
                 * Calls ListRoom.
                 * @param request ListRoomInput message or plain object
                 * @param callback Node-style callback called with the error, if any, and ListRoomOutput
                 */
                public listRoom(request: stmp.examples.room.IListRoomInput, callback: stmp.examples.room.RoomService.ListRoomCallback): void;

                /**
                 * Calls ListRoom.
                 * @param request ListRoomInput message or plain object
                 * @returns Promise
                 */
                public listRoom(request: stmp.examples.room.IListRoomInput): Promise<stmp.examples.room.ListRoomOutput>;

                /**
                 * Calls JoinRoom.
                 * @param request JoinRoomInput message or plain object
                 * @param callback Node-style callback called with the error, if any, and RoomModel
                 */
                public joinRoom(request: stmp.examples.room.IJoinRoomInput, callback: stmp.examples.room.RoomService.JoinRoomCallback): void;

                /**
                 * Calls JoinRoom.
                 * @param request JoinRoomInput message or plain object
                 * @returns Promise
                 */
                public joinRoom(request: stmp.examples.room.IJoinRoomInput): Promise<stmp.examples.room.RoomModel>;

                /**
                 * Calls ExitRoom.
                 * @param request ExitRoomInput message or plain object
                 * @param callback Node-style callback called with the error, if any, and Empty
                 */
                public exitRoom(request: stmp.examples.room.IExitRoomInput, callback: stmp.examples.room.RoomService.ExitRoomCallback): void;

                /**
                 * Calls ExitRoom.
                 * @param request ExitRoomInput message or plain object
                 * @returns Promise
                 */
                public exitRoom(request: stmp.examples.room.IExitRoomInput): Promise<google.protobuf.Empty>;
            }

            namespace RoomService {

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#createRoom}.
                 * @param error Error, if any
                 * @param [response] RoomModel
                 */
                type CreateRoomCallback = (error: (Error|null), response?: stmp.examples.room.RoomModel) => void;

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#listRoom}.
                 * @param error Error, if any
                 * @param [response] ListRoomOutput
                 */
                type ListRoomCallback = (error: (Error|null), response?: stmp.examples.room.ListRoomOutput) => void;

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#joinRoom}.
                 * @param error Error, if any
                 * @param [response] RoomModel
                 */
                type JoinRoomCallback = (error: (Error|null), response?: stmp.examples.room.RoomModel) => void;

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#exitRoom}.
                 * @param error Error, if any
                 * @param [response] Empty
                 */
                type ExitRoomCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;
            }

            /** Properties of a UserEnterEvent. */
            interface IUserEnterEvent {

                /** UserEnterEvent roomId */
                roomId?: (number|Long|null);

                /** UserEnterEvent user */
                user?: (stmp.examples.room.IUserModel|null);
            }

            /** Represents a UserEnterEvent. */
            class UserEnterEvent implements IUserEnterEvent {

                /**
                 * Constructs a new UserEnterEvent.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IUserEnterEvent);

                /** UserEnterEvent roomId. */
                public roomId: (number|Long);

                /** UserEnterEvent user. */
                public user?: (stmp.examples.room.IUserModel|null);

                /**
                 * Creates a new UserEnterEvent instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns UserEnterEvent instance
                 */
                public static create(properties?: stmp.examples.room.IUserEnterEvent): stmp.examples.room.UserEnterEvent;

                /**
                 * Encodes the specified UserEnterEvent message. Does not implicitly {@link stmp.examples.room.UserEnterEvent.verify|verify} messages.
                 * @param message UserEnterEvent message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IUserEnterEvent, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified UserEnterEvent message, length delimited. Does not implicitly {@link stmp.examples.room.UserEnterEvent.verify|verify} messages.
                 * @param message UserEnterEvent message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IUserEnterEvent, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a UserEnterEvent message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns UserEnterEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.UserEnterEvent;

                /**
                 * Decodes a UserEnterEvent message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns UserEnterEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.UserEnterEvent;

                /**
                 * Verifies a UserEnterEvent message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a UserEnterEvent message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns UserEnterEvent
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.UserEnterEvent;

                /**
                 * Creates a plain object from a UserEnterEvent message. Also converts values to other types if specified.
                 * @param message UserEnterEvent
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.UserEnterEvent, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this UserEnterEvent to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Properties of a UserExitEvent. */
            interface IUserExitEvent {

                /** UserExitEvent roomId */
                roomId?: (number|Long|null);

                /** UserExitEvent userId */
                userId?: (number|Long|null);
            }

            /** Represents a UserExitEvent. */
            class UserExitEvent implements IUserExitEvent {

                /**
                 * Constructs a new UserExitEvent.
                 * @param [properties] Properties to set
                 */
                constructor(properties?: stmp.examples.room.IUserExitEvent);

                /** UserExitEvent roomId. */
                public roomId: (number|Long);

                /** UserExitEvent userId. */
                public userId: (number|Long);

                /**
                 * Creates a new UserExitEvent instance using the specified properties.
                 * @param [properties] Properties to set
                 * @returns UserExitEvent instance
                 */
                public static create(properties?: stmp.examples.room.IUserExitEvent): stmp.examples.room.UserExitEvent;

                /**
                 * Encodes the specified UserExitEvent message. Does not implicitly {@link stmp.examples.room.UserExitEvent.verify|verify} messages.
                 * @param message UserExitEvent message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encode(message: stmp.examples.room.IUserExitEvent, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Encodes the specified UserExitEvent message, length delimited. Does not implicitly {@link stmp.examples.room.UserExitEvent.verify|verify} messages.
                 * @param message UserExitEvent message or plain object to encode
                 * @param [writer] Writer to encode to
                 * @returns Writer
                 */
                public static encodeDelimited(message: stmp.examples.room.IUserExitEvent, writer?: $protobuf.Writer): $protobuf.Writer;

                /**
                 * Decodes a UserExitEvent message from the specified reader or buffer.
                 * @param reader Reader or buffer to decode from
                 * @param [length] Message length if known beforehand
                 * @returns UserExitEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): stmp.examples.room.UserExitEvent;

                /**
                 * Decodes a UserExitEvent message from the specified reader or buffer, length delimited.
                 * @param reader Reader or buffer to decode from
                 * @returns UserExitEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): stmp.examples.room.UserExitEvent;

                /**
                 * Verifies a UserExitEvent message.
                 * @param message Plain object to verify
                 * @returns `null` if valid, otherwise the reason why it is not
                 */
                public static verify(message: { [k: string]: any }): (string|null);

                /**
                 * Creates a UserExitEvent message from a plain object. Also converts values to their respective internal types.
                 * @param object Plain object
                 * @returns UserExitEvent
                 */
                public static fromObject(object: { [k: string]: any }): stmp.examples.room.UserExitEvent;

                /**
                 * Creates a plain object from a UserExitEvent message. Also converts values to other types if specified.
                 * @param message UserExitEvent
                 * @param [options] Conversion options
                 * @returns Plain object
                 */
                public static toObject(message: stmp.examples.room.UserExitEvent, options?: $protobuf.IConversionOptions): { [k: string]: any };

                /**
                 * Converts this UserExitEvent to JSON.
                 * @returns JSON object
                 */
                public toJSON(): { [k: string]: any };
            }

            /** Represents a RoomEvents */
            class RoomEvents extends $protobuf.rpc.Service {

                /**
                 * Constructs a new RoomEvents service.
                 * @param rpcImpl RPC implementation
                 * @param [requestDelimited=false] Whether requests are length-delimited
                 * @param [responseDelimited=false] Whether responses are length-delimited
                 */
                constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

                /**
                 * Creates new RoomEvents service using the specified rpc implementation.
                 * @param rpcImpl RPC implementation
                 * @param [requestDelimited=false] Whether requests are length-delimited
                 * @param [responseDelimited=false] Whether responses are length-delimited
                 * @returns RPC service. Useful where requests and/or responses are streamed.
                 */
                public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): RoomEvents;

                /**
                 * Calls UserEnter.
                 * @param request UserEnterEvent message or plain object
                 * @param callback Node-style callback called with the error, if any, and Empty
                 */
                public userEnter(request: stmp.examples.room.IUserEnterEvent, callback: stmp.examples.room.RoomEvents.UserEnterCallback): void;

                /**
                 * Calls UserEnter.
                 * @param request UserEnterEvent message or plain object
                 * @returns Promise
                 */
                public userEnter(request: stmp.examples.room.IUserEnterEvent): Promise<google.protobuf.Empty>;

                /**
                 * Calls UserExit.
                 * @param request UserExitEvent message or plain object
                 * @param callback Node-style callback called with the error, if any, and Empty
                 */
                public userExit(request: stmp.examples.room.IUserExitEvent, callback: stmp.examples.room.RoomEvents.UserExitCallback): void;

                /**
                 * Calls UserExit.
                 * @param request UserExitEvent message or plain object
                 * @returns Promise
                 */
                public userExit(request: stmp.examples.room.IUserExitEvent): Promise<google.protobuf.Empty>;
            }

            namespace RoomEvents {

                /**
                 * Callback as used by {@link stmp.examples.room.RoomEvents#userEnter}.
                 * @param error Error, if any
                 * @param [response] Empty
                 */
                type UserEnterCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;

                /**
                 * Callback as used by {@link stmp.examples.room.RoomEvents#userExit}.
                 * @param error Error, if any
                 * @param [response] Empty
                 */
                type UserExitCallback = (error: (Error|null), response?: google.protobuf.Empty) => void;
            }
        }
    }
}

/** Namespace google. */
export namespace google {

    /** Namespace protobuf. */
    namespace protobuf {

        /** Properties of an Empty. */
        interface IEmpty {
        }

        /** Represents an Empty. */
        class Empty implements IEmpty {

            /**
             * Constructs a new Empty.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.IEmpty);

            /**
             * Creates a new Empty instance using the specified properties.
             * @param [properties] Properties to set
             * @returns Empty instance
             */
            public static create(properties?: google.protobuf.IEmpty): google.protobuf.Empty;

            /**
             * Encodes the specified Empty message. Does not implicitly {@link google.protobuf.Empty.verify|verify} messages.
             * @param message Empty message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.IEmpty, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified Empty message, length delimited. Does not implicitly {@link google.protobuf.Empty.verify|verify} messages.
             * @param message Empty message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: google.protobuf.IEmpty, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes an Empty message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Empty
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.Empty;

            /**
             * Decodes an Empty message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns Empty
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): google.protobuf.Empty;

            /**
             * Verifies an Empty message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates an Empty message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Empty
             */
            public static fromObject(object: { [k: string]: any }): google.protobuf.Empty;

            /**
             * Creates a plain object from an Empty message. Also converts values to other types if specified.
             * @param message Empty
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: google.protobuf.Empty, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Empty to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }
    }
}
