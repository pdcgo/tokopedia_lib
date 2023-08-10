/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from "axios";
import { useState } from "react";
import type { SdkConfig, Response as ErrResponse } from "./sdk_types";

const isDev = import.meta.env.DEV

const client = axios.create({
    baseURL: "http://localhost:5000",
    timeout: 30_000,
});

export type UseQueryOptions<D = any, E = Error> = {
    signal?: AbortSignal
    onSuccess?: (data: D) => void
    onError?: (err: E) => void
}

export type SenderConfigs<Method = unknown, Path = unknown, Payload = unknown, Params = unknown> = {
    method: Method
    path: Path
    payload?: Payload
    params?: Params
}

function useRequestRaw<T extends keyof SdkConfig, K extends SdkConfig>(_key: T, options?: UseQueryOptions<K[T]['response'], ErrResponse>) {
    async function sender(config: SenderConfigs<K[T]['method'], K[T]['path'], K[T]['payload'], K[T]['params']>, senderOptions?: UseQueryOptions<K[T]['response'], ErrResponse>) {
        try {
            const { data } = await client({
                method: config.method,
                data: config.payload,
                url: config.path,
                params: config.params,
                signal: senderOptions?.signal
            });

            if (data?.error) {
                const err = { error: data.error, msg: data.msg }

                options?.onError?.(err)
                senderOptions?.onError?.(err)
            } else {
                options?.onSuccess?.(data as K[T]['response'])
                senderOptions?.onSuccess?.(data as K[T]['response'])
            }
        } catch (error) {
            if (isDev) console.log(error)

            const err = { error: String(error), msg: "Unpredictable Error Type" }

            options?.onError?.(err)
            senderOptions?.onError?.(err)
        }
    }

    return { sender }
}

function useRequest<T extends keyof SdkConfig, K extends SdkConfig>(_key: T, options?: UseQueryOptions<K[T]['response'], ErrResponse>) {
    const [pending, setPending] = useState(false)
    const [response, setResponse] = useState<K[T]['response'] | null>(null)
    const [error, setError] = useState<ErrResponse | null>(null)

    async function sender(config: SenderConfigs<K[T]['method'], K[T]['path'], K[T]['payload'], K[T]['params']>, senderOptions?: UseQueryOptions<K[T]['response'], ErrResponse>) {
        setPending(true)

        try {
            const { data } = await client({
                method: config.method,
                data: config.payload,
                url: config.path,
                params: config.params,
                signal: senderOptions?.signal
            });

            if (data?.error) {
                const err = { error: data.error, msg: data.msg }

                setError(err)
                setResponse(null)
                options?.onError?.(err)
                senderOptions?.onError?.(err)
            } else {
                setError(null)
                setResponse(data as K[T]['response'])
                options?.onSuccess?.(data as K[T]['response'])
                senderOptions?.onSuccess?.(data as K[T]['response'])
            }
        } catch (error) {
            if (isDev) console.log(error)

            const err = { error: String(error), msg: "Unpredictable Error Type" }

            setResponse(null)
            setError(err)
            options?.onError?.(err)
            senderOptions?.onError?.(err)
        } finally {
            setPending(false)
        }

        return {
            pending,
            response,
            error,
        }
    }

    return {
        sender,
        pending,
        response,
        error,
    }
}

export { useRequest, useRequestRaw }

