/* eslint-disable @typescript-eslint/no-explicit-any*/

import axios from "axios";
import { useState } from "react";
import { ClientReturn, Clients, MaybeNull, SendOptions, Target, clients } from "./newapisdk";

export interface MutationClientReturn<Data, Query, Body, Err = Error> extends Omit<ClientReturn<Data, Query, Err>, "send">{
    mutate(a: SendOptions<Data, Query, Err>, b?: Partial<Body>): void
}

export type Mutate<K extends Target> = MutationClientReturn<
    Clients[K]["response"],
    Clients[K]["query"],
    Clients[K]["body"]
>["mutate"]

export function useMutation<
    K extends Target,
    R extends Clients[K]["response"],
    Q extends Clients[K]["query"],
    B extends Clients[K]["body"],
>(action: K, options?: SendOptions<R, Q>): MutationClientReturn<R, Q, B> {
    const uri = clients[action].url;
    const method = clients[action].method;
    const queryOptions = options;

    const [pending, setPending] = useState(false);
    const [data, setData] = useState<MaybeNull<R>>(null);
    const [error, setError] = useState<MaybeNull<Error>>(null);

    async function mutate(options: SendOptions<R, Q> | undefined = queryOptions, body?: Partial<B>) {
        setPending(true);

        const query = queryOptions?.query || options?.query;

        try {
            const { data } = await axios({
                method,
                url: uri,
                data: body,
                ...(query
                    ? {
                          params: query
                      }
                    : {})
            });

            queryOptions?.onSuccess?.(data);
            options?.onSuccess?.(data);
            setData(data);
            setError(null);
        } catch (e) {
            queryOptions?.onError?.(e as any);
            options?.onError?.(e as any);
            setError(e as any);
            setData(null);
        } finally {
            setPending(false);
        }
    }

    return {
        data,
        error,
        pending,
        mutate
    };
}
