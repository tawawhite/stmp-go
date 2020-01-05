/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.stmp = (function() {

    /**
     * Namespace stmp.
     * @exports stmp
     * @namespace
     */
    var stmp = {};

    stmp.examples = (function() {

        /**
         * Namespace examples.
         * @memberof stmp
         * @namespace
         */
        var examples = {};

        examples.room = (function() {

            /**
             * Namespace room.
             * @memberof stmp.examples
             * @namespace
             */
            var room = {};

            room.UserModel = (function() {

                /**
                 * Properties of a UserModel.
                 * @memberof stmp.examples.room
                 * @interface IUserModel
                 * @property {string|null} [name] UserModel name
                 * @property {string|null} [room] UserModel room
                 * @property {stmp.examples.room.UserModel.Status|null} [status] UserModel status
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
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserModel name.
                 * @member {string} name
                 * @memberof stmp.examples.room.UserModel
                 * @instance
                 */
                UserModel.prototype.name = "";

                /**
                 * UserModel room.
                 * @member {string} room
                 * @memberof stmp.examples.room.UserModel
                 * @instance
                 */
                UserModel.prototype.room = "";

                /**
                 * UserModel status.
                 * @member {stmp.examples.room.UserModel.Status} status
                 * @memberof stmp.examples.room.UserModel
                 * @instance
                 */
                UserModel.prototype.status = 0;

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
                    if (message.name != null && message.hasOwnProperty("name"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
                    if (message.room != null && message.hasOwnProperty("room"))
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.room);
                    if (message.status != null && message.hasOwnProperty("status"))
                        writer.uint32(/* id 3, wireType 0 =*/24).int32(message.status);
                    return writer;
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.UserModel();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.name = reader.string();
                            break;
                        case 2:
                            message.room = reader.string();
                            break;
                        case 3:
                            message.status = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Status enum.
                 * @name stmp.examples.room.UserModel.Status
                 * @enum {string}
                 * @property {number} Offline=0 Offline value
                 * @property {number} Online=1 Online value
                 * @property {number} Chatting=2 Chatting value
                 * @property {number} ChattingOffline=3 ChattingOffline value
                 */
                UserModel.Status = (function() {
                    var valuesById = {}, values = Object.create(valuesById);
                    values[valuesById[0] = "Offline"] = 0;
                    values[valuesById[1] = "Online"] = 1;
                    values[valuesById[2] = "Chatting"] = 2;
                    values[valuesById[3] = "ChattingOffline"] = 3;
                    return values;
                })();

                return UserModel;
            })();

            room.LoginInput = (function() {

                /**
                 * Properties of a LoginInput.
                 * @memberof stmp.examples.room
                 * @interface ILoginInput
                 * @property {string|null} [name] LoginInput name
                 */

                /**
                 * Constructs a new LoginInput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a LoginInput.
                 * @implements ILoginInput
                 * @constructor
                 * @param {stmp.examples.room.ILoginInput=} [properties] Properties to set
                 */
                function LoginInput(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * LoginInput name.
                 * @member {string} name
                 * @memberof stmp.examples.room.LoginInput
                 * @instance
                 */
                LoginInput.prototype.name = "";

                /**
                 * Encodes the specified LoginInput message. Does not implicitly {@link stmp.examples.room.LoginInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.LoginInput
                 * @static
                 * @param {stmp.examples.room.ILoginInput} message LoginInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                LoginInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.name != null && message.hasOwnProperty("name"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
                    return writer;
                };

                /**
                 * Decodes a LoginInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.LoginInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.LoginInput} LoginInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                LoginInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.LoginInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
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

                return LoginInput;
            })();

            room.ListInput = (function() {

                /**
                 * Properties of a ListInput.
                 * @memberof stmp.examples.room
                 * @interface IListInput
                 * @property {number|Long|null} [limit] ListInput limit
                 * @property {number|Long|null} [offset] ListInput offset
                 */

                /**
                 * Constructs a new ListInput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a ListInput.
                 * @implements IListInput
                 * @constructor
                 * @param {stmp.examples.room.IListInput=} [properties] Properties to set
                 */
                function ListInput(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ListInput limit.
                 * @member {number|Long} limit
                 * @memberof stmp.examples.room.ListInput
                 * @instance
                 */
                ListInput.prototype.limit = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListInput offset.
                 * @member {number|Long} offset
                 * @memberof stmp.examples.room.ListInput
                 * @instance
                 */
                ListInput.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified ListInput message. Does not implicitly {@link stmp.examples.room.ListInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.ListInput
                 * @static
                 * @param {stmp.examples.room.IListInput} message ListInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.limit != null && message.hasOwnProperty("limit"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.limit);
                    if (message.offset != null && message.hasOwnProperty("offset"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int64(message.offset);
                    return writer;
                };

                /**
                 * Decodes a ListInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.ListInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.ListInput} ListInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ListInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
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

                return ListInput;
            })();

            room.ListUserOutput = (function() {

                /**
                 * Properties of a ListUserOutput.
                 * @memberof stmp.examples.room
                 * @interface IListUserOutput
                 * @property {number|Long|null} [total] ListUserOutput total
                 * @property {Array.<stmp.examples.room.IUserModel>|null} [users] ListUserOutput users
                 */

                /**
                 * Constructs a new ListUserOutput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a ListUserOutput.
                 * @implements IListUserOutput
                 * @constructor
                 * @param {stmp.examples.room.IListUserOutput=} [properties] Properties to set
                 */
                function ListUserOutput(properties) {
                    this.users = [];
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ListUserOutput total.
                 * @member {number|Long} total
                 * @memberof stmp.examples.room.ListUserOutput
                 * @instance
                 */
                ListUserOutput.prototype.total = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListUserOutput users.
                 * @member {Array.<stmp.examples.room.IUserModel>} users
                 * @memberof stmp.examples.room.ListUserOutput
                 * @instance
                 */
                ListUserOutput.prototype.users = $util.emptyArray;

                /**
                 * Encodes the specified ListUserOutput message. Does not implicitly {@link stmp.examples.room.ListUserOutput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.ListUserOutput
                 * @static
                 * @param {stmp.examples.room.IListUserOutput} message ListUserOutput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListUserOutput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.total != null && message.hasOwnProperty("total"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.total);
                    if (message.users != null && message.users.length)
                        for (var i = 0; i < message.users.length; ++i)
                            $root.stmp.examples.room.UserModel.encode(message.users[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
                };

                /**
                 * Decodes a ListUserOutput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.ListUserOutput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.ListUserOutput} ListUserOutput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListUserOutput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ListUserOutput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.total = reader.int64();
                            break;
                        case 2:
                            if (!(message.users && message.users.length))
                                message.users = [];
                            message.users.push($root.stmp.examples.room.UserModel.decode(reader, reader.uint32()));
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return ListUserOutput;
            })();

            room.UserService = (function() {

                /**
                 * Constructs a new UserService service.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a UserService
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function UserService(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (UserService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = UserService;

                /**
                 * Callback as used by {@link stmp.examples.room.UserService#listUser}.
                 * @memberof stmp.examples.room.UserService
                 * @typedef ListUserCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.room.ListUserOutput} [response] ListUserOutput
                 */

                /**
                 * Calls ListUser.
                 * @function listUser
                 * @memberof stmp.examples.room.UserService
                 * @instance
                 * @param {stmp.examples.room.IListInput} request ListInput message or plain object
                 * @param {stmp.examples.room.UserService.ListUserCallback} callback Node-style callback called with the error, if any, and ListUserOutput
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(UserService.prototype.listUser = function listUser(request, callback) {
                    return this.rpcCall(listUser, $root.stmp.examples.room.ListInput, $root.stmp.examples.room.ListUserOutput, request, callback);
                }, "name", { value: "ListUser" });

                /**
                 * Calls ListUser.
                 * @function listUser
                 * @memberof stmp.examples.room.UserService
                 * @instance
                 * @param {stmp.examples.room.IListInput} request ListInput message or plain object
                 * @returns {Promise<stmp.examples.room.ListUserOutput>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.room.UserService#login}.
                 * @memberof stmp.examples.room.UserService
                 * @typedef LoginCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.room.UserModel} [response] UserModel
                 */

                /**
                 * Calls Login.
                 * @function login
                 * @memberof stmp.examples.room.UserService
                 * @instance
                 * @param {stmp.examples.room.ILoginInput} request LoginInput message or plain object
                 * @param {stmp.examples.room.UserService.LoginCallback} callback Node-style callback called with the error, if any, and UserModel
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(UserService.prototype.login = function login(request, callback) {
                    return this.rpcCall(login, $root.stmp.examples.room.LoginInput, $root.stmp.examples.room.UserModel, request, callback);
                }, "name", { value: "Login" });

                /**
                 * Calls Login.
                 * @function login
                 * @memberof stmp.examples.room.UserService
                 * @instance
                 * @param {stmp.examples.room.ILoginInput} request LoginInput message or plain object
                 * @returns {Promise<stmp.examples.room.UserModel>} Promise
                 * @variation 2
                 */

                return UserService;
            })();

            room.UserEvents = (function() {

                /**
                 * Constructs a new UserEvents service.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a UserEvents
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function UserEvents(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (UserEvents.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = UserEvents;

                /**
                 * Callback as used by {@link stmp.examples.room.UserEvents#statusUpdated}.
                 * @memberof stmp.examples.room.UserEvents
                 * @typedef StatusUpdatedCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {google.protobuf.Empty} [response] Empty
                 */

                /**
                 * Calls StatusUpdated.
                 * @function statusUpdated
                 * @memberof stmp.examples.room.UserEvents
                 * @instance
                 * @param {stmp.examples.room.IUserModel} request UserModel message or plain object
                 * @param {stmp.examples.room.UserEvents.StatusUpdatedCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(UserEvents.prototype.statusUpdated = function statusUpdated(request, callback) {
                    return this.rpcCall(statusUpdated, $root.stmp.examples.room.UserModel, $root.google.protobuf.Empty, request, callback);
                }, "name", { value: "StatusUpdated" });

                /**
                 * Calls StatusUpdated.
                 * @function statusUpdated
                 * @memberof stmp.examples.room.UserEvents
                 * @instance
                 * @param {stmp.examples.room.IUserModel} request UserModel message or plain object
                 * @returns {Promise<google.protobuf.Empty>} Promise
                 * @variation 2
                 */

                return UserEvents;
            })();

            room.ChatMessageModel = (function() {

                /**
                 * Properties of a ChatMessageModel.
                 * @memberof stmp.examples.room
                 * @interface IChatMessageModel
                 * @property {string|null} [room] ChatMessageModel room
                 * @property {string|null} [user] ChatMessageModel user
                 * @property {string|null} [content] ChatMessageModel content
                 * @property {number|Long|null} [createdAt] ChatMessageModel createdAt
                 */

                /**
                 * Constructs a new ChatMessageModel.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a ChatMessageModel.
                 * @implements IChatMessageModel
                 * @constructor
                 * @param {stmp.examples.room.IChatMessageModel=} [properties] Properties to set
                 */
                function ChatMessageModel(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ChatMessageModel room.
                 * @member {string} room
                 * @memberof stmp.examples.room.ChatMessageModel
                 * @instance
                 */
                ChatMessageModel.prototype.room = "";

                /**
                 * ChatMessageModel user.
                 * @member {string} user
                 * @memberof stmp.examples.room.ChatMessageModel
                 * @instance
                 */
                ChatMessageModel.prototype.user = "";

                /**
                 * ChatMessageModel content.
                 * @member {string} content
                 * @memberof stmp.examples.room.ChatMessageModel
                 * @instance
                 */
                ChatMessageModel.prototype.content = "";

                /**
                 * ChatMessageModel createdAt.
                 * @member {number|Long} createdAt
                 * @memberof stmp.examples.room.ChatMessageModel
                 * @instance
                 */
                ChatMessageModel.prototype.createdAt = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified ChatMessageModel message. Does not implicitly {@link stmp.examples.room.ChatMessageModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.ChatMessageModel
                 * @static
                 * @param {stmp.examples.room.IChatMessageModel} message ChatMessageModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ChatMessageModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.room != null && message.hasOwnProperty("room"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.room);
                    if (message.user != null && message.hasOwnProperty("user"))
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.user);
                    if (message.content != null && message.hasOwnProperty("content"))
                        writer.uint32(/* id 3, wireType 2 =*/26).string(message.content);
                    if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                        writer.uint32(/* id 4, wireType 0 =*/32).int64(message.createdAt);
                    return writer;
                };

                /**
                 * Decodes a ChatMessageModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.ChatMessageModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.ChatMessageModel} ChatMessageModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ChatMessageModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ChatMessageModel();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.room = reader.string();
                            break;
                        case 2:
                            message.user = reader.string();
                            break;
                        case 3:
                            message.content = reader.string();
                            break;
                        case 4:
                            message.createdAt = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return ChatMessageModel;
            })();

            room.RoomModel = (function() {

                /**
                 * Properties of a RoomModel.
                 * @memberof stmp.examples.room
                 * @interface IRoomModel
                 * @property {string|null} [name] RoomModel name
                 * @property {Object.<string,stmp.examples.room.IUserModel>|null} [users] RoomModel users
                 * @property {Array.<stmp.examples.room.IChatMessageModel>|null} [messages] RoomModel messages
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
                    this.messages = [];
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

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
                 * RoomModel messages.
                 * @member {Array.<stmp.examples.room.IChatMessageModel>} messages
                 * @memberof stmp.examples.room.RoomModel
                 * @instance
                 */
                RoomModel.prototype.messages = $util.emptyArray;

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
                    if (message.name != null && message.hasOwnProperty("name"))
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                    if (message.users != null && message.hasOwnProperty("users"))
                        for (var keys = Object.keys(message.users), i = 0; i < keys.length; ++i) {
                            writer.uint32(/* id 3, wireType 2 =*/26).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]);
                            $root.stmp.examples.room.UserModel.encode(message.users[keys[i]], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim().ldelim();
                        }
                    if (message.messages != null && message.messages.length)
                        for (var i = 0; i < message.messages.length; ++i)
                            $root.stmp.examples.room.ChatMessageModel.encode(message.messages[i], writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
                    return writer;
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.RoomModel(), key;
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 2:
                            message.name = reader.string();
                            break;
                        case 3:
                            reader.skip().pos++;
                            if (message.users === $util.emptyObject)
                                message.users = {};
                            key = reader.string();
                            reader.pos++;
                            message.users[key] = $root.stmp.examples.room.UserModel.decode(reader, reader.uint32());
                            break;
                        case 4:
                            if (!(message.messages && message.messages.length))
                                message.messages = [];
                            message.messages.push($root.stmp.examples.room.ChatMessageModel.decode(reader, reader.uint32()));
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
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
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.CreateRoomInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
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

                return CreateRoomInput;
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
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
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
                        for (var i = 0; i < message.rooms.length; ++i)
                            $root.stmp.examples.room.RoomModel.encode(message.rooms[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ListRoomOutput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
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

                return ListRoomOutput;
            })();

            room.JoinRoomInput = (function() {

                /**
                 * Properties of a JoinRoomInput.
                 * @memberof stmp.examples.room
                 * @interface IJoinRoomInput
                 * @property {string|null} [room] JoinRoomInput room
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
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * JoinRoomInput room.
                 * @member {string} room
                 * @memberof stmp.examples.room.JoinRoomInput
                 * @instance
                 */
                JoinRoomInput.prototype.room = "";

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
                    if (message.room != null && message.hasOwnProperty("room"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.room);
                    return writer;
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.JoinRoomInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.room = reader.string();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return JoinRoomInput;
            })();

            room.ExitRoomInput = (function() {

                /**
                 * Properties of an ExitRoomInput.
                 * @memberof stmp.examples.room
                 * @interface IExitRoomInput
                 * @property {string|null} [room] ExitRoomInput room
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
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ExitRoomInput room.
                 * @member {string} room
                 * @memberof stmp.examples.room.ExitRoomInput
                 * @instance
                 */
                ExitRoomInput.prototype.room = "";

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
                    if (message.room != null && message.hasOwnProperty("room"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.room);
                    return writer;
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.ExitRoomInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.room = reader.string();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return ExitRoomInput;
            })();

            room.SendMessageInput = (function() {

                /**
                 * Properties of a SendMessageInput.
                 * @memberof stmp.examples.room
                 * @interface ISendMessageInput
                 * @property {string|null} [room] SendMessageInput room
                 * @property {string|null} [content] SendMessageInput content
                 */

                /**
                 * Constructs a new SendMessageInput.
                 * @memberof stmp.examples.room
                 * @classdesc Represents a SendMessageInput.
                 * @implements ISendMessageInput
                 * @constructor
                 * @param {stmp.examples.room.ISendMessageInput=} [properties] Properties to set
                 */
                function SendMessageInput(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * SendMessageInput room.
                 * @member {string} room
                 * @memberof stmp.examples.room.SendMessageInput
                 * @instance
                 */
                SendMessageInput.prototype.room = "";

                /**
                 * SendMessageInput content.
                 * @member {string} content
                 * @memberof stmp.examples.room.SendMessageInput
                 * @instance
                 */
                SendMessageInput.prototype.content = "";

                /**
                 * Encodes the specified SendMessageInput message. Does not implicitly {@link stmp.examples.room.SendMessageInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.room.SendMessageInput
                 * @static
                 * @param {stmp.examples.room.ISendMessageInput} message SendMessageInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                SendMessageInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.room != null && message.hasOwnProperty("room"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.room);
                    if (message.content != null && message.hasOwnProperty("content"))
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.content);
                    return writer;
                };

                /**
                 * Decodes a SendMessageInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.room.SendMessageInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.room.SendMessageInput} SendMessageInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                SendMessageInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.SendMessageInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.room = reader.string();
                            break;
                        case 2:
                            message.content = reader.string();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return SendMessageInput;
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
                 * @param {stmp.examples.room.IListInput} request ListInput message or plain object
                 * @param {stmp.examples.room.RoomService.ListRoomCallback} callback Node-style callback called with the error, if any, and ListRoomOutput
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.listRoom = function listRoom(request, callback) {
                    return this.rpcCall(listRoom, $root.stmp.examples.room.ListInput, $root.stmp.examples.room.ListRoomOutput, request, callback);
                }, "name", { value: "ListRoom" });

                /**
                 * Calls ListRoom.
                 * @function listRoom
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.IListInput} request ListInput message or plain object
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

                /**
                 * Callback as used by {@link stmp.examples.room.RoomService#sendMessage}.
                 * @memberof stmp.examples.room.RoomService
                 * @typedef SendMessageCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {google.protobuf.Empty} [response] Empty
                 */

                /**
                 * Calls SendMessage.
                 * @function sendMessage
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.ISendMessageInput} request SendMessageInput message or plain object
                 * @param {stmp.examples.room.RoomService.SendMessageCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.sendMessage = function sendMessage(request, callback) {
                    return this.rpcCall(sendMessage, $root.stmp.examples.room.SendMessageInput, $root.google.protobuf.Empty, request, callback);
                }, "name", { value: "SendMessage" });

                /**
                 * Calls SendMessage.
                 * @function sendMessage
                 * @memberof stmp.examples.room.RoomService
                 * @instance
                 * @param {stmp.examples.room.ISendMessageInput} request SendMessageInput message or plain object
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
                 * @property {string|null} [room] UserEnterEvent room
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
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserEnterEvent room.
                 * @member {string} room
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @instance
                 */
                UserEnterEvent.prototype.room = "";

                /**
                 * UserEnterEvent user.
                 * @member {stmp.examples.room.IUserModel|null|undefined} user
                 * @memberof stmp.examples.room.UserEnterEvent
                 * @instance
                 */
                UserEnterEvent.prototype.user = null;

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
                    if (message.room != null && message.hasOwnProperty("room"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.room);
                    if (message.user != null && message.hasOwnProperty("user"))
                        $root.stmp.examples.room.UserModel.encode(message.user, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.UserEnterEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.room = reader.string();
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

                return UserEnterEvent;
            })();

            room.UserExitEvent = (function() {

                /**
                 * Properties of a UserExitEvent.
                 * @memberof stmp.examples.room
                 * @interface IUserExitEvent
                 * @property {string|null} [room] UserExitEvent room
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
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserExitEvent room.
                 * @member {string} room
                 * @memberof stmp.examples.room.UserExitEvent
                 * @instance
                 */
                UserExitEvent.prototype.room = "";

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
                    if (message.room != null && message.hasOwnProperty("room"))
                        writer.uint32(/* id 1, wireType 2 =*/10).string(message.room);
                    return writer;
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
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.room.UserExitEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.room = reader.string();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
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

                /**
                 * Callback as used by {@link stmp.examples.room.RoomEvents#newMessage}.
                 * @memberof stmp.examples.room.RoomEvents
                 * @typedef NewMessageCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {google.protobuf.Empty} [response] Empty
                 */

                /**
                 * Calls NewMessage.
                 * @function newMessage
                 * @memberof stmp.examples.room.RoomEvents
                 * @instance
                 * @param {stmp.examples.room.IChatMessageModel} request ChatMessageModel message or plain object
                 * @param {stmp.examples.room.RoomEvents.NewMessageCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.newMessage = function newMessage(request, callback) {
                    return this.rpcCall(newMessage, $root.stmp.examples.room.ChatMessageModel, $root.google.protobuf.Empty, request, callback);
                }, "name", { value: "NewMessage" });

                /**
                 * Calls NewMessage.
                 * @function newMessage
                 * @memberof stmp.examples.room.RoomEvents
                 * @instance
                 * @param {stmp.examples.room.IChatMessageModel} request ChatMessageModel message or plain object
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

$root.google = (function() {

    /**
     * Namespace google.
     * @exports google
     * @namespace
     */
    var google = {};

    google.protobuf = (function() {

        /**
         * Namespace protobuf.
         * @memberof google
         * @namespace
         */
        var protobuf = {};

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
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

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
                var end = length === undefined ? reader.len : reader.pos + length, message = new $root.google.protobuf.Empty();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            return Empty;
        })();

        return protobuf;
    })();

    return google;
})();

module.exports = $root;
