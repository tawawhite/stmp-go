/*!
 * Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
 * @since 2020-01-02 18:11:53
 *
 * fake stmp module wait for implement
 */

declare module 'stmp' {

    export interface Context {
    }

    export interface MethodMetadata<I = any, O = any> {
        method: string;
        action: string;
        binAction: Uint8Array;
        input: new () => I,
        output: new () => O,
    }

    export function registerMethodAction<I, O>(method: string, action: string, input: new () => I, output: new() => O): void;

    export type Payload = string | Uint8Array
    export type MiddlewareFunc = (ctx: Context, method: MethodMetadata, payload: Payload) => void | Promise<void>
    export type InterceptFunc = (ctx: Context, method: MethodMetadata, payload: Payload) => boolean | Promise<boolean>
    export type HandlerFunc<I = any, O = any> = (ctx: Context, input: I, output: O) => void | Promise<void>

    export class Router {
        middleware(handler: MiddlewareFunc): void

        intercept(handler: InterceptFunc): void

        register<I, O>(inst: any, method: string, handler: HandlerFunc<I, O>): void

        unregister(inst: any, method: string): void
    }

    export class Server extends Router {
        broadcast<I>(sendOptions: SendOptions<I>, filter?: (conn: Connection) => boolean): void

        notify<I>(sendOptions: SendOptions<I>, ...conns: Connection[]): void
    }

    export class SendOptions<I> {
        input: I;
        method: string;

        constructor(method: string, input: I)
    }

    export interface CallOptions {
        notify?: boolean;
    }

    export class Connection extends Router {
        invoke<I, O>(send: SendOptions<I>, options?: CallOptions): Promise<O>
    }

    export type Header = Record<string, string[]>

    export interface DialOptions {
        header: Header
    }

    export function dialWebSocket(addr: string, options?: Partial<DialOptions>): Promise<Connection>;
}
