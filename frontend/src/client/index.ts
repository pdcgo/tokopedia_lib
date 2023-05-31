/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from "axios";
import { useState } from "react";

const client = axios.create({
    baseURL: "http://localhost:8080/tokopedia/"
});


type UseQueryOptions<D = any, E = Error> = {
    onSuccess?: (data: D) => void
    onError?: (err: E) => void
}

function useQuery() {
    const [pending, setPending] = useState(false)
    const [response, setResponse] = useState(null)
    const [error, setError] = useState(null)

    async function sender() {
        setPending(true)
        setError(null)
        setResponse(null)

        try {

            return
        } catch (error) {

            return
        } finally {

            setPending(false)
        }
    }

    return {
        sender,
        pending,
        response,
        error,
    }
}

export { useQuery };

