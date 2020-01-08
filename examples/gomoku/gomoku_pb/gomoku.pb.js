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

        examples.gomoku = (function() {

            /**
             * Namespace gomoku.
             * @memberof stmp.examples
             * @namespace
             */
            var gomoku = {};

            gomoku.Empty = (function() {

                /**
                 * Properties of an Empty.
                 * @memberof stmp.examples.gomoku
                 * @interface IEmpty
                 */

                /**
                 * Constructs a new Empty.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents an Empty.
                 * @implements IEmpty
                 * @constructor
                 * @param {stmp.examples.gomoku.IEmpty=} [properties] Properties to set
                 */
                function Empty(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * Encodes the specified Empty message. Does not implicitly {@link stmp.examples.gomoku.Empty.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.Empty
                 * @static
                 * @param {stmp.examples.gomoku.IEmpty} message Empty message or plain object to encode
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
                 * @memberof stmp.examples.gomoku.Empty
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.Empty} Empty
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                Empty.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.Empty();
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

            gomoku.PlayerModel = (function() {

                /**
                 * Properties of a PlayerModel.
                 * @memberof stmp.examples.gomoku
                 * @interface IPlayerModel
                 * @property {number|Long|null} [id] PlayerModel id
                 * @property {string|null} [name] PlayerModel name
                 * @property {stmp.examples.gomoku.PlayerModel.Status|null} [status] PlayerModel status
                 * @property {number|Long|null} [roomId] PlayerModel roomId
                 * @property {number|null} [seat] PlayerModel seat
                 * @property {number|Long|null} [gameId] PlayerModel gameId
                 * @property {number|Long|null} [readyTimeout] PlayerModel readyTimeout
                 */

                /**
                 * Constructs a new PlayerModel.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a PlayerModel.
                 * @implements IPlayerModel
                 * @constructor
                 * @param {stmp.examples.gomoku.IPlayerModel=} [properties] Properties to set
                 */
                function PlayerModel(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * PlayerModel id.
                 * @member {number|Long} id
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @instance
                 */
                PlayerModel.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * PlayerModel name.
                 * @member {string} name
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @instance
                 */
                PlayerModel.prototype.name = "";

                /**
                 * PlayerModel status.
                 * @member {stmp.examples.gomoku.PlayerModel.Status} status
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @instance
                 */
                PlayerModel.prototype.status = 0;

                /**
                 * PlayerModel roomId.
                 * @member {number|Long} roomId
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @instance
                 */
                PlayerModel.prototype.roomId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * PlayerModel seat.
                 * @member {number} seat
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @instance
                 */
                PlayerModel.prototype.seat = 0;

                /**
                 * PlayerModel gameId.
                 * @member {number|Long} gameId
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @instance
                 */
                PlayerModel.prototype.gameId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * PlayerModel readyTimeout.
                 * @member {number|Long} readyTimeout
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @instance
                 */
                PlayerModel.prototype.readyTimeout = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified PlayerModel message. Does not implicitly {@link stmp.examples.gomoku.PlayerModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @static
                 * @param {stmp.examples.gomoku.IPlayerModel} message PlayerModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                PlayerModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.id != null && message.hasOwnProperty("id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
                    if (message.name != null && message.hasOwnProperty("name"))
                        writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                    if (message.status != null && message.hasOwnProperty("status"))
                        writer.uint32(/* id 3, wireType 0 =*/24).int32(message.status);
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        writer.uint32(/* id 4, wireType 0 =*/32).int64(message.roomId);
                    if (message.seat != null && message.hasOwnProperty("seat"))
                        writer.uint32(/* id 5, wireType 0 =*/40).int32(message.seat);
                    if (message.gameId != null && message.hasOwnProperty("gameId"))
                        writer.uint32(/* id 6, wireType 0 =*/48).int64(message.gameId);
                    if (message.readyTimeout != null && message.hasOwnProperty("readyTimeout"))
                        writer.uint32(/* id 7, wireType 0 =*/56).int64(message.readyTimeout);
                    return writer;
                };

                /**
                 * Decodes a PlayerModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.PlayerModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.PlayerModel} PlayerModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                PlayerModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.PlayerModel();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.id = reader.int64();
                            break;
                        case 2:
                            message.name = reader.string();
                            break;
                        case 3:
                            message.status = reader.int32();
                            break;
                        case 4:
                            message.roomId = reader.int64();
                            break;
                        case 5:
                            message.seat = reader.int32();
                            break;
                        case 6:
                            message.gameId = reader.int64();
                            break;
                        case 7:
                            message.readyTimeout = reader.int64();
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
                 * @name stmp.examples.gomoku.PlayerModel.Status
                 * @enum {string}
                 * @property {number} Reserved=0 Reserved value
                 * @property {number} Free=1 Free value
                 * @property {number} Standby=2 Standby value
                 * @property {number} Unready=3 Unready value
                 * @property {number} Ready=4 Ready value
                 * @property {number} Playing=5 Playing value
                 */
                PlayerModel.Status = (function() {
                    var valuesById = {}, values = Object.create(valuesById);
                    values[valuesById[0] = "Reserved"] = 0;
                    values[valuesById[1] = "Free"] = 1;
                    values[valuesById[2] = "Standby"] = 2;
                    values[valuesById[3] = "Unready"] = 3;
                    values[valuesById[4] = "Ready"] = 4;
                    values[valuesById[5] = "Playing"] = 5;
                    return values;
                })();

                return PlayerModel;
            })();

            gomoku.RoomModel = (function() {

                /**
                 * Properties of a RoomModel.
                 * @memberof stmp.examples.gomoku
                 * @interface IRoomModel
                 * @property {number|Long|null} [id] RoomModel id
                 * @property {Object.<string,number|Long>|null} [players] RoomModel players
                 * @property {number|Long|null} [gameId] RoomModel gameId
                 * @property {Array.<number|Long>|null} [spectators] RoomModel spectators
                 */

                /**
                 * Constructs a new RoomModel.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a RoomModel.
                 * @implements IRoomModel
                 * @constructor
                 * @param {stmp.examples.gomoku.IRoomModel=} [properties] Properties to set
                 */
                function RoomModel(properties) {
                    this.players = {};
                    this.spectators = [];
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * RoomModel id.
                 * @member {number|Long} id
                 * @memberof stmp.examples.gomoku.RoomModel
                 * @instance
                 */
                RoomModel.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * RoomModel players.
                 * @member {Object.<string,number|Long>} players
                 * @memberof stmp.examples.gomoku.RoomModel
                 * @instance
                 */
                RoomModel.prototype.players = $util.emptyObject;

                /**
                 * RoomModel gameId.
                 * @member {number|Long} gameId
                 * @memberof stmp.examples.gomoku.RoomModel
                 * @instance
                 */
                RoomModel.prototype.gameId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * RoomModel spectators.
                 * @member {Array.<number|Long>} spectators
                 * @memberof stmp.examples.gomoku.RoomModel
                 * @instance
                 */
                RoomModel.prototype.spectators = $util.emptyArray;

                /**
                 * Encodes the specified RoomModel message. Does not implicitly {@link stmp.examples.gomoku.RoomModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.RoomModel
                 * @static
                 * @param {stmp.examples.gomoku.IRoomModel} message RoomModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                RoomModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.id != null && message.hasOwnProperty("id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
                    if (message.players != null && message.hasOwnProperty("players"))
                        for (var keys = Object.keys(message.players), i = 0; i < keys.length; ++i)
                            writer.uint32(/* id 2, wireType 2 =*/18).fork().uint32(/* id 1, wireType 0 =*/8).int32(keys[i]).uint32(/* id 2, wireType 0 =*/16).int64(message.players[keys[i]]).ldelim();
                    if (message.gameId != null && message.hasOwnProperty("gameId"))
                        writer.uint32(/* id 3, wireType 0 =*/24).int64(message.gameId);
                    if (message.spectators != null && message.spectators.length) {
                        writer.uint32(/* id 4, wireType 2 =*/34).fork();
                        for (var i = 0; i < message.spectators.length; ++i)
                            writer.int64(message.spectators[i]);
                        writer.ldelim();
                    }
                    return writer;
                };

                /**
                 * Decodes a RoomModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.RoomModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.RoomModel} RoomModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                RoomModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.RoomModel(), key;
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.id = reader.int64();
                            break;
                        case 2:
                            reader.skip().pos++;
                            if (message.players === $util.emptyObject)
                                message.players = {};
                            key = reader.int32();
                            reader.pos++;
                            message.players[key] = reader.int64();
                            break;
                        case 3:
                            message.gameId = reader.int64();
                            break;
                        case 4:
                            if (!(message.spectators && message.spectators.length))
                                message.spectators = [];
                            if ((tag & 7) === 2) {
                                var end2 = reader.uint32() + reader.pos;
                                while (reader.pos < end2)
                                    message.spectators.push(reader.int64());
                            } else
                                message.spectators.push(reader.int64());
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Reasons enum.
                 * @name stmp.examples.gomoku.RoomModel.Reasons
                 * @enum {string}
                 * @property {number} Reserved=0 Reserved value
                 * @property {number} InRoomAlready=1 InRoomAlready value
                 */
                RoomModel.Reasons = (function() {
                    var valuesById = {}, values = Object.create(valuesById);
                    values[valuesById[0] = "Reserved"] = 0;
                    values[valuesById[1] = "InRoomAlready"] = 1;
                    return values;
                })();

                return RoomModel;
            })();

            gomoku.HandModel = (function() {

                /**
                 * Properties of a HandModel.
                 * @memberof stmp.examples.gomoku
                 * @interface IHandModel
                 * @property {number|null} [x] HandModel x
                 * @property {number|null} [y] HandModel y
                 * @property {number|null} [t] HandModel t
                 */

                /**
                 * Constructs a new HandModel.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a HandModel.
                 * @implements IHandModel
                 * @constructor
                 * @param {stmp.examples.gomoku.IHandModel=} [properties] Properties to set
                 */
                function HandModel(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * HandModel x.
                 * @member {number} x
                 * @memberof stmp.examples.gomoku.HandModel
                 * @instance
                 */
                HandModel.prototype.x = 0;

                /**
                 * HandModel y.
                 * @member {number} y
                 * @memberof stmp.examples.gomoku.HandModel
                 * @instance
                 */
                HandModel.prototype.y = 0;

                /**
                 * HandModel t.
                 * @member {number} t
                 * @memberof stmp.examples.gomoku.HandModel
                 * @instance
                 */
                HandModel.prototype.t = 0;

                /**
                 * Encodes the specified HandModel message. Does not implicitly {@link stmp.examples.gomoku.HandModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.HandModel
                 * @static
                 * @param {stmp.examples.gomoku.IHandModel} message HandModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                HandModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.x != null && message.hasOwnProperty("x"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int32(message.x);
                    if (message.y != null && message.hasOwnProperty("y"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int32(message.y);
                    if (message.t != null && message.hasOwnProperty("t"))
                        writer.uint32(/* id 3, wireType 0 =*/24).int32(message.t);
                    return writer;
                };

                /**
                 * Decodes a HandModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.HandModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.HandModel} HandModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                HandModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.HandModel();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.x = reader.int32();
                            break;
                        case 2:
                            message.y = reader.int32();
                            break;
                        case 3:
                            message.t = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return HandModel;
            })();

            gomoku.ApplyModel = (function() {

                /**
                 * Properties of an ApplyModel.
                 * @memberof stmp.examples.gomoku
                 * @interface IApplyModel
                 */

                /**
                 * Constructs a new ApplyModel.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents an ApplyModel.
                 * @implements IApplyModel
                 * @constructor
                 * @param {stmp.examples.gomoku.IApplyModel=} [properties] Properties to set
                 */
                function ApplyModel(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * Encodes the specified ApplyModel message. Does not implicitly {@link stmp.examples.gomoku.ApplyModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.ApplyModel
                 * @static
                 * @param {stmp.examples.gomoku.IApplyModel} message ApplyModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ApplyModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    return writer;
                };

                /**
                 * Decodes an ApplyModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.ApplyModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.ApplyModel} ApplyModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ApplyModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.ApplyModel();
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

                /**
                 * Kind enum.
                 * @name stmp.examples.gomoku.ApplyModel.Kind
                 * @enum {string}
                 * @property {number} Reserved=0 Reserved value
                 * @property {number} GiveUp=1 GiveUp value
                 * @property {number} Draw=2 Draw value
                 */
                ApplyModel.Kind = (function() {
                    var valuesById = {}, values = Object.create(valuesById);
                    values[valuesById[0] = "Reserved"] = 0;
                    values[valuesById[1] = "GiveUp"] = 1;
                    values[valuesById[2] = "Draw"] = 2;
                    return values;
                })();

                return ApplyModel;
            })();

            gomoku.GomokuModel = (function() {

                /**
                 * Properties of a GomokuModel.
                 * @memberof stmp.examples.gomoku
                 * @interface IGomokuModel
                 * @property {number|Long|null} [id] GomokuModel id
                 * @property {number|Long|null} [roomId] GomokuModel roomId
                 * @property {number|Long|null} [playerBlack] GomokuModel playerBlack
                 * @property {number|Long|null} [playerWhite] GomokuModel playerWhite
                 * @property {number|null} [seatBlack] GomokuModel seatBlack
                 * @property {number|null} [seatWhite] GomokuModel seatWhite
                 * @property {Array.<stmp.examples.gomoku.IHandModel>|null} [history] GomokuModel history
                 * @property {number|Long|null} [createdAt] GomokuModel createdAt
                 * @property {stmp.examples.gomoku.GomokuModel.Result|null} [result] GomokuModel result
                 * @property {number|Long|null} [winner] GomokuModel winner
                 */

                /**
                 * Constructs a new GomokuModel.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a GomokuModel.
                 * @implements IGomokuModel
                 * @constructor
                 * @param {stmp.examples.gomoku.IGomokuModel=} [properties] Properties to set
                 */
                function GomokuModel(properties) {
                    this.history = [];
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * GomokuModel id.
                 * @member {number|Long} id
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * GomokuModel roomId.
                 * @member {number|Long} roomId
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.roomId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * GomokuModel playerBlack.
                 * @member {number|Long} playerBlack
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.playerBlack = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * GomokuModel playerWhite.
                 * @member {number|Long} playerWhite
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.playerWhite = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * GomokuModel seatBlack.
                 * @member {number} seatBlack
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.seatBlack = 0;

                /**
                 * GomokuModel seatWhite.
                 * @member {number} seatWhite
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.seatWhite = 0;

                /**
                 * GomokuModel history.
                 * @member {Array.<stmp.examples.gomoku.IHandModel>} history
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.history = $util.emptyArray;

                /**
                 * GomokuModel createdAt.
                 * @member {number|Long} createdAt
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.createdAt = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * GomokuModel result.
                 * @member {stmp.examples.gomoku.GomokuModel.Result} result
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.result = 0;

                /**
                 * GomokuModel winner.
                 * @member {number|Long} winner
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @instance
                 */
                GomokuModel.prototype.winner = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified GomokuModel message. Does not implicitly {@link stmp.examples.gomoku.GomokuModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @static
                 * @param {stmp.examples.gomoku.IGomokuModel} message GomokuModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                GomokuModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.id != null && message.hasOwnProperty("id"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int64(message.roomId);
                    if (message.playerBlack != null && message.hasOwnProperty("playerBlack"))
                        writer.uint32(/* id 3, wireType 0 =*/24).int64(message.playerBlack);
                    if (message.playerWhite != null && message.hasOwnProperty("playerWhite"))
                        writer.uint32(/* id 4, wireType 0 =*/32).int64(message.playerWhite);
                    if (message.seatBlack != null && message.hasOwnProperty("seatBlack"))
                        writer.uint32(/* id 5, wireType 0 =*/40).int32(message.seatBlack);
                    if (message.seatWhite != null && message.hasOwnProperty("seatWhite"))
                        writer.uint32(/* id 6, wireType 0 =*/48).int32(message.seatWhite);
                    if (message.history != null && message.history.length)
                        for (var i = 0; i < message.history.length; ++i)
                            $root.stmp.examples.gomoku.HandModel.encode(message.history[i], writer.uint32(/* id 7, wireType 2 =*/58).fork()).ldelim();
                    if (message.createdAt != null && message.hasOwnProperty("createdAt"))
                        writer.uint32(/* id 8, wireType 0 =*/64).int64(message.createdAt);
                    if (message.result != null && message.hasOwnProperty("result"))
                        writer.uint32(/* id 9, wireType 0 =*/72).int32(message.result);
                    if (message.winner != null && message.hasOwnProperty("winner"))
                        writer.uint32(/* id 10, wireType 0 =*/80).int64(message.winner);
                    return writer;
                };

                /**
                 * Decodes a GomokuModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.GomokuModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.GomokuModel} GomokuModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                GomokuModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.GomokuModel();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.id = reader.int64();
                            break;
                        case 2:
                            message.roomId = reader.int64();
                            break;
                        case 3:
                            message.playerBlack = reader.int64();
                            break;
                        case 4:
                            message.playerWhite = reader.int64();
                            break;
                        case 5:
                            message.seatBlack = reader.int32();
                            break;
                        case 6:
                            message.seatWhite = reader.int32();
                            break;
                        case 7:
                            if (!(message.history && message.history.length))
                                message.history = [];
                            message.history.push($root.stmp.examples.gomoku.HandModel.decode(reader, reader.uint32()));
                            break;
                        case 8:
                            message.createdAt = reader.int64();
                            break;
                        case 9:
                            message.result = reader.int32();
                            break;
                        case 10:
                            message.winner = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                /**
                 * Result enum.
                 * @name stmp.examples.gomoku.GomokuModel.Result
                 * @enum {string}
                 * @property {number} Playing=0 Playing value
                 * @property {number} Win=1 Win value
                 * @property {number} Draw=2 Draw value
                 * @property {number} ApplyGiveUp=3 ApplyGiveUp value
                 * @property {number} ApplyDraw=4 ApplyDraw value
                 * @property {number} UserStepTimeout=5 UserStepTimeout value
                 * @property {number} UserTotalTimeout=6 UserTotalTimeout value
                 */
                GomokuModel.Result = (function() {
                    var valuesById = {}, values = Object.create(valuesById);
                    values[valuesById[0] = "Playing"] = 0;
                    values[valuesById[1] = "Win"] = 1;
                    values[valuesById[2] = "Draw"] = 2;
                    values[valuesById[3] = "ApplyGiveUp"] = 3;
                    values[valuesById[4] = "ApplyDraw"] = 4;
                    values[valuesById[5] = "UserStepTimeout"] = 5;
                    values[valuesById[6] = "UserTotalTimeout"] = 6;
                    return values;
                })();

                return GomokuModel;
            })();

            gomoku.FullRoomModel = (function() {

                /**
                 * Properties of a FullRoomModel.
                 * @memberof stmp.examples.gomoku
                 * @interface IFullRoomModel
                 * @property {stmp.examples.gomoku.IRoomModel|null} [room] FullRoomModel room
                 * @property {Object.<string,stmp.examples.gomoku.IPlayerModel>|null} [players] FullRoomModel players
                 * @property {stmp.examples.gomoku.IGomokuModel|null} [game] FullRoomModel game
                 */

                /**
                 * Constructs a new FullRoomModel.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a FullRoomModel.
                 * @implements IFullRoomModel
                 * @constructor
                 * @param {stmp.examples.gomoku.IFullRoomModel=} [properties] Properties to set
                 */
                function FullRoomModel(properties) {
                    this.players = {};
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * FullRoomModel room.
                 * @member {stmp.examples.gomoku.IRoomModel|null|undefined} room
                 * @memberof stmp.examples.gomoku.FullRoomModel
                 * @instance
                 */
                FullRoomModel.prototype.room = null;

                /**
                 * FullRoomModel players.
                 * @member {Object.<string,stmp.examples.gomoku.IPlayerModel>} players
                 * @memberof stmp.examples.gomoku.FullRoomModel
                 * @instance
                 */
                FullRoomModel.prototype.players = $util.emptyObject;

                /**
                 * FullRoomModel game.
                 * @member {stmp.examples.gomoku.IGomokuModel|null|undefined} game
                 * @memberof stmp.examples.gomoku.FullRoomModel
                 * @instance
                 */
                FullRoomModel.prototype.game = null;

                /**
                 * Encodes the specified FullRoomModel message. Does not implicitly {@link stmp.examples.gomoku.FullRoomModel.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.FullRoomModel
                 * @static
                 * @param {stmp.examples.gomoku.IFullRoomModel} message FullRoomModel message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                FullRoomModel.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.room != null && message.hasOwnProperty("room"))
                        $root.stmp.examples.gomoku.RoomModel.encode(message.room, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
                    if (message.players != null && message.hasOwnProperty("players"))
                        for (var keys = Object.keys(message.players), i = 0; i < keys.length; ++i) {
                            writer.uint32(/* id 2, wireType 2 =*/18).fork().uint32(/* id 1, wireType 0 =*/8).int64(keys[i]);
                            $root.stmp.examples.gomoku.PlayerModel.encode(message.players[keys[i]], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim().ldelim();
                        }
                    if (message.game != null && message.hasOwnProperty("game"))
                        $root.stmp.examples.gomoku.GomokuModel.encode(message.game, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
                    return writer;
                };

                /**
                 * Decodes a FullRoomModel message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.FullRoomModel
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.FullRoomModel} FullRoomModel
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                FullRoomModel.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.FullRoomModel(), key;
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.room = $root.stmp.examples.gomoku.RoomModel.decode(reader, reader.uint32());
                            break;
                        case 2:
                            reader.skip().pos++;
                            if (message.players === $util.emptyObject)
                                message.players = {};
                            key = reader.int64();
                            reader.pos++;
                            message.players[typeof key === "object" ? $util.longToHash(key) : key] = $root.stmp.examples.gomoku.PlayerModel.decode(reader, reader.uint32());
                            break;
                        case 3:
                            message.game = $root.stmp.examples.gomoku.GomokuModel.decode(reader, reader.uint32());
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return FullRoomModel;
            })();

            gomoku.ListInput = (function() {

                /**
                 * Properties of a ListInput.
                 * @memberof stmp.examples.gomoku
                 * @interface IListInput
                 * @property {number|Long|null} [limit] ListInput limit
                 * @property {number|Long|null} [offset] ListInput offset
                 */

                /**
                 * Constructs a new ListInput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a ListInput.
                 * @implements IListInput
                 * @constructor
                 * @param {stmp.examples.gomoku.IListInput=} [properties] Properties to set
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
                 * @memberof stmp.examples.gomoku.ListInput
                 * @instance
                 */
                ListInput.prototype.limit = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListInput offset.
                 * @member {number|Long} offset
                 * @memberof stmp.examples.gomoku.ListInput
                 * @instance
                 */
                ListInput.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified ListInput message. Does not implicitly {@link stmp.examples.gomoku.ListInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.ListInput
                 * @static
                 * @param {stmp.examples.gomoku.IListInput} message ListInput message or plain object to encode
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
                 * @memberof stmp.examples.gomoku.ListInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.ListInput} ListInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.ListInput();
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

            gomoku.ListRoomOutput = (function() {

                /**
                 * Properties of a ListRoomOutput.
                 * @memberof stmp.examples.gomoku
                 * @interface IListRoomOutput
                 * @property {number|Long|null} [total] ListRoomOutput total
                 * @property {Array.<stmp.examples.gomoku.IRoomModel>|null} [rooms] ListRoomOutput rooms
                 */

                /**
                 * Constructs a new ListRoomOutput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a ListRoomOutput.
                 * @implements IListRoomOutput
                 * @constructor
                 * @param {stmp.examples.gomoku.IListRoomOutput=} [properties] Properties to set
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
                 * @memberof stmp.examples.gomoku.ListRoomOutput
                 * @instance
                 */
                ListRoomOutput.prototype.total = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListRoomOutput rooms.
                 * @member {Array.<stmp.examples.gomoku.IRoomModel>} rooms
                 * @memberof stmp.examples.gomoku.ListRoomOutput
                 * @instance
                 */
                ListRoomOutput.prototype.rooms = $util.emptyArray;

                /**
                 * Encodes the specified ListRoomOutput message. Does not implicitly {@link stmp.examples.gomoku.ListRoomOutput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.ListRoomOutput
                 * @static
                 * @param {stmp.examples.gomoku.IListRoomOutput} message ListRoomOutput message or plain object to encode
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
                            $root.stmp.examples.gomoku.RoomModel.encode(message.rooms[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
                };

                /**
                 * Decodes a ListRoomOutput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.ListRoomOutput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.ListRoomOutput} ListRoomOutput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListRoomOutput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.ListRoomOutput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.total = reader.int64();
                            break;
                        case 2:
                            if (!(message.rooms && message.rooms.length))
                                message.rooms = [];
                            message.rooms.push($root.stmp.examples.gomoku.RoomModel.decode(reader, reader.uint32()));
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

            gomoku.LookonRoomInput = (function() {

                /**
                 * Properties of a LookonRoomInput.
                 * @memberof stmp.examples.gomoku
                 * @interface ILookonRoomInput
                 * @property {number|Long|null} [roomId] LookonRoomInput roomId
                 */

                /**
                 * Constructs a new LookonRoomInput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a LookonRoomInput.
                 * @implements ILookonRoomInput
                 * @constructor
                 * @param {stmp.examples.gomoku.ILookonRoomInput=} [properties] Properties to set
                 */
                function LookonRoomInput(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * LookonRoomInput roomId.
                 * @member {number|Long} roomId
                 * @memberof stmp.examples.gomoku.LookonRoomInput
                 * @instance
                 */
                LookonRoomInput.prototype.roomId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified LookonRoomInput message. Does not implicitly {@link stmp.examples.gomoku.LookonRoomInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.LookonRoomInput
                 * @static
                 * @param {stmp.examples.gomoku.ILookonRoomInput} message LookonRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                LookonRoomInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.roomId);
                    return writer;
                };

                /**
                 * Decodes a LookonRoomInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.LookonRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.LookonRoomInput} LookonRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                LookonRoomInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.LookonRoomInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.roomId = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return LookonRoomInput;
            })();

            gomoku.JoinRoomInput = (function() {

                /**
                 * Properties of a JoinRoomInput.
                 * @memberof stmp.examples.gomoku
                 * @interface IJoinRoomInput
                 * @property {number|Long|null} [roomId] JoinRoomInput roomId
                 * @property {number|null} [preferSeat] JoinRoomInput preferSeat
                 * @property {number|null} [seat] JoinRoomInput seat
                 */

                /**
                 * Constructs a new JoinRoomInput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a JoinRoomInput.
                 * @implements IJoinRoomInput
                 * @constructor
                 * @param {stmp.examples.gomoku.IJoinRoomInput=} [properties] Properties to set
                 */
                function JoinRoomInput(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * JoinRoomInput roomId.
                 * @member {number|Long} roomId
                 * @memberof stmp.examples.gomoku.JoinRoomInput
                 * @instance
                 */
                JoinRoomInput.prototype.roomId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * JoinRoomInput preferSeat.
                 * @member {number} preferSeat
                 * @memberof stmp.examples.gomoku.JoinRoomInput
                 * @instance
                 */
                JoinRoomInput.prototype.preferSeat = 0;

                /**
                 * JoinRoomInput seat.
                 * @member {number} seat
                 * @memberof stmp.examples.gomoku.JoinRoomInput
                 * @instance
                 */
                JoinRoomInput.prototype.seat = 0;

                /**
                 * Encodes the specified JoinRoomInput message. Does not implicitly {@link stmp.examples.gomoku.JoinRoomInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.JoinRoomInput
                 * @static
                 * @param {stmp.examples.gomoku.IJoinRoomInput} message JoinRoomInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                JoinRoomInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.roomId != null && message.hasOwnProperty("roomId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.roomId);
                    if (message.preferSeat != null && message.hasOwnProperty("preferSeat"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int32(message.preferSeat);
                    if (message.seat != null && message.hasOwnProperty("seat"))
                        writer.uint32(/* id 3, wireType 0 =*/24).int32(message.seat);
                    return writer;
                };

                /**
                 * Decodes a JoinRoomInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.JoinRoomInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.JoinRoomInput} JoinRoomInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                JoinRoomInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.JoinRoomInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.roomId = reader.int64();
                            break;
                        case 2:
                            message.preferSeat = reader.int32();
                            break;
                        case 3:
                            message.seat = reader.int32();
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

            gomoku.RoomService = (function() {

                /**
                 * Constructs a new RoomService service.
                 * @memberof stmp.examples.gomoku
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
                 * Callback as used by {@link stmp.examples.gomoku.RoomService#matchRoom}.
                 * @memberof stmp.examples.gomoku.RoomService
                 * @typedef MatchRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.FullRoomModel} [response] FullRoomModel
                 */

                /**
                 * Calls MatchRoom.
                 * @function matchRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @param {stmp.examples.gomoku.RoomService.MatchRoomCallback} callback Node-style callback called with the error, if any, and FullRoomModel
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.matchRoom = function matchRoom(request, callback) {
                    return this.rpcCall(matchRoom, $root.stmp.examples.gomoku.Empty, $root.stmp.examples.gomoku.FullRoomModel, request, callback);
                }, "name", { value: "MatchRoom" });

                /**
                 * Calls MatchRoom.
                 * @function matchRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @returns {Promise<stmp.examples.gomoku.FullRoomModel>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomService#listRoom}.
                 * @memberof stmp.examples.gomoku.RoomService
                 * @typedef ListRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.ListRoomOutput} [response] ListRoomOutput
                 */

                /**
                 * Calls ListRoom.
                 * @function listRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IListInput} request ListInput message or plain object
                 * @param {stmp.examples.gomoku.RoomService.ListRoomCallback} callback Node-style callback called with the error, if any, and ListRoomOutput
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.listRoom = function listRoom(request, callback) {
                    return this.rpcCall(listRoom, $root.stmp.examples.gomoku.ListInput, $root.stmp.examples.gomoku.ListRoomOutput, request, callback);
                }, "name", { value: "ListRoom" });

                /**
                 * Calls ListRoom.
                 * @function listRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IListInput} request ListInput message or plain object
                 * @returns {Promise<stmp.examples.gomoku.ListRoomOutput>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomService#lookonRoom}.
                 * @memberof stmp.examples.gomoku.RoomService
                 * @typedef LookonRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.FullRoomModel} [response] FullRoomModel
                 */

                /**
                 * Calls LookonRoom.
                 * @function lookonRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.ILookonRoomInput} request LookonRoomInput message or plain object
                 * @param {stmp.examples.gomoku.RoomService.LookonRoomCallback} callback Node-style callback called with the error, if any, and FullRoomModel
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.lookonRoom = function lookonRoom(request, callback) {
                    return this.rpcCall(lookonRoom, $root.stmp.examples.gomoku.LookonRoomInput, $root.stmp.examples.gomoku.FullRoomModel, request, callback);
                }, "name", { value: "LookonRoom" });

                /**
                 * Calls LookonRoom.
                 * @function lookonRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.ILookonRoomInput} request LookonRoomInput message or plain object
                 * @returns {Promise<stmp.examples.gomoku.FullRoomModel>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomService#joinRoom}.
                 * @memberof stmp.examples.gomoku.RoomService
                 * @typedef JoinRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.FullRoomModel} [response] FullRoomModel
                 */

                /**
                 * Calls JoinRoom.
                 * @function joinRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IJoinRoomInput} request JoinRoomInput message or plain object
                 * @param {stmp.examples.gomoku.RoomService.JoinRoomCallback} callback Node-style callback called with the error, if any, and FullRoomModel
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.joinRoom = function joinRoom(request, callback) {
                    return this.rpcCall(joinRoom, $root.stmp.examples.gomoku.JoinRoomInput, $root.stmp.examples.gomoku.FullRoomModel, request, callback);
                }, "name", { value: "JoinRoom" });

                /**
                 * Calls JoinRoom.
                 * @function joinRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IJoinRoomInput} request JoinRoomInput message or plain object
                 * @returns {Promise<stmp.examples.gomoku.FullRoomModel>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomService#ready}.
                 * @memberof stmp.examples.gomoku.RoomService
                 * @typedef ReadyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls Ready.
                 * @function ready
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @param {stmp.examples.gomoku.RoomService.ReadyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.ready = function ready(request, callback) {
                    return this.rpcCall(ready, $root.stmp.examples.gomoku.Empty, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "Ready" });

                /**
                 * Calls Ready.
                 * @function ready
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomService#unready}.
                 * @memberof stmp.examples.gomoku.RoomService
                 * @typedef UnreadyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls Unready.
                 * @function unready
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @param {stmp.examples.gomoku.RoomService.UnreadyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.unready = function unready(request, callback) {
                    return this.rpcCall(unready, $root.stmp.examples.gomoku.Empty, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "Unready" });

                /**
                 * Calls Unready.
                 * @function unready
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomService#exitRoom}.
                 * @memberof stmp.examples.gomoku.RoomService
                 * @typedef ExitRoomCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls ExitRoom.
                 * @function exitRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @param {stmp.examples.gomoku.RoomService.ExitRoomCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomService.prototype.exitRoom = function exitRoom(request, callback) {
                    return this.rpcCall(exitRoom, $root.stmp.examples.gomoku.Empty, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "ExitRoom" });

                /**
                 * Calls ExitRoom.
                 * @function exitRoom
                 * @memberof stmp.examples.gomoku.RoomService
                 * @instance
                 * @param {stmp.examples.gomoku.IEmpty} request Empty message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                return RoomService;
            })();

            gomoku.UserJoinEvent = (function() {

                /**
                 * Properties of a UserJoinEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserJoinEvent
                 * @property {number|Long|null} [userId] UserJoinEvent userId
                 * @property {number|null} [seat] UserJoinEvent seat
                 * @property {number|null} [readyTimeout] UserJoinEvent readyTimeout
                 */

                /**
                 * Constructs a new UserJoinEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserJoinEvent.
                 * @implements IUserJoinEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserJoinEvent=} [properties] Properties to set
                 */
                function UserJoinEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserJoinEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserJoinEvent
                 * @instance
                 */
                UserJoinEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * UserJoinEvent seat.
                 * @member {number} seat
                 * @memberof stmp.examples.gomoku.UserJoinEvent
                 * @instance
                 */
                UserJoinEvent.prototype.seat = 0;

                /**
                 * UserJoinEvent readyTimeout.
                 * @member {number} readyTimeout
                 * @memberof stmp.examples.gomoku.UserJoinEvent
                 * @instance
                 */
                UserJoinEvent.prototype.readyTimeout = 0;

                /**
                 * Encodes the specified UserJoinEvent message. Does not implicitly {@link stmp.examples.gomoku.UserJoinEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserJoinEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserJoinEvent} message UserJoinEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserJoinEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    if (message.seat != null && message.hasOwnProperty("seat"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int32(message.seat);
                    if (message.readyTimeout != null && message.hasOwnProperty("readyTimeout"))
                        writer.uint32(/* id 3, wireType 0 =*/24).int32(message.readyTimeout);
                    return writer;
                };

                /**
                 * Decodes a UserJoinEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserJoinEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserJoinEvent} UserJoinEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserJoinEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserJoinEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
                            break;
                        case 2:
                            message.seat = reader.int32();
                            break;
                        case 3:
                            message.readyTimeout = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserJoinEvent;
            })();

            gomoku.UserReadyEvent = (function() {

                /**
                 * Properties of a UserReadyEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserReadyEvent
                 * @property {number|Long|null} [userId] UserReadyEvent userId
                 */

                /**
                 * Constructs a new UserReadyEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserReadyEvent.
                 * @implements IUserReadyEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserReadyEvent=} [properties] Properties to set
                 */
                function UserReadyEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserReadyEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserReadyEvent
                 * @instance
                 */
                UserReadyEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified UserReadyEvent message. Does not implicitly {@link stmp.examples.gomoku.UserReadyEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserReadyEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserReadyEvent} message UserReadyEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserReadyEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    return writer;
                };

                /**
                 * Decodes a UserReadyEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserReadyEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserReadyEvent} UserReadyEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserReadyEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserReadyEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserReadyEvent;
            })();

            gomoku.UserUnreadyEvent = (function() {

                /**
                 * Properties of a UserUnreadyEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserUnreadyEvent
                 * @property {number|Long|null} [userId] UserUnreadyEvent userId
                 * @property {number|null} [readyTimeout] UserUnreadyEvent readyTimeout
                 */

                /**
                 * Constructs a new UserUnreadyEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserUnreadyEvent.
                 * @implements IUserUnreadyEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserUnreadyEvent=} [properties] Properties to set
                 */
                function UserUnreadyEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserUnreadyEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserUnreadyEvent
                 * @instance
                 */
                UserUnreadyEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * UserUnreadyEvent readyTimeout.
                 * @member {number} readyTimeout
                 * @memberof stmp.examples.gomoku.UserUnreadyEvent
                 * @instance
                 */
                UserUnreadyEvent.prototype.readyTimeout = 0;

                /**
                 * Encodes the specified UserUnreadyEvent message. Does not implicitly {@link stmp.examples.gomoku.UserUnreadyEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserUnreadyEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserUnreadyEvent} message UserUnreadyEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserUnreadyEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    if (message.readyTimeout != null && message.hasOwnProperty("readyTimeout"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int32(message.readyTimeout);
                    return writer;
                };

                /**
                 * Decodes a UserUnreadyEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserUnreadyEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserUnreadyEvent} UserUnreadyEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserUnreadyEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserUnreadyEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
                            break;
                        case 2:
                            message.readyTimeout = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserUnreadyEvent;
            })();

            gomoku.UserLookonEvent = (function() {

                /**
                 * Properties of a UserLookonEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserLookonEvent
                 * @property {number|Long|null} [userId] UserLookonEvent userId
                 */

                /**
                 * Constructs a new UserLookonEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserLookonEvent.
                 * @implements IUserLookonEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserLookonEvent=} [properties] Properties to set
                 */
                function UserLookonEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserLookonEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserLookonEvent
                 * @instance
                 */
                UserLookonEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified UserLookonEvent message. Does not implicitly {@link stmp.examples.gomoku.UserLookonEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserLookonEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserLookonEvent} message UserLookonEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserLookonEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    return writer;
                };

                /**
                 * Decodes a UserLookonEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserLookonEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserLookonEvent} UserLookonEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserLookonEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserLookonEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserLookonEvent;
            })();

            gomoku.UserExitEvent = (function() {

                /**
                 * Properties of a UserExitEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserExitEvent
                 * @property {number|Long|null} [userId] UserExitEvent userId
                 */

                /**
                 * Constructs a new UserExitEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserExitEvent.
                 * @implements IUserExitEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserExitEvent=} [properties] Properties to set
                 */
                function UserExitEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserExitEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserExitEvent
                 * @instance
                 */
                UserExitEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified UserExitEvent message. Does not implicitly {@link stmp.examples.gomoku.UserExitEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserExitEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserExitEvent} message UserExitEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserExitEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    return writer;
                };

                /**
                 * Decodes a UserExitEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserExitEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserExitEvent} UserExitEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserExitEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserExitEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
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

            gomoku.RoomEvents = (function() {

                /**
                 * Constructs a new RoomEvents service.
                 * @memberof stmp.examples.gomoku
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
                 * Callback as used by {@link stmp.examples.gomoku.RoomEvents#userJoin}.
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @typedef UserJoinCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserJoin.
                 * @function userJoin
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserJoinEvent} request UserJoinEvent message or plain object
                 * @param {stmp.examples.gomoku.RoomEvents.UserJoinCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.userJoin = function userJoin(request, callback) {
                    return this.rpcCall(userJoin, $root.stmp.examples.gomoku.UserJoinEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserJoin" });

                /**
                 * Calls UserJoin.
                 * @function userJoin
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserJoinEvent} request UserJoinEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomEvents#userReady}.
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @typedef UserReadyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserReady.
                 * @function userReady
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserReadyEvent} request UserReadyEvent message or plain object
                 * @param {stmp.examples.gomoku.RoomEvents.UserReadyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.userReady = function userReady(request, callback) {
                    return this.rpcCall(userReady, $root.stmp.examples.gomoku.UserReadyEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserReady" });

                /**
                 * Calls UserReady.
                 * @function userReady
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserReadyEvent} request UserReadyEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomEvents#userUnready}.
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @typedef UserUnreadyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserUnready.
                 * @function userUnready
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserUnreadyEvent} request UserUnreadyEvent message or plain object
                 * @param {stmp.examples.gomoku.RoomEvents.UserUnreadyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.userUnready = function userUnready(request, callback) {
                    return this.rpcCall(userUnready, $root.stmp.examples.gomoku.UserUnreadyEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserUnready" });

                /**
                 * Calls UserUnready.
                 * @function userUnready
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserUnreadyEvent} request UserUnreadyEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomEvents#userLookon}.
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @typedef UserLookonCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserLookon.
                 * @function userLookon
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserLookonEvent} request UserLookonEvent message or plain object
                 * @param {stmp.examples.gomoku.RoomEvents.UserLookonCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.userLookon = function userLookon(request, callback) {
                    return this.rpcCall(userLookon, $root.stmp.examples.gomoku.UserLookonEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserLookon" });

                /**
                 * Calls UserLookon.
                 * @function userLookon
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserLookonEvent} request UserLookonEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.RoomEvents#userExit}.
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @typedef UserExitCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserExit.
                 * @function userExit
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserExitEvent} request UserExitEvent message or plain object
                 * @param {stmp.examples.gomoku.RoomEvents.UserExitCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(RoomEvents.prototype.userExit = function userExit(request, callback) {
                    return this.rpcCall(userExit, $root.stmp.examples.gomoku.UserExitEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserExit" });

                /**
                 * Calls UserExit.
                 * @function userExit
                 * @memberof stmp.examples.gomoku.RoomEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserExitEvent} request UserExitEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                return RoomEvents;
            })();

            gomoku.ApplyInput = (function() {

                /**
                 * Properties of an ApplyInput.
                 * @memberof stmp.examples.gomoku
                 * @interface IApplyInput
                 * @property {stmp.examples.gomoku.ApplyModel.Kind|null} [kind] ApplyInput kind
                 */

                /**
                 * Constructs a new ApplyInput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents an ApplyInput.
                 * @implements IApplyInput
                 * @constructor
                 * @param {stmp.examples.gomoku.IApplyInput=} [properties] Properties to set
                 */
                function ApplyInput(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ApplyInput kind.
                 * @member {stmp.examples.gomoku.ApplyModel.Kind} kind
                 * @memberof stmp.examples.gomoku.ApplyInput
                 * @instance
                 */
                ApplyInput.prototype.kind = 0;

                /**
                 * Encodes the specified ApplyInput message. Does not implicitly {@link stmp.examples.gomoku.ApplyInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.ApplyInput
                 * @static
                 * @param {stmp.examples.gomoku.IApplyInput} message ApplyInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ApplyInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.kind != null && message.hasOwnProperty("kind"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int32(message.kind);
                    return writer;
                };

                /**
                 * Decodes an ApplyInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.ApplyInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.ApplyInput} ApplyInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ApplyInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.ApplyInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.kind = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return ApplyInput;
            })();

            gomoku.ReplyInput = (function() {

                /**
                 * Properties of a ReplyInput.
                 * @memberof stmp.examples.gomoku
                 * @interface IReplyInput
                 * @property {stmp.examples.gomoku.ApplyModel.Kind|null} [kind] ReplyInput kind
                 * @property {boolean|null} [accept] ReplyInput accept
                 */

                /**
                 * Constructs a new ReplyInput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a ReplyInput.
                 * @implements IReplyInput
                 * @constructor
                 * @param {stmp.examples.gomoku.IReplyInput=} [properties] Properties to set
                 */
                function ReplyInput(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ReplyInput kind.
                 * @member {stmp.examples.gomoku.ApplyModel.Kind} kind
                 * @memberof stmp.examples.gomoku.ReplyInput
                 * @instance
                 */
                ReplyInput.prototype.kind = 0;

                /**
                 * ReplyInput accept.
                 * @member {boolean} accept
                 * @memberof stmp.examples.gomoku.ReplyInput
                 * @instance
                 */
                ReplyInput.prototype.accept = false;

                /**
                 * Encodes the specified ReplyInput message. Does not implicitly {@link stmp.examples.gomoku.ReplyInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.ReplyInput
                 * @static
                 * @param {stmp.examples.gomoku.IReplyInput} message ReplyInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ReplyInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.kind != null && message.hasOwnProperty("kind"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int32(message.kind);
                    if (message.accept != null && message.hasOwnProperty("accept"))
                        writer.uint32(/* id 2, wireType 0 =*/16).bool(message.accept);
                    return writer;
                };

                /**
                 * Decodes a ReplyInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.ReplyInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.ReplyInput} ReplyInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ReplyInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.ReplyInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.kind = reader.int32();
                            break;
                        case 2:
                            message.accept = reader.bool();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return ReplyInput;
            })();

            gomoku.GomokuService = (function() {

                /**
                 * Constructs a new GomokuService service.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a GomokuService
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function GomokuService(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (GomokuService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = GomokuService;

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuService#play}.
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @typedef PlayCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls Play.
                 * @function play
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @instance
                 * @param {stmp.examples.gomoku.IHandModel} request HandModel message or plain object
                 * @param {stmp.examples.gomoku.GomokuService.PlayCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuService.prototype.play = function play(request, callback) {
                    return this.rpcCall(play, $root.stmp.examples.gomoku.HandModel, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "Play" });

                /**
                 * Calls Play.
                 * @function play
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @instance
                 * @param {stmp.examples.gomoku.IHandModel} request HandModel message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuService#apply}.
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @typedef ApplyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls Apply.
                 * @function apply
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @instance
                 * @param {stmp.examples.gomoku.IApplyInput} request ApplyInput message or plain object
                 * @param {stmp.examples.gomoku.GomokuService.ApplyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuService.prototype.apply = function apply(request, callback) {
                    return this.rpcCall(apply, $root.stmp.examples.gomoku.ApplyInput, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "Apply" });

                /**
                 * Calls Apply.
                 * @function apply
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @instance
                 * @param {stmp.examples.gomoku.IApplyInput} request ApplyInput message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuService#reply}.
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @typedef ReplyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls Reply.
                 * @function reply
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @instance
                 * @param {stmp.examples.gomoku.IReplyInput} request ReplyInput message or plain object
                 * @param {stmp.examples.gomoku.GomokuService.ReplyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuService.prototype.reply = function reply(request, callback) {
                    return this.rpcCall(reply, $root.stmp.examples.gomoku.ReplyInput, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "Reply" });

                /**
                 * Calls Reply.
                 * @function reply
                 * @memberof stmp.examples.gomoku.GomokuService
                 * @instance
                 * @param {stmp.examples.gomoku.IReplyInput} request ReplyInput message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                return GomokuService;
            })();

            gomoku.UserPlayEvent = (function() {

                /**
                 * Properties of a UserPlayEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserPlayEvent
                 * @property {number|Long|null} [userId] UserPlayEvent userId
                 * @property {stmp.examples.gomoku.IHandModel|null} [hand] UserPlayEvent hand
                 */

                /**
                 * Constructs a new UserPlayEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserPlayEvent.
                 * @implements IUserPlayEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserPlayEvent=} [properties] Properties to set
                 */
                function UserPlayEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserPlayEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserPlayEvent
                 * @instance
                 */
                UserPlayEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * UserPlayEvent hand.
                 * @member {stmp.examples.gomoku.IHandModel|null|undefined} hand
                 * @memberof stmp.examples.gomoku.UserPlayEvent
                 * @instance
                 */
                UserPlayEvent.prototype.hand = null;

                /**
                 * Encodes the specified UserPlayEvent message. Does not implicitly {@link stmp.examples.gomoku.UserPlayEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserPlayEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserPlayEvent} message UserPlayEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserPlayEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    if (message.hand != null && message.hasOwnProperty("hand"))
                        $root.stmp.examples.gomoku.HandModel.encode(message.hand, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
                };

                /**
                 * Decodes a UserPlayEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserPlayEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserPlayEvent} UserPlayEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserPlayEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserPlayEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
                            break;
                        case 2:
                            message.hand = $root.stmp.examples.gomoku.HandModel.decode(reader, reader.uint32());
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserPlayEvent;
            })();

            gomoku.UserApplyEvent = (function() {

                /**
                 * Properties of a UserApplyEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserApplyEvent
                 * @property {stmp.examples.gomoku.ApplyModel.Kind|null} [kind] UserApplyEvent kind
                 */

                /**
                 * Constructs a new UserApplyEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserApplyEvent.
                 * @implements IUserApplyEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserApplyEvent=} [properties] Properties to set
                 */
                function UserApplyEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserApplyEvent kind.
                 * @member {stmp.examples.gomoku.ApplyModel.Kind} kind
                 * @memberof stmp.examples.gomoku.UserApplyEvent
                 * @instance
                 */
                UserApplyEvent.prototype.kind = 0;

                /**
                 * Encodes the specified UserApplyEvent message. Does not implicitly {@link stmp.examples.gomoku.UserApplyEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserApplyEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserApplyEvent} message UserApplyEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserApplyEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.kind != null && message.hasOwnProperty("kind"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int32(message.kind);
                    return writer;
                };

                /**
                 * Decodes a UserApplyEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserApplyEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserApplyEvent} UserApplyEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserApplyEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserApplyEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.kind = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserApplyEvent;
            })();

            gomoku.UserReplyEvent = (function() {

                /**
                 * Properties of a UserReplyEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserReplyEvent
                 * @property {stmp.examples.gomoku.ApplyModel.Kind|null} [kind] UserReplyEvent kind
                 * @property {boolean|null} [accepted] UserReplyEvent accepted
                 */

                /**
                 * Constructs a new UserReplyEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserReplyEvent.
                 * @implements IUserReplyEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserReplyEvent=} [properties] Properties to set
                 */
                function UserReplyEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserReplyEvent kind.
                 * @member {stmp.examples.gomoku.ApplyModel.Kind} kind
                 * @memberof stmp.examples.gomoku.UserReplyEvent
                 * @instance
                 */
                UserReplyEvent.prototype.kind = 0;

                /**
                 * UserReplyEvent accepted.
                 * @member {boolean} accepted
                 * @memberof stmp.examples.gomoku.UserReplyEvent
                 * @instance
                 */
                UserReplyEvent.prototype.accepted = false;

                /**
                 * Encodes the specified UserReplyEvent message. Does not implicitly {@link stmp.examples.gomoku.UserReplyEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserReplyEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserReplyEvent} message UserReplyEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserReplyEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.kind != null && message.hasOwnProperty("kind"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int32(message.kind);
                    if (message.accepted != null && message.hasOwnProperty("accepted"))
                        writer.uint32(/* id 2, wireType 0 =*/16).bool(message.accepted);
                    return writer;
                };

                /**
                 * Decodes a UserReplyEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserReplyEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserReplyEvent} UserReplyEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserReplyEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserReplyEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.kind = reader.int32();
                            break;
                        case 2:
                            message.accepted = reader.bool();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserReplyEvent;
            })();

            gomoku.UserDisconnectedEvent = (function() {

                /**
                 * Properties of a UserDisconnectedEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserDisconnectedEvent
                 * @property {number|Long|null} [userId] UserDisconnectedEvent userId
                 * @property {number|null} [waitTimeout] UserDisconnectedEvent waitTimeout
                 */

                /**
                 * Constructs a new UserDisconnectedEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserDisconnectedEvent.
                 * @implements IUserDisconnectedEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserDisconnectedEvent=} [properties] Properties to set
                 */
                function UserDisconnectedEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserDisconnectedEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserDisconnectedEvent
                 * @instance
                 */
                UserDisconnectedEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * UserDisconnectedEvent waitTimeout.
                 * @member {number} waitTimeout
                 * @memberof stmp.examples.gomoku.UserDisconnectedEvent
                 * @instance
                 */
                UserDisconnectedEvent.prototype.waitTimeout = 0;

                /**
                 * Encodes the specified UserDisconnectedEvent message. Does not implicitly {@link stmp.examples.gomoku.UserDisconnectedEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserDisconnectedEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserDisconnectedEvent} message UserDisconnectedEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserDisconnectedEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    if (message.waitTimeout != null && message.hasOwnProperty("waitTimeout"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int32(message.waitTimeout);
                    return writer;
                };

                /**
                 * Decodes a UserDisconnectedEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserDisconnectedEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserDisconnectedEvent} UserDisconnectedEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserDisconnectedEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserDisconnectedEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
                            break;
                        case 2:
                            message.waitTimeout = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserDisconnectedEvent;
            })();

            gomoku.UserConnectedEvent = (function() {

                /**
                 * Properties of a UserConnectedEvent.
                 * @memberof stmp.examples.gomoku
                 * @interface IUserConnectedEvent
                 * @property {number|Long|null} [userId] UserConnectedEvent userId
                 */

                /**
                 * Constructs a new UserConnectedEvent.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a UserConnectedEvent.
                 * @implements IUserConnectedEvent
                 * @constructor
                 * @param {stmp.examples.gomoku.IUserConnectedEvent=} [properties] Properties to set
                 */
                function UserConnectedEvent(properties) {
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * UserConnectedEvent userId.
                 * @member {number|Long} userId
                 * @memberof stmp.examples.gomoku.UserConnectedEvent
                 * @instance
                 */
                UserConnectedEvent.prototype.userId = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * Encodes the specified UserConnectedEvent message. Does not implicitly {@link stmp.examples.gomoku.UserConnectedEvent.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.UserConnectedEvent
                 * @static
                 * @param {stmp.examples.gomoku.IUserConnectedEvent} message UserConnectedEvent message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                UserConnectedEvent.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.userId != null && message.hasOwnProperty("userId"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.userId);
                    return writer;
                };

                /**
                 * Decodes a UserConnectedEvent message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.UserConnectedEvent
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.UserConnectedEvent} UserConnectedEvent
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                UserConnectedEvent.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.UserConnectedEvent();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.userId = reader.int64();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return UserConnectedEvent;
            })();

            gomoku.GomokuEvents = (function() {

                /**
                 * Constructs a new GomokuEvents service.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a GomokuEvents
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function GomokuEvents(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (GomokuEvents.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = GomokuEvents;

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuEvents#gameStart}.
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @typedef GameStartCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls GameStart.
                 * @function gameStart
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IGomokuModel} request GomokuModel message or plain object
                 * @param {stmp.examples.gomoku.GomokuEvents.GameStartCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuEvents.prototype.gameStart = function gameStart(request, callback) {
                    return this.rpcCall(gameStart, $root.stmp.examples.gomoku.GomokuModel, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "GameStart" });

                /**
                 * Calls GameStart.
                 * @function gameStart
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IGomokuModel} request GomokuModel message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuEvents#userPlay}.
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @typedef UserPlayCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserPlay.
                 * @function userPlay
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserPlayEvent} request UserPlayEvent message or plain object
                 * @param {stmp.examples.gomoku.GomokuEvents.UserPlayCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuEvents.prototype.userPlay = function userPlay(request, callback) {
                    return this.rpcCall(userPlay, $root.stmp.examples.gomoku.UserPlayEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserPlay" });

                /**
                 * Calls UserPlay.
                 * @function userPlay
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserPlayEvent} request UserPlayEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuEvents#userApply}.
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @typedef UserApplyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserApply.
                 * @function userApply
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserApplyEvent} request UserApplyEvent message or plain object
                 * @param {stmp.examples.gomoku.GomokuEvents.UserApplyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuEvents.prototype.userApply = function userApply(request, callback) {
                    return this.rpcCall(userApply, $root.stmp.examples.gomoku.UserApplyEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserApply" });

                /**
                 * Calls UserApply.
                 * @function userApply
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserApplyEvent} request UserApplyEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuEvents#userReply}.
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @typedef UserReplyCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserReply.
                 * @function userReply
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserReplyEvent} request UserReplyEvent message or plain object
                 * @param {stmp.examples.gomoku.GomokuEvents.UserReplyCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuEvents.prototype.userReply = function userReply(request, callback) {
                    return this.rpcCall(userReply, $root.stmp.examples.gomoku.UserReplyEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserReply" });

                /**
                 * Calls UserReply.
                 * @function userReply
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserReplyEvent} request UserReplyEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuEvents#userDisconnected}.
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @typedef UserDisconnectedCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserDisconnected.
                 * @function userDisconnected
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserDisconnectedEvent} request UserDisconnectedEvent message or plain object
                 * @param {stmp.examples.gomoku.GomokuEvents.UserDisconnectedCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuEvents.prototype.userDisconnected = function userDisconnected(request, callback) {
                    return this.rpcCall(userDisconnected, $root.stmp.examples.gomoku.UserDisconnectedEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserDisconnected" });

                /**
                 * Calls UserDisconnected.
                 * @function userDisconnected
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserDisconnectedEvent} request UserDisconnectedEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuEvents#userConnected}.
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @typedef UserConnectedCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls UserConnected.
                 * @function userConnected
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserConnectedEvent} request UserConnectedEvent message or plain object
                 * @param {stmp.examples.gomoku.GomokuEvents.UserConnectedCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuEvents.prototype.userConnected = function userConnected(request, callback) {
                    return this.rpcCall(userConnected, $root.stmp.examples.gomoku.UserConnectedEvent, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "UserConnected" });

                /**
                 * Calls UserConnected.
                 * @function userConnected
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IUserConnectedEvent} request UserConnectedEvent message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.GomokuEvents#gameOver}.
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @typedef GameOverCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls GameOver.
                 * @function gameOver
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IGomokuModel} request GomokuModel message or plain object
                 * @param {stmp.examples.gomoku.GomokuEvents.GameOverCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(GomokuEvents.prototype.gameOver = function gameOver(request, callback) {
                    return this.rpcCall(gameOver, $root.stmp.examples.gomoku.GomokuModel, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "GameOver" });

                /**
                 * Calls GameOver.
                 * @function gameOver
                 * @memberof stmp.examples.gomoku.GomokuEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IGomokuModel} request GomokuModel message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                return GomokuEvents;
            })();

            gomoku.LoginInput = (function() {

                /**
                 * Properties of a LoginInput.
                 * @memberof stmp.examples.gomoku
                 * @interface ILoginInput
                 * @property {string|null} [name] LoginInput name
                 */

                /**
                 * Constructs a new LoginInput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a LoginInput.
                 * @implements ILoginInput
                 * @constructor
                 * @param {stmp.examples.gomoku.ILoginInput=} [properties] Properties to set
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
                 * @memberof stmp.examples.gomoku.LoginInput
                 * @instance
                 */
                LoginInput.prototype.name = "";

                /**
                 * Encodes the specified LoginInput message. Does not implicitly {@link stmp.examples.gomoku.LoginInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.LoginInput
                 * @static
                 * @param {stmp.examples.gomoku.ILoginInput} message LoginInput message or plain object to encode
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
                 * @memberof stmp.examples.gomoku.LoginInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.LoginInput} LoginInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                LoginInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.LoginInput();
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

            gomoku.ListPlayerInput = (function() {

                /**
                 * Properties of a ListPlayerInput.
                 * @memberof stmp.examples.gomoku
                 * @interface IListPlayerInput
                 * @property {number|Long|null} [limit] ListPlayerInput limit
                 * @property {number|Long|null} [offset] ListPlayerInput offset
                 * @property {Array.<number|Long>|null} [ids] ListPlayerInput ids
                 */

                /**
                 * Constructs a new ListPlayerInput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a ListPlayerInput.
                 * @implements IListPlayerInput
                 * @constructor
                 * @param {stmp.examples.gomoku.IListPlayerInput=} [properties] Properties to set
                 */
                function ListPlayerInput(properties) {
                    this.ids = [];
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ListPlayerInput limit.
                 * @member {number|Long} limit
                 * @memberof stmp.examples.gomoku.ListPlayerInput
                 * @instance
                 */
                ListPlayerInput.prototype.limit = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListPlayerInput offset.
                 * @member {number|Long} offset
                 * @memberof stmp.examples.gomoku.ListPlayerInput
                 * @instance
                 */
                ListPlayerInput.prototype.offset = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListPlayerInput ids.
                 * @member {Array.<number|Long>} ids
                 * @memberof stmp.examples.gomoku.ListPlayerInput
                 * @instance
                 */
                ListPlayerInput.prototype.ids = $util.emptyArray;

                /**
                 * Encodes the specified ListPlayerInput message. Does not implicitly {@link stmp.examples.gomoku.ListPlayerInput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.ListPlayerInput
                 * @static
                 * @param {stmp.examples.gomoku.IListPlayerInput} message ListPlayerInput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListPlayerInput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.limit != null && message.hasOwnProperty("limit"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.limit);
                    if (message.offset != null && message.hasOwnProperty("offset"))
                        writer.uint32(/* id 2, wireType 0 =*/16).int64(message.offset);
                    if (message.ids != null && message.ids.length) {
                        writer.uint32(/* id 3, wireType 2 =*/26).fork();
                        for (var i = 0; i < message.ids.length; ++i)
                            writer.int64(message.ids[i]);
                        writer.ldelim();
                    }
                    return writer;
                };

                /**
                 * Decodes a ListPlayerInput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.ListPlayerInput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.ListPlayerInput} ListPlayerInput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListPlayerInput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.ListPlayerInput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.limit = reader.int64();
                            break;
                        case 2:
                            message.offset = reader.int64();
                            break;
                        case 3:
                            if (!(message.ids && message.ids.length))
                                message.ids = [];
                            if ((tag & 7) === 2) {
                                var end2 = reader.uint32() + reader.pos;
                                while (reader.pos < end2)
                                    message.ids.push(reader.int64());
                            } else
                                message.ids.push(reader.int64());
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return ListPlayerInput;
            })();

            gomoku.ListPlayerOutput = (function() {

                /**
                 * Properties of a ListPlayerOutput.
                 * @memberof stmp.examples.gomoku
                 * @interface IListPlayerOutput
                 * @property {number|Long|null} [total] ListPlayerOutput total
                 * @property {Array.<stmp.examples.gomoku.IPlayerModel>|null} [players] ListPlayerOutput players
                 */

                /**
                 * Constructs a new ListPlayerOutput.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a ListPlayerOutput.
                 * @implements IListPlayerOutput
                 * @constructor
                 * @param {stmp.examples.gomoku.IListPlayerOutput=} [properties] Properties to set
                 */
                function ListPlayerOutput(properties) {
                    this.players = [];
                    if (properties)
                        for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                            if (properties[keys[i]] != null)
                                this[keys[i]] = properties[keys[i]];
                }

                /**
                 * ListPlayerOutput total.
                 * @member {number|Long} total
                 * @memberof stmp.examples.gomoku.ListPlayerOutput
                 * @instance
                 */
                ListPlayerOutput.prototype.total = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

                /**
                 * ListPlayerOutput players.
                 * @member {Array.<stmp.examples.gomoku.IPlayerModel>} players
                 * @memberof stmp.examples.gomoku.ListPlayerOutput
                 * @instance
                 */
                ListPlayerOutput.prototype.players = $util.emptyArray;

                /**
                 * Encodes the specified ListPlayerOutput message. Does not implicitly {@link stmp.examples.gomoku.ListPlayerOutput.verify|verify} messages.
                 * @function encode
                 * @memberof stmp.examples.gomoku.ListPlayerOutput
                 * @static
                 * @param {stmp.examples.gomoku.IListPlayerOutput} message ListPlayerOutput message or plain object to encode
                 * @param {$protobuf.Writer} [writer] Writer to encode to
                 * @returns {$protobuf.Writer} Writer
                 */
                ListPlayerOutput.encode = function encode(message, writer) {
                    if (!writer)
                        writer = $Writer.create();
                    if (message.total != null && message.hasOwnProperty("total"))
                        writer.uint32(/* id 1, wireType 0 =*/8).int64(message.total);
                    if (message.players != null && message.players.length)
                        for (var i = 0; i < message.players.length; ++i)
                            $root.stmp.examples.gomoku.PlayerModel.encode(message.players[i], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                    return writer;
                };

                /**
                 * Decodes a ListPlayerOutput message from the specified reader or buffer.
                 * @function decode
                 * @memberof stmp.examples.gomoku.ListPlayerOutput
                 * @static
                 * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                 * @param {number} [length] Message length if known beforehand
                 * @returns {stmp.examples.gomoku.ListPlayerOutput} ListPlayerOutput
                 * @throws {Error} If the payload is not a reader or valid buffer
                 * @throws {$protobuf.util.ProtocolError} If required fields are missing
                 */
                ListPlayerOutput.decode = function decode(reader, length) {
                    if (!(reader instanceof $Reader))
                        reader = $Reader.create(reader);
                    var end = length === undefined ? reader.len : reader.pos + length, message = new $root.stmp.examples.gomoku.ListPlayerOutput();
                    while (reader.pos < end) {
                        var tag = reader.uint32();
                        switch (tag >>> 3) {
                        case 1:
                            message.total = reader.int64();
                            break;
                        case 2:
                            if (!(message.players && message.players.length))
                                message.players = [];
                            message.players.push($root.stmp.examples.gomoku.PlayerModel.decode(reader, reader.uint32()));
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                        }
                    }
                    return message;
                };

                return ListPlayerOutput;
            })();

            gomoku.PlayerService = (function() {

                /**
                 * Constructs a new PlayerService service.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a PlayerService
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function PlayerService(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (PlayerService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = PlayerService;

                /**
                 * Callback as used by {@link stmp.examples.gomoku.PlayerService#login}.
                 * @memberof stmp.examples.gomoku.PlayerService
                 * @typedef LoginCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.PlayerModel} [response] PlayerModel
                 */

                /**
                 * Calls Login.
                 * @function login
                 * @memberof stmp.examples.gomoku.PlayerService
                 * @instance
                 * @param {stmp.examples.gomoku.ILoginInput} request LoginInput message or plain object
                 * @param {stmp.examples.gomoku.PlayerService.LoginCallback} callback Node-style callback called with the error, if any, and PlayerModel
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(PlayerService.prototype.login = function login(request, callback) {
                    return this.rpcCall(login, $root.stmp.examples.gomoku.LoginInput, $root.stmp.examples.gomoku.PlayerModel, request, callback);
                }, "name", { value: "Login" });

                /**
                 * Calls Login.
                 * @function login
                 * @memberof stmp.examples.gomoku.PlayerService
                 * @instance
                 * @param {stmp.examples.gomoku.ILoginInput} request LoginInput message or plain object
                 * @returns {Promise<stmp.examples.gomoku.PlayerModel>} Promise
                 * @variation 2
                 */

                /**
                 * Callback as used by {@link stmp.examples.gomoku.PlayerService#listUser}.
                 * @memberof stmp.examples.gomoku.PlayerService
                 * @typedef ListUserCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.ListPlayerOutput} [response] ListPlayerOutput
                 */

                /**
                 * Calls ListUser.
                 * @function listUser
                 * @memberof stmp.examples.gomoku.PlayerService
                 * @instance
                 * @param {stmp.examples.gomoku.IListPlayerInput} request ListPlayerInput message or plain object
                 * @param {stmp.examples.gomoku.PlayerService.ListUserCallback} callback Node-style callback called with the error, if any, and ListPlayerOutput
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(PlayerService.prototype.listUser = function listUser(request, callback) {
                    return this.rpcCall(listUser, $root.stmp.examples.gomoku.ListPlayerInput, $root.stmp.examples.gomoku.ListPlayerOutput, request, callback);
                }, "name", { value: "ListUser" });

                /**
                 * Calls ListUser.
                 * @function listUser
                 * @memberof stmp.examples.gomoku.PlayerService
                 * @instance
                 * @param {stmp.examples.gomoku.IListPlayerInput} request ListPlayerInput message or plain object
                 * @returns {Promise<stmp.examples.gomoku.ListPlayerOutput>} Promise
                 * @variation 2
                 */

                return PlayerService;
            })();

            gomoku.PlayerEvents = (function() {

                /**
                 * Constructs a new PlayerEvents service.
                 * @memberof stmp.examples.gomoku
                 * @classdesc Represents a PlayerEvents
                 * @extends $protobuf.rpc.Service
                 * @constructor
                 * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                 * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                 * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                 */
                function PlayerEvents(rpcImpl, requestDelimited, responseDelimited) {
                    $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                }

                (PlayerEvents.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = PlayerEvents;

                /**
                 * Callback as used by {@link stmp.examples.gomoku.PlayerEvents#statusUpdated}.
                 * @memberof stmp.examples.gomoku.PlayerEvents
                 * @typedef StatusUpdatedCallback
                 * @type {function}
                 * @param {Error|null} error Error, if any
                 * @param {stmp.examples.gomoku.Empty} [response] Empty
                 */

                /**
                 * Calls StatusUpdated.
                 * @function statusUpdated
                 * @memberof stmp.examples.gomoku.PlayerEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IPlayerModel} request PlayerModel message or plain object
                 * @param {stmp.examples.gomoku.PlayerEvents.StatusUpdatedCallback} callback Node-style callback called with the error, if any, and Empty
                 * @returns {undefined}
                 * @variation 1
                 */
                Object.defineProperty(PlayerEvents.prototype.statusUpdated = function statusUpdated(request, callback) {
                    return this.rpcCall(statusUpdated, $root.stmp.examples.gomoku.PlayerModel, $root.stmp.examples.gomoku.Empty, request, callback);
                }, "name", { value: "StatusUpdated" });

                /**
                 * Calls StatusUpdated.
                 * @function statusUpdated
                 * @memberof stmp.examples.gomoku.PlayerEvents
                 * @instance
                 * @param {stmp.examples.gomoku.IPlayerModel} request PlayerModel message or plain object
                 * @returns {Promise<stmp.examples.gomoku.Empty>} Promise
                 * @variation 2
                 */

                return PlayerEvents;
            })();

            return gomoku;
        })();

        return examples;
    })();

    return stmp;
})();

module.exports = $root;
