/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-01 22:10:15
 */

import * as pb from './room.pb'
import {CallOptions, Connection, Context, SendOptions, Server} from 'stmp'

export namespace STMP {
    // the top level namespace always is STMP
    // the child namespaces is the package in .proto file
    // Uppercase is used to avoid conflicts with package stmp's global namespace
    export namespace stmp {
        namespace examples {
            namespace room {
                class UserServiceServer {
                    /**
                     * register impl to srv, the impl's methods will be bound automatically
                     * @param srv
                     * @param inst
                     */
                    static register(srv: Server, inst: UserServiceServer): void

                    /**
                     * unregister impl from srv
                     * @param srv
                     * @param inst
                     */
                    static unregister(srv: Server, inst: UserServiceServer): void

                    /**
                     * unimplemented method listUser
                     *
                     * @param ctx
                     * @param input
                     * @param output
                     */
                    listUser(ctx: Context, input: pb.stmp.examples.room.ListUserInput, output: pb.stmp.examples.room.ListUserOutput): void
                }

                /**
                 * create send context for server side and client side
                 */
                class UserServiceBuilder {
                    static listUser(data: pb.stmp.examples.room.IListUserInput): SendOptions<pb.stmp.examples.room.ListUserInput>
                }

                /**
                 * the client wrapper for connection
                 */
                class UserServiceClient {
                    readonly conn: Connection;

                    constructor(conn: Connection)

                    listUser(data: pb.stmp.examples.room.IListUserInput, options?: CallOptions): Promise<pb.stmp.examples.room.ListUserOutput>
                }
            }
        }
    }
}
