/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-01 22:10:13
 */

import {PayloadMap, registerMethodAction, notifyOptions} from 'stmp'
import * as pb from './room.pb'

const root = Object.create(null);

export default root

/**
 * @template {T}
 *
 * @callback Factory
 * @param {T} ns
 */

/**
 * @template {T}
 *
 * @param {*} root
 * @param {string} ns
 * @param {Factory.<T>} factories
 */
function initNamespace(root, ns, ...factories) {
    for (const item of ns.split('.')) {
        root = (root[item] = root[item] || Object.create(null))
    }
    for (const factory of factories) {
        factory(root)
    }
}

initNamespace(root, 'stmp.examples.room', (room) => {
    // each service takes a factory
    registerMethodAction("stmp.examples.room.UserService.ListUser", "1001", pb.stmp.examples.room.ListUserInput, pb.google.protobuf.Empty);

    class UserServiceServer {
        /**
         * @param {import("stmp").Server} srv
         * @param {UserServiceServer} inst
         */
        static register(srv, inst) {
            srv.register(inst, "stmp.examples.room.UserService.ListUser", inst.listUser);
        }

        /**
         * @param {import("stmp").Server} srv
         * @param {UserServiceServer} inst
         */
        static unregister(srv, inst) {
            srv.unregister(inst, "stmp.examples.room.UserService.ListUser");
        }

        /**
         * @param {import("stmp").Context} ctx
         * @param {pb.stmp.examples.room.ListUserInput} input
         * @param {pb.stmp.examples.room.ListUserOutput} output
         */
        listUser(ctx, input, output) {
            throw new Error("not implemented")
        }
    }

    class UserServiceBroadcaster {
        /**
         * @param {pb.stmp.examples.room.IListUserInput} input
         * @param {import("stmp").Connection} conn
         * @param {Partial.<import("stmp").CallOptions>} [options]
         * @returns {Promise<pb.stmp.examples.room.ListUserOutput>}
         */
        static listUser(input, conn, options) {
            return conn.invoke("stmp.examples.room.UserService.ListUser", input, options)
        }

        /**
         *
         * @param {pb.stmp.examples.room.IListUserInput} input
         * @param {Set.<import("stmp").Connection>} conns
         */
        static listUserForSet(input, conns) {
            const pm = new PayloadMap(input);
            for (const conn of conns) {
                conn.call("stmp.examples.room.UserService.ListUser", pm.get(conn), notifyOptions)
            }
        }

        /**
         * @param {pb.stmp.examples.room.IListUserInput} input
         * @param {import("stmp").Server} srv
         * @param {import("stmp").ConnFilter} [filter]
         */
        static broadcastListUser(input, srv, filter) {
            return srv.broadcast("stmp.examples.room.UserService.ListUser", input, filter)
        }

        static listUserMethod() {
            return "stmp.examples.room.UserService.ListUser"
        }

        static listUserAction() {
            return "1001"
        }
    }

    class UserServiceClient {
        /**
         * @param {import("stmp").Connection} conn
         */
        constructor(conn) {
            this.conn = conn
        }

        /**
         * @param {pb.stmp.examples.room.IListUserInput} input
         * @param {Partial.<import("stmp").CallOptions>} options
         * @returns {Promise.<pb.stmp.examples.room.ListUserOutput>}
         */
        listUser(input, options) {
            return this.conn.invoke("stmp.examples.room.UserService.ListUser", input, options)
        }
    }
});
