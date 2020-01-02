/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-01 22:10:13
 */

import {registerMethodAction, SendOptions} from 'stmp'
import * as pb from './room.pb'

export const STMP = Object.create(null);

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

initNamespace(STMP, 'stmp.examples.room', (room) => {
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

    class UserServiceBuilder {
        /**
         * @param {pb.stmp.examples.room.IListUserInput} input
         * @returns {import("stmp").SendOptions.<pb.stmp.examples.room.ListUserInput>}
         */
        static listUser(input) {
            return new SendOptions("stmp.examples.room.UserService.ListUser", input)
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
         * @param {import("stmp").CallOptions} options
         * @returns {Promise.<pb.stmp.examples.room.ListUserOutput>}
         */
        listUser(input, options) {
            return this.conn.invoke(UserServiceBuilder.listUser(input), options)
        }
    }
});
