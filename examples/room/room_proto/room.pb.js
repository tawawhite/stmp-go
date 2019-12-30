/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
import * as $protobuf from "protobufjs/minimal";

// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const stmp = $root.stmp = (() => {

    /**
     * Namespace stmp.
     * @exports stmp
     * @namespace
     */
    const stmp = {};

    stmp.examples = (function() {

        /**
         * Namespace examples.
         * @memberof stmp
         * @namespace
         */
        const examples = {};

        examples.room = (function() {

            /**
             * Namespace room.
             * @memberof stmp.examples
             * @namespace
             */
            const room = {};

            room.UserModel = (function() {

                /**
                 * Properties of a UserModel.
                 * @memberof stmp.examples.room
                 * @interface IUserModel
                 * @property {number|Long|null} [id] UserModel id
                 * @property {string|null} [nickname] UserModel nickname
                 * @property {string|null} [avatar] UserModel avatar
                 */

                /**
                 * Constructs a new UserModel.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a UserModel.
                 * @implements IUserModel
                 * @constructor
                 * @param {stmp.examples.room.IUserModel=} [properties] Properties to set
                 */
                function UserModel(properties) {
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserModel id.
                 * @member {number|Long} id
                 * @memberof stmp.examples.room.UserModel
                 * @instance
                 */
                UserModel.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * UserModel nickname.
                 * @member {string} nickname
                 * @memberof stmp.examples.room.UserModel
                 * @instance
                 */
                UserModel.prototype.nickname = "";

                /**
                 * UserModel avatar.
                 * @member {string} avatar
                 * @memberof stmp.examples.room.UserModel
                 * @instance
                 */
                UserModel.prototype.avatar = "";

                /**
                 * Creates a new UserModel instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {stmp.examples.room.IUserModel=} [properties] Properties to set
                 * @returns {stmp.examples.room.UserModel} UserModel instance
                 */
                UserModel.create = function create(properties) {
                    return new UserModel(properties);
                };

                /**
                 * Encodes the specified UserModel message. Does not implicitly {@link stmp.examples.room.UserModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {stmp.examples.room.IUserModel} message UserModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.id != null && message.hasOwnProperty("id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
                    if (message.nickname != null && message.hasOwnProperty("nickname"))
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.nickname);
                    if (message.avatar != null && message.hasOwnProperty("avatar"))
                        writer.uint32(/* id 3, wireType 2 =*/26).string(message.avatar);
                    return writer;
                };

                /**
                 * Encodes the specified UserModel message, length delimited. Does not implicitly {@link stmp.examples.room.UserModel.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {stmp.examples.room.IUserModel} message UserModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserModel.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a UserModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.UserModel} UserModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.UserModel();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.id = reader.int64();
                            break;
                        case 2:
                            message.nickname = reader.string();
                            break;
                        case 3:
                            message.avatar = reader.string();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a UserModel message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.UserModel} UserModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserModel.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a UserModel message.
                 * @function verify
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                UserModel.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (!$util.isInteger(message.id) && !(message.id && $util.isInteger(message.id.low) && $util.isInteger(message.id.high)))
                            return "id: integer|Long expected";
                    if (message.nickname != null && message.hasOwnProperty("nickname"))
                        if (!$util.isString(message.nickname))
                            return "nickname: string expected";
                    if (message.avatar != null && message.hasOwnProperty("avatar"))
                        if (!$util.isString(message.avatar))
                            return "avatar: string expected";
                    return null;
                };

                /**
                 * Creates a UserModel message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.UserModel} UserModel
                 */
                UserModel.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.UserModel)
                        return object;
                    let message = new $root.stmp.examples.room.UserModel();
                    if (object.id != null)
                        if ($util.Long)
                            (message.id = $util.Long.fromValue(object.id)).unsigned = false;
                        else if (typeof object.id === "string")
                            message.id = parseInt(object.id, 10);
                        else if (typeof object.id === "number")
                            message.id = object.id;
                        else if (typeof object.id === "object")
                            message.id = new $util.LongBits(object.id.low >>> 0, object.id.high >>> 0).toNumber();
                    if (object.nickname != null)
                        message.nickname = String(object.nickname);
                    if (object.avatar != null)
                        message.avatar = String(object.avatar);
                    return message;
                };

                /**
                 * Creates a plain object from a UserModel message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.UserModel
                 * @static
                 * @param {stmp.examples.room.UserModel} message UserModel
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                UserModel.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.defaults) {
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.id = options.longs === String ? "0" : 0;
                        object.nickname = "";
                        object.avatar = "";
                    }
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (typeof message.id === "number")
                            object.id = options.longs === String ? String(message.id) : message.id;
                        else
                            object.id = options.longs === String ? $util.Long.prototype.toString.call(message.id) : options.longs === Number ? new $util.LongBits(message.id.low >>> 0, message.id.high >>> 0).toNumber() : message.id;
                    if (message.nickname != null && message.hasOwnProperty("nickname"))
                        object.nickname = message.nickname;
                    if (message.avatar != null && message.hasOwnProperty("avatar"))
                        object.avatar = message.avatar;
                    return object;
                };

                /**
                 * Converts this UserModel to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.UserModel
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                UserModel.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return UserModel;
            })();

            room.RoomModel = (function() {

                /**
                 * Properties of a RoomModel.
                 * @memberof stmp.examples.room
                 * @interface IRoomModel
                 * @property {number|Long|null} [id] RoomModel id
                 * @property {string|null} [name] RoomModel name
                 * @property {Object.<string,stmp.examples.room.IUserModel>|null} [users] RoomModel users
                 */

                /**
                 * Constructs a new RoomModel.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a RoomModel.
                 * @implements IRoomModel
                 * @constructor
                 * @param {stmp.examples.room.IRoomModel=} [properties] Properties to set
                 */
                function RoomModel(properties) {
                    this.users = {};
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * RoomModel id.
                 * @member {number|Long} id
                 * @memberof stmp.examples.room.RoomModel
                 * @instance
                 */
                RoomModel.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * RoomModel name.
                 * @member {string} name
                 * @memberof stmp.examples.room.RoomModel
                 * @instance
                 */
                RoomModel.prototype.name = "";

                /**
                 * RoomModel users.
                 * @member {Object.<string,stmp.examples.room.IUserModel>} users
                 * @memberof stmp.examples.room.RoomModel
                 * @instance
                 */
                RoomModel.prototype.users = $util.emptyObject;

                /**
                 * Creates a new RoomModel instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {stmp.examples.room.IRoomModel=} [properties] Properties to set
                 * @returns {stmp.examples.room.RoomModel} RoomModel instance
                 */
                RoomModel.create = function create(properties) {
                    return new RoomModel(properties);
                };

                /**
                 * Encodes the specified RoomModel message. Does not implicitly {@link stmp.examples.room.RoomModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {stmp.examples.room.IRoomModel} message RoomModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                RoomModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.id != null && message.hasOwnProperty("id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
                    if (message.name != null && message.hasOwnProperty("name"))
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                    if (message.users != null && message.hasOwnProperty("users"))
                        for (let keys = Object.keys(message.users), i = 0; i < keys.length; ++i) {
                            writer.uint32(/* id 3, wireType 2 =*/26).fork().uint32(/* id 1, wireType 0 =*/8).int64(keys[i]);
                            $root.stmp.examples.room.UserModel.encode(message.users[keys[i]], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim().ldelim();
                        }
                    return writer;
                };

                /**
                 * Encodes the specified RoomModel message, length delimited. Does not implicitly {@link stmp.examples.room.RoomModel.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {stmp.examples.room.IRoomModel} message RoomModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                RoomModel.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a RoomModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.RoomModel} RoomModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                RoomModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.RoomModel(), key;
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.id = reader.int64();
                            break;
                        case 2:
                            message.name = reader.string();
                            break;
                        case 3:
                            reader.skip().pos++;
                            if (message.users === $util.emptyObject)
                                message.users = {};
                            key = reader.int64();
                            reader.pos++;
                            message.users[typeof key === "object" ? $util.longToHash(key) : key] = $root.stmp.examples.room.UserModel.decode(reader, reader.uint32());
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a RoomModel message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.RoomModel} RoomModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                RoomModel.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a RoomModel message.
                 * @function verify
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                RoomModel.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (!$util.isInteger(message.id) && !(message.id && $util.isInteger(message.id.low) && $util.isInteger(message.id.high)))
                            return "id: integer|Long expected";
                    if (message.name != null && message.hasOwnProperty("name"))
                        if (!$util.isString(message.name))
                            return "name: string expected";
                    if (message.users != null && message.hasOwnProperty("users")) {
                        if (!$util.isObject(message.users))
                            return "users: object expected";
                        let key = Object.keys(message.users);
                        for (let i = 0; i < key.length; ++i) {
                            if (!$util.key64Re.test(key[i]))
                                return "users: integer|Long key{k:int64} expected";
                            {
                                let error = $root.stmp.examples.room.UserModel.verify(message.users[key[i]]);
                                if (error)
                                    return "users." + error;
                            }
                        }
                    }
                    return null;
                };

                /**
                 * Creates a RoomModel message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.RoomModel} RoomModel
                 */
                RoomModel.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.RoomModel)
                        return object;
                    let message = new $root.stmp.examples.room.RoomModel();
                    if (object.id != null)
                        if ($util.Long)
                            (message.id = $util.Long.fromValue(object.id)).unsigned = false;
                        else if (typeof object.id === "string")
                            message.id = parseInt(object.id, 10);
                        else if (typeof object.id === "number")
                            message.id = object.id;
                        else if (typeof object.id === "object")
                            message.id = new $util.LongBits(object.id.low >>> 0, object.id.high >>> 0).toNumber();
                    if (object.name != null)
                        message.name = String(object.name);
                    if (object.users) {
                        if (typeof object.users !== "object")
                            throw TypeError(".stmp.examples.room.RoomModel.users: object expected");
                        message.users = {};
                        for (let keys = Object.keys(object.users), i = 0; i < keys.length; ++i) {
                            if (typeof object.users[keys[i]] !== "object")
                                throw TypeError(".stmp.examples.room.RoomModel.users: object expected");
                            message.users[keys[i]] = $root.stmp.examples.room.UserModel.fromObject(object.users[keys[i]]);
                        }
                    }
                    return message;
                };

                /**
                 * Creates a plain object from a RoomModel message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.RoomModel
                 * @static
                 * @param {stmp.examples.room.RoomModel} message RoomModel
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                RoomModel.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.objects || options.defaults)
                        object.users = {};
                    if (options.defaults) {
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.id = options.longs === String ? "0" : 0;
                        object.name = "";
                    }
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (typeof message.id === "number")
                            object.id = options.longs === String ? String(message.id) : message.id;
                        else
                            object.id = options.longs === String ? $util.Long.prototype.toString.call(message.id) : options.longs === Number ? new $util.LongBits(message.id.low >>> 0, message.id.high >>> 0).toNumber() : message.id;
                    if (message.name != null && message.hasOwnProperty("name"))
                        object.name = message.name;
                    let keys2;
                    if (message.users && (keys2 = Object.keys(message.users)).length) {
                        object.users = {};
                        for (let j = 0; j < keys2.length; ++j)
                            object.users[keys2[j]] = $root.stmp.examples.room.UserModel.toObject(message.users[keys2[j]], options);
                    }
                    return object;
                };

                /**
                 * Converts this RoomModel to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.RoomModel
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                RoomModel.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return RoomModel;
            })();

            room.CreateRoomInput = (function() {

                /**
                 * Properties of a CreateRoomInput.
                 * @memberof stmp.examples.room
                 * @interface ICreateRoomInput
                 * @property {string|null} [name] CreateRoomInput name
                 */

                /**
                 * Constructs a new CreateRoomInput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a CreateRoomInput.
                 * @implements ICreateRoomInput
                 * @constructor
                 * @param {stmp.examples.room.ICreateRoomInput=} [properties] Properties to set
                 */
                function CreateRoomInput(properties) {
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * CreateRoomInput name.
                 * @member {string} name
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @instance
                 */
                CreateRoomInput.prototype.name = "";

                /**
                 * Creates a new CreateRoomInput instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {stmp.examples.room.ICreateRoomInput=} [properties] Properties to set
                 * @returns {stmp.examples.room.CreateRoomInput} CreateRoomInput instance
                 */
                CreateRoomInput.create = function create(properties) {
                    return new CreateRoomInput(properties);
                };

                /**
                 * Encodes the specified CreateRoomInput message. Does not implicitly {@link stmp.examples.room.CreateRoomInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {stmp.examples.room.ICreateRoomInput} message CreateRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                CreateRoomInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.name != null && message.hasOwnProperty("name"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
                    return writer;
                };

                /**
                 * Encodes the specified CreateRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.CreateRoomInput.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {stmp.examples.room.ICreateRoomInput} message CreateRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                CreateRoomInput.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a CreateRoomInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.CreateRoomInput} CreateRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                CreateRoomInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.CreateRoomInput();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.name = reader.string();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a CreateRoomInput message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.CreateRoomInput} CreateRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                CreateRoomInput.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a CreateRoomInput message.
                 * @function verify
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                CreateRoomInput.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.name != null && message.hasOwnProperty("name"))
                        if (!$util.isString(message.name))
                            return "name: string expected";
                    return null;
                };

                /**
                 * Creates a CreateRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.CreateRoomInput} CreateRoomInput
                 */
                CreateRoomInput.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.CreateRoomInput)
                        return object;
                    let message = new $root.stmp.examples.room.CreateRoomInput();
                    if (object.name != null)
                        message.name = String(object.name);
                    return message;
                };

                /**
                 * Creates a plain object from a CreateRoomInput message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @static
                 * @param {stmp.examples.room.CreateRoomInput} message CreateRoomInput
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                CreateRoomInput.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.defaults)
                        object.name = "";
                    if (message.name != null && message.hasOwnProperty("name"))
                        object.name = message.name;
                    return object;
                };

                /**
                 * Converts this CreateRoomInput to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.CreateRoomInput
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                CreateRoomInput.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return CreateRoomInput;
            })();

            room.ListRoomInput = (function() {

                /**
                 * Properties of a ListRoomInput.
                 * @memberof stmp.examples.room
                 * @interface IListRoomInput
                 * @property {number|Long|null} [limit] ListRoomInput limit
                 * @property {number|Long|null} [offset] ListRoomInput offset
                 */

                /**
                 * Constructs a new ListRoomInput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a ListRoomInput.
                 * @implements IListRoomInput
                 * @constructor
                 * @param {stmp.examples.room.IListRoomInput=} [properties] Properties to set
                 */
                function ListRoomInput(properties) {
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ListRoomInput limit.
                 * @member {number|Long} limit
                 * @memberof stmp.examples.room.ListRoomInput
                 * @instance
                 */
                ListRoomInput.prototype.limit = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListRoomInput offset.
                 * @member {number|Long} offset
                 * @memberof stmp.examples.room.ListRoomInput
                 * @instance
                 */
                ListRoomInput.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Creates a new ListRoomInput instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {stmp.examples.room.IListRoomInput=} [properties] Properties to set
                 * @returns {stmp.examples.room.ListRoomInput} ListRoomInput instance
                 */
                ListRoomInput.create = function create(properties) {
                    return new ListRoomInput(properties);
                };

                /**
                 * Encodes the specified ListRoomInput message. Does not implicitly {@link stmp.examples.room.ListRoomInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {stmp.examples.room.IListRoomInput} message ListRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListRoomInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.limit != null && message.hasOwnProperty("limit"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.limit);
                    if (message.offset != null && message.hasOwnProperty("offset"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int64(message.offset);
                    return writer;
                };

                /**
                 * Encodes the specified ListRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.ListRoomInput.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {stmp.examples.room.IListRoomInput} message ListRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListRoomInput.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a ListRoomInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.ListRoomInput} ListRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListRoomInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ListRoomInput();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.limit = reader.int64();
                            break;
                        case 2:
                            message.offset = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a ListRoomInput message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.ListRoomInput} ListRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListRoomInput.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a ListRoomInput message.
                 * @function verify
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                ListRoomInput.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.limit != null && message.hasOwnProperty("limit"))
                        if (!$util.isInteger(message.limit) && !(message.limit && $util.isInteger(message.limit.low) && $util.isInteger(message.limit.high)))
                            return "limit: integer|Long expected";
                    if (message.offset != null && message.hasOwnProperty("offset"))
                        if (!$util.isInteger(message.offset) && !(message.offset && $util.isInteger(message.offset.low) && $util.isInteger(message.offset.high)))
                            return "offset: integer|Long expected";
                    return null;
                };

                /**
                 * Creates a ListRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.ListRoomInput} ListRoomInput
                 */
                ListRoomInput.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.ListRoomInput)
                        return object;
                    let message = new $root.stmp.examples.room.ListRoomInput();
                    if (object.limit != null)
                        if ($util.Long)
                            (message.limit = $util.Long.fromValue(object.limit)).unsigned = false;
                        else if (typeof object.limit === "string")
                            message.limit = parseInt(object.limit, 10);
                        else if (typeof object.limit === "number")
                            message.limit = object.limit;
                        else if (typeof object.limit === "object")
                            message.limit = new $util.LongBits(object.limit.low >>> 0, object.limit.high >>> 0).toNumber();
                    if (object.offset != null)
                        if ($util.Long)
                            (message.offset = $util.Long.fromValue(object.offset)).unsigned = false;
                        else if (typeof object.offset === "string")
                            message.offset = parseInt(object.offset, 10);
                        else if (typeof object.offset === "number")
                            message.offset = object.offset;
                        else if (typeof object.offset === "object")
                            message.offset = new $util.LongBits(object.offset.low >>> 0, object.offset.high >>> 0).toNumber();
                    return message;
                };

                /**
                 * Creates a plain object from a ListRoomInput message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.ListRoomInput
                 * @static
                 * @param {stmp.examples.room.ListRoomInput} message ListRoomInput
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                ListRoomInput.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.defaults) {
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.limit = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.limit = options.longs === String ? "0" : 0;
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.offset = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.offset = options.longs === String ? "0" : 0;
                    }
                    if (message.limit != null && message.hasOwnProperty("limit"))
                        if (typeof message.limit === "number")
                            object.limit = options.longs === String ? String(message.limit) : message.limit;
                        else
                            object.limit = options.longs === String ? $util.Long.prototype.toString.call(message.limit) : options.longs === Number ? new $util.LongBits(message.limit.low >>> 0, message.limit.high >>> 0).toNumber() : message.limit;
                    if (message.offset != null && message.hasOwnProperty("offset"))
                        if (typeof message.offset === "number")
                            object.offset = options.longs === String ? String(message.offset) : message.offset;
                        else
                            object.offset = options.longs === String ? $util.Long.prototype.toString.call(message.offset) : options.longs === Number ? new $util.LongBits(message.offset.low >>> 0, message.offset.high >>> 0).toNumber() : message.offset;
                    return object;
                };

                /**
                 * Converts this ListRoomInput to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.ListRoomInput
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                ListRoomInput.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return ListRoomInput;
            })();

            room.ListRoomOutput = (function() {

                /**
                 * Properties of a ListRoomOutput.
                 * @memberof stmp.examples.room
                 * @interface IListRoomOutput
                 * @property {number|Long|null} [total] ListRoomOutput total
                 * @property {Array.<stmp.examples.room.IRoomModel>|null} [rooms] ListRoomOutput rooms
                 */

                /**
                 * Constructs a new ListRoomOutput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a ListRoomOutput.
                 * @implements IListRoomOutput
                 * @constructor
                 * @param {stmp.examples.room.IListRoomOutput=} [properties] Properties to set
                 */
                function ListRoomOutput(properties) {
                    this.rooms = [];
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ListRoomOutput total.
                 * @member {number|Long} total
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @instance
                 */
                ListRoomOutput.prototype.total = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListRoomOutput rooms.
                 * @member {Array.<stmp.examples.room.IRoomModel>} rooms
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @instance
                 */
                ListRoomOutput.prototype.rooms = $util.emptyArray;

                /**
                 * Creates a new ListRoomOutput instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {stmp.examples.room.IListRoomOutput=} [properties] Properties to set
                 * @returns {stmp.examples.room.ListRoomOutput} ListRoomOutput instance
                 */
                ListRoomOutput.create = function create(properties) {
                    return new ListRoomOutput(properties);
                };

                /**
                 * Encodes the specified ListRoomOutput message. Does not implicitly {@link stmp.examples.room.ListRoomOutput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {stmp.examples.room.IListRoomOutput} message ListRoomOutput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListRoomOutput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.total != null && message.hasOwnProperty("total"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.total);
                    if (message.rooms != null && message.rooms.length)
                        for (let i = 0; i < message.rooms.length; ++i)
                            $root.stmp.examples.room.RoomModel.encode(message.rooms[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
                };

                /**
                 * Encodes the specified ListRoomOutput message, length delimited. Does not implicitly {@link stmp.examples.room.ListRoomOutput.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {stmp.examples.room.IListRoomOutput} message ListRoomOutput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListRoomOutput.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a ListRoomOutput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.ListRoomOutput} ListRoomOutput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListRoomOutput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ListRoomOutput();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.total = reader.int64();
                            break;
                        case 2:
                            if (!(message.rooms && message.rooms.length))
                                message.rooms = [];
                            message.rooms.push($root.stmp.examples.room.RoomModel.decode(reader, reader.uint32()));
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a ListRoomOutput message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.ListRoomOutput} ListRoomOutput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListRoomOutput.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a ListRoomOutput message.
                 * @function verify
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                ListRoomOutput.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.total != null && message.hasOwnProperty("total"))
                        if (!$util.isInteger(message.total) && !(message.total && $util.isInteger(message.total.low) && $util.isInteger(message.total.high)))
                            return "total: integer|Long expected";
                    if (message.rooms != null && message.hasOwnProperty("rooms")) {
                        if (!Array.isArray(message.rooms))
                            return "rooms: array expected";
                        for (let i = 0; i < message.rooms.length; ++i) {
                            let error = $root.stmp.examples.room.RoomModel.verify(message.rooms[i]);
                            if (error)
                                return "rooms." + error;
                        }
                    }
                    return null;
                };

                /**
                 * Creates a ListRoomOutput message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.ListRoomOutput} ListRoomOutput
                 */
                ListRoomOutput.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.ListRoomOutput)
                        return object;
                    let message = new $root.stmp.examples.room.ListRoomOutput();
                    if (object.total != null)
                        if ($util.Long)
                            (message.total = $util.Long.fromValue(object.total)).unsigned = false;
                        else if (typeof object.total === "string")
                            message.total = parseInt(object.total, 10);
                        else if (typeof object.total === "number")
                            message.total = object.total;
                        else if (typeof object.total === "object")
                            message.total = new $util.LongBits(object.total.low >>> 0, object.total.high >>> 0).toNumber();
                    if (object.rooms) {
                        if (!Array.isArray(object.rooms))
                            throw TypeError(".stmp.examples.room.ListRoomOutput.rooms: array expected");
                        message.rooms = [];
                        for (let i = 0; i < object.rooms.length; ++i) {
                            if (typeof object.rooms[i] !== "object")
                                throw TypeError(".stmp.examples.room.ListRoomOutput.rooms: object expected");
                            message.rooms[i] = $root.stmp.examples.room.RoomModel.fromObject(object.rooms[i]);
                        }
                    }
                    return message;
                };

                /**
                 * Creates a plain object from a ListRoomOutput message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @static
                 * @param {stmp.examples.room.ListRoomOutput} message ListRoomOutput
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                ListRoomOutput.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.arrays || options.defaults)
                        object.rooms = [];
                    if (options.defaults)
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.total = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.total = options.longs === String ? "0" : 0;
                    if (message.total != null && message.hasOwnProperty("total"))
                        if (typeof message.total === "number")
                            object.total = options.longs === String ? String(message.total) : message.total;
                        else
                            object.total = options.longs === String ? $util.Long.prototype.toString.call(message.total) : options.longs === Number ? new $util.LongBits(message.total.low >>> 0, message.total.high >>> 0).toNumber() : message.total;
                    if (message.rooms && message.rooms.length) {
                        object.rooms = [];
                        for (let j = 0; j < message.rooms.length; ++j)
                            object.rooms[j] = $root.stmp.examples.room.RoomModel.toObject(message.rooms[j], options);
                    }
                    return object;
                };

                /**
                 * Converts this ListRoomOutput to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.ListRoomOutput
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                ListRoomOutput.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return ListRoomOutput;
            })();

            room.JoinRoomInput = (function() {

                /**
                 * Properties of a JoinRoomInput.
                 * @memberof stmp.examples.room
                 * @interface IJoinRoomInput
                 * @property {number|Long|null} [id] JoinRoomInput id
                 */

                /**
                 * Constructs a new JoinRoomInput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a JoinRoomInput.
                 * @implements IJoinRoomInput
                 * @constructor
                 * @param {stmp.examples.room.IJoinRoomInput=} [properties] Properties to set
                 */
                function JoinRoomInput(properties) {
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * JoinRoomInput id.
                 * @member {number|Long} id
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @instance
                 */
                JoinRoomInput.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Creates a new JoinRoomInput instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {stmp.examples.room.IJoinRoomInput=} [properties] Properties to set
                 * @returns {stmp.examples.room.JoinRoomInput} JoinRoomInput instance
                 */
                JoinRoomInput.create = function create(properties) {
                    return new JoinRoomInput(properties);
                };

                /**
                 * Encodes the specified JoinRoomInput message. Does not implicitly {@link stmp.examples.room.JoinRoomInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {stmp.examples.room.IJoinRoomInput} message JoinRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                JoinRoomInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.id != null && message.hasOwnProperty("id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
                    return writer;
                };

                /**
                 * Encodes the specified JoinRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.JoinRoomInput.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {stmp.examples.room.IJoinRoomInput} message JoinRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                JoinRoomInput.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a JoinRoomInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.JoinRoomInput} JoinRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                JoinRoomInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.JoinRoomInput();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.id = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a JoinRoomInput message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.JoinRoomInput} JoinRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                JoinRoomInput.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a JoinRoomInput message.
                 * @function verify
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                JoinRoomInput.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (!$util.isInteger(message.id) && !(message.id && $util.isInteger(message.id.low) && $util.isInteger(message.id.high)))
                            return "id: integer|Long expected";
                    return null;
                };

                /**
                 * Creates a JoinRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.JoinRoomInput} JoinRoomInput
                 */
                JoinRoomInput.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.JoinRoomInput)
                        return object;
                    let message = new $root.stmp.examples.room.JoinRoomInput();
                    if (object.id != null)
                        if ($util.Long)
                            (message.id = $util.Long.fromValue(object.id)).unsigned = false;
                        else if (typeof object.id === "string")
                            message.id = parseInt(object.id, 10);
                        else if (typeof object.id === "number")
                            message.id = object.id;
                        else if (typeof object.id === "object")
                            message.id = new $util.LongBits(object.id.low >>> 0, object.id.high >>> 0).toNumber();
                    return message;
                };

                /**
                 * Creates a plain object from a JoinRoomInput message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @static
                 * @param {stmp.examples.room.JoinRoomInput} message JoinRoomInput
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                JoinRoomInput.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.defaults)
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.id = options.longs === String ? "0" : 0;
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (typeof message.id === "number")
                            object.id = options.longs === String ? String(message.id) : message.id;
                        else
                            object.id = options.longs === String ? $util.Long.prototype.toString.call(message.id) : options.longs === Number ? new $util.LongBits(message.id.low >>> 0, message.id.high >>> 0).toNumber() : message.id;
                    return object;
                };

                /**
                 * Converts this JoinRoomInput to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                JoinRoomInput.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return JoinRoomInput;
            })();

            room.ExitRoomInput = (function() {

                /**
                 * Properties of an ExitRoomInput.
                 * @memberof stmp.examples.room
                 * @interface IExitRoomInput
                 * @property {number|Long|null} [id] ExitRoomInput id
                 */

                /**
                 * Constructs a new ExitRoomInput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents an ExitRoomInput.
                 * @implements IExitRoomInput
                 * @constructor
                 * @param {stmp.examples.room.IExitRoomInput=} [properties] Properties to set
                 */
                function ExitRoomInput(properties) {
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ExitRoomInput id.
                 * @member {number|Long} id
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @instance
                 */
                ExitRoomInput.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Creates a new ExitRoomInput instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {stmp.examples.room.IExitRoomInput=} [properties] Properties to set
                 * @returns {stmp.examples.room.ExitRoomInput} ExitRoomInput instance
                 */
                ExitRoomInput.create = function create(properties) {
                    return new ExitRoomInput(properties);
                };

                /**
                 * Encodes the specified ExitRoomInput message. Does not implicitly {@link stmp.examples.room.ExitRoomInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {stmp.examples.room.IExitRoomInput} message ExitRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ExitRoomInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.id != null && message.hasOwnProperty("id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
                    return writer;
                };

                /**
                 * Encodes the specified ExitRoomInput message, length delimited. Does not implicitly {@link stmp.examples.room.ExitRoomInput.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {stmp.examples.room.IExitRoomInput} message ExitRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ExitRoomInput.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes an ExitRoomInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.ExitRoomInput} ExitRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ExitRoomInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ExitRoomInput();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.id = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes an ExitRoomInput message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.ExitRoomInput} ExitRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ExitRoomInput.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies an ExitRoomInput message.
                 * @function verify
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                ExitRoomInput.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (!$util.isInteger(message.id) && !(message.id && $util.isInteger(message.id.low) && $util.isInteger(message.id.high)))
                            return "id: integer|Long expected";
                    return null;
                };

                /**
                 * Creates an ExitRoomInput message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.ExitRoomInput} ExitRoomInput
                 */
                ExitRoomInput.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.ExitRoomInput)
                        return object;
                    let message = new $root.stmp.examples.room.ExitRoomInput();
                    if (object.id != null)
                        if ($util.Long)
                            (message.id = $util.Long.fromValue(object.id)).unsigned = false;
                        else if (typeof object.id === "string")
                            message.id = parseInt(object.id, 10);
                        else if (typeof object.id === "number")
                            message.id = object.id;
                        else if (typeof object.id === "object")
                            message.id = new $util.LongBits(object.id.low >>> 0, object.id.high >>> 0).toNumber();
                    return message;
                };

                /**
                 * Creates a plain object from an ExitRoomInput message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @static
                 * @param {stmp.examples.room.ExitRoomInput} message ExitRoomInput
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                ExitRoomInput.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.defaults)
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.id = options.longs === String ? "0" : 0;
                    if (message.id != null && message.hasOwnProperty("id"))
                        if (typeof message.id === "number")
                            object.id = options.longs === String ? String(message.id) : message.id;
                        else
                            object.id = options.longs === String ? $util.Long.prototype.toString.call(message.id) : options.longs === Number ? new $util.LongBits(message.id.low >>> 0, message.id.high >>> 0).toNumber() : message.id;
                    return object;
                };

                /**
                 * Converts this ExitRoomInput to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                ExitRoomInput.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return ExitRoomInput;
            })();

            room.RoomService = (function() {

                /**
                 * Constructs a new RoomService service.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a RoomService
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function RoomService(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (RoomService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = RoomService;

                /**
                 * Creates new RoomService service using the specified rpc implementation.
                 * @function create
                 * @memberof stmp.examples.room.RoomService
                 * @static
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 * @returns {RoomService} RPC service. Useful where requests and/or responses are streamed.
                 */
                RoomService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
                    return new this(rpcImpl, requestDelimited, responseDelimited);
                };

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#createRoom}.
                 * @memberof stmp.examples.room.RoomService
                 * @typedef CreateRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.room.RoomModel} [response] RoomModel
                 */

                /**
                 * Calls CreateRoom.
                 * @function createRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.ICreateRoomInput} request CreateRoomInput message or plain object
                 * @param {stmp.examples.room.RoomService.CreateRoomCallback} callback Node-style callback called with the error, if any, and RoomModel
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.createRoom = function createRoom(request, callback) {
                    return this.rpcCall(createRoom, $root.stmp.examples.room.CreateRoomInput, $root.stmp.examples.room.RoomModel, request, callback);
                }, "name", { value: "CreateRoom" });

                /**
                 * Calls CreateRoom.
                 * @function createRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.ICreateRoomInput} request CreateRoomInput message or plain object
                 * @returns {Promise<stmp.examples.room.RoomModel>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#listRoom}.
                 * @memberof stmp.examples.room.RoomService
                 * @typedef ListRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.room.ListRoomOutput} [response] ListRoomOutput
                 */

                /**
                 * Calls ListRoom.
                 * @function listRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.IListRoomInput} request ListRoomInput message or plain object
                 * @param {stmp.examples.room.RoomService.ListRoomCallback} callback Node-style callback called with the error, if any, and ListRoomOutput
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.listRoom = function listRoom(request, callback) {
                    return this.rpcCall(listRoom, $root.stmp.examples.room.ListRoomInput, $root.stmp.examples.room.ListRoomOutput, request, callback);
                }, "name", { value: "ListRoom" });

                /**
                 * Calls ListRoom.
                 * @function listRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.IListRoomInput} request ListRoomInput message or plain object
                 * @returns {Promise<stmp.examples.room.ListRoomOutput>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#joinRoom}.
                 * @memberof stmp.examples.room.RoomService
                 * @typedef JoinRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.room.RoomModel} [response] RoomModel
                 */

                /**
                 * Calls JoinRoom.
                 * @function joinRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.IJoinRoomInput} request JoinRoomInput message or plain object
                 * @param {stmp.examples.room.RoomService.JoinRoomCallback} callback Node-style callback called with the error, if any, and RoomModel
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.joinRoom = function joinRoom(request, callback) {
                    return this.rpcCall(joinRoom, $root.stmp.examples.room.JoinRoomInput, $root.stmp.examples.room.RoomModel, request, callback);
                }, "name", { value: "JoinRoom" });

                /**
                 * Calls JoinRoom.
                 * @function joinRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.IJoinRoomInput} request JoinRoomInput message or plain object
                 * @returns {Promise<stmp.examples.room.RoomModel>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#exitRoom}.
                 * @memberof stmp.examples.room.RoomService
                 * @typedef ExitRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {google.protobuf.Empty} [response] Empty
                 */

                /**
                 * Calls ExitRoom.
                 * @function exitRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.IExitRoomInput} request ExitRoomInput message or plain object
                 * @param {stmp.examples.room.RoomService.ExitRoomCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.exitRoom = function exitRoom(request, callback) {
                    return this.rpcCall(exitRoom, $root.stmp.examples.room.ExitRoomInput, $root.google.protobuf.Empty, request, callback);
                }, "name", { value: "ExitRoom" });

                /**
                 * Calls ExitRoom.
                 * @function exitRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.IExitRoomInput} request ExitRoomInput message or plain object
                 * @returns {Promise<google.protobuf.Empty>} Promise
                 * @variation 2
                 */

                return RoomService;
            })();

            room.UserEnterEvent = (function() {

                /**
                 * Properties of a UserEnterEvent.
                 * @memberof stmp.examples.room
                 * @interface IUserEnterEvent
                 * @property {number|Long|null} [roomId] UserEnterEvent roomId
                 * @property {stmp.examples.room.IUserModel|null} [user] UserEnterEvent user
                 */

                /**
                 * Constructs a new UserEnterEvent.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a UserEnterEvent.
                 * @implements IUserEnterEvent
                 * @constructor
                 * @param {stmp.examples.room.IUserEnterEvent=} [properties] Properties to set
                 */
                function UserEnterEvent(properties) {
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserEnterEvent roomId.
                 * @member {number|Long} roomId
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @instance
                 */
                UserEnterEvent.prototype.roomId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * UserEnterEvent user.
                 * @member {stmp.examples.room.IUserModel|null|undefined} user
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @instance
                 */
                UserEnterEvent.prototype.user = null;

                /**
                 * Creates a new UserEnterEvent instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {stmp.examples.room.IUserEnterEvent=} [properties] Properties to set
                 * @returns {stmp.examples.room.UserEnterEvent} UserEnterEvent instance
                 */
                UserEnterEvent.create = function create(properties) {
                    return new UserEnterEvent(properties);
                };

                /**
                 * Encodes the specified UserEnterEvent message. Does not implicitly {@link stmp.examples.room.UserEnterEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {stmp.examples.room.IUserEnterEvent} message UserEnterEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserEnterEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.roomId);
                    if (message.user != null && message.hasOwnProperty("user"))
                        $root.stmp.examples.room.UserModel.encode(message.user, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
                };

                /**
                 * Encodes the specified UserEnterEvent message, length delimited. Does not implicitly {@link stmp.examples.room.UserEnterEvent.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {stmp.examples.room.IUserEnterEvent} message UserEnterEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserEnterEvent.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a UserEnterEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.UserEnterEvent} UserEnterEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserEnterEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.UserEnterEvent();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.roomId = reader.int64();
                            break;
                        case 2:
                            message.user = $root.stmp.examples.room.UserModel.decode(reader, reader.uint32());
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a UserEnterEvent message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.UserEnterEvent} UserEnterEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserEnterEvent.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a UserEnterEvent message.
                 * @function verify
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                UserEnterEvent.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        if (!$util.isInteger(message.roomId) && !(message.roomId && $util.isInteger(message.roomId.low) && $util.isInteger(message.roomId.high)))
                            return "roomId: integer|Long expected";
                    if (message.user != null && message.hasOwnProperty("user")) {
                        let error = $root.stmp.examples.room.UserModel.verify(message.user);
                        if (error)
                            return "user." + error;
                    }
                    return null;
                };

                /**
                 * Creates a UserEnterEvent message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.UserEnterEvent} UserEnterEvent
                 */
                UserEnterEvent.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.UserEnterEvent)
                        return object;
                    let message = new $root.stmp.examples.room.UserEnterEvent();
                    if (object.roomId != null)
                        if ($util.Long)
                            (message.roomId = $util.Long.fromValue(object.roomId)).unsigned = false;
                        else if (typeof object.roomId === "string")
                            message.roomId = parseInt(object.roomId, 10);
                        else if (typeof object.roomId === "number")
                            message.roomId = object.roomId;
                        else if (typeof object.roomId === "object")
                            message.roomId = new $util.LongBits(object.roomId.low >>> 0, object.roomId.high >>> 0).toNumber();
                    if (object.user != null) {
                        if (typeof object.user !== "object")
                            throw TypeError(".stmp.examples.room.UserEnterEvent.user: object expected");
                        message.user = $root.stmp.examples.room.UserModel.fromObject(object.user);
                    }
                    return message;
                };

                /**
                 * Creates a plain object from a UserEnterEvent message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @static
                 * @param {stmp.examples.room.UserEnterEvent} message UserEnterEvent
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                UserEnterEvent.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.defaults) {
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.roomId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.roomId = options.longs === String ? "0" : 0;
                        object.user = null;
                    }
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        if (typeof message.roomId === "number")
                            object.roomId = options.longs === String ? String(message.roomId) : message.roomId;
                        else
                            object.roomId = options.longs === String ? $util.Long.prototype.toString.call(message.roomId) : options.longs === Number ? new $util.LongBits(message.roomId.low >>> 0, message.roomId.high >>> 0).toNumber() : message.roomId;
                    if (message.user != null && message.hasOwnProperty("user"))
                        object.user = $root.stmp.examples.room.UserModel.toObject(message.user, options);
                    return object;
                };

                /**
                 * Converts this UserEnterEvent to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                UserEnterEvent.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return UserEnterEvent;
            })();

            room.UserExitEvent = (function() {

                /**
                 * Properties of a UserExitEvent.
                 * @memberof stmp.examples.room
                 * @interface IUserExitEvent
                 * @property {number|Long|null} [roomId] UserExitEvent roomId
                 * @property {number|Long|null} [userId] UserExitEvent userId
                 */

                /**
                 * Constructs a new UserExitEvent.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a UserExitEvent.
                 * @implements IUserExitEvent
                 * @constructor
                 * @param {stmp.examples.room.IUserExitEvent=} [properties] Properties to set
                 */
                function UserExitEvent(properties) {
                    if (properties)
                        for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserExitEvent roomId.
                 * @member {number|Long} roomId
                 * @memberof stmp.examples.room.UserExitEvent
                 * @instance
                 */
                UserExitEvent.prototype.roomId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * UserExitEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.room.UserExitEvent
                 * @instance
                 */
                UserExitEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Creates a new UserExitEvent instance using the specified properties.
                 * @function create
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {stmp.examples.room.IUserExitEvent=} [properties] Properties to set
                 * @returns {stmp.examples.room.UserExitEvent} UserExitEvent instance
                 */
                UserExitEvent.create = function create(properties) {
                    return new UserExitEvent(properties);
                };

                /**
                 * Encodes the specified UserExitEvent message. Does not implicitly {@link stmp.examples.room.UserExitEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {stmp.examples.room.IUserExitEvent} message UserExitEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserExitEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.roomId);
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int64(message.userId);
                    return writer;
                };

                /**
                 * Encodes the specified UserExitEvent message, length delimited. Does not implicitly {@link stmp.examples.room.UserExitEvent.verify|verify} messages.
                 * @function encodeDelimited
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {stmp.examples.room.IUserExitEvent} message UserExitEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserExitEvent.encodeDelimited = function encodeDelimited(message, writer) {
                    return this.encode(message, writer).ldelim();
                };

                /**
                 * Decodes a UserExitEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.UserExitEvent} UserExitEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserExitEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    let end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.UserExitEvent();
                    while (reader.pos < end) {
                        let tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.roomId = reader.int64();
                            break;
                        case 2:
                            message.userId = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Decodes a UserExitEvent message from the specified reader or buffer, length delimited.
                 * @function decodeDelimited
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @returns {stmp.examples.room.UserExitEvent} UserExitEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserExitEvent.decodeDelimited = function decodeDelimited(reader) {
                    if (!(reader instanceof $Reader))
                        reader = new $Reader(reader);
                    return this.decode(reader, reader.uint32());
                };

                /**
                 * Verifies a UserExitEvent message.
                 * @function verify
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {Object.<string,*>} message Plain object to verify
                 * @returns {string|null} `null` if valid, otherwise the reason why it is not
                 */
                UserExitEvent.verify = function verify(message) {
                    if (typeof message !== "object" || message === null)
                        return "object expected";
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        if (!$util.isInteger(message.roomId) && !(message.roomId && $util.isInteger(message.roomId.low) && $util.isInteger(message.roomId.high)))
                            return "roomId: integer|Long expected";
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        if (!$util.isInteger(message.userId) && !(message.userId && $util.isInteger(message.userId.low) && $util.isInteger(message.userId.high)))
                            return "userId: integer|Long expected";
                    return null;
                };

                /**
                 * Creates a UserExitEvent message from a plain object. Also converts values to their respective internal types.
                 * @function fromObject
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {Object.<string,*>} object Plain object
                 * @returns {stmp.examples.room.UserExitEvent} UserExitEvent
                 */
                UserExitEvent.fromObject = function fromObject(object) {
                    if (object instanceof $root.stmp.examples.room.UserExitEvent)
                        return object;
                    let message = new $root.stmp.examples.room.UserExitEvent();
                    if (object.roomId != null)
                        if ($util.Long)
                            (message.roomId = $util.Long.fromValue(object.roomId)).unsigned = false;
                        else if (typeof object.roomId === "string")
                            message.roomId = parseInt(object.roomId, 10);
                        else if (typeof object.roomId === "number")
                            message.roomId = object.roomId;
                        else if (typeof object.roomId === "object")
                            message.roomId = new $util.LongBits(object.roomId.low >>> 0, object.roomId.high >>> 0).toNumber();
                    if (object.userId != null)
                        if ($util.Long)
                            (message.userId = $util.Long.fromValue(object.userId)).unsigned = false;
                        else if (typeof object.userId === "string")
                            message.userId = parseInt(object.userId, 10);
                        else if (typeof object.userId === "number")
                            message.userId = object.userId;
                        else if (typeof object.userId === "object")
                            message.userId = new $util.LongBits(object.userId.low >>> 0, object.userId.high >>> 0).toNumber();
                    return message;
                };

                /**
                 * Creates a plain object from a UserExitEvent message. Also converts values to other types if specified.
                 * @function toObject
                 * @memberof stmp.examples.room.UserExitEvent
                 * @static
                 * @param {stmp.examples.room.UserExitEvent} message UserExitEvent
                 * @param {$protobuf.IConversionOptions} [options] Conversion options
                 * @returns {Object.<string,*>} Plain object
                 */
                UserExitEvent.toObject = function toObject(message, options) {
                    if (!options)
                        options = {};
                    let object = {};
                    if (options.defaults) {
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.roomId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.roomId = options.longs === String ? "0" : 0;
                        if ($util.Long) {
                            let long = new $util.Long(0, 0, false);
                            object.userId = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                        } else
                            object.userId = options.longs === String ? "0" : 0;
                    }
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        if (typeof message.roomId === "number")
                            object.roomId = options.longs === String ? String(message.roomId) : message.roomId;
                        else
                            object.roomId = options.longs === String ? $util.Long.prototype.toString.call(message.roomId) : options.longs === Number ? new $util.LongBits(message.roomId.low >>> 0, message.roomId.high >>> 0).toNumber() : message.roomId;
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        if (typeof message.userId === "number")
                            object.userId = options.longs === String ? String(message.userId) : message.userId;
                        else
                            object.userId = options.longs === String ? $util.Long.prototype.toString.call(message.userId) : options.longs === Number ? new $util.LongBits(message.userId.low >>> 0, message.userId.high >>> 0).toNumber() : message.userId;
                    return object;
                };

                /**
                 * Converts this UserExitEvent to JSON.
                 * @function toJSON
                 * @memberof stmp.examples.room.UserExitEvent
                 * @instance
                 * @returns {Object.<string,*>} JSON object
                 */
                UserExitEvent.prototype.toJSON = function toJSON() {
                    return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                };

                return UserExitEvent;
            })();

            room.RoomEvents = (function() {

                /**
                 * Constructs a new RoomEvents service.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a RoomEvents
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function RoomEvents(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (RoomEvents.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = RoomEvents;

                /**
                 * Creates new RoomEvents service using the specified rpc implementation.
                 * @function create
                 * @memberof stmp.examples.room.RoomEvents
                 * @static
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 * @returns {RoomEvents} RPC service. Useful where requests and/or responses are streamed.
                 */
                RoomEvents.create = function create(rpcImpl, requestDelimited, responseDelimited) {
                    return new this(rpcImpl, requestDelimited, responseDelimited);
                };

                /**
                 * Callback as used by {@link stmp.examples.room.RoomEvents#userEnter}.
                 * @memberof stmp.examples.room.RoomEvents
                 * @typedef UserEnterCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {google.protobuf.Empty} [response] Empty
                 */

                /**
                 * Calls UserEnter.
                 * @function userEnter
                 * @memberof stmp.examples.room.RoomEvents
                 * @instance
                 * @param {stmp.examples.room.IUserEnterEvent} request UserEnterEvent message or plain object
                 * @param {stmp.examples.room.RoomEvents.UserEnterCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.userEnter = function userEnter(request, callback) {
                    return this.rpcCall(userEnter, $root.stmp.examples.room.UserEnterEvent, $root.google.protobuf.Empty, request, callback);
                }, "name", { value: "UserEnter" });

                /**
                 * Calls UserEnter.
                 * @function userEnter
                 * @memberof stmp.examples.room.RoomEvents
                 * @instance
                 * @param {stmp.examples.room.IUserEnterEvent} request UserEnterEvent message or plain object
                 * @returns {Promise<google.protobuf.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.room.RoomEvents#userExit}.
                 * @memberof stmp.examples.room.RoomEvents
                 * @typedef UserExitCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {google.protobuf.Empty} [response] Empty
                 */

                /**
                 * Calls UserExit.
                 * @function userExit
                 * @memberof stmp.examples.room.RoomEvents
                 * @instance
                 * @param {stmp.examples.room.IUserExitEvent} request UserExitEvent message or plain object
                 * @param {stmp.examples.room.RoomEvents.UserExitCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.userExit = function userExit(request, callback) {
                    return this.rpcCall(userExit, $root.stmp.examples.room.UserExitEvent, $root.google.protobuf.Empty, request, callback);
                }, "name", { value: "UserExit" });

                /**
                 * Calls UserExit.
                 * @function userExit
                 * @memberof stmp.examples.room.RoomEvents
                 * @instance
                 * @param {stmp.examples.room.IUserExitEvent} request UserExitEvent message or plain object
                 * @returns {Promise<google.protobuf.Empty>} Promise
                 * @variation 2
                 */

                return RoomEvents;
            })();

            return room;
        })();

        return examples;
    })();

    return stmp;
})();

