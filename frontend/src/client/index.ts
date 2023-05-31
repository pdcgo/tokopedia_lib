/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from "axios";
import { useState } from "react";
import type { SdkConfig, Response as ErrResponse } from "./sdk_types";

const isDev = import.meta.env.DEV

const client = axios.create({
    baseURL: "http://localhost:8080"
});

type UseQueryOptions<D = any, E = Error> = {
    onSuccess?: (data: D) => void
    onError?: (err: E) => void
}

type SenderConfigs<Method = unknown, Path = unknown, Payload = unknown, Params = unknown> = {
    method: Method
    path: Path
    payload?: Payload
    params?: Params
}

function useRequest<T extends keyof SdkConfig, K extends SdkConfig>(_key: T, options?: UseQueryOptions<K[T]['response'], ErrResponse>) {
    const [pending, setPending] = useState(false)
    const [response, setResponse] = useState<K[T]['response'] | null>(null)
    const [error, setError] = useState<ErrResponse | null>(null)

    async function sender(config: SenderConfigs<K[T]['method'], K[T]['path'], K[T]['payload'], K[T]['params']>) {
        setPending(true)
        
        try {
            const { data } = await client({
                method: config.method,
                data: config.payload,
                url: config.path,
                params: config.params,
            });
            
            if (data.error) {
                const err = { error: data.error, msg: data.msg }
                
                setError(err)
                setResponse(null)
                options?.onError?.(err)
            } else {
                setError(null)
                setResponse(data as K[T]['response'])
                options?.onSuccess?.(data as K[T]['response'])
            }
        } catch (error) {
            if (isDev) console.log(error)

            const err = { error: String(error), msg: "Unpredictable Error Type" }
            
            setResponse(null)
            setError(err)
            options?.onError?.(err)
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

export { useRequest }

