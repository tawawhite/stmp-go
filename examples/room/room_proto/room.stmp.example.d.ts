/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-01 22:10:15
 */

import {CallOptions, Connection, Context, SendOptions, Server} from 'stmp'
import pb from "./room.pb";

export = stmp

declare namespace stmp {
    // the top level namespace always is stmp
    // the child namespaces is the package in .proto file
    namespace stmp {
        namespace examples {
            namespace room {
                import ListUserInput = pb.stmp.examples.room.ListUserInput;
                import ListUserOutput = pb.stmp.examples.room.ListUserOutput;
                import IListUserInput = pb.stmp.examples.room.IListUserInput;

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
                    listUser(ctx: Context, input: ListUserInput, output: ListUserOutput): void | Promise<void>
                }

                /**
                 * create send context for server side and client side
                 */
                class UserServiceBuilder {
                    static listUser(data: IListUserInput): SendOptions<ListUserInput>
                }

                /**
                 * the client wrapper for connection
                 */
                class UserServiceClient {
                    readonly conn: Connection;

                    constructor(conn: Connection)

                    listUser(data: IListUserInput, options?: CallOptions): Promise<ListUserOutput>
                }
            }
        }
    }
}