export const google = $root.google = (() => {

    /**
     * Namespace google.
     * @exports google
     * @namespace
     */
    const google = {};

    google.protobuf = (function() {

        /**
         * Namespace protobuf.
         * @memberof google
         * @namespace
         */
        const protobuf = {};

        protobuf.Empty = (function() {

            /**
             * Properties of an Empty.
             * @memberof google.protobuf
             * @interface IEmpty
             */

            /**
             * Constructs a new Empty.
             * @memberof google.protobuf
             * @classdesc Represents an Empty.
             * @implements IEmpty
             * @constructor
             * @param {google.protobuf.IEmpty=} [properties] Properties to set
             */
            function Empty(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Creates a new Empty instance using the specified properties.
             * @function create
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.IEmpty=} [properties] Properties to set
             * @returns {google.protobuf.Empty} Empty instance
             */
            Empty.create = function create(properties) {
                return new Empty(properties);
            };

            /**
             * Encodes the specified Empty message. Does not implicitly {@link google.protobuf.Empty.verify|verify} messages.
             * @function encode
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.IEmpty} message Empty message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Empty.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                return writer;
            };

            /**
             * Encodes the specified Empty message, length delimited. Does not implicitly {@link google.protobuf.Empty.verify|verify} messages.
             * @function encodeDelimited
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.IEmpty} message Empty message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Empty.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes an Empty message from the specified reader or buffer.
             * @function decode
             * @memberof google.protobuf.Empty
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {google.protobuf.Empty} Empty
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Empty.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.google.protobuf.Empty();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Decodes an Empty message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof google.protobuf.Empty
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {google.protobuf.Empty} Empty
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Empty.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies an Empty message.
             * @function verify
             * @memberof google.protobuf.Empty
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Empty.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                return null;
            };

            /**
             * Creates an Empty message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof google.protobuf.Empty
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {google.protobuf.Empty} Empty
             */
            Empty.fromObject = function fromObject(object) {
                if (object instanceof $root.google.protobuf.Empty)
                    return object;
                return new $root.google.protobuf.Empty();
            };

            /**
             * Creates a plain object from an Empty message. Also converts values to other types if specified.
             * @function toObject
             * @memberof google.protobuf.Empty
             * @static
             * @param {google.protobuf.Empty} message Empty
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Empty.toObject = function toObject() {
                return {};
            };

            /**
             * Converts this Empty to JSON.
             * @function toJSON
             * @memberof google.protobuf.Empty
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Empty.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Empty;
        })();

        return protobuf;
    })();

    return google;
})();

export { $root as default };
