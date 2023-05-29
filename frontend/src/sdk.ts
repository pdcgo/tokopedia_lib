import { AxiosInstance, AxiosResponse } from "axios";



class ClientSdk {
    client!: AxiosInstance

}
const clientSdk = new ClientSdk()

export function SetClient(client: AxiosInstance) {
    clientSdk.client = client
}


export async function GetPing(query: any) {
    let res = await clientSdk.client.get<any, AxiosResponse<any, any>, any>('/ping', {
        params: query,
    });
    return res.data;
}
